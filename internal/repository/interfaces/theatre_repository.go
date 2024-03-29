package repository

import (
	"events/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//go:generate mockgen -source=theatre_repository.go -destination=mocks/theatre_repository_mock.go

type TheatreRepository interface {
	GetAllPerformances(page, pageSize int) ([]*domain.GetPerformanceResponse, error)
	GetTotalPerformancesCount() (int, error)
	GetPerformanceByID(id primitive.ObjectID) (*domain.GetPerformanceResponse, error)
	CreatePerformance(request *domain.CreatePerformanceRequest) (*domain.CreatePerformanceResponse, error)
	UpdatePerformance(id primitive.ObjectID, request *domain.UpdatePerformanceRequest) (*domain.UpdatePerformanceResponse, error)
	DeletePerformance(id primitive.ObjectID) error
	SearchPerformances(query string, page int, pageSize int) ([]*domain.GetPerformanceResponse, error)
	FilterPerformancesByTags(tags []string, page int, pageSize int) ([]*domain.GetPerformanceResponse, error)
}
