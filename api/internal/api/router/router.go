package router

import (
	"context"
	"net/http"

	"go_core/api/internal/api/rest/health"
	"go_core/api/pkg/httpserv"
)

// Router defines the routes & handlers of the app
type Router struct {
	ctx                  context.Context
	corsOrigins          []string
	isGQLIntrospectionOn bool
	healthRESTHandler    health.Handler
}

// Handler returns the Handler for use by the server
func (rtr Router) Handler() http.Handler {
	return httpserv.Handler(
		rtr.healthRESTHandler.CheckReadiness(),
	)
}
