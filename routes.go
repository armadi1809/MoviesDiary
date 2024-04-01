package main

import (
	"net/http"

	"github.com/armadi1809/moviesdiary/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nedpals/supabase-go"
)

func routes(sbClient *supabase.Client) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Group(func(authenticated chi.Router) {
		authenticated.Use(handlers.WithAuth(sbClient))
		authenticated.Get("/", handlers.HomeHandler())
	})

	r.Get("/login/google", handlers.GoogleLoginHandler(sbClient))
	r.Get("/auth/callback", handlers.HandleAuthCallback)
	r.Handle("/public/*", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	return r
}
