package auth

import "github.com/vektah/gqlparser/gqlerror"

var (
	ErrUnauthorized = gqlerror.Errorf("user is not authorized")
	ErrNotAdmin     = gqlerror.Errorf("user is not an admin")
	ErrNoToken      = gqlerror.Errorf("no user token present")
)
