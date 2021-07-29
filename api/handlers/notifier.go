package handlers

import (
	"encoding/json"
	"github.com/odpf/siren/domain"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strings"
)

// Notify handler
func Notify(notifierServices domain.NotifierServices, logger *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		provider := r.URL.Query().Get("provider")
		switch provider {
		case "slack":
			var payload domain.SlackMessage
			err := json.NewDecoder(r.Body).Decode(&payload)
			if err != nil {
				badRequest(w, err, logger)
				return
			}
			v := validator.New()
			_ = v.RegisterValidation("receiverTypeChecker", func(fl validator.FieldLevel) bool {
				return fl.Field().Interface().(string) == "user" || fl.Field().Interface().(string) == "channel"
			})
			err = v.Struct(payload)
			if err != nil {
				if _, ok := err.(*validator.InvalidValidationError); ok {
					logger.Error("invalid validation error")
					internalServerError(w, err, logger)
					return
				}
				badRequest(w, err, logger)
				return
			}
			result, err := notifierServices.Slack.Notify(&payload)
			if err != nil {
				if isBadRequest(err) {
					badRequest(w, err, logger)
					return
				}
				internalServerError(w, err, logger)
				return
			}
			returnJSON(w, result)
			return
		case "":
			badRequest(w, errors.New("provider not given in query params"), logger)
			return
		default:
			badRequest(w, errors.New("unrecognized provider"), logger)
			return
		}
	}
}

func isBadRequest(err error) bool {
	return strings.Contains(err.Error(), "failed to get id for") ||
		strings.Contains(err.Error(), "app is not part of")
}
