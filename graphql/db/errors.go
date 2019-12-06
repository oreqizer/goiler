package db

import "errors"

var (
	ErrDefault         = errors.New("database error")
	ErrFetchingResults = errors.New("database error while fetching results")
)
