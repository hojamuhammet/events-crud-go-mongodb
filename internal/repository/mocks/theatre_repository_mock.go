// Code generated by MockGen. DO NOT EDIT.
// Source: theatre_repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	domain "events/internal/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// MockTheatreRepository is a mock of TheatreRepository interface.
type MockTheatreRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTheatreRepositoryMockRecorder
}

// MockTheatreRepositoryMockRecorder is the mock recorder for MockTheatreRepository.
type MockTheatreRepositoryMockRecorder struct {
	mock *MockTheatreRepository
}

// NewMockTheatreRepository creates a new mock instance.
func NewMockTheatreRepository(ctrl *gomock.Controller) *MockTheatreRepository {
	mock := &MockTheatreRepository{ctrl: ctrl}
	mock.recorder = &MockTheatreRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTheatreRepository) EXPECT() *MockTheatreRepositoryMockRecorder {
	return m.recorder
}

// CreatePerformance mocks base method.
func (m *MockTheatreRepository) CreatePerformance(request *domain.CreatePerformanceRequest) (*domain.CreatePerformanceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePerformance", request)
	ret0, _ := ret[0].(*domain.CreatePerformanceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePerformance indicates an expected call of CreatePerformance.
func (mr *MockTheatreRepositoryMockRecorder) CreatePerformance(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePerformance", reflect.TypeOf((*MockTheatreRepository)(nil).CreatePerformance), request)
}

// DeletePerformance mocks base method.
func (m *MockTheatreRepository) DeletePerformance(id primitive.ObjectID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePerformance", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePerformance indicates an expected call of DeletePerformance.
func (mr *MockTheatreRepositoryMockRecorder) DeletePerformance(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePerformance", reflect.TypeOf((*MockTheatreRepository)(nil).DeletePerformance), id)
}

// FilterPerformancesByTags mocks base method.
func (m *MockTheatreRepository) FilterPerformancesByTags(tags []string, page, pageSize int) ([]*domain.GetPerformanceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilterPerformancesByTags", tags, page, pageSize)
	ret0, _ := ret[0].([]*domain.GetPerformanceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterPerformancesByTags indicates an expected call of FilterPerformancesByTags.
func (mr *MockTheatreRepositoryMockRecorder) FilterPerformancesByTags(tags, page, pageSize interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterPerformancesByTags", reflect.TypeOf((*MockTheatreRepository)(nil).FilterPerformancesByTags), tags, page, pageSize)
}

// GetAllPerformances mocks base method.
func (m *MockTheatreRepository) GetAllPerformances(page, pageSize int) ([]*domain.GetPerformanceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllPerformances", page, pageSize)
	ret0, _ := ret[0].([]*domain.GetPerformanceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllPerformances indicates an expected call of GetAllPerformances.
func (mr *MockTheatreRepositoryMockRecorder) GetAllPerformances(page, pageSize interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllPerformances", reflect.TypeOf((*MockTheatreRepository)(nil).GetAllPerformances), page, pageSize)
}

// GetPerformanceByID mocks base method.
func (m *MockTheatreRepository) GetPerformanceByID(id primitive.ObjectID) (*domain.GetPerformanceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPerformanceByID", id)
	ret0, _ := ret[0].(*domain.GetPerformanceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPerformanceByID indicates an expected call of GetPerformanceByID.
func (mr *MockTheatreRepositoryMockRecorder) GetPerformanceByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPerformanceByID", reflect.TypeOf((*MockTheatreRepository)(nil).GetPerformanceByID), id)
}

// GetTotalPerformancesCount mocks base method.
func (m *MockTheatreRepository) GetTotalPerformancesCount() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotalPerformancesCount")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTotalPerformancesCount indicates an expected call of GetTotalPerformancesCount.
func (mr *MockTheatreRepositoryMockRecorder) GetTotalPerformancesCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotalPerformancesCount", reflect.TypeOf((*MockTheatreRepository)(nil).GetTotalPerformancesCount))
}

// SearchPerformances mocks base method.
func (m *MockTheatreRepository) SearchPerformances(query string, page, pageSize int) ([]*domain.GetPerformanceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchPerformances", query, page, pageSize)
	ret0, _ := ret[0].([]*domain.GetPerformanceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchPerformances indicates an expected call of SearchPerformances.
func (mr *MockTheatreRepositoryMockRecorder) SearchPerformances(query, page, pageSize interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchPerformances", reflect.TypeOf((*MockTheatreRepository)(nil).SearchPerformances), query, page, pageSize)
}

// UpdatePerformance mocks base method.
func (m *MockTheatreRepository) UpdatePerformance(id primitive.ObjectID, request *domain.UpdatePerformanceRequest) (*domain.UpdatePerformanceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePerformance", id, request)
	ret0, _ := ret[0].(*domain.UpdatePerformanceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePerformance indicates an expected call of UpdatePerformance.
func (mr *MockTheatreRepositoryMockRecorder) UpdatePerformance(id, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePerformance", reflect.TypeOf((*MockTheatreRepository)(nil).UpdatePerformance), id, request)
}