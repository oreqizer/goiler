package mutation

import (
	"context"
	"github.com/getsentry/raven-go"
	"github.com/oreqizer/goiler/generated"
	"github.com/oreqizer/goiler/graphql/auth"
	"github.com/oreqizer/goiler/graphql/db"
	"github.com/oreqizer/goiler/graphql/schemas"
	"github.com/oreqizer/goiler/models"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"time"
)

// UpsertAccount adds a new account
func (Mutation) UpsertAccount(
	ctx context.Context,
	input generated.UpsertAccountInput,
) (*generated.UpsertAccountPayload, error) {
	dbi := db.GetDB(ctx)
	a := auth.GetAuth(ctx)
	if a == nil {
		return nil, auth.ErrNoToken
	}

	model := models.Account{
		AuthID:  a.AuthID,
		Name:    input.Name,
		Surname: input.Surname,
		Email:   input.Email,
	}

	if err := model.Upsert(ctx, dbi, true, []string{"auth_id"}, boil.Infer(), boil.Infer()); err != nil {
		raven.CaptureError(err, nil)
		return nil, db.ErrDefault
	}

	res := generated.UpsertAccountPayload{
		Account:          &schemas.Account{Account: model},
		ClientMutationID: input.ClientMutationID,
	}

	return &res, nil
}

// DeleteAccount marks an account as deleted
func (Mutation) DeleteAccount(
	ctx context.Context,
	input generated.DeleteAccountInput,
) (*generated.DeleteAccountPayload, error) {
	dbi := db.GetDB(ctx)
	a, err := auth.GetAuthAccount(ctx)
	if err != nil {
		return nil, err
	}

	if !a.Account.DeletedAt.Valid {
		tx, err := dbi.BeginTx(ctx, nil)
		if err != nil {
			raven.CaptureError(err, nil)
			return nil, db.ErrDefault
		}

		// Delete the account
		a.Account.DeletedAt = null.TimeFrom(time.Now())

		if _, err := a.Account.Update(ctx, tx, boil.Infer()); err != nil {
			_ = tx.Rollback()
			raven.CaptureError(err, nil)
			return nil, db.ErrDefault
		}

		if err := tx.Commit(); err != nil {
			raven.CaptureError(err, nil)
			return nil, db.ErrDefault
		}
	}

	node := schemas.Account{Account: *a.Account}
	res := generated.DeleteAccountPayload{
		DeletedID:        node.ID(),
		ClientMutationID: input.ClientMutationID,
	}

	return &res, nil
}
