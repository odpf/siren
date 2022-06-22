package v1beta1

import (
	"context"
	"testing"

	"github.com/odpf/salt/log"
	"github.com/odpf/siren/core/template"
	sirenv1beta1 "github.com/odpf/siren/internal/server/proto/odpf/siren/v1beta1"
	"github.com/odpf/siren/internal/server/v1beta1/mocks"
	"github.com/odpf/siren/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestGRPCServer_ListTemplates(t *testing.T) {
	t.Run("should return list of all templates", func(t *testing.T) {
		mockedTemplateService := &mocks.TemplateService{}
		dummyGRPCServer := GRPCServer{
			templateService: mockedTemplateService,
			logger:          log.NewNoop(),
		}
		dummyReq := &sirenv1beta1.ListTemplatesRequest{}
		dummyResult := []template.Template{
			{
				ID:   1,
				Name: "foo",
				Body: "bar",
				Tags: []string{"foo", "bar"},
				Variables: []template.Variable{
					{
						Name:        "foo",
						Type:        "bar",
						Default:     "",
						Description: "",
					},
				},
			},
		}

		mockedTemplateService.EXPECT().Index("").
			Return(dummyResult, nil).Once()
		res, err := dummyGRPCServer.ListTemplates(context.Background(), dummyReq)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(res.GetTemplates()))
		assert.Equal(t, "foo", res.GetTemplates()[0].GetName())
		assert.Equal(t, "bar", res.GetTemplates()[0].GetBody())
		assert.Equal(t, 1, len(res.GetTemplates()[0].GetVariables()))
		assert.Equal(t, "foo", res.GetTemplates()[0].GetVariables()[0].GetName())
	})

	t.Run("should return list of all templates matching particular tag", func(t *testing.T) {
		mockedTemplateService := &mocks.TemplateService{}
		dummyGRPCServer := GRPCServer{
			templateService: mockedTemplateService,
			logger:          log.NewNoop(),
		}
		dummyReq := &sirenv1beta1.ListTemplatesRequest{
			Tag: "foo",
		}

		dummyResult := []template.Template{
			{
				ID:   1,
				Name: "foo",
				Body: "bar",
				Tags: []string{"foo", "bar"},
				Variables: []template.Variable{
					{
						Name:        "foo",
						Type:        "bar",
						Default:     "",
						Description: "",
					},
				},
			},
		}

		mockedTemplateService.EXPECT().Index("foo").
			Return(dummyResult, nil).Once()
		res, err := dummyGRPCServer.ListTemplates(context.Background(), dummyReq)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(res.GetTemplates()))
		assert.Equal(t, "foo", res.GetTemplates()[0].GetName())
		assert.Equal(t, "bar", res.GetTemplates()[0].GetBody())
		assert.Equal(t, 1, len(res.GetTemplates()[0].GetVariables()))
		assert.Equal(t, "foo", res.GetTemplates()[0].GetVariables()[0].GetName())
	})

	t.Run("should return error Internal if getting templates failed", func(t *testing.T) {
		mockedTemplateService := &mocks.TemplateService{}
		dummyGRPCServer := GRPCServer{
			templateService: mockedTemplateService,
			logger:          log.NewNoop(),
		}
		dummyReq := &sirenv1beta1.ListTemplatesRequest{
			Tag: "foo",
		}
		mockedTemplateService.EXPECT().Index("foo").
			Return(nil, errors.New("random error")).Once()
		res, err := dummyGRPCServer.ListTemplates(context.Background(), dummyReq)
		assert.Nil(t, res)
		assert.EqualError(t, err, "rpc error: code = Internal desc = some unexpected error occurred")
	})
}

