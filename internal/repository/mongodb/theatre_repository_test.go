package repository_test

import (
	"errors"
	"events/internal/domain"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"go.mongodb.org/mongo-driver/bson/primitive"

	mock_repo "events/internal/repository/mocks"
)

func TestGetAllPerformances(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repo.NewMockTheatreRepository(ctrl)

	tests := []struct {
		name     string
		page     int
		pageSize int
		want     []*domain.GetPerformanceResponse
		wantErr  bool
		err      error
	}{
		{
			name:     "Successful retrieval",
			page:     1,
			pageSize: 10,
			want: []*domain.GetPerformanceResponse{
				{
					ID:          primitive.NewObjectID(),
					Cover:       "cover1.jpg",
					Name:        "Test Performance 1",
					Description: "This is a test performance 1",
					Duration:    "120 mins",
					Age:         "18+",
					Categories:  []string{"Action", "Adventure"},
					Tags:        []string{"test", "performance"},
					Media:       []string{"media1", "media2"},
				},
				{
					ID:          primitive.NewObjectID(),
					Cover:       "cover2.jpg",
					Name:        "Test Performance 2",
					Description: "This is a test performance 2",
					Duration:    "150 mins",
					Age:         "18+",
					Categories:  []string{"Comedy", "Drama"},
					Tags:        []string{"test", "performance"},
					Media:       []string{"media3", "media4"},
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name:     "Error retrieving performance list",
			page:     1,
			pageSize: 10,
			want:     nil,
			wantErr:  true,
			err:      errors.New("error retrieving performance list"),
		},
		{
			name:     "Error decoding performance",
			page:     1,
			pageSize: 10,
			want:     nil,
			wantErr:  true,
			err:      errors.New("error decoding performance"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.EXPECT().GetAllPerformances(tt.page, tt.pageSize).Return(tt.want, tt.err)

			got, err := mockRepo.GetAllPerformances(tt.page, tt.pageSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllPerformances() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllPerformances() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPerformanceByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repo.NewMockTheatreRepository(ctrl)

	tests := []struct {
		name    string
		id      primitive.ObjectID
		want    *domain.GetPerformanceResponse
		wantErr bool
		err     error
	}{
		{
			name: "Successful retrieval",
			id:   primitive.NewObjectID(),
			want: &domain.GetPerformanceResponse{
				ID:          primitive.NewObjectID(),
				Cover:       "cover.jpg",
				Name:        "Test Performance",
				Description: "This is a test performance",
				Duration:    "120 mins",
				Age:         "18+",
				Categories:  []string{"Action", "Adventure"},
				Tags:        []string{"test", "performance"},
				Media:       []string{"media1", "media2"},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name:    "No documents found",
			id:      primitive.NewObjectID(),
			want:    nil,
			wantErr: false,
			err:     nil,
		},
		{
			name:    "Error getting performance by ID",
			id:      primitive.NewObjectID(),
			want:    nil,
			wantErr: true,
			err:     errors.New("error getting performance by ID"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.EXPECT().GetPerformanceByID(tt.id).Return(tt.want, tt.err)

			got, err := mockRepo.GetPerformanceByID(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPerformanceByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPerformanceByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreatePerformance(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repo.NewMockTheatreRepository(ctrl)

	tests := []struct {
		name    string
		request *domain.CreatePerformanceRequest
		want    *domain.CreatePerformanceResponse
		wantErr bool
		err     error
	}{
		{
			name: "Successful creation",
			request: &domain.CreatePerformanceRequest{
				Cover:       "cover.jpg",
				Name:        "Test Performance",
				Description: "This is a test performance",
				Duration:    "120 mins",
				Age:         "18+",
				Categories:  []string{"Action", "Adventure"},
				Tags:        []string{"test", "performance"},
				Media:       []string{"media1", "media2"},
			},
			want: &domain.CreatePerformanceResponse{
				ID:          primitive.NewObjectID(),
				Cover:       "cover.jpg",
				Name:        "Test Performance",
				Description: "This is a test performance",
				Duration:    "120 mins",
				Age:         "18+",
				Categories:  []string{"Action", "Adventure"},
				Tags:        []string{"test", "performance"},
				Media:       []string{"media1", "media2"},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Error inserting performance document",
			request: &domain.CreatePerformanceRequest{
				Cover:       "cover.jpg",
				Name:        "Test Performance",
				Description: "This is a test performance",
				Duration:    "120 mins",
				Age:         "18+",
				Categories:  []string{"Action", "Adventure"},
				Tags:        []string{"test", "performance"},
				Media:       []string{"media1", "media2"},
			},
			want:    nil,
			wantErr: true,
			err:     errors.New("error inserting performance document"),
		},
		{
			name: "Error getting inserted performance ID",
			request: &domain.CreatePerformanceRequest{
				Cover:       "cover.jpg",
				Name:        "Test Performance",
				Description: "This is a test performance",
				Duration:    "120 mins",
				Age:         "18+",
				Categories:  []string{"Action", "Adventure"},
				Tags:        []string{"test", "performance"},
				Media:       []string{"media1", "media2"},
			},
			want:    nil,
			wantErr: true,
			err:     errors.New("error getting inserted performance ID"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.EXPECT().CreatePerformance(tt.request).Return(tt.want, tt.err)

			got, err := mockRepo.CreatePerformance(tt.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreatePerformance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreatePerformance() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdatePerformance(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repo.NewMockTheatreRepository(ctrl)

	tests := []struct {
		name    string
		id      primitive.ObjectID
		update  *domain.UpdatePerformanceRequest
		want    *domain.UpdatePerformanceResponse
		wantErr bool
		err     error
	}{
		{
			name: "Successful update",
			id:   primitive.NewObjectID(),
			update: &domain.UpdatePerformanceRequest{
				Cover:       "new_cover.jpg",
				Name:        "New Performance",
				Description: "This is a new performance",
				Duration:    "150 mins",
				Age:         "18+",
				Categories:  []string{"Drama", "Thriller"},
				Tags:        []string{"new", "performance"},
				Media:       []string{"media3", "media4"},
			},
			want: &domain.UpdatePerformanceResponse{
				ID:          primitive.NewObjectID(),
				Cover:       "new_cover.jpg",
				Name:        "New Performance",
				Description: "This is a new performance",
				Duration:    "150 mins",
				Age:         "18+",
				Categories:  []string{"Drama", "Thriller"},
				Tags:        []string{"new", "performance"},
				Media:       []string{"media3", "media4"},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name:    "Error updating performance",
			id:      primitive.NewObjectID(),
			update:  &domain.UpdatePerformanceRequest{},
			want:    nil,
			wantErr: true,
			err:     errors.New("error updating performance"),
		},
		{
			name:    "Error fetching updated performance",
			id:      primitive.NewObjectID(),
			update:  &domain.UpdatePerformanceRequest{},
			want:    nil,
			wantErr: true,
			err:     errors.New("error fetching updated performance"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.EXPECT().UpdatePerformance(tt.id, tt.update).Return(tt.want, tt.err)

			got, err := mockRepo.UpdatePerformance(tt.id, tt.update)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdatePerformance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdatePerformance() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchPerformances(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repo.NewMockTheatreRepository(ctrl)

	tests := []struct {
		name     string
		query    string
		page     int
		pageSize int
		want     []*domain.GetPerformanceResponse
		wantErr  bool
		err      error
	}{
		{
			name:     "Successful search",
			query:    "Test Performance",
			page:     1,
			pageSize: 10,
			want: []*domain.GetPerformanceResponse{
				{
					ID:          primitive.NewObjectID(),
					Cover:       "cover1.jpg",
					Name:        "Test Performance 1",
					Description: "This is a test performance 1",
					Duration:    "120 mins",
					Age:         "18+",
					Categories:  []string{"Action", "Adventure"},
					Tags:        []string{"test", "performance"},
					Media:       []string{"media1", "media2"},
				},
				{
					ID:          primitive.NewObjectID(),
					Cover:       "cover2.jpg",
					Name:        "Test Performance 2",
					Description: "This is a test performance 2",
					Duration:    "150 mins",
					Age:         "18+",
					Categories:  []string{"Comedy", "Drama"},
					Tags:        []string{"test", "performance"},
					Media:       []string{"media3", "media4"},
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name:     "Error retrieving performance list",
			query:    "Test Performance",
			page:     1,
			pageSize: 10,
			want:     nil,
			wantErr:  true,
			err:      errors.New("error retrieving performance list"),
		},
		{
			name:     "Error decoding performance",
			query:    "Test Performance",
			page:     1,
			pageSize: 10,
			want:     nil,
			wantErr:  true,
			err:      errors.New("error decoding performance"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.EXPECT().SearchPerformances(tt.query, tt.page, tt.pageSize).Return(tt.want, tt.err)

			got, err := mockRepo.SearchPerformances(tt.query, tt.page, tt.pageSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchPerformances() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchPerformances() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterPerformancesByTags(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repo.NewMockTheatreRepository(ctrl)

	tests := []struct {
		name     string
		tags     []string
		page     int
		pageSize int
		want     []*domain.GetPerformanceResponse
		wantErr  bool
		err      error
	}{
		{
			name:     "Successful filter",
			tags:     []string{"Action", "Adventure"},
			page:     1,
			pageSize: 10,
			want: []*domain.GetPerformanceResponse{
				{
					ID:          primitive.NewObjectID(),
					Cover:       "cover1.jpg",
					Name:        "Test Performance 1",
					Description: "This is a test performance 1",
					Duration:    "120 mins",
					Age:         "18+",
					Categories:  []string{"Action", "Adventure"},
					Tags:        []string{"test", "performance"},
					Media:       []string{"media1", "media2"},
				},
				{
					ID:          primitive.NewObjectID(),
					Cover:       "cover2.jpg",
					Name:        "Test Performance 2",
					Description: "This is a test performance 2",
					Duration:    "150 mins",
					Age:         "18+",
					Categories:  []string{"Comedy", "Drama"},
					Tags:        []string{"test", "performance"},
					Media:       []string{"media3", "media4"},
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name:     "Error retrieving performance list",
			tags:     []string{"Action", "Adventure"},
			page:     1,
			pageSize: 10,
			want:     nil,
			wantErr:  true,
			err:      errors.New("error retrieving performance list"),
		},
		{
			name:     "Error decoding performance",
			tags:     []string{"Action", "Adventure"},
			page:     1,
			pageSize: 10,
			want:     nil,
			wantErr:  true,
			err:      errors.New("error decoding performance"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.EXPECT().FilterPerformancesByTags(tt.tags, tt.page, tt.pageSize).Return(tt.want, tt.err)

			got, err := mockRepo.FilterPerformancesByTags(tt.tags, tt.page, tt.pageSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("FilterPerformancesByTags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterPerformancesByTags() got = %v, want %v", got, tt.want)
			}
		})
	}
}
