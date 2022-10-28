package user

import (
	"context"

	"go_core/api/internal/repository"
)

// ApiRestController represents the specification of this pkg
type ApiRestController interface {
	CreateObject(context.Context, string) (int, error)
}

// New initializes a new Controller instance and returns it
func New(repo repository.Registry) ApiRestController {
	return impl{repo: repo}
}

type impl struct {
	repo repository.Registry
}
