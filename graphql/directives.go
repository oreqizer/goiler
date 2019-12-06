package graphql

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/oreqizer/goiler/generated"
	"github.com/oreqizer/goiler/graphql/auth"
)

// Directives mount directive implementations
// https://gqlgen.com/reference/directives/
func Directives(c *generated.Config) {
	c.Directives.HasRole = func(
		ctx context.Context,
		obj interface{},
		next graphql.Resolver,
		role generated.Role,
	) (interface{}, error) {
		if role == generated.RoleAdmin {
			_, err := auth.GetAuthAdmin(ctx)
			if err != nil {
				return nil, err
			}
		}

		if role == generated.RoleUser {
			_, err := auth.GetAuthAccount(ctx)
			if err != nil {
				return nil, err
			}
		}

		return next(ctx)
	}
}
