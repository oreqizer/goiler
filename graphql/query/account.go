package query

import (
	"context"
	"github.com/getsentry/raven-go"
	"github.com/oreqizer/goiler/graphql/auth"
	"github.com/oreqizer/goiler/graphql/db"
	"github.com/oreqizer/goiler/graphql/schemas"
	"github.com/oreqizer/goiler/models"
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

func (Query) Accounts(ctx context.Context) ([]*schemas.Account, error) {
	dbi := db.GetDB(ctx)

	res, err := models.Accounts(
		db.QueryNotDeleted,
	).All(ctx, dbi)
	if err != nil {
		raven.CaptureError(err, nil)
		return nil, db.ErrFetchingResults
	}

	return schemas.Accounts(res).ToSlice(), nil
}