func TestGRPCServer_GetTemplate(t *testing.T) {
	t.Run("should return template by name", func(t *testing.T) {
		mockedTemplateService := &mocks.TemplateService{}
		dummyGRPCServer := GRPCServer{
			templateService: mockedTemplateService,
			logger:          log.NewNoop(),
		}
		dummyReq := &sirenv1beta1.GetTemplateRequest{
			Name: "foo",
		}
		dummyResult := &template.Template{
			ID:   1,
			Name: "foo",
			Body: "bar",
			Tags: []string{"foo", "bar"},
			Variables: []template.Variable{
				{
					Name:        "foo",
					Type:        "bar",
					Default:     "",
					Description: "",
				},
			},
		}

		mockedTemplateService.EXPECT().GetByName("foo").
			Return(dummyResult, nil).Once()
		res, err := dummyGRPCServer.GetTemplate(context.Background(), dummyReq)
		assert.Nil(t, err)
		assert.Equal(t, uint64(1), res.GetTemplate().GetId())
		assert.Equal(t, "foo", res.GetTemplate().GetName())
		assert.Equal(t, "bar", res.GetTemplate().GetBody())
		assert.Equal(t, "foo", res.GetTemplate().GetVariables()[0].GetName())
		mockedTemplateService.AssertCalled(t, "GetByName", dummyReq.Name)
	})

	t.Run("should return error Not Found if template does not exist", func(t *testing.T) {
		mockedTemplateService := &mocks.TemplateService{}
		dummyGRPCServer := GRPCServer{
			templateService: mockedTemplateService,
			logger:          log.NewNoop(),
		}
		dummyReq := &sirenv1beta1.GetTemplateRequest{
			Name: "foo",
		}
		mockedTemplateService.EXPECT().GetByName("foo").
			Return(nil, errors.ErrNotFound).Once()
		res, err := dummyGRPCServer.GetTemplate(context.Background(), dummyReq)
		assert.Nil(t, res)
		assert.EqualError(t, err, "rpc error: code = NotFound desc = requested entity not found")
	})

	t.Run("should return error Internal if getting template by name failed", func(t *testing.T) {
		mockedTemplateService := &mocks.TemplateService{}
		dummyGRPCServer := GRPCServer{
			templateService: mockedTemplateService,
			logger:          log.NewNoop(),
		}
		dummyReq := &sirenv1beta1.GetTemplateRequest{
			Name: "foo",
		}
		mockedTemplateService.EXPECT().GetByName("foo").
			Return(nil, errors.New("random error")).Once()
		res, err := dummyGRPCServer.GetTemplate(context.Background(), dummyReq)
		assert.Nil(t, res)
		assert.EqualError(t, err, "rpc error: code = Internal desc = some unexpected error occurred")
	})
}

func TestGRPCServer_UpsertTemplate(t *testing.T) {
	dummyReq := &sirenv1beta1.UpsertTemplateRequest{
		Id:   1,
		Name: "foo",
		Body: "bar",
		Tags: []string{"foo", "bar"},
		Variables: []*sirenv1beta1.TemplateVariables{
			{
				Name:        "foo",
				Type:        "bar",
				Default:     "",
				Description: "",
			},
		},
	}
	template := &template.Template{
		ID:   1,
		Name: "foo",
		Body: "bar",
		Tags: []string{"foo", "bar"},
		Variables: []template.Variable{
			{
				Name:        "foo",
				Type:        "bar",
				Default:     "",
				Description: "",
			},
		},
	}

	t.Run("should return template by name", func(t *testing.T) {
		mockedTemplateService := &mocks.TemplateService{}
		dummyGRPCServer := GRPCServer{
			templateService: mockedTemplateService,
			logger:          log.NewNoop(),
		}

		mockedTemplateService.EXPECT().Upsert(template).Return(nil).Once()
		res, err := dummyGRPCServer.UpsertTemplate(context.Background(), dummyReq)
		assert.Nil(t, err)
		assert.Equal(t, uint64(1), res.GetId())
		mockedTemplateService.AssertCalled(t, "Upsert", template)
	})

	t.Run("should return error AlreadyExists if upsert template return err conflict", func(t *testing.T) {
		mockedTemplateService := &mocks.TemplateService{}
		dummyGRPCServer := GRPCServer{
			templateService: mockedTemplateService,
			logger:          log.NewNoop(),
		}
		mockedTemplateService.EXPECT().Upsert(template).Return(errors.ErrConflict).Once()
		res, err := dummyGRPCServer.UpsertTemplate(context.Background(), dummyReq)
		assert.Nil(t, res)
		assert.EqualError(t, err, "rpc error: code = AlreadyExists desc = an entity with conflicting identifier exists")
	})

	t.Run("should return error Internal if upsert template failed", func(t *testing.T) {
		mockedTemplateService := &mocks.TemplateService{}
		dummyGRPCServer := GRPCServer{
			templateService: mockedTemplateService,
			logger:          log.NewNoop(),
		}
		mockedTemplateService.EXPECT().Upsert(template).Return(errors.New("random error")).Once()
		res, err := dummyGRPCServer.UpsertTemplate(context.Background(), dummyReq)
		assert.Nil(t, res)
		assert.EqualError(t, err, "rpc error: code = Internal desc = some unexpected error occurred")
	})
}

