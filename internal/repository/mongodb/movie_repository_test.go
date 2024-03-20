package repository_test

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"events/internal/domain"
	mock_repository "events/internal/repository/mocks"
)

func TestGetAllMovies(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockMovieRepository(ctrl)

	tests := []struct {
		name     string
		page     int
		pageSize int
		want     []*domain.GetMovieResponse
		wantErr  bool
		err      error
	}{
		{
			name:     "Successful retrieval",
			page:     1,
			pageSize: 10,
			want: []*domain.GetMovieResponse{
				{
					ID:           primitive.NewObjectID(),
					Cover:        "cover1.jpg",
					Name:         "Test Movie 1",
					OriginalName: "Test Movie Original 1",
					Description:  "This is a test movie 1",
					Duration:     "120 mins",
					ReleaseDate:  time.Now(),
					Age:          "18+",
					Categories:   []string{"Action", "Adventure"},
					Tags:         []string{"test", "movie"},
					Media:        []string{"media1", "media2"},
				},
				{
					ID:           primitive.NewObjectID(),
					Cover:        "cover2.jpg",
					Name:         "Test Movie 2",
					OriginalName: "Test Movie Original 2",
					Description:  "This is a test movie 2",
					Duration:     "150 mins",
					ReleaseDate:  time.Now(),
					Age:          "18+",
					Categories:   []string{"Comedy", "Drama"},
					Tags:         []string{"test", "movie"},
					Media:        []string{"media3", "media4"},
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name:     "Error retrieving movie list",
			page:     1,
			pageSize: 10,
			want:     nil,
			wantErr:  true,
			err:      errors.New("error retrieving movie list"),
		},
		{
			name:     "Error decoding movie",
			page:     1,
			pageSize: 10,
			want:     nil,
			wantErr:  true,
			err:      errors.New("error decoding movie"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.EXPECT().GetAllMovies(tt.page, tt.pageSize).Return(tt.want, tt.err)

			got, err := mockRepo.GetAllMovies(tt.page, tt.pageSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllMovies() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllMovies() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMovieByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockMovieRepository(ctrl)

	id := primitive.NewObjectID()
	expectedMovie := &domain.GetMovieResponse{
		ID:           id,
		Cover:        "cover",
		Name:         "name",
		OriginalName: "originalName",
		Description:  "description",
		Duration:     "duration",
		ReleaseDate:  time.Now(),
		Age:          "age",
		Categories:   []string{"category1", "category2"},
		Tags:         []string{"tag1", "tag2"},
		Media:        []string{"media1", "media2"},
	}

	testCases := []struct {
		name          string
		expectedMovie *domain.GetMovieResponse
		expectedErr   error
	}{
		{
			name:          "Success",
			expectedMovie: expectedMovie,
			expectedErr:   nil,
		},
		{
			name:          "Error",
			expectedMovie: nil,
			expectedErr:   errors.New("error getting movie by ID"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo.EXPECT().GetMovieByID(id).Return(tc.expectedMovie, tc.expectedErr).Times(1)

			movie, err := mockRepo.GetMovieByID(id)

			if tc.expectedErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedMovie, movie)
			}
		})
	}
}

func TestCreateMovie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockMovieRepository(ctrl)

	tests := []struct {
		name    string
		request *domain.CreateMovieRequest
		want    *domain.CreateMovieResponse
		wantErr bool
		err     error
	}{
		{
			name: "Successful creation",
			request: &domain.CreateMovieRequest{
				Cover:        "cover.jpg",
				Name:         "Test Movie",
				OriginalName: "Test Movie Original",
				Description:  "This is a test movie",
				Duration:     "120 mins",
				ReleaseDate:  time.Now(),
				Age:          "18+",
				Categories:   []string{"Action", "Adventure"},
				Tags:         []string{"test", "movie"},
				Media:        []string{"media1", "media2"},
			},
			want: &domain.CreateMovieResponse{
				ID:           primitive.NewObjectID(),
				Cover:        "cover.jpg",
				Name:         "Test Movie",
				OriginalName: "Test Movie Original",
				Description:  "This is a test movie",
				Duration:     "120 mins",
				ReleaseDate:  time.Now(),
				Age:          "18+",
				Categories:   []string{"Action", "Adventure"},
				Tags:         []string{"test", "movie"},
				Media:        []string{"media1", "media2"},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Error inserting movie document",
			request: &domain.CreateMovieRequest{
				Cover:        "cover.jpg",
				Name:         "Test Movie",
				OriginalName: "Test Movie Original",
				Description:  "This is a test movie",
				Duration:     "120 mins",
				ReleaseDate:  time.Now(),
				Age:          "18+",
				Categories:   []string{"Action", "Adventure"},
				Tags:         []string{"test", "movie"},
				Media:        []string{"media1", "media2"},
			},
			want:    nil,
			wantErr: true,
			err:     errors.New("error inserting movie document"),
		},
		{
			name: "Error getting inserted movie ID",
			request: &domain.CreateMovieRequest{
				Cover:        "cover.jpg",
				Name:         "Test Movie",
				OriginalName: "Test Movie Original",
				Description:  "This is a test movie",
				Duration:     "120 mins",
				ReleaseDate:  time.Now(),
				Age:          "18+",
				Categories:   []string{"Action", "Adventure"},
				Tags:         []string{"test", "movie"},
				Media:        []string{"media1", "media2"},
			},
			want:    nil,
			wantErr: true,
			err:     errors.New("error getting inserted movie ID"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.EXPECT().CreateMovie(tt.request).Return(tt.want, tt.err)

			got, err := mockRepo.CreateMovie(tt.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateMovie() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateMovie() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateMovie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockMovieRepository(ctrl)

	tests := []struct {
		name    string
		id      primitive.ObjectID
		request *domain.UpdateMovieRequest
		want    *domain.UpdateMovieResponse
		wantErr bool
		err     error
	}{
		{
			name: "Successful update",
			id:   primitive.NewObjectID(),
			request: &domain.UpdateMovieRequest{
				Cover:        "new_cover.jpg",
				Name:         "New Test Movie",
				OriginalName: "New Test Movie Original",
				Description:  "This is a new test movie",
				Duration:     "150 mins",
				ReleaseDate:  time.Now(),
				Age:          "18+",
				Categories:   []string{"Action", "Adventure"},
				Tags:         []string{"new_test", "movie"},
				Media:        []string{"new_media1", "new_media2"},
			},
			want: &domain.UpdateMovieResponse{
				ID:           primitive.NewObjectID(),
				Cover:        "new_cover.jpg",
				Name:         "New Test Movie",
				OriginalName: "New Test Movie Original",
				Description:  "This is a new test movie",
				Duration:     "150 mins",
				ReleaseDate:  time.Now(),
				Age:          "18+",
				Categories:   []string{"Action", "Adventure"},
				Tags:         []string{"new_test", "movie"},
				Media:        []string{"new_media1", "new_media2"},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Error updating movie",
			id:   primitive.NewObjectID(),
			request: &domain.UpdateMovieRequest{
				Cover:        "new_cover.jpg",
				Name:         "New Test Movie",
				OriginalName: "New Test Movie Original",
				Description:  "This is a new test movie",
				Duration:     "150 mins",
				ReleaseDate:  time.Now(),
				Age:          "18+",
				Categories:   []string{"Action", "Adventure"},
				Tags:         []string{"new_test", "movie"},
				Media:        []string{"new_media1", "new_media2"},
			},
			want:    nil,
			wantErr: true,
			err:     errors.New("error updating movie"),
		},
		{
			name: "Error fetching updated movie",
			id:   primitive.NewObjectID(),
			request: &domain.UpdateMovieRequest{
				Cover:        "new_cover.jpg",
				Name:         "New Test Movie",
				OriginalName: "New Test Movie Original",
				Description:  "This is a new test movie",
				Duration:     "150 mins",
				ReleaseDate:  time.Now(),
				Age:          "18+",
				Categories:   []string{"Action", "Adventure"},
				Tags:         []string{"new_test", "movie"},
				Media:        []string{"new_media1", "new_media2"},
			},
			want:    nil,
			wantErr: true,
			err:     errors.New("error fetching updated movie"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.EXPECT().UpdateMovie(tt.id, tt.request).Return(tt.want, tt.err)

			got, err := mockRepo.UpdateMovie(tt.id, tt.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateMovie() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateMovie() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteMovie(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockMovieRepository(ctrl)

	id := primitive.NewObjectID()

	testCases := []struct {
		name        string
		expectedErr error
	}{
		{
			name:        "Success",
			expectedErr: nil,
		},
		{
			name:        "Error",
			expectedErr: errors.New("error deleting movie"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo.EXPECT().DeleteMovie(id).Return(tc.expectedErr).Times(1)

			err := mockRepo.DeleteMovie(id)

			if tc.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedErr, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestSearchMovies(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockMovieRepository(ctrl)

	tests := []struct {
		name     string
		query    string
		page     int
		pageSize int
		want     []*domain.GetMovieResponse
		wantErr  bool
		err      error
	}{
		{
			name:     "Successful search",
			query:    "Test Movie",
			page:     1,
			pageSize: 10,
			want: []*domain.GetMovieResponse{
				{
					ID:           primitive.NewObjectID(),
					Cover:        "cover1.jpg",
					Name:         "Test Movie 1",
					OriginalName: "Test Movie Original 1",
					Description:  "This is a test movie 1",
					Duration:     "120 mins",
					ReleaseDate:  time.Now(),
					Age:          "18+",
					Categories:   []string{"Action", "Adventure"},
					Tags:         []string{"test", "movie"},
					Media:        []string{"media1", "media2"},
				},
				{
					ID:           primitive.NewObjectID(),
					Cover:        "cover2.jpg",
					Name:         "Test Movie 2",
					OriginalName: "Test Movie Original 2",
					Description:  "This is a test movie 2",
					Duration:     "150 mins",
					ReleaseDate:  time.Now(),
					Age:          "18+",
					Categories:   []string{"Comedy", "Drama"},
					Tags:         []string{"test", "movie"},
					Media:        []string{"media3", "media4"},
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name:     "Error retrieving movie list",
			query:    "Test Movie",
			page:     1,
			pageSize: 10,
			want:     nil,
			wantErr:  true,
			err:      errors.New("error retrieving movie list"),
		},
		{
			name:     "Error decoding movie",
			query:    "Test Movie",
			page:     1,
			pageSize: 10,
			want:     nil,
			wantErr:  true,
			err:      errors.New("error decoding movie"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.EXPECT().SearchMovies(tt.query, tt.page, tt.pageSize).Return(tt.want, tt.err)

			got, err := mockRepo.SearchMovies(tt.query, tt.page, tt.pageSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchMovies() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchMovies() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterMoviesByTags(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockMovieRepository(ctrl)

	tests := []struct {
		name     string
		tags     []string
		page     int
		pageSize int
		want     []*domain.GetMovieResponse
		wantErr  bool
		err      error
	}{
		{
			name:     "Successful filter",
			tags:     []string{"Action", "Adventure"},
			page:     1,
			pageSize: 10,
			want: []*domain.GetMovieResponse{
				{
					ID:           primitive.NewObjectID(),
					Cover:        "cover1.jpg",
					Name:         "Test Movie 1",
					OriginalName: "Test Movie Original 1",
					Description:  "This is a test movie 1",
					Duration:     "120 mins",
					ReleaseDate:  time.Now(),
					Age:          "18+",
					Categories:   []string{"Action", "Adventure"},
					Tags:         []string{"test", "movie"},
					Media:        []string{"media1", "media2"},
				},
				{
					ID:           primitive.NewObjectID(),
					Cover:        "cover2.jpg",
					Name:         "Test Movie 2",
					OriginalName: "Test Movie Original 2",
					Description:  "This is a test movie 2",
					Duration:     "150 mins",
					ReleaseDate:  time.Now(),
					Age:          "18+",
					Categories:   []string{"Comedy", "Drama"},
					Tags:         []string{"test", "movie"},
					Media:        []string{"media3", "media4"},
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name:     "Error retrieving movie list",
			tags:     []string{"Action", "Adventure"},
			page:     1,
			pageSize: 10,
			want:     nil,
			wantErr:  true,
			err:      errors.New("error retrieving movie list"),
		},
		{
			name:     "Error decoding movie",
			tags:     []string{"Action", "Adventure"},
			page:     1,
			pageSize: 10,
			want:     nil,
			wantErr:  true,
			err:      errors.New("error decoding movie"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo.EXPECT().FilterMoviesByTags(tt.tags, tt.page, tt.pageSize).Return(tt.want, tt.err)

			got, err := mockRepo.FilterMoviesByTags(tt.tags, tt.page, tt.pageSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("FilterMoviesByTags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterMoviesByTags() got = %v, want %v", got, tt.want)
			}
		})
	}
}
