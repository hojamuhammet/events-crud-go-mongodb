// Code generated by MockGen. DO NOT EDIT.
// Source: movie_repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	domain "events/internal/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// MockMovieRepository is a mock of MovieRepository interface.
type MockMovieRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMovieRepositoryMockRecorder
}

// MockMovieRepositoryMockRecorder is the mock recorder for MockMovieRepository.
type MockMovieRepositoryMockRecorder struct {
	mock *MockMovieRepository
}

// NewMockMovieRepository creates a new mock instance.
func NewMockMovieRepository(ctrl *gomock.Controller) *MockMovieRepository {
	mock := &MockMovieRepository{ctrl: ctrl}
	mock.recorder = &MockMovieRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMovieRepository) EXPECT() *MockMovieRepositoryMockRecorder {
	return m.recorder
}

// CreateMovie mocks base method.
func (m *MockMovieRepository) CreateMovie(request *domain.CreateMovieRequest) (*domain.CreateMovieResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMovie", request)
	ret0, _ := ret[0].(*domain.CreateMovieResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMovie indicates an expected call of CreateMovie.
func (mr *MockMovieRepositoryMockRecorder) CreateMovie(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMovie", reflect.TypeOf((*MockMovieRepository)(nil).CreateMovie), request)
}

// DeleteMovie mocks base method.
func (m *MockMovieRepository) DeleteMovie(id primitive.ObjectID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMovie", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMovie indicates an expected call of DeleteMovie.
func (mr *MockMovieRepositoryMockRecorder) DeleteMovie(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMovie", reflect.TypeOf((*MockMovieRepository)(nil).DeleteMovie), id)
}

// FilterMoviesByTags mocks base method.
func (m *MockMovieRepository) FilterMoviesByTags(tags []string, page, pageSize int) ([]*domain.GetMovieResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilterMoviesByTags", tags, page, pageSize)
	ret0, _ := ret[0].([]*domain.GetMovieResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterMoviesByTags indicates an expected call of FilterMoviesByTags.
func (mr *MockMovieRepositoryMockRecorder) FilterMoviesByTags(tags, page, pageSize interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterMoviesByTags", reflect.TypeOf((*MockMovieRepository)(nil).FilterMoviesByTags), tags, page, pageSize)
}

// GetAllMovies mocks base method.
func (m *MockMovieRepository) GetAllMovies(page, pageSize int) ([]*domain.GetMovieResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllMovies", page, pageSize)
	ret0, _ := ret[0].([]*domain.GetMovieResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllMovies indicates an expected call of GetAllMovies.
func (mr *MockMovieRepositoryMockRecorder) GetAllMovies(page, pageSize interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllMovies", reflect.TypeOf((*MockMovieRepository)(nil).GetAllMovies), page, pageSize)
}

// GetMovieByID mocks base method.
func (m *MockMovieRepository) GetMovieByID(id primitive.ObjectID) (*domain.GetMovieResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMovieByID", id)
	ret0, _ := ret[0].(*domain.GetMovieResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMovieByID indicates an expected call of GetMovieByID.
func (mr *MockMovieRepositoryMockRecorder) GetMovieByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMovieByID", reflect.TypeOf((*MockMovieRepository)(nil).GetMovieByID), id)
}

// GetTotalMoviesCount mocks base method.
func (m *MockMovieRepository) GetTotalMoviesCount() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotalMoviesCount")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTotalMoviesCount indicates an expected call of GetTotalMoviesCount.
func (mr *MockMovieRepositoryMockRecorder) GetTotalMoviesCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotalMoviesCount", reflect.TypeOf((*MockMovieRepository)(nil).GetTotalMoviesCount))
}

// SearchMovies mocks base method.
func (m *MockMovieRepository) SearchMovies(query string, page, pageSize int) ([]*domain.GetMovieResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchMovies", query, page, pageSize)
	ret0, _ := ret[0].([]*domain.GetMovieResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchMovies indicates an expected call of SearchMovies.
func (mr *MockMovieRepositoryMockRecorder) SearchMovies(query, page, pageSize interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchMovies", reflect.TypeOf((*MockMovieRepository)(nil).SearchMovies), query, page, pageSize)
}

// UpdateMovie mocks base method.
func (m *MockMovieRepository) UpdateMovie(id primitive.ObjectID, request *domain.UpdateMovieRequest) (*domain.UpdateMovieResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMovie", id, request)
	ret0, _ := ret[0].(*domain.UpdateMovieResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateMovie indicates an expected call of UpdateMovie.
func (mr *MockMovieRepositoryMockRecorder) UpdateMovie(id, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMovie", reflect.TypeOf((*MockMovieRepository)(nil).UpdateMovie), id, request)
}
