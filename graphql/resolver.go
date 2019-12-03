package graphql

import (
	"github.com/oreqizer/goiler/graphql/mutation"
	"github.com/oreqizer/goiler/graphql/query"

	"github.com/oreqizer/goiler/generated"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutation.Mutation{}
}

func (r *Resolver) Query() generated.QueryResolver {
	return &query.Query{}
}
