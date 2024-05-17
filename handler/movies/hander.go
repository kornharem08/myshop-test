package movies

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/kornharem08/myshop-test/pkg/models"
	"github.com/kornharem08/myshop-test/pkg/movies"
	"go.mongodb.org/mongo-driver/mongo"
)

type IHandler interface {
	ListMovies(c *fiber.Ctx) error
	GetMovieByID(c *fiber.Ctx) error
	DeleteMovieByID(c *fiber.Ctx) error
	CreateMovie(c *fiber.Ctx) error
}

type Handler struct {
	movieService movies.IService
}

func NewHandler(db *mongo.Database) IHandler {
	return &Handler{
		movieService: movies.NewService(db),
	}
}

func (h Handler) ListMovies(c *fiber.Ctx) error {
	pageStr := c.Query("page") // Get the value of the "page" query parameter
	limitStr := c.Query("limit")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid value for 'page'",
		})
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid value for 'limit'",
		})
	}

	responses, err := h.movieService.ListMovies(page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "unexpected error",
		})
	}

	return c.JSON(responses)
}

func (h Handler) GetMovieByID(c *fiber.Ctx) error {

	id := c.Params("id")
	if id == "" {
		id = c.Query("id")
	}

	// Check if movie ID is empty
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing movie ID",
		})
	}
	responses, err := h.movieService.GetMovieByID(id)
	if err != nil {
		// Handle movie not found error
		if errors.Is(err, models.ErrMovieNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Movie not found",
			})
		}
		// Return generic error message for internal server errors
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	return c.JSON(responses)
}

func (h Handler) CreateMovie(c *fiber.Ctx) error {
	// Parse request body into movie struct
	var movie models.Movie
	if err := c.BodyParser(&movie); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Call service method to create movie
	if err := h.movieService.CreateMovie(movie); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create movie",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Movie created successfully",
	})
}

func (h Handler) DeleteMovieByID(c *fiber.Ctx) error {
	// Retrieve movie ID from path parameter or query parameter
	id := c.Params("id")
	if id == "" {
		id = c.Query("id")
	}

	// Check if movie ID is empty
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing movie ID",
		})
	}

	// Call service method to delete movie by ID
	if err := h.movieService.DeleteMovieByID(id); err != nil {
		// Handle movie not found error
		if errors.Is(err, models.ErrMovieNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Movie not found",
			})
		}
		// Return generic error message for internal server errors
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete movie",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Movie deleted successfully",
	})
}
