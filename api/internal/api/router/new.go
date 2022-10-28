package router

import (
	"context"

	"go_core/api/internal/api/rest/health"
	"go_core/api/internal/controller/system"
)

// New creates and returns a new Router instance
func New(
	ctx context.Context,
	corsOrigin []string,
	isGQLIntrospectionOn bool,
	systemCtrl system.Controller,
) Router {
	return Router{
		ctx:                  ctx,
		corsOrigins:          corsOrigin,
		isGQLIntrospectionOn: isGQLIntrospectionOn,
		healthRESTHandler:    health.New(systemCtrl),
	}
}
