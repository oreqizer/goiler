package query

import (
	"context"
	"github.com/oreqizer/goiler/graphql/auth"
	"github.com/oreqizer/goiler/graphql/schemas"
)

func (Query) Account(ctx context.Context) (*schemas.Account, error) {
	a, err := auth.GetAuth(ctx)
	if err == auth.ErrorNotInContext {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	if a.Account == nil {
		return nil, nil
	}

	return &schemas.Account{Account: *a.Account}, nil
}
