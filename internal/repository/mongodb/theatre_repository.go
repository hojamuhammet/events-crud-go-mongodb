package repository

import (
	"context"
	"errors"
	"events/internal/domain"
	"events/pkg/lib/utils"
	"log/slog"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBTheatreRepository struct {
	collection *mongo.Collection
}

func NewMongoDBTheatreRepository(collection *mongo.Collection) *MongoDBTheatreRepository {
	return &MongoDBTheatreRepository{
		collection: collection,
	}
}

func (r *MongoDBTheatreRepository) GetAllPerformances(page, pageSize int) ([]*domain.GetPerformanceResponse, error) {
	skip := (page - 1) * pageSize

	filter := bson.M{} // Empty filter to retrieve all documents

	opts := options.Find().
		SetSkip(int64(skip)).
		SetLimit(int64(pageSize))

	cursor, err := r.collection.Find(context.Background(), filter, opts)
	if err != nil {
		slog.Error("error retrieving performance list", utils.Err(err))
		return nil, err
	}
	defer cursor.Close(context.Background())

	var performances []*domain.GetPerformanceResponse
	for cursor.Next(context.Background()) {
		var performance domain.GetPerformanceResponse
		if err := cursor.Decode(&performance); err != nil {
			slog.Error("Error decoding performance: ", utils.Err(err))
			return nil, err
		}
		performances = append(performances, &performance)
	}

	return performances, nil
}

func (r *MongoDBTheatreRepository) GetPerformanceByID(id primitive.ObjectID) (*domain.GetPerformanceResponse, error) {
	filter := bson.M{"_id": id}

	var performance domain.GetPerformanceResponse

	err := r.collection.FindOne(context.Background(), filter).Decode(&performance)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		slog.Error("Error getting performance by ID: %v", utils.Err(err))
		return nil, err
	}
	return &performance, nil
}

func (r *MongoDBTheatreRepository) CreatePerformance(theatre *domain.CreatePerformanceRequest) (*domain.CreatePerformanceResponse, error) {
	t := domain.CreatePerformanceResponse{
		Cover:       theatre.Cover,
		Name:        theatre.Name,
		Description: theatre.Description,
		Duration:    theatre.Duration,
		Age:         theatre.Age,
		Categories:  theatre.Categories,
		Tags:        theatre.Tags,
		Media:       theatre.Media,
	}

	result, err := r.collection.InsertOne(context.Background(), t)
	if err != nil {
		slog.Error("error inserting performance document: %v", utils.Err(err))
		return nil, err
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		slog.Error("error getting inserted performance ID")
		return nil, errors.New("error getting inserted performance ID")
	}

	t.ID = insertedID

	return &t, nil
}

func (r *MongoDBTheatreRepository) UpdatePerformance(id primitive.ObjectID, update *domain.UpdatePerformanceRequest) (*domain.UpdatePerformanceResponse, error) {
	updateFields := bson.M{}

	if update.Cover != "" {
		updateFields["cover"] = update.Cover
	}
	if update.Name != "" {
		updateFields["name"] = update.Name
	}
	if update.Description != "" {
		updateFields["description"] = update.Description
	}
	if update.Duration != "" {
		updateFields["duration"] = update.Duration
	}
	if update.Age != "" {
		updateFields["age"] = update.Age
	}
	if len(update.Categories) > 0 {
		updateFields["categories"] = update.Categories
	}
	if len(update.Tags) > 0 {
		updateFields["tags"] = update.Tags
	}
	if len(update.Media) > 0 {
		updateFields["media"] = update.Media
	}
	filter := bson.M{"_id": id}

	result, err := r.collection.UpdateOne(context.Background(), filter, bson.M{"$set": updateFields})
	if err != nil {
		slog.Error("error updating performance: ", utils.Err(err))
		return nil, err
	}

	if result.ModifiedCount == 0 {
		slog.Warn("No performance document were modified")
		return nil, errors.New("performance not found")
	}

	updatePerformance, err := r.GetPerformanceByID(id)
	if err != nil {
		slog.Error("error fetching updated performance: ", utils.Err(err))
		return nil, err
	}

	updateResponse := &domain.UpdatePerformanceResponse{
		ID:          updatePerformance.ID,
		Cover:       updatePerformance.Cover,
		Name:        updatePerformance.Name,
		Description: updatePerformance.Description,
		Duration:    updatePerformance.Duration,
		Age:         updatePerformance.Age,
		Categories:  updatePerformance.Categories,
		Tags:        updatePerformance.Tags,
		Media:       updatePerformance.Media,
	}
	return updateResponse, nil
}

func (r *MongoDBTheatreRepository) DeletePerformance(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}

	result, err := r.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		slog.Error("Error deleting performance: ", utils.Err(err))
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("performance not found")
	}

	return nil
}

func (r *MongoDBTheatreRepository) SearchPerformances(query string, page int, pageSize int) ([]*domain.GetPerformanceResponse, error) {
	offset := (page - 1) * pageSize

	options := options.Find().SetSkip(int64(offset)).SetLimit(int64(pageSize))

	filter := bson.M{
		"$or": []interface{}{
			bson.M{"name": bson.M{"$regex": query, "$options": "i"}},
			bson.M{"description": bson.M{"$regex": query, "$options": "i"}},
		},
	}

	cursor, err := r.collection.Find(context.Background(), filter, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var performances []*domain.GetPerformanceResponse

	for cursor.Next(context.Background()) {
		var performance domain.GetPerformanceResponse
		if err := cursor.Decode(&performance); err != nil {
			return nil, err
		}
		performances = append(performances, &performance)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return performances, nil
}
