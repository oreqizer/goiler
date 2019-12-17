package db

import "github.com/vektah/gqlparser/gqlerror"

var (
	// ErrDefault is a default database error
	ErrDefault = gqlerror.Errorf("database error")
	// ErrFetchingResults is an error for failed result fetching
	ErrFetchingResults = gqlerror.Errorf("database error while fetching results")
)
