package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	//create a router mux
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(app.enableCORS)

	mux.Post("/authenticate", app.authenticate)
	mux.Get("/refresh", app.refreshToken)
	mux.Get("/logout", app.logout)
	mux.Get("/", app.Home)

	mux.Get("/movies", app.AllMovies)
	mux.Get("/movies/{id}", app.GetMovie)

	mux.Get("/genres", app.AllGenres)

	mux.Route("/admin", func(r chi.Router) {
		r.Use(app.authRequired)

		r.Get("/movies", app.MovieCatalog)
		r.Get("/movies/{id}", app.MovieForEdit)
		r.Put("/movies/0", app.InsertMovie)
	})
	return mux
}
