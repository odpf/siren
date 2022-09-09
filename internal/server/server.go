package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-openapi/runtime/middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/odpf/salt/log"
	"github.com/odpf/salt/mux"
	"github.com/odpf/siren/internal/api"
	"github.com/odpf/siren/internal/api/v1beta1"
	"github.com/odpf/siren/pkg/zaputil"
	swagger "github.com/odpf/siren/proto"
	sirenv1beta1 "github.com/odpf/siren/proto/odpf/siren/v1beta1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/newrelic/go-agent/v3/integrations/nrgrpc"
	"github.com/newrelic/go-agent/v3/newrelic"
)

const defaultGracePeriod = 5 * time.Second

type Config struct {
	Host string `mapstructure:"host" default:"localhost"`
	Port int    `mapstructure:"port" default:"8080"`
}

func (cfg Config) addr() string { return fmt.Sprintf("%s:%d", cfg.Host, cfg.Port) }

// RunServer runs the application server
func RunServer(
	c Config,
	logger log.Logger,
	nr *newrelic.Application,
	apiDeps *api.Deps) error {

	var err error

	// init grpc server
	zapLogger, err := zaputil.GRPCZapLogger(logger)
	if err != nil {
		return err
	}
	loggerOpts := []grpc_zap.Option{
		grpc_zap.WithLevels(zaputil.GRPCCodeToLevel),
		grpc_zap.WithTimestampFormat(time.RFC3339Nano),
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(),
			grpc_ctxtags.UnaryServerInterceptor(),
			nrgrpc.UnaryServerInterceptor(nr),
			grpc_validator.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(zapLogger, loggerOpts...),
		)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_recovery.StreamServerInterceptor(),
			grpc_ctxtags.StreamServerInterceptor(),
			nrgrpc.StreamServerInterceptor(nr),
			grpc_validator.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(zapLogger, loggerOpts...),
		)),
	)

	// init http proxy
	grpcDialCtx, grpcDialCancel := context.WithTimeout(context.Background(), time.Second*5)
	defer grpcDialCancel()

	grpcConn, err := grpc.DialContext(grpcDialCtx, c.addr(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	httpGateway := runtime.NewServeMux(
		runtime.WithErrorHandler(runtime.DefaultHTTPErrorHandler),
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.HTTPBodyMarshaler{
			Marshaler: &runtime.JSONPb{
				MarshalOptions: protojson.MarshalOptions{
					UseProtoNames: true,
				},
				UnmarshalOptions: protojson.UnmarshalOptions{
					DiscardUnknown: true,
				},
			},
		}),
		runtime.WithHealthEndpointAt(grpc_health_v1.NewHealthClient(grpcConn), "/ping"),
	)

	reflection.Register(grpcServer)

	runtimeCtx, runtimeCancel := context.WithCancel(context.Background())
	defer runtimeCancel()

	sirenServiceRPC := v1beta1.NewGRPCServer(
		nr,
		logger,
		apiDeps,
	)
	grpcServer.RegisterService(&sirenv1beta1.SirenService_ServiceDesc, sirenServiceRPC)
	grpcServer.RegisterService(&grpc_health_v1.Health_ServiceDesc, sirenServiceRPC)
	if err := sirenv1beta1.RegisterSirenServiceHandler(runtimeCtx, httpGateway, grpcConn); err != nil {
		return err
	}

	baseMux := http.NewServeMux()
	baseMux.HandleFunc("/siren.swagger.yaml", func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.FS(swagger.File)).ServeHTTP(w, r)
	})
	baseMux.Handle("/documentation", middleware.SwaggerUI(middleware.SwaggerUIOpts{
		SpecURL: "/siren.swagger.yaml",
		Path:    "documentation",
	}, http.NotFoundHandler()))
	baseMux.Handle("/", httpGateway)

	logger.Info("server is running", "host", c.Host, "port", c.Port)
	mux.Serve(runtimeCtx, c.addr(),
		mux.WithHTTP(&http.Server{Handler: baseMux}),
		mux.WithGRPC(grpcServer),
		mux.WithGracePeriod(defaultGracePeriod),
	)

	return nil
}
