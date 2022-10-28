package router

import (
	"context"
	"net/http"
	"sort"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

func TestRouter_Handler(t *testing.T) {
	// Given:
	r := New(
		context.Background(),
		[]string{},
		true,
		nil,
	)

	expectedRoutes := []string{
		"POST /api/add-friend",
		"POST /api/block",
		"POST /api/common-friend",
		"POST /api/create-user",
		"POST /api/friend-list",
		"POST /api/subscribe",
		"POST /api/update-receiver",
	}
	sort.Strings(expectedRoutes)
	var routesFound []string

	handler, ok := r.Handler().(chi.Router)
	if !ok {
		require.FailNow(t, "handler is not a chi router")
	}

	// When:
	err := chi.Walk(
		handler,
		func(method string, route string, _ http.Handler, _ ...func(http.Handler) http.Handler) error {
			routesFound = append(routesFound, method+" "+route) // TODO: Add middleware stack check also if possible.
			return nil
		},
	)
	require.NoError(t, err)
	sort.Strings(routesFound)

	// Then:
	require.EqualValues(t, expectedRoutes, routesFound)
}
