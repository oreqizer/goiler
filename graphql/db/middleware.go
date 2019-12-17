package db

import (
	"context"
	"database/sql"
	"net/http"
)

type key struct {
	name string
}

var keyDB = &key{"db"}

// Middleware creates a db middleware
func Middleware(db *sql.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), keyDB, db)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// GetDB retrieves the DB instance from the context
func GetDB(ctx context.Context) *sql.DB {
	DB := ctx.Value(keyDB).(*sql.DB)

	return DB
}
