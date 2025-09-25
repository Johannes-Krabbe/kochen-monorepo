-- name: CreateRecipe :one
INSERT INTO recipes (owner_id, visibility, slug, content)
VALUES ($1, $2, $3, $4)
RETURNING id, owner_id, visibility, slug, content, created_at, updated_at;

-- name: GetRecipeBySlug :one
SELECT *
FROM recipes
WHERE slug = $1;
