package auth

import "github.com/vektah/gqlparser/gqlerror"

var (
	// ErrUnauthorized means user is not authorized
	ErrUnauthorized = gqlerror.Errorf("user is not authorized")
	// ErrNotAdmin means user is not an admin
	ErrNotAdmin = gqlerror.Errorf("user is not an admin")
	// ErrNoToken means user has no token present
	ErrNoToken = gqlerror.Errorf("no user token present")
)
