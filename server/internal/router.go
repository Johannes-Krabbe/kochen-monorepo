package internal

import (
	"net/http"

	apiHandlers "github.com/Johannes-Krabbe/kochen-monorepo/server/internal/handlers/api"
	authApiHandlers "github.com/Johannes-Krabbe/kochen-monorepo/server/internal/handlers/api/auth"
	"github.com/Johannes-Krabbe/kochen-monorepo/server/internal/handlers/ui/componentHandlers"
	"github.com/Johannes-Krabbe/kochen-monorepo/server/internal/handlers/ui/pageHandlers"
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

	// === UI ===
	// Pages
	r.Get("/", pageHandlers.GetIndex)

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
