package repository

import (
	"events/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TheatreRepository interface {
	GetAllPerformances(page, pageSize int) ([]*domain.GetPerformanceResponse, error)
	GetPerformance(id primitive.ObjectID) (*domain.GetPerformanceResponse, error)
	CreatePerformance(request *domain.CreatePerformanceRequest) (*domain.CreatePerformanceResponse, error)
	UpdatePerformance(id primitive.ObjectID, request *domain.UpdatePerformanceRequest) (*domain.UpdatePerformanceResponse, error)
	DeletePerformance(id primitive.ObjectID) error
	SearchPerformances(query string, page int, pageSize int) ([]*domain.GetPerformanceResponse, error)
}
