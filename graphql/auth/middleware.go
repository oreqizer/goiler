package auth

import (
	"context"
	"database/sql"
	firebase "firebase.google.com/go"
	"github.com/getsentry/raven-go"
	"github.com/oreqizer/goiler/graphql/db"
	"github.com/oreqizer/goiler/models"
	. "github.com/volatiletech/sqlboiler/queries/qm"
	"net/http"
	"strings"
)

type key struct {
	name string
}

var keyAuth = &key{"auth"}

func Middleware(dbi *sql.DB, fb *firebase.App) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			client, err := fb.Auth(ctx)
			if err != nil {
				raven.CaptureError(err, nil)
				http.Error(w, "auth provider error", http.StatusInternalServerError)
				return
			}

			header := strings.ReplaceAll(r.Header.Get("Authorization"), "Bearer ", "")
			var authID string
			if header == "" {
				if strings.Contains(r.Host, "localhost") {
					// Development account
					authID = "TEST"
				} else {
					authID = ""
				}
			} else {
				token, err := client.VerifyIDToken(ctx, header)
				if err != nil {
					raven.CaptureError(err, nil)
					http.Error(w, "auth error", http.StatusUnauthorized)
					return
				}
				authID = token.Subject
			}

			// Allow unauthenticated users in
			if authID == "" {
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			res, err := models.Accounts(db.QueryNotDeleted, Where("auth_id = ?", authID)).One(ctx, dbi)
			if err == sql.ErrNoRows {
				// No account yet, but has auth
				ctx = context.WithValue(r.Context(), keyAuth, &Auth{
					AuthID: authID,
				})

				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
			if err != nil {
				raven.CaptureError(err, nil)
				http.Error(w, "database error", http.StatusInternalServerError)
				return
			}

			ctx = context.WithValue(r.Context(), keyAuth, &Auth{
				AuthID:  authID,
				Account: res,
			})

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetAuth(ctx context.Context) *Auth {
	auth, ok := ctx.Value(keyAuth).(*Auth)
	if !ok {
		return nil
	}

	return auth
}

func GetAuthAccount(ctx context.Context) (*Auth, error) {
	auth := GetAuth(ctx)

	if auth == nil || auth.Account == nil {
		return nil, ErrUnauthorized
	}

	return auth, nil
}

func GetAuthAdmin(ctx context.Context) (*Auth, error) {
	auth, err := GetAuthAccount(ctx)
	if err != nil || !auth.Account.IsAdmin {
		return nil, ErrNotAdmin
	}

	return auth, nil
}
