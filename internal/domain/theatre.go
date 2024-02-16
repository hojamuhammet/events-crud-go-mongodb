package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type CommonPerformanceRequest struct {
	Cover       string   `json:"cover" bson:"cover"`
	Name        string   `json:"name" bson:"name"`
	Description string   `json:"description" bson:"description"`
	Duration    string   `json:"duration" bson:"duration"`
	Age         string   `json:"age" bson:"age"`
	Categories  []string `json:"categories" bson:"categories"`
	Tags        []string `json:"tags" bson:"tags"`
	Media       []string `json:"media" bson:"media"`
}

type CommonPerformanceResponse struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Cover       string             `json:"cover" bson:"cover"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Duration    string             `json:"duration" bson:"duration"`
	Age         string             `json:"age" bson:"age"`
	Categories  []string           `json:"categories" bson:"categories"`
	Tags        []string           `json:"tags" bson:"tags"`
	Media       []string           `json:"media" bson:"media"`
}

type GetPerformanceResponse CommonPerformanceResponse
type CreatePerformanceRequest CommonPerformanceRequest
type CreatePerformanceResponse CommonPerformanceResponse
type UpdatePerformanceRequest CommonPerformanceRequest
type UpdatePerformanceResponse CommonPerformanceResponse
