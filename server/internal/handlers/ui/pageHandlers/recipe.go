package pageHandlers

import (
	"net/http"

	"github.com/Johannes-Krabbe/kochen-monorepo/server/internal/services"
	"github.com/Johannes-Krabbe/kochen-monorepo/server/internal/ui/pages"
	"github.com/Johannes-Krabbe/kochen-monorepo/server/internal/utils"
	"github.com/go-chi/chi/v5"
)

func GetRecipe(recipeService *services.RecipeService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		slug := chi.URLParam(r, "slug")
		if slug == "" {
			http.Error(w, "Recipe slug is required", http.StatusBadRequest)
			return
		}

		recipe, content, err := recipeService.GetBySlug(slug)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		component := pages.RecipePage(recipe, content)
		utils.Render(w, r, component)
	}
}

