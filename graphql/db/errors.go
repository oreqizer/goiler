package db

import "github.com/vektah/gqlparser/gqlerror"

var (
	ErrDefault         = gqlerror.Errorf("database error")
	ErrFetchingResults = gqlerror.Errorf("database error while fetching results")
)
