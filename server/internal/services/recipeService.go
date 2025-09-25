package services

import (
	"context"
	"encoding/json"
	"fmt"

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

func (s *RecipeService) Create(ownerID uuid.UUID, visibility sqlc.Visibility, slug string, content string) (sqlc.Recipe, error) {
	id := uuid.Must(uuid.NewV7())

	// Validate that content is valid RecipeContent JSON
	var recipeContent RecipeContent
	if err := json.Unmarshal([]byte(content), &recipeContent); err != nil {
		return sqlc.Recipe{}, fmt.Errorf("invalid recipe content JSON: %w", err)
	}

	params := sqlc.CreateRecipeParams{
		ID:         id,
		OwnerID:    ownerID,
		Visibility: visibility,
		Slug:       slug,
		Content:    pqtype.NullRawMessage{RawMessage: []byte(content), Valid: content != ""},
	}

	return s.queries.CreateRecipe(context.Background(), params)
}

func (s *RecipeService) GetBySlug(slug string) (sqlc.Recipe, RecipeContent, error) {
	recipe, err := s.queries.GetRecipeBySlug(context.Background(), slug)
	if err != nil {
		return sqlc.Recipe{}, RecipeContent{}, err
	}

	var content RecipeContent
	if recipe.Content.Valid {
		if err := json.Unmarshal(recipe.Content.RawMessage, &content); err != nil {
			return recipe, RecipeContent{}, fmt.Errorf("failed to parse recipe content: %w", err)
		}
	}

	return recipe, content, nil
}

// recipe.content type:
type Ingredient struct {
	ID     string   `json:"id"`
	Item   string   `json:"item"`
	Amount *float32 `json:"amount,omitempty"`
	Unit   *string  `json:"unit,omitempty"`
	Notes  *string  `json:"notes,omitempty"`
	Link   *string  `json:"link,omitempty"`
}

type Instruction struct {
	Step        int    `json:"step"`
	Description string `json:"description"`
}

type RecipeContent struct {
	Title           string        `json:"title"`
	Description     string        `json:"description"`
	Servings        int           `json:"servings"`
	PrepTimeMinutes int           `json:"prepTimeMinutes"`
	CookTime        int           `json:"cookTimeMinutes"`
	Ingredients     []Ingredient  `json:"ingredients"`
	Instructions    []Instruction `json:"instructions"`
}
