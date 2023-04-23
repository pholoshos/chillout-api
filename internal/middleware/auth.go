// internal/middleware/auth.go

package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/myuser/my-api/internal/server"
	"github.com/myuser/my-api/pkg/utils"
)

// Authenticate middleware validates the Authorization header
// and sets the user ID in the request context.
func Authenticate(s *server.Server, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			s.Logger.Error("Authorization header missing")
			utils.WriteJSONError(w, errors.New("Unauthorized"), http.StatusUnauthorized)
			return
		}

		authHeaderParts := strings.SplitN(authHeader, " ", 2)
		if len(authHeaderParts) != 2 {
