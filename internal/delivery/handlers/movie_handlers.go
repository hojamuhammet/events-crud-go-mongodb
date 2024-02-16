package handlers

import (
	"encoding/json"
	"events/internal/domain"
	"events/internal/service"
	"events/pkg/lib/error"
	"events/pkg/lib/status"
	"events/pkg/lib/utils"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MovieHandler struct {
	MovieService *service.MovieService
	Router       *chi.Mux
}

type StatusMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (h *MovieHandler) GetAllMoviesHandler(w http.ResponseWriter, r *http.Request) {
	page := 1      // Default page if not provided
	pageSize := 10 // Default page size, adjust as needed

	pageStr := r.URL.Query().Get("page")
	if pageStr != "" {
		pageNum, err := strconv.Atoi(pageStr)
		if err != nil || pageNum < 1 {
			utils.RespondWithErrorJSON(w, status.BadRequest, error.InvalidRequestFormat)
			return
		}
		page = pageNum
	}

	movies, err := h.MovieService.GetAllMovies(page, pageSize)
	if err != nil {
		slog.Error("Error getting movies: ", utils.Err(err))
		utils.RespondWithErrorJSON(w, status.InternalServerError, error.InternalServerError)
		return
	}

	var prevPage interface{}
	if page > 1 {
		prevPage = page - 1
	} else {
		prevPage = nil // No previous page
	}

	var nextPage interface{}
	if len(movies) == pageSize {
		nextPage = page + 1
	} else {
		nextPage = nil
	}

	// Prepare pagination info
	pagination := map[string]interface{}{
		"current_page": page,
		"prev_page":    prevPage,
		"next_page":    nextPage,
	}

	// Respond with the retrieved movies and pagination info
	responseData := map[string]interface{}{
		"movies":     movies,
		"pagination": pagination,
	}

	utils.RespondWithJSON(w, status.OK, responseData)
}

func (h *MovieHandler) GetMovieHandler(w http.ResponseWriter, r *http.Request) {
	movieID := chi.URLParam(r, "id")

	objectID, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		slog.Error("Invalid movie ID: ", utils.Err(err))
		utils.RespondWithErrorJSON(w, status.BadRequest, error.InvalidMovieID)
		return
	}

	movie, err := h.MovieService.GetMovieByID(objectID)
	if err != nil {
		slog.Error("Error getting movie by ID: ", utils.Err(err))
		utils.RespondWithErrorJSON(w, status.InternalServerError, error.InternalServerError)
		return
	}

	if movie == nil {
		utils.RespondWithErrorJSON(w, status.NotFound, error.MovieNotFound)
		return
	}

	utils.RespondWithJSON(w, status.OK, movie)
}

