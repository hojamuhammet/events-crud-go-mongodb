package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommonMovieRequest struct {
	Cover        string    `json:"cover" bson:"cover"`
	Name         string    `json:"name" bson:"name"`
	OriginalName string    `json:"originalName" bson:"originalName"`
	Description  string    `json:"description" bson:"description"`
	Duration     string    `json:"duration" bson:"duration"`
	ReleaseDate  time.Time `json:"releaseDate" bson:"releaseDate"`
	Age          string    `json:"age" bson:"age"`
	Categories   []string  `json:"categories" bson:"categories"`
	Tags         []string  `json:"tags" bson:"tags"`
	Media        []string  `json:"media" bson:"media"`
}

type CommonMovieResponse struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Cover        string             `json:"cover" bson:"cover"`
	Name         string             `json:"name" bson:"name"`
	OriginalName string             `json:"originalName" bson:"originalName"`
	Description  string             `json:"description" bson:"description"`
	Duration     string             `json:"duration" bson:"duration"`
	ReleaseDate  time.Time          `json:"releaseDate" bson:"releaseDate"`
	Age          string             `json:"age" bson:"age"`
	Categories   []string           `json:"categories" bson:"categories"`
	Tags         []string           `json:"tags" bson:"tags"`
	Media        []string           `json:"media" bson:"media"`
}

type GetMovieResponse CommonMovieResponse
type CreateMovieRequest CommonMovieRequest
type CreateMovieResponse CommonMovieResponse
type UpdateMovieRequest CommonMovieRequest
type UpdateMovieResponse CommonMovieResponse
