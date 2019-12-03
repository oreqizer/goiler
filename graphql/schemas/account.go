package schemas

import (
	"github.com/oreqizer/go-relay"
	"github.com/oreqizer/goiler/models"
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

func (s Accounts) ToSlice() []Account {
	ns := make([]Account, len(s))
	for i, v := range s {
		ns[i] = Account{Account: *v}
	}

	return ns
}
