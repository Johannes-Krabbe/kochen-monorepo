package internal

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/Johannes-Krabbe/kochen-monorepo/server/internal/handlers"
	"github.com/Johannes-Krabbe/kochen-monorepo/server/internal/handlers/auth"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"github.com/Johannes-Krabbe/kochen-monorepo/server/docs"
)

func NewRouter(db *DB) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	authHandler := auth.NewAuthHandler(db.Queries)

	r.Get("/health", handlers.HealthCheck)
	
	r.Route("/v1", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Post("/signup", authHandler.Signup)
			r.Post("/login", authHandler.Login)
		})
	})

	r.Get("/swagger/doc.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(docs.SwaggerInfo.ReadDoc()))
	})
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	return r
}
