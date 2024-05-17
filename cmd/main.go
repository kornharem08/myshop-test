package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kornharem08/myshop-test/handler/movies"
	"github.com/kornharem08/myshop-test/pkg/mongo"
)

func main() {

	client, err := mongo.ConnectMongoDB("mongodb://localhost:27017/")
	if err != nil {
		log.Fatalf("Failed to create repository: %v", err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()
	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")                           // Allow requests from any origin
		c.Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE")        // Allow specified methods
		c.Set("Access-Control-Allow-Headers", "Content-Type,Authorization") // Allow specified headers
		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusOK)
		}
		return c.Next()
	})
	db := client.Database("mflix")
	movieHandler := movies.NewHandler(db)
	// Define routes
	app.Get("/movies", movieHandler.ListMovies)
	app.Get("/movie", movieHandler.GetMovieByID)
	app.Post("/movie", movieHandler.CreateMovie)
	app.Delete("/movie", movieHandler.DeleteMovieByID)
	app.Listen(":8000")
}
