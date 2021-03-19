package api

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/odpf/siren/api/handlers"
	"github.com/odpf/siren/service"
)

// New initializes the service router
func New(container *service.Container) *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	r.Use(logger)

	// Route => handler
	r.Methods("GET").Path("/ping").Handler(handlers.Ping())

	r.Methods("GET").Path("/swagger.yaml").Handler(handlers.SwaggerFile())
	r.Methods("GET").Path("/documentation").Handler(middleware.SwaggerUI(middleware.SwaggerUIOpts{
		SpecURL: "/swagger.yaml",
		Path:    "documentation",
	}, r.NotFoundHandler))

	r.Methods("PUT").Path("/templates").Handler(handlers.UpsertTemplates(container.TemplatesService))
	r.Methods("GET").Path("/templates").Handler(handlers.IndexTemplates(container.TemplatesService))
	r.Methods("GET").Path("/templates/{name}").Handler(handlers.GetTemplates(container.TemplatesService))
	r.Methods("DELETE").Path("/templates/{name}").Handler(handlers.DeleteTemplates(container.TemplatesService))
	r.Methods("POST").Path("/templates/{name}/render").Handler(handlers.RenderTemplates(container.TemplatesService))
	r.Methods("PUT").Path("/alertingCredentials/teams/{teamName}").Handler(handlers.UpdateAlertCredentials(container.AlertmanagerService))

	r.Methods("PUT").Path("/rules").Handler(handlers.UpsertRule(container.RulesService))
	r.Methods("GET").Path("/rules").Handler(handlers.GetRules(container.RulesService))

	return r
}
