package auth

import "github.com/vektah/gqlparser/gqlerror"

var (
	ErrNotAdmin = gqlerror.Errorf("user is not an admin")
	ErrNoToken  = gqlerror.Errorf("no user token present")
	ErrNotOwner = gqlerror.Errorf("user is not the trip owner")
)
