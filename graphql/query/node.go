package query

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/getsentry/raven-go"
	"github.com/oreqizer/go-relay"
	"github.com/oreqizer/goiler/generated"
	"github.com/oreqizer/goiler/graphql/auth"
	"github.com/oreqizer/goiler/graphql/db"
	"github.com/oreqizer/goiler/graphql/schemas"
	"github.com/oreqizer/goiler/models"
)

func (Query) Node(ctx context.Context, id string) (generated.Node, error) {
	dbi, err := db.GetDB(ctx)
	if err != nil {
		return nil, err
	}

	local := relay.FromGlobalID(id)
	if local == nil {
		raven.CaptureError(fmt.Errorf("invalid global ID '%s'", id), nil)
		return nil, db.ErrFetchingResults
	}

	switch local.Type {
	case schemas.TypeAccount:
		return GetAccount(ctx, dbi, local.ID)
	}

	raven.CaptureError(fmt.Errorf("unknown type '%s'", local.Type), nil)
	return nil, db.ErrFetchingResults
}

func GetAccount(ctx context.Context, dbi *sql.DB, id string) (*schemas.Account, error) {
	a, err := auth.GetAuthAccount(ctx)
	if err != nil {
		return nil, err
	}

	res, err := models.FindAccount(ctx, dbi, id)
	if err != nil {
		raven.CaptureError(err, nil)
		return nil, db.ErrFetchingResults
	}

	if res.DeletedAt.Valid {
		return nil, nil
	}

	if !a.Account.IsAdmin && a.Account.ID != id {
		return nil, nil
	}

	return &schemas.Account{Account: *res}, nil
}
