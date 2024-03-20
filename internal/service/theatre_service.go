package service

import (
	"events/internal/domain"
	repository "events/internal/repository/interfaces"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TheatreService struct {
	TheatreService repository.TheatreRepository
}

func NewTheatreService(theatreRepository repository.TheatreRepository) *TheatreService {
	return &TheatreService{TheatreService: theatreRepository}
}

func (s *TheatreService) GetAllPerformances(page, pageSize int) ([]*domain.GetPerformanceResponse, error) {
	return s.TheatreService.GetAllPerformances(page, pageSize)
}

func (s *TheatreService) GetTotalPerformancesCount() (int, error) {
	return s.TheatreService.GetTotalPerformancesCount()
}

func (s *TheatreService) GetPerformanceByID(id primitive.ObjectID) (*domain.GetPerformanceResponse, error) {
	return s.TheatreService.GetPerformanceByID(id)
}

func (s *TheatreService) CreatePerformance(request *domain.CreatePerformanceRequest) (*domain.CreatePerformanceResponse, error) {
	return s.TheatreService.CreatePerformance(request)
}

func (s *TheatreService) UpdatePerformance(id primitive.ObjectID, request *domain.UpdatePerformanceRequest) (*domain.UpdatePerformanceResponse, error) {
	return s.TheatreService.UpdatePerformance(id, request)
}

func (s *TheatreService) DeletePerformance(id primitive.ObjectID) error {
	return s.TheatreService.DeletePerformance(id)
}

func (s *TheatreService) SearchPerformances(query string, page int, pageSize int) ([]*domain.GetPerformanceResponse, error) {
	return s.TheatreService.SearchPerformances(query, page, pageSize)
}

func (s *TheatreService) FilterPerformancesByTags(tags []string, page int, pageSize int) ([]*domain.GetPerformanceResponse, error) {
	return s.TheatreService.FilterPerformancesByTags(tags, page, pageSize)
}
