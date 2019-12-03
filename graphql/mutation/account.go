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

func (Mutation) AddAccount(
	ctx context.Context,
	input generated.AddAccountInput,
) (*generated.AddAccountPayload, error) {
	dbi, err := db.GetDB(ctx)
	if err != nil {
		return nil, err
	}

	a, err := auth.GetAuth(ctx)
	if err != nil {
		return nil, err
	}

	model := models.Account{
		AuthID:    a.AuthID,
		Name:      input.Name,
		Surname:   input.Surname,
		Email:     input.Email,
		DeletedAt: null.TimeFromPtr(nil),
	}

	if err := model.Upsert(ctx, dbi, true, []string{"auth_id"}, boil.Infer(), boil.Infer()); err != nil {
		raven.CaptureError(err, nil)
		return nil, db.ErrDefault
	}

	res := generated.AddAccountPayload{
		Account:          &schemas.Account{Account: model},
		ClientMutationID: input.ClientMutationID,
	}

	return &res, nil
}

func (Mutation) EditAccount(
	ctx context.Context,
	input generated.EditAccountInput,
) (*generated.EditAccountPayload, error) {
	dbi, err := db.GetDB(ctx)
	if err != nil {
		return nil, err
	}

	a, err := auth.GetAuthAccount(ctx)
	if err != nil {
		return nil, err
	}

	a.Account.Name = input.Name
	a.Account.Surname = input.Surname

	if _, err := a.Account.Update(ctx, dbi, boil.Infer()); err != nil {
		raven.CaptureError(err, nil)
		return nil, db.ErrDefault
	}

	res := generated.EditAccountPayload{
		Account:          &schemas.Account{Account: *a.Account},
		ClientMutationID: input.ClientMutationID,
	}

	return &res, nil
}

func (Mutation) DeleteAccount(
	ctx context.Context,
	input generated.DeleteAccountInput,
) (*generated.DeleteAccountPayload, error) {
	dbi, err := db.GetDB(ctx)
	if err != nil {
		return nil, err
	}

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
