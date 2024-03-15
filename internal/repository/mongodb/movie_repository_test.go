package repository_test

import (
	"errors"
	"events/internal/domain"
	"events/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestGetAllMovies(t *testing.T) {
	mockRepo := new(mocks.MockMongoDBMovieRepository)
	movies := []*domain.GetMovieResponse{
		{
			ID:           primitive.NewObjectID(),
			Cover:        "cover1.jpg",
			Name:         "Movie Name 1",
			OriginalName: "Original Movie Name 1",
			Description:  "This is a description of the first movie.",
			Duration:     "120 minutes",
			ReleaseDate:  time.Now(),
			Age:          "PG-13",
			Categories:   []string{"Action", "Adventure"},
			Tags:         []string{"exciting", "thriller"},
			Media:        []string{"media1.mp4", "media2.mp4"},
		},
		{
			ID:           primitive.NewObjectID(),
			Cover:        "cover2.jpg",
			Name:         "Movie Name 2",
			OriginalName: "Original Movie Name 2",
			Description:  "This is a description of the second movie.",
			Duration:     "150 minutes",
			ReleaseDate:  time.Now(),
			Age:          "R",
			Categories:   []string{"Drama", "Thriller"},
			Tags:         []string{"intense", "gripping"},
			Media:        []string{"media3.mp4", "media4.mp4"},
		},
	}

	tests := []struct {
		name     string
		mock     func()
		wantErr  bool
		wantResp []*domain.GetMovieResponse
	}{
		{
			name: "success",
			mock: func() {
				mockRepo.On("GetAllMovies", 1, 10).Return(movies, nil).Once()
			},
			wantErr:  false,
			wantResp: movies,
		},
		{
			name: "find error",
			mock: func() {
				mockRepo.On("GetAllMovies", 1, 10).Return(nil, errors.New("find error")).Once()
			},
			wantErr:  true,
			wantResp: nil,
		},
		{
			name: "decode error",
			mock: func() {
				mockRepo.On("GetAllMovies", 1, 10).Return(nil, errors.New("decode error")).Once()
			},
			wantErr:  true,
			wantResp: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			result, err := mockRepo.GetAllMovies(1, 10)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.wantResp, result)

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestGetMovieByID(t *testing.T) {
	mockRepo := new(mocks.MockMongoDBMovieRepository)
	id := primitive.NewObjectID()
	movie := &domain.GetMovieResponse{ID: id}

	tests := []struct {
		name     string
		mock     func()
		wantErr  bool
		wantResp *domain.GetMovieResponse
	}{
		{
			name: "success",
			mock: func() {
				mockRepo.On("GetMovieByID", id).Return(movie, nil).Once()
			},
			wantErr:  false,
			wantResp: movie,
		},
		{
			name: "not found",
			mock: func() {
				mockRepo.On("GetMovieByID", id).Return(nil, errors.New("not found")).Once()
			},
			wantErr:  true,
			wantResp: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			result, err := mockRepo.GetMovieByID(id)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.wantResp, result)

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestCreateMovie(t *testing.T) {
	mockRepo := new(mocks.MockMongoDBMovieRepository)
	id := primitive.NewObjectID()
	movie := &domain.CreateMovieRequest{
		Cover:        "cover.jpg",
		Name:         "Movie Name",
		OriginalName: "Original Movie Name",
		Description:  "This is a description of the movie.",
		Duration:     "120 minutes",
		ReleaseDate:  time.Now(),
		Age:          "13+",
		Categories:   []string{"Action", "Adventure"},
		Tags:         []string{"exciting", "thriller"},
		Media:        []string{"media1.mp4", "media2.mp4"},
	}

	response := &domain.CreateMovieResponse{
		ID:           id,
		Cover:        "cover.jpg",
		Name:         "Movie Name",
		OriginalName: "Original Movie Name",
		Description:  "This is a description of the movie.",
		Duration:     "120 minutes",
		ReleaseDate:  time.Now(),
		Age:          "13+",
		Categories:   []string{"Action", "Adventure"},
		Tags:         []string{"exciting", "thriller"},
		Media:        []string{"media1.mp4", "media2.mp4"},
	}

	tests := []struct {
		name     string
		mock     func()
		wantErr  bool
		wantResp *domain.CreateMovieResponse
	}{
		{
			name: "success",
			mock: func() {
				mockRepo.On("CreateMovie", movie).Return(response, nil).Once()
			},
			wantErr:  false,
			wantResp: response,
		},
		{
			name: "insert error",
			mock: func() {
				mockRepo.On("CreateMovie", movie).Return(nil, mongo.WriteException{}).Once()
			},
			wantErr:  true,
			wantResp: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			result, err := mockRepo.CreateMovie(movie)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.wantResp, result)

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestUpdateMovie(t *testing.T) {
	mockRepo := new(mocks.MockMongoDBMovieRepository)
	id := primitive.NewObjectID()
	movie := &domain.UpdateMovieRequest{
		Cover:        "cover.jpg",
		Name:         "Movie Name",
		OriginalName: "Original Movie Name",
		Description:  "This is a description of the movie.",
		Duration:     "120 minutes",
		ReleaseDate:  time.Now(),
		Age:          "13+",
		Categories:   []string{"Action", "Adventure"},
		Tags:         []string{"exciting", "thriller"},
		Media:        []string{"media1.mp4", "media2.mp4"},
	}

	response := &domain.UpdateMovieResponse{
		ID:           id,
		Cover:        "cover.jpg",
		Name:         "Movie Name",
		OriginalName: "Original Movie Name",
		Description:  "This is a description of the movie.",
		Duration:     "120 minutes",
		ReleaseDate:  time.Now(),
		Age:          "13+",
		Categories:   []string{"Action", "Adventure"},
		Tags:         []string{"exciting", "thriller"},
		Media:        []string{"media1.mp4", "media2.mp4"},
	}

	tests := []struct {
		name     string
		mock     func()
		wantErr  bool
		wantResp *domain.UpdateMovieResponse
	}{
		{
			name: "success",
			mock: func() {
				mockRepo.On("UpdateMovie", id, movie).Return(response, nil).Once()
			},
			wantErr:  false,
			wantResp: response,
		},
		{
			name: "update error",
			mock: func() {
				mockRepo.On("UpdateMovie", id, movie).Return(nil, mongo.WriteException{}).Once()
			},
			wantErr:  true,
			wantResp: nil,
		},
		{
			name: "fetch updated movie error",
			mock: func() {
				mockRepo.On("UpdateMovie", id, movie).Return(nil, errors.New("error fetching updated movie")).Once()
			},
			wantErr:  true,
			wantResp: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			result, err := mockRepo.UpdateMovie(id, movie)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.wantResp, result)

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestDeleteMovie(t *testing.T) {
	mockRepo := new(mocks.MockMongoDBMovieRepository)
	id := primitive.NewObjectID()

	tests := []struct {
		name    string
		mock    func()
		wantErr bool
	}{
		{
			name: "success",
			mock: func() {
				mockRepo.On("DeleteMovie", id).Return(nil).Once()
			},
			wantErr: false,
		},
		{
			name: "delete error",
			mock: func() {
				mockRepo.On("DeleteMovie", id).Return(mongo.WriteException{}).Once()
			},
			wantErr: true,
		},
		{
			name: "not found",
			mock: func() {
				mockRepo.On("DeleteMovie", id).Return(errors.New("movie not found")).Once()
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			err := mockRepo.DeleteMovie(id)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestSearchMovies(t *testing.T) {
	mockRepo := new(mocks.MockMongoDBMovieRepository)
	movies := []*domain.GetMovieResponse{
		{
			ID:           primitive.NewObjectID(),
			Cover:        "cover1.jpg",
			Name:         "Movie Name 1",
			OriginalName: "Original Movie Name 1",
			Description:  "This is a description of the first movie.",
			Duration:     "120 minutes",
			ReleaseDate:  time.Now(),
			Age:          "PG-13",
			Categories:   []string{"Action", "Adventure"},
			Tags:         []string{"exciting", "thriller"},
			Media:        []string{"media1.mp4", "media2.mp4"},
		},
		{
			ID:           primitive.NewObjectID(),
			Cover:        "cover2.jpg",
			Name:         "Movie Name 2",
			OriginalName: "Original Movie Name 2",
			Description:  "This is a description of the second movie.",
			Duration:     "150 minutes",
			ReleaseDate:  time.Now(),
			Age:          "R",
			Categories:   []string{"Drama", "Thriller"},
			Tags:         []string{"intense", "gripping"},
			Media:        []string{"media3.mp4", "media4.mp4"},
		},
	}

	tests := []struct {
		name     string
		mock     func()
		wantErr  bool
		wantResp []*domain.GetMovieResponse
	}{
		{
			name: "success",
			mock: func() {
				mockRepo.On("SearchMovies", "query", 1, 10).Return(movies, nil).Once()
			},
			wantErr:  false,
			wantResp: movies,
		},
		{
			name: "find error",
			mock: func() {
				mockRepo.On("SearchMovies", "query", 1, 10).Return(nil, errors.New("find error")).Once()
			},
			wantErr:  true,
			wantResp: nil,
		},
		{
			name: "decode error",
			mock: func() {
				mockRepo.On("SearchMovies", "query", 1, 10).Return(nil, errors.New("decode error")).Once()
			},
			wantErr:  true,
			wantResp: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			result, err := mockRepo.SearchMovies("query", 1, 10)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.wantResp, result)

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestFilterMoviesByTags(t *testing.T) {
	mockRepo := new(mocks.MockMongoDBMovieRepository)
	movies := []*domain.GetMovieResponse{
		{
			ID:           primitive.NewObjectID(),
			Cover:        "cover1.jpg",
			Name:         "Movie Name 1",
			OriginalName: "Original Movie Name 1",
			Description:  "This is a description of the first movie.",
			Duration:     "120 minutes",
			ReleaseDate:  time.Now(),
			Age:          "PG-13",
			Categories:   []string{"Action", "Adventure"},
			Tags:         []string{"exciting", "thriller"},
			Media:        []string{"media1.mp4", "media2.mp4"},
		},
		{
			ID:           primitive.NewObjectID(),
			Cover:        "cover2.jpg",
			Name:         "Movie Name 2",
			OriginalName: "Original Movie Name 2",
			Description:  "This is a description of the second movie.",
			Duration:     "150 minutes",
			ReleaseDate:  time.Now(),
			Age:          "R",
			Categories:   []string{"Drama", "Thriller"},
			Tags:         []string{"intense", "gripping"},
			Media:        []string{"media3.mp4", "media4.mp4"},
		},
	}

	tests := []struct {
		name     string
		mock     func()
		wantErr  bool
		wantResp []*domain.GetMovieResponse
	}{
		{
			name: "success",
			mock: func() {
				mockRepo.On("FilterMoviesByTags", []string{"tag1", "tag2"}, 1, 10).Return(movies, nil).Once()
			},
			wantErr:  false,
			wantResp: movies,
		},
		{
			name: "find error",
			mock: func() {
				mockRepo.On("FilterMoviesByTags", []string{"tag1", "tag2"}, 1, 10).Return(nil, errors.New("find error")).Once()
			},
			wantErr:  true,
			wantResp: nil,
		},
		{
			name: "decode error",
			mock: func() {
				mockRepo.On("FilterMoviesByTags", []string{"tag1", "tag2"}, 1, 10).Return(nil, errors.New("decode error")).Once()
			},
			wantErr:  true,
			wantResp: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			result, err := mockRepo.FilterMoviesByTags([]string{"tag1", "tag2"}, 1, 10)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.wantResp, result)

			mockRepo.AssertExpectations(t)
		})
	}
}
