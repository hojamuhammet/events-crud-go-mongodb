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

type MongoDBMovieRepository struct {
	collection *mongo.Collection
}

func NewMongoDBMovieRepository(collection *mongo.Collection) *MongoDBMovieRepository {
	return &MongoDBMovieRepository{
		collection: collection,
	}
}

func (r *MongoDBMovieRepository) GetAllMovies(page, pageSize int) ([]*domain.GetMovieResponse, error) {
	skip := (page - 1) * pageSize

	filter := bson.M{}

	opts := options.Find().
		SetSkip(int64(skip)).
		SetLimit(int64(pageSize))

	cursor, err := r.collection.Find(context.Background(), filter, opts)
	if err != nil {
		slog.Error("error retrieving movie list", utils.Err(err))
		return nil, err
	}
	defer cursor.Close(context.Background())

	var movies []*domain.GetMovieResponse
	for cursor.Next(context.Background()) {
		var movie domain.GetMovieResponse
		if err := cursor.Decode(&movie); err != nil {
			slog.Error("Error decoding movie: ", utils.Err(err))
			return nil, err
		}
		movies = append(movies, &movie)
	}

	return movies, nil
}

func (r *MongoDBMovieRepository) GetMovieByID(id primitive.ObjectID) (*domain.GetMovieResponse, error) {
	filter := bson.M{"_id": id}

	var movie domain.GetMovieResponse

	err := r.collection.FindOne(context.Background(), filter).Decode(&movie)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		slog.Error("error getting movie by ID: %v", utils.Err(err))
		return nil, err
	}
	return &movie, nil
}

func (r *MongoDBMovieRepository) CreateMovie(movie *domain.CreateMovieRequest) (*domain.CreateMovieResponse, error) {
	m := domain.CreateMovieResponse{
		Cover:        movie.Cover,
		Name:         movie.Name,
		OriginalName: movie.OriginalName,
		Description:  movie.Description,
		Duration:     movie.Duration,
		ReleaseDate:  movie.ReleaseDate,
		Age:          movie.Age,
		Tags:         movie.Tags,
		Categories:   movie.Categories,
		Media:        movie.Media,
	}

	result, err := r.collection.InsertOne(context.Background(), m)
	if err != nil {
		slog.Error("error inserting movie document: %v", utils.Err(err))
		return nil, err
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		slog.Error("error getting inserted movie ID")
		return nil, errors.New("error getting inserted movie ID")
	}

	m.ID = insertedID

	return &m, nil
}

func (r *MongoDBMovieRepository) UpdateMovie(id primitive.ObjectID, movie *domain.UpdateMovieRequest) (*domain.UpdateMovieResponse, error) {
	update := bson.M{
		"$set": bson.M{
			"cover":        movie.Cover,
			"name":         movie.Name,
			"originalName": movie.OriginalName,
			"description":  movie.Description,
			"duration":     movie.Duration,
			"releaseDate":  movie.ReleaseDate,
			"age":          movie.Age,
			"categories":   movie.Categories,
			"tags":         movie.Tags,
			"media":        movie.Media,
		},
	}

	filter := bson.M{"_id": id}

	_, err := r.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		slog.Error("error updating movie: ", utils.Err(err))
		return nil, err
	}

	updatedMovie, err := r.GetMovieByID(id)
	if err != nil {
		slog.Error("error fetching updated movie: ", utils.Err(err))
		return nil, err
	}

	updateResponse := &domain.UpdateMovieResponse{
		ID:           updatedMovie.ID,
		Cover:        updatedMovie.Cover,
		Name:         updatedMovie.Name,
		OriginalName: updatedMovie.OriginalName,
		Description:  updatedMovie.Description,
		Duration:     updatedMovie.Duration,
		ReleaseDate:  updatedMovie.ReleaseDate,
		Age:          updatedMovie.Age,
		Categories:   updatedMovie.Categories,
		Tags:         updatedMovie.Tags,
		Media:        updatedMovie.Media,
	}

	return updateResponse, nil
}

func (r *MongoDBMovieRepository) DeleteMovie(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}

	result, err := r.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		slog.Error("Error deleting movie: ", utils.Err(err))
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("movie not found")
	}

	return nil
}

func (r *MongoDBMovieRepository) SearchMovies(query string, page int, pageSize int) ([]*domain.GetMovieResponse, error) {
	offset := (page - 1) * pageSize

	options := options.Find().SetSkip(int64(offset)).SetLimit(int64(pageSize))

	filter := bson.M{
		"$or": []interface{}{
			bson.M{"name": bson.M{"$regex": query, "$options": "i"}},
			bson.M{"originalName": bson.M{"$regex": query, "$options": "i"}},
		},
	}

	cursor, err := r.collection.Find(context.Background(), filter, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var movies []*domain.GetMovieResponse

	for cursor.Next(context.Background()) {
		var movie domain.GetMovieResponse
		if err := cursor.Decode(&movie); err != nil {
			return nil, err
		}
		movies = append(movies, &movie)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return movies, nil
}

func (r *MongoDBMovieRepository) FilterMoviesByTags(tags []string, page int, pageSize int) ([]*domain.GetMovieResponse, error) {
	offset := (page - 1) * pageSize

	var tagConditions []bson.M
	for _, tag := range tags {
		tagConditions = append(tagConditions, bson.M{"tags": tag})
	}

	filter := bson.M{"$and": tagConditions}

	options := options.Find().SetSkip(int64(offset)).SetLimit(int64(pageSize))

	cursor, err := r.collection.Find(context.Background(), filter, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var movies []*domain.GetMovieResponse
	for cursor.Next(context.Background()) {
		var movie domain.GetMovieResponse
		if err := cursor.Decode(&movie); err != nil {
			return nil, err
		}
		movies = append(movies, &movie)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return movies, nil
}
