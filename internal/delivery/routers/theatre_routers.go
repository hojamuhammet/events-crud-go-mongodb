package routes

import (
	"events/internal/delivery/handlers"
	"events/internal/service"

	"github.com/go-chi/chi/v5"
)

func SetupTheatreRouter(theatreRouter *chi.Mux, theatreService *service.TheatreService) {
	theatreHandler := handlers.TheatreHandler{
		Router:         theatreRouter,
		TheatreService: theatreService,
	}

	theatreRouter.Get("/", theatreHandler.GetAllPerformances)
	theatreRouter.Get("/{id}", theatreHandler.GetPerformanceByID)
	theatreRouter.Post("/", theatreHandler.CreatePerformanceHandler)
	theatreRouter.Put("/{id}", theatreHandler.UpdatePerformanceHandler)
	theatreRouter.Delete("/{id}", theatreHandler.DeletePerformance)
	theatreRouter.Get("/search", theatreHandler.SearchPerfomancesHandler)
}