func (h *MovieHandler) CreateMovieHandler(w http.ResponseWriter, r *http.Request) {
	var createMovieRequest domain.CreateMovieRequest
	err := json.NewDecoder(r.Body).Decode(&createMovieRequest)
	if err != nil {
		utils.RespondWithErrorJSON(w, status.BadRequest, error.InvalidRequestBody)
		return
	}

	movie, err := h.MovieService.CreateMovie(&createMovieRequest)
	if err != nil {
		slog.Error("Error creating movie: ", utils.Err(err))
		utils.RespondWithErrorJSON(w, status.InternalServerError, fmt.Sprintf("Error creating movie: %v", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(movie)
}

func (h *MovieHandler) UpdateMovieHandler(w http.ResponseWriter, r *http.Request) {
	movieID := chi.URLParam(r, "id")

	objectID, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		slog.Error("Invalid movie ID: ", utils.Err(err))
		utils.RespondWithErrorJSON(w, status.BadRequest, error.InvalidMovieID)
		return
	}

	var updateMovieRequest domain.UpdateMovieRequest
	err = json.NewDecoder(r.Body).Decode(&updateMovieRequest)
	if err != nil {
		utils.RespondWithErrorJSON(w, status.BadRequest, error.InvalidRequestBody)
		return
	}

	movie, err := h.MovieService.UpdateMovie(objectID, &updateMovieRequest)
	if err != nil {
		slog.Error("Error updating movie: ", utils.Err(err))
		utils.RespondWithErrorJSON(w, status.InternalServerError, fmt.Sprintf("Error updating movie: %v", err))
		return
	}

	if err != nil {
		if err.Error() == "movie not found" {
			utils.RespondWithErrorJSON(w, status.NotFound, error.MovieNotFound)
		} else {
			slog.Error("Error updating movie:", utils.Err(err))
			utils.RespondWithErrorJSON(w, status.InternalServerError, error.InternalServerError)
		}
		return
	}

	utils.RespondWithJSON(w, status.OK, movie)
}

func (h *MovieHandler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	movieID := chi.URLParam(r, "id")

	objectID, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		slog.Error("Invalid movie ID: ", utils.Err(err))
		utils.RespondWithErrorJSON(w, status.BadRequest, error.InvalidMovieID)
		return
	}

	err = h.MovieService.DeleteMovie(objectID)
	if err != nil {
		if err.Error() == "movie not found" {
			utils.RespondWithErrorJSON(w, status.NotFound, error.MovieNotFound)
		} else {
			slog.Error("Error deleting movie:", utils.Err(err))
			utils.RespondWithErrorJSON(w, status.InternalServerError, error.InternalServerError)
		}
		return
	}

	response := StatusMessage{
		Code:    200,
		Message: "Movie deleted successfully",
	}

	utils.RespondWithJSON(w, status.OK, response)
}

func (h *MovieHandler) SearchMoviesHandler(w http.ResponseWriter, r *http.Request) {
	page := 1      // Default page if not provided
	pageSize := 10 // Default page size, adjust as needed

	pageStr := r.URL.Query().Get("page")
	if pageStr != "" {
		pageNum, err := strconv.Atoi(pageStr)
		if err != nil || pageNum < 1 {
			utils.RespondWithErrorJSON(w, status.BadRequest, error.InvalidRequestFormat)
			return
		}
		page = pageNum
	}

	query := r.URL.Query().Get("query")

	movies, err := h.MovieService.SearchMovies(query, page, pageSize)
	if err != nil {
		slog.Error("Error searching movies: ", utils.Err(err))
		utils.RespondWithErrorJSON(w, status.InternalServerError, error.InternalServerError)
		return
	}

	var prevPage interface{}
	if page > 1 {
		prevPage = page - 1
	} else {
		prevPage = nil // No previous page
	}

	var nextPage interface{}
	if len(movies) == pageSize {
		nextPage = page + 1
	} else {
		nextPage = nil
	}

	pagination := map[string]interface{}{
		"current_page": page,
		"prev_page":    prevPage,
		"next_page":    nextPage,
	}

	responseData := map[string]interface{}{
		"movies":     movies,
		"pagination": pagination,
	}

	utils.RespondWithJSON(w, status.OK, responseData)
}

func (h *MovieHandler) FilterByTagsHandler(w http.ResponseWriter, r *http.Request) {
	page := 1      // Default page if not provided
	pageSize := 10 // Default page size, adjust as needed
	queryTags := r.URL.Query()["tags"]

	pageStr := r.URL.Query().Get("page")
	if pageStr != "" {
		pageNum, err := strconv.Atoi(pageStr)
		if err != nil || pageNum < 1 {
			utils.RespondWithErrorJSON(w, status.BadRequest, error.InvalidRequestFormat)
			return
		}
		page = pageNum
	}

	if len(queryTags) == 0 {
		utils.RespondWithErrorJSON(w, status.BadRequest, error.MissingTags)
		return
	}

	movies, err := h.MovieService.FilterByTags(queryTags, page, pageSize)
	if err != nil {
		slog.Error("Error filtering movies by tags: ", utils.Err(err))
		utils.RespondWithErrorJSON(w, status.InternalServerError, error.InternalServerError)
		return
	}

	var prevPage interface{}
	if page > 1 {
		prevPage = page - 1
	} else {
		prevPage = nil // No previous page
	}

	var nextPage interface{}
	if len(movies) == pageSize {
		nextPage = page + 1
	} else {
		nextPage = nil
	}

	pagination := map[string]interface{}{
		"current_page": page,
		"prev_page":    prevPage,
		"next_page":    nextPage,
	}

	responseData := map[string]interface{}{
		"movies":     movies,
		"pagination": pagination,
	}

	utils.RespondWithJSON(w, status.OK, responseData)
}
