package auth

import "errors"

var (
	ErrorNotAdmin     = errors.New("user is not an admin")
	ErrorNotOwner     = errors.New("user is not the trip owner")
	ErrorNotInContext = errors.New("user not in context")
)
