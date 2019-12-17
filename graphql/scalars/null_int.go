package scalars

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/volatiletech/null"
)

// MarshalNullInt allows marshalling 'null.Int'
func MarshalNullInt(ni null.Int) graphql.Marshaler {
	if !ni.Valid {
		return graphql.Null
	}
	return graphql.MarshalInt(ni.Int)
}

// UnmarshalNullInt allows unmarshalling 'null.Int'
func UnmarshalNullInt(v interface{}) (null.Int, error) {
	if v == nil {
		return null.Int{Valid: false}, nil
	}
	i, err := graphql.UnmarshalInt(v)
	return null.Int{Int: i}, err
}
