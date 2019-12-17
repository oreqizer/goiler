package auth

import "github.com/oreqizer/goiler/models"

// Auth holds information about user's authentication
type Auth struct {
	AuthID  string
	Account *models.Account
}