func TestGRPCServer_DeleteTemplate(t *testing.T) {
	t.Run("should delete template", func(t *testing.T) {
		mockedTemplateService := &mocks.TemplateService{}
		dummyGRPCServer := GRPCServer{
			templateService: mockedTemplateService,
			logger:          log.NewNoop(),
		}
		dummyReq := &sirenv1beta1.DeleteTemplateRequest{
			Name: "foo",
		}

		mockedTemplateService.EXPECT().Delete("foo").
			Return(nil).Once()
		res, err := dummyGRPCServer.DeleteTemplate(context.Background(), dummyReq)
		assert.Nil(t, err)
		assert.Equal(t, &sirenv1beta1.DeleteTemplateResponse{}, res)
		mockedTemplateService.AssertCalled(t, "Delete", "foo")
	})

	t.Run("should return error Internal if deleting template failed", func(t *testing.T) {
		mockedTemplateService := &mocks.TemplateService{}
		dummyGRPCServer := GRPCServer{
			templateService: mockedTemplateService,
			logger:          log.NewNoop(),
		}
		dummyReq := &sirenv1beta1.DeleteTemplateRequest{
			Name: "foo",
		}
		mockedTemplateService.EXPECT().Delete("foo").
			Return(errors.New("random error")).Once()
		res, err := dummyGRPCServer.DeleteTemplate(context.Background(), dummyReq)
		assert.Nil(t, res)
		assert.EqualError(t, err, "rpc error: code = Internal desc = some unexpected error occurred")
	})
}

func TestGRPCServer_RenderTemplate(t *testing.T) {
	dummyReq := &sirenv1beta1.RenderTemplateRequest{
		Name: "foo",
		Variables: map[string]string{
			"foo": "bar",
		},
	}

	t.Run("should render template", func(t *testing.T) {
		mockedTemplateService := &mocks.TemplateService{}
		dummyGRPCServer := GRPCServer{
			templateService: mockedTemplateService,
			logger:          log.NewNoop(),
		}

		mockedTemplateService.EXPECT().Render("foo", dummyReq.GetVariables()).
			Return("random", nil).Once()
		res, err := dummyGRPCServer.RenderTemplate(context.Background(), dummyReq)
		assert.Nil(t, err)
		assert.Equal(t, "random", res.GetBody())
		mockedTemplateService.AssertCalled(t, "Render", "foo", dummyReq.GetVariables())
	})

	t.Run("should return error Internal if rendering template failed", func(t *testing.T) {
		mockedTemplateService := &mocks.TemplateService{}
		dummyGRPCServer := GRPCServer{
			templateService: mockedTemplateService,
			logger:          log.NewNoop(),
		}
		mockedTemplateService.EXPECT().Render("foo", dummyReq.GetVariables()).
			Return("", errors.New("random error")).Once()
		res, err := dummyGRPCServer.RenderTemplate(context.Background(), dummyReq)
		assert.Empty(t, res)
		assert.EqualError(t, err, "rpc error: code = Internal desc = some unexpected error occurred")
	})
}
