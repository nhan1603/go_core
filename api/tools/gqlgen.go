//go:build tools
// +build tools

package tools

import (
	_ "go_core/api/pkg/httpserv/gql/scalar"

	_ "github.com/99designs/gqlgen"
	_ "github.com/99designs/gqlgen/graphql"
	_ "github.com/99designs/gqlgen/graphql/introspection"
)
