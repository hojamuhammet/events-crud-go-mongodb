package mocks

import (
	"events/internal/domain"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockMongoDBMovieRepository struct {
	mock.Mock
}

func (m *MockMongoDBMovieRepository) GetAllMovies(page, pageSize int) ([]*domain.GetMovieResponse, error) {
	args := m.Called(page, pageSize)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.GetMovieResponse), args.Error(1)
}

func (m *MockMongoDBMovieRepository) GetMovieByID(id primitive.ObjectID) (*domain.GetMovieResponse, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.GetMovieResponse), args.Error(1)
}

func (m *MockMongoDBMovieRepository) CreateMovie(movie *domain.CreateMovieRequest) (*domain.CreateMovieResponse, error) {
	args := m.Called(movie)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.CreateMovieResponse), args.Error(1)
}

func (m *MockMongoDBMovieRepository) UpdateMovie(id primitive.ObjectID, update *domain.UpdateMovieRequest) (*domain.UpdateMovieResponse, error) {
	args := m.Called(id, update)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.UpdateMovieResponse), args.Error(1)
}

func (m *MockMongoDBMovieRepository) DeleteMovie(id primitive.ObjectID) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockMongoDBMovieRepository) SearchMovies(query string, page, pageSize int) ([]*domain.GetMovieResponse, error) {
	args := m.Called(query, page, pageSize)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.GetMovieResponse), args.Error(1)
}

func (m *MockMongoDBMovieRepository) FilterMoviesByTags(tags []string, page, pageSize int) ([]*domain.GetMovieResponse, error) {
	args := m.Called(tags, page, pageSize)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.GetMovieResponse), args.Error(1)
}
