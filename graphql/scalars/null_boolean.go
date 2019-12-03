package scalars

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/volatiletech/null"
)

func MarshalNullBoolean(ni null.Bool) graphql.Marshaler {
	if !ni.Valid {
		return graphql.Null
	}
	return graphql.MarshalBoolean(ni.Bool)
}

func UnmarshalNullBoolean(v interface{}) (null.Bool, error) {
	if v == nil {
		return null.Bool{Valid: false}, nil
	}
	i, err := graphql.UnmarshalBoolean(v)
	return null.Bool{Bool: i}, err
}
