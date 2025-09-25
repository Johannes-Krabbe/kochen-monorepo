package services

import (
	"context"

	sqlc "github.com/Johannes-Krabbe/kochen-monorepo/server/internal/database/sqlc"
	"github.com/google/uuid"
	"github.com/sqlc-dev/pqtype"
)

type RecipeService struct {
	queries *sqlc.Queries
}

func NewRecipeService(queries *sqlc.Queries) *RecipeService {
	return &RecipeService{queries: queries}
}

func (s *RecipeService) Create(ownerID uuid.UUID, visibility sqlc.Visibility, slug string, content []byte) (sqlc.Recipe, error) {
	id := uuid.Must(uuid.NewV7())

	params := sqlc.CreateRecipeParams{
		ID:         id,
		OwnerID:    ownerID,
		Visibility: visibility,
		Slug:       slug,
		Content:    pqtype.NullRawMessage{RawMessage: content, Valid: content != nil},
	}

	return s.queries.CreateRecipe(context.Background(), params)
}
