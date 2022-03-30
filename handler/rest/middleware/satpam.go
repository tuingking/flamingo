package middleware

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/tuingking/flamingo/infra/contextkey"
	"github.com/tuingking/flamingo/internal/auth"
)

var (
	ErrNoTokenFound = errors.New("no token found")
	ErrUnauthorized = errors.New("token is unauthorized")
	ErrAlgoInvalid  = errors.New("algorithm mismatch")
	ErrExpired      = errors.New("token is expired")
)

// Satpam is middleware to verify access token taken from `Authorization: Bearer ...`
func Satpam(authsvc auth.Service) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			accessToken := r.Header.Get("Authorization")
			claims, err := authsvc.VerifyAccessToken(accessToken)
			if err != nil {
				logrus.Error(errors.Wrap(err, "[SATPAM] VerifyAccessToken"))
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), contextkey.Identity, claims)

			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}
