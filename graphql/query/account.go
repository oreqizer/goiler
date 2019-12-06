package query

import (
	"context"
	"github.com/oreqizer/goiler/graphql/auth"
	"github.com/oreqizer/goiler/graphql/schemas"
)

func (Query) Account(ctx context.Context) (*schemas.Account, error) {
	a := auth.GetAuth(ctx)
	if a == nil {
		return nil, nil
	}

	if a.Account == nil {
		return nil, nil
	}

	return &schemas.Account{Account: *a.Account}, nil
}
