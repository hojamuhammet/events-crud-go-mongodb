package service

import (
	"events/internal/domain"
	"events/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MovieService struct {
	MovieRepository repository.MovieRepository
}

func NewMovieService(movieRepository repository.MovieRepository) *MovieService {
	return &MovieService{MovieRepository: movieRepository}
}

func (s *MovieService) GetAllMovies(page, pageSize int) ([]*domain.GetMovieResponse, error) {
	return s.MovieRepository.GetAllMovies(page, pageSize)
}

func (s *MovieService) GetMovieByID(id primitive.ObjectID) (*domain.GetMovieResponse, error) {
	return s.MovieRepository.GetMovieByID(id)
}

func (s *MovieService) CreateMovie(request *domain.CreateMovieRequest) (*domain.CreateMovieResponse, error) {
	return s.MovieRepository.CreateMovie(request)
}

func (s *MovieService) UpdateMovie(id primitive.ObjectID, update *domain.UpdateMovieRequest) (*domain.UpdateMovieResponse, error) {
	return s.MovieRepository.UpdateMovie(id, update)
}

func (s *MovieService) DeleteMovie(id primitive.ObjectID) error {
	return s.MovieRepository.DeleteMovie(id)
}

func (s *MovieService) SearchMovies(query string, page int, pageSize int) ([]*domain.GetMovieResponse, error) {
	return s.MovieRepository.SearchMovies(query, page, pageSize)
}

func (s *MovieService) FilterMoviesByTags(tags []string, page int, pageSize int) ([]*domain.GetMovieResponse, error) {
	return s.MovieRepository.FilterMoviesByTags(tags, page, pageSize)
}
