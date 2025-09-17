package authApiHandlers

import (
	sqlc "github.com/Johannes-Krabbe/kochen-monorepo/server/internal/database/sqlc"
)

type AuthHandler struct {
	queries *sqlc.Queries
}

func NewAuthHandler(queries *sqlc.Queries) *AuthHandler {
	return &AuthHandler{queries: queries}
}
