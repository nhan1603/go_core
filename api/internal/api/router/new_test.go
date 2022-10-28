package router

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	// Given:
	ctx := context.Background()
	corsOrigin := []string{"*"}

	// When:
	r := New(ctx, corsOrigin, true, nil)

	// Then:
	require.Equal(t, ctx, r.ctx)
	require.Equal(t, corsOrigin, r.corsOrigins)
	require.NotNil(t, r.healthRESTHandler)
	require.True(t, r.isGQLIntrospectionOn)
}
