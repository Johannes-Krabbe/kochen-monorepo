package auth

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Johannes-Krabbe/kochen-monorepo/server/internal/utils/errors"
	"github.com/Johannes-Krabbe/kochen-monorepo/server/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// Login godoc
// @Summary User login
// @Description Authenticate user with username/email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param login body LoginRequest true "Login request"
// @Success 200 {object} AuthResponse
// @Failure 400 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /v1/auth/login [post]
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Login == "" || req.Password == "" {
		errors.WriteJSONError(w, 400, errors.Validation)
		return
	}

	user, err := h.queries.GetUserByEmailOrUsername(context.Background(), req.Login)
	if err != nil {
		errors.WriteJSONError(w, 400, errors.InvalidCredentials)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		errors.WriteJSONError(w, 400, errors.InvalidCredentials)
		return
	}

	token, err := utils.GenerateJWT(user.ID.String())
	if err != nil {
		errors.WriteJSONError(w, 400, errors.InternalServerError)
		return
	}

	response := AuthResponse{
		Token: token,
		User: UserProfile{
			ID:       user.ID.String(),
			Username: user.Username,
			Email:    user.Email,
			Name:     user.Name,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
