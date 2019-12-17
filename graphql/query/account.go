package query

import (
	"context"
	"github.com/getsentry/raven-go"
	"github.com/oreqizer/go-relaygen/relay"
	"github.com/oreqizer/goiler/graphql/auth"
	"github.com/oreqizer/goiler/graphql/db"
	"github.com/oreqizer/goiler/graphql/schemas"
	"github.com/oreqizer/goiler/models"
)

// Account retrieves the currently logged-in account, nil otherwise
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

// Accounts lists all accounts for admins
func (Query) Accounts(
	ctx context.Context,
	after *string,
	first *int,
	before *string,
	last *int,
) (*schemas.AccountConnection, error) {
	dbi := db.GetDB(ctx)

	res, err := models.Accounts(
		db.QueryNotDeleted,
	).All(ctx, dbi)
	if err != nil {
		raven.CaptureError(err, nil)
		return nil, db.ErrFetchingResults
	}

	args := relay.ConnectionArgs{
		After:  after,
		First:  first,
		Before: before,
		Last:   last,
	}

	return schemas.AccountConnectionFromArray(schemas.Accounts(res).ToSlice(), &args), nil
}
