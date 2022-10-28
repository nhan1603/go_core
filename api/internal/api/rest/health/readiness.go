package health

import (
	"context"
	"errors"
	"net/http"

	"go_core/api/pkg/httpserv"
)

// CheckReadiness checks for system readiness
func (h Handler) CheckReadiness() http.HandlerFunc {
	return httpserv.ErrHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		err := h.systemCtrl.CheckReadiness(r.Context())

		if errors.Is(err, context.Canceled) {
			return nil
		}

		return err
	})
}
