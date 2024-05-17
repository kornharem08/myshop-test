package movies

import (
	"github.com/kornharem08/myshop-test/pkg/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type IService interface {
	ListMovies(page, limit int) ([]models.Movie, error)
	GetMovieByID(id string) (models.Movie, error)
	DeleteMovieByID(id string) error
	CreateMovie(movie models.Movie) error
}

type Service struct {
	Repository IRepository
}

func NewService(db *mongo.Database) Service {
	return Service{
		Repository: NewRepository(db),
	}
}

func (s Service) ListMovies(page, limit int) ([]models.Movie, error) {
	return s.Repository.FindAll(page, limit)
}

func (s Service) GetMovieByID(id string) (models.Movie, error) {
	return s.Repository.FindByID(id)
}

func (s Service) CreateMovie(movie models.Movie) error {
	// Add any business logic/validation here before creating the movie
	return s.Repository.CreateMovie(movie)
}

func (s Service) DeleteMovieByID(id string) error {
	// Add any business logic/validation here before deleting the movie
	return s.Repository.DeleteMovieByID(id)
}
