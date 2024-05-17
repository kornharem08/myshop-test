package movies

import (
	"context"
	"errors"
	"fmt"

	"github.com/kornharem08/myshop-test/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IRepository interface {
	FindAll(page, limit int) ([]models.Movie, error)
	FindByID(id string) (models.Movie, error)
	CreateMovie(movie models.Movie) error
	DeleteMovieByID(id string) error
}

type Repository struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) Repository {
	return Repository{
		collection: db.Collection("movies"),
	}
}

func (r Repository) FindAll(page, limit int) ([]models.Movie, error) {
	var movies []models.Movie
	fmt.Println(page, limit)
	// Create a context
	ctx := context.TODO()

	skip := int64((page - 1) * limit)

	// Define options to customize the query
	options := options.Find()
	options.SetLimit(int64(limit))
	options.SetSkip(skip)
	// Execute the query
	cursor, err := r.collection.Find(ctx, bson.D{}, options)
	if err != nil {
		return nil, err
	}

	// Iterate through the cursor and decode each document into a Movie object
	for cursor.Next(ctx) {
		var movie models.Movie
		if err := cursor.Decode(&movie); err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	// Check for any errors during cursor iteration
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// Close the cursor
	cursor.Close(ctx)

	return movies, nil
}

func (r Repository) FindByID(id string) (models.Movie, error) {
	var movie models.Movie

	// Create a context
	ctx := context.TODO()
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Movie{}, err
	}
	// Define filter to find the movie by its ID
	filter := bson.M{"_id": objectId}
	fmt.Println(id)
	// Execute the query
	if err := r.collection.FindOne(ctx, filter).Decode(&movie); err != nil {
		fmt.Println(err)
		// Check if document not found
		if errors.Is(err, mongo.ErrNoDocuments) {
			return models.Movie{}, models.ErrMovieNotFound
		}
		return movie, err
	}

	return movie, nil
}

func (r Repository) CreateMovie(movie models.Movie) error {
	// Create a context
	ctx := context.TODO()

	// Insert the movie document into the collection
	if _, err := r.collection.InsertOne(ctx, movie); err != nil {
		return err
	}

	return nil
}

func (r Repository) DeleteMovieByID(id string) error {
	// Create a context
	ctx := context.TODO()

	// Convert ID string to ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	// Define filter to find the movie by its ID
	filter := bson.M{"_id": objectID}

	// Delete the movie document from the collection
	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	// Check if no document was deleted
	if result.DeletedCount == 0 {
		return models.ErrMovieNotFound
	}

	return nil
}
