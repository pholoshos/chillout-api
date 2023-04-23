// internal/handlers/auth.go

package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/myuser/my-api/internal/server"
	"github.com/myuser/my-api/pkg/utils"
)

// AuthHandler handles authentication requests.
type AuthHandler struct {
	s *server.Server
}

// NewAuthHandler returns a new instance of AuthHandler.
func NewAuthHandler(s *server.Server) *AuthHandler {
	return &AuthHandler{s: s}
}

// RegisterRoutes registers the authentication routes.
func (h *AuthHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/auth", h.authenticate).Methods("POST")
}

// authenticate handles the /auth POST route.
func (h *AuthHandler) authenticate(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := utils.DecodeJSONBody(r, &credentials); err != nil {
