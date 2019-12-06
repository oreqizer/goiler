package auth

import "errors"

var (
	ErrNotAdmin = errors.New("user is not an admin")
	ErrNoToken  = errors.New("no user token present")
	ErrNotOwner = errors.New("user is not the trip owner")
)
