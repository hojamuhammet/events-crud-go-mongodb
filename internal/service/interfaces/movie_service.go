package service

import (
	"events/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//go:generate mockgen -source=movie_service.go -destination=mocks/movie_service_mock.go

type MovieService interface {
	GetAllMovies(page, pageSize int) ([]*domain.GetMovieResponse, error)
	GetTotalMoviesCount() (int, error)
	GetMovieByID(id primitive.ObjectID) (*domain.GetMovieResponse, error)
	CreateMovie(request *domain.CreateMovieRequest) (*domain.CreateMovieResponse, error)
	UpdateMovie(id primitive.ObjectID, request *domain.UpdateMovieRequest) (*domain.UpdateMovieResponse, error)
	DeleteMovie(id primitive.ObjectID) error
	SearchMovies(query string, page int, pageSize int) ([]*domain.GetMovieResponse, error)
	FilterMoviesByTags(tags []string, page int, pageSize int) ([]*domain.GetMovieResponse, error)
}
