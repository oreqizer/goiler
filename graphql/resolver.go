package graphql

import (
	"github.com/oreqizer/goiler/graphql/mutation"
	"github.com/oreqizer/goiler/graphql/query"

	"github.com/oreqizer/goiler/generated"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

// Resolver is the entrypoint to the schema
type Resolver struct{}

// Mutation holds all mutations
func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutation.Mutation{}
}

// Query holds all queries
func (r *Resolver) Query() generated.QueryResolver {
	return &query.Query{}
}
