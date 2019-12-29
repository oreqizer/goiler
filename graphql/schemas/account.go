//go:generate go run github.com/vektah/dataloaden AccountLoader string *github.com/oreqizer/goiler/graphql/schemas.Account
//go:generate go run github.com/oreqizer/go-relaygen Account Account.ID

package schemas

import (
	"context"
	"github.com/getsentry/raven-go"
	"github.com/lib/pq"
	"github.com/oreqizer/go-relaygen/relay"
	"github.com/oreqizer/goiler/graphql/db"
	"github.com/oreqizer/goiler/graphql/slices"
	"github.com/oreqizer/goiler/models"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// TypeAccount is the name of the Account type
const TypeAccount = "Account"

// Account holds information about an account
type Account struct {
	models.Account
}

// IsNode satisfies the Node interface
func (Account) IsNode() {}

// ID satisfies the ID interface
func (a *Account) ID() string {
	return relay.ToGlobalID(TypeAccount, a.Account.ID)
}

// Accounts is a slice of model accounts
type Accounts []*models.Account

// ToSlice converts model accounts to gql accounts
func (s Accounts) ToSlice() []*Account {
	ns := make([]*Account, len(s))
	for i, v := range s {
		ns[i] = &Account{Account: *v}
	}

	return ns
}

const queryAccountLoader = `(
SELECT * FROM unnest(? :: UUID[]) WITH ORDINALITY
) AS ids (id, ordering) ON ids.id = account.id`

// MakeAccountLoader creates an account loader
func MakeAccountLoader(ctx context.Context) *AccountLoader {
	return NewAccountLoader(AccountLoaderConfig{
		Fetch: func(keys []string) (accounts []*Account, errors []error) {
			dbi := db.GetDB(ctx)

			res, err := models.Accounts(
				db.QueryNotDeleted,
				qm.InnerJoin(queryAccountLoader, pq.StringArray(keys)),
				qm.OrderBy("ids.ordering"),
			).All(ctx, dbi)
			if err != nil {
				raven.CaptureError(err, nil)
				return nil, []error{db.ErrLoader}
			}

			return Accounts(res).ToSlice(), nil
		},
		Wait: LoaderWait,
	})
}

// GetAccountLoader retrieves the account loader from the context
func GetAccountLoader(ctx context.Context) *AccountLoader {
	acc := ctx.Value(keyAccount).(*AccountLoader)

	return acc
}
