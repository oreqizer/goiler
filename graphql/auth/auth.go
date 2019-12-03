package auth

import "github.com/oreqizer/goiler/models"

type Auth struct {
	AuthID  string
	Account *models.Account
}
