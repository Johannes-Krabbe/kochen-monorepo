package auth

import (
	"context"
	"encoding/json"
	"net/http"
	"regexp"
	"strings"

	sqlc "github.com/Johannes-Krabbe/kochen-monorepo/server/internal/database/sqlc"
	"github.com/Johannes-Krabbe/kochen-monorepo/server/internal/utils"
	"github.com/Johannes-Krabbe/kochen-monorepo/server/internal/utils/errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	queries *sqlc.Queries
}

func NewAuthHandler(queries *sqlc.Queries) *AuthHandler {
	return &AuthHandler{queries: queries}
}

type SignupRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string      `json:"token"`
	User  UserProfile `json:"user"`
}

type UserProfile struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}

var usernameRegex = regexp.MustCompile(`^[a-z0-9_-]+$`)

// Signup godoc
// @Summary User signup
// @Description Create a new user account with username, email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param signup body SignupRequest true "Signup request"
// @Success 201 {object} AuthResponse
// @Failure 400 {object} errors.ErrorResponse
// @Failure 500 {object} errors.ErrorResponse
// @Router /v1/auth/signup [post]
func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	var req SignupRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errors.WriteJSONError(w, 400, errors.Validation, "Invalid JSON")
		return
	}

	if req.Username == "" || req.Email == "" || req.Password == "" {
		errors.WriteJSONError(w, 400, errors.Validation, "Username, email, and password are required")
		return
	}

	req.Username = strings.ToLower(req.Username)
	if !usernameRegex.MatchString(req.Username) {
		errors.WriteJSONError(w, 400, errors.Validation, "Username must contain only lowercase letters, numbers, underscores, and dashes")
		return
	}

	if len(req.Password) < 8 {
		errors.WriteJSONError(w, 400, errors.Validation, "Password must be at least 8 characters long")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		errors.WriteJSONError(w, 500, errors.InternalServerError, "Failed to process password")
		return
	}

	userID := uuid.Must(uuid.NewV7())

	params := sqlc.CreateUserParams{
		ID:           userID,
		Email:        req.Email,
		Name:         req.Username,
		Username:     req.Username,
		PasswordHash: string(hashedPassword),
	}

	user, err := h.queries.CreateUser(context.Background(), params)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			// TODO handle username taken and email taken
			


			http.Error(w, "Username or email already exists", http.StatusConflict)
			return
		}
		errors.WriteJSONError(w, 500, errors.InternalServerError, "Failed to create user")
		return
	}

	token, err := utils.GenerateJWT(user.ID.String())
	if err != nil {
		errors.WriteJSONError(w, 500, errors.InternalServerError, "Failed to generate token")
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
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
