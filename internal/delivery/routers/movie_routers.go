package routes

import (
	"events/internal/delivery/handlers"
	"events/internal/service"

	"github.com/go-chi/chi/v5"
)

func SetupMovieRouter(movieRouter *chi.Mux, movieService *service.MovieService) {
	movieHandler := handlers.MovieHandler{
		Router:       movieRouter,
		MovieService: movieService,
	}

	movieRouter.Get("/", movieHandler.GetAllMoviesHandler)
	movieRouter.Get("/{id}", movieHandler.GetMovieByIDHandler)
	movieRouter.Post("/", movieHandler.CreateMovieHandler)
	movieRouter.Put("/{id}", movieHandler.UpdateMovieHandler)
	movieRouter.Delete("/{id}", movieHandler.DeleteMovieHandler)
	movieRouter.Get("/search", movieHandler.SearchMoviesHandler)
	movieRouter.Get("/filter/tags", movieHandler.FilterMoviesByTagsHandler)
}
