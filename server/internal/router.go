package internal

import (
	"net/http"

	apiHandlers "github.com/Johannes-Krabbe/kochen-monorepo/server/internal/handlers/api"
	authApiHandlers "github.com/Johannes-Krabbe/kochen-monorepo/server/internal/handlers/api/auth"
	"github.com/Johannes-Krabbe/kochen-monorepo/server/internal/handlers/ui/componentHandlers"
	"github.com/Johannes-Krabbe/kochen-monorepo/server/internal/handlers/ui/pageHandlers"
	"github.com/Johannes-Krabbe/kochen-monorepo/server/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(db *DB) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	// Static files
	fileServer := http.FileServer(http.Dir("./internal/ui/static/"))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	// Favicon
	r.Get("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./internal/ui/static/favicon.ico")
	})

	// === UI ===
	// Services
	recipeService := services.NewRecipeService(db.Queries)

	// Pages
	r.Get("/", pageHandlers.GetIndex)
	r.Get("/recipe/{slug}", pageHandlers.GetRecipe(recipeService))

	// Components
	r.Route("/component", func(r chi.Router) {
		r.Post("/increase", componentHandlers.PostIncrease)
	})

	// === API ===
	authHandler := authApiHandlers.NewAuthHandler(db.Queries)

	r.Route("/api", func(r chi.Router) {
		r.Get("/health", apiHandlers.HealthCheck)

		r.Route("/v1", func(r chi.Router) {
			r.Route("/auth", func(r chi.Router) {
				r.Post("/signup", authHandler.Signup)
				r.Post("/login", authHandler.Login)
			})

		})
	})

	return r
}
