package main

import (
	"events/internal/config"
	routes "events/internal/delivery/routers"
	repository "events/internal/repository/mongodb"
	"events/internal/service"
	"events/pkg/database"
	"events/pkg/lib/utils"
	"events/pkg/logger"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
)

func main() {
	cfg := config.LoadConfig()

	log := logger.SetupLogger(cfg.Env)

	slog.Info("Starting the server...", slog.String("env", cfg.Env))
	slog.Debug("Debug messages are enabled") // If env is set to prod, debug messages are going to be disabled

	if err := database.InitDB(cfg); err != nil {
		log.Error("Error setting up MongoDB: %v", err)
	}
	defer database.Close()

	mainRouter := chi.NewRouter()

	movieRouter := chi.NewRouter()

	mainRouter.Route("/api/movie", func(r chi.Router) {
		r.Mount("/", movieRouter)
	})

	movieCollection := database.GetDB().Collection("movies")
	movieRepository := repository.NewMongoDBMovieRepository(movieCollection)
	movieService := service.NewMovieService(movieRepository)
	routes.SetupMovieRouter(movieRouter, movieService)

	theatreRouter := chi.NewRouter()

	mainRouter.Route("/api/performance", func(r chi.Router) {
		r.Mount("/", theatreRouter)
	})

	theatreCollection := database.GetDB().Collection("theatre")
	theatreRepository := repository.NewMongoDBTheatreRepository(theatreCollection)
	theatreService := service.NewTheatreService(theatreRepository) // SOLVE THIS PROBLEM WITH POINTER
	routes.SetupTheatreRouter(theatreRouter, theatreService)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-stop
		log.Info("Shutting down the server gracefully...")

		database.Close()
		os.Exit(0)
	}()

	if err := http.ListenAndServe(cfg.Server.Address, mainRouter); err != nil {
		slog.Error("Server failed to start:", utils.Err(err))
	}
}
