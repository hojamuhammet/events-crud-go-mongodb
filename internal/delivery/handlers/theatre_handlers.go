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

type TheatreHandler struct {
	TheatreService *service.TheatreService
	Router         *chi.Mux
}

func (h *TheatreHandler) GetAllPerformances(w http.ResponseWriter, r *http.Request) {
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

	performances, err := h.TheatreService.GetAllPerformances(page, pageSize)
	if err != nil {
		slog.Error("Error getting performances: ", utils.Err(err))
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
	if len(performances) == pageSize {
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
		"performances": performances,
		"pagination":   pagination,
	}

	utils.RespondWithJSON(w, status.OK, responseData)
}

func (h *TheatreHandler) GetPerformanceByID(w http.ResponseWriter, r *http.Request) {
	movieID := chi.URLParam(r, "id")

	objectID, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		slog.Error("Invalid performance ID: ", utils.Err(err))
		utils.RespondWithErrorJSON(w, status.BadRequest, error.InvalidPerformanceID)
		return
	}

	performance, err := h.TheatreService.TheatreService.GetPerformanceByID(objectID)
	if err != nil {
		slog.Error("Error getting performance by ID: ", utils.Err(err))
		utils.RespondWithErrorJSON(w, status.InternalServerError, error.InternalServerError)
		return
	}

	if performance == nil {
		utils.RespondWithErrorJSON(w, status.NotFound, error.PerformanceNotFound)
	}

	utils.RespondWithJSON(w, status.OK, performance)
}

func (h *TheatreHandler) CreatePerformanceHandler(w http.ResponseWriter, r *http.Request) {
	var createMovieRequest domain.CreatePerformanceRequest
	err := json.NewDecoder(r.Body).Decode(&createMovieRequest)
	if err != nil {
		utils.RespondWithErrorJSON(w, status.BadRequest, error.InvalidRequestBody)
		return
	}

	movie, err := h.TheatreService.CreatePerformance(&createMovieRequest)
	if err != nil {
		slog.Error("Error creating movie: ", utils.Err(err))
		utils.RespondWithErrorJSON(w, status.InternalServerError, fmt.Sprintf("Error creating movie: %v", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(movie)
}

func (h *TheatreHandler) UpdatePerformanceHandler(w http.ResponseWriter, r *http.Request) {
	movieID := chi.URLParam(r, "id")

	objectID, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		slog.Error("Invalid performance ID: ", utils.Err(err))
		utils.RespondWithErrorJSON(w, status.BadRequest, error.InvalidPerformanceID)
		return
	}

	var updatePerformanceRequest domain.UpdatePerformanceRequest
	err = json.NewDecoder(r.Body).Decode(&updatePerformanceRequest)
	if err != nil {
		utils.RespondWithErrorJSON(w, status.BadRequest, error.InvalidRequestBody)
		return
	}

	performance, err := h.TheatreService.UpdatePerformance(objectID, &updatePerformanceRequest)
	if err != nil {
		if err.Error() == "performance not found" {
			utils.RespondWithErrorJSON(w, status.NotFound, error.PerformanceNotFound)
		} else {
			slog.Error("Error deleting performance:", utils.Err(err))
			utils.RespondWithErrorJSON(w, status.InternalServerError, error.InternalServerError)
		}
		return
	}

	utils.RespondWithJSON(w, status.OK, performance)
}

func (h *TheatreHandler) DeletePerformance(w http.ResponseWriter, r *http.Request) {
	movieID := chi.URLParam(r, "id")

	objectID, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		slog.Error("Invalid performance ID: ", utils.Err(err))
		utils.RespondWithErrorJSON(w, status.BadRequest, error.InvalidPerformanceID)
		return
	}

	err = h.TheatreService.DeletePerformance(objectID)
	if err != nil {
		if err.Error() == "performance not found" {
			utils.RespondWithErrorJSON(w, status.NotFound, error.PerformanceNotFound)
		} else {
			slog.Error("Error deleting performance:", utils.Err(err))
			utils.RespondWithErrorJSON(w, status.InternalServerError, error.InternalServerError)
		}
		return
	}

	response := StatusMessage{
		Code:    200,
		Message: "Performance deleted successfully",
	}

	utils.RespondWithJSON(w, status.OK, response)
}

func (h *TheatreHandler) SearchPerfomancesHandler(w http.ResponseWriter, r *http.Request) {
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

	movies, err := h.TheatreService.SearchPerformances(query, page, pageSize)
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
