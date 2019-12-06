package schemas

import (
	"context"
	"github.com/getsentry/raven-go"
	"github.com/oreqizer/go-relay"
	"github.com/oreqizer/goiler/graphql/db"
	"github.com/oreqizer/goiler/graphql/slices"
	"github.com/oreqizer/goiler/models"
	. "github.com/volatiletech/sqlboiler/queries/qm"
)

const TypeAccount = "Account"

type Account struct {
	models.Account
}

func (Account) IsNode() {}

func (a *Account) ID() string {
	return relay.ToGlobalID(TypeAccount, a.Account.ID)
}

type Accounts []*models.Account

func (s Accounts) ToSlice() []*Account {
	ns := make([]*Account, len(s))
	for i, v := range s {
		ns[i] = &Account{Account: *v}
	}

	return ns
}

func MakeAccountLoader(ctx context.Context) *AccountLoader {
	return NewAccountLoader(AccountLoaderConfig{
		Fetch: func(keys []string) (accounts []*Account, errors []error) {
			dbi := db.GetDB(ctx)

			res, err := models.Accounts(
				db.QueryNotDeleted,
				WhereIn("id in ?", slices.StringsToInterfaces(keys)...),
			).All(ctx, dbi)
			if err != nil {
				raven.CaptureError(err, nil)
				return nil, []error{db.ErrFetchingResults}
			}

			return Accounts(res).ToSlice(), nil
		},
	})
}

func GetAccountLoader(ctx context.Context) *AccountLoader {
	acc := ctx.Value(keyAccount).(*AccountLoader)

	return acc
}
