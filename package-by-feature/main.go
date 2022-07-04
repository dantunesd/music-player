package main

import (
	"context"
	"log"
	"music-player/internal/playlist"
	"music-player/internal/song"
	"music-player/internal/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://root:example@localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	if err := mongoClient.Ping(context.Background(), nil); err != nil {
		log.Fatal(err)
	}

	songRepository := song.RepositoryImpl{MongoClient: mongoClient}
	songService := song.ServiceImpl{Repository: &songRepository}
	songHandler := song.Handler{Service: &songService}

	userRepository := user.RepositoryImpl{MongoClient: mongoClient}
	userService := user.ServiceImpl{Repository: &userRepository}
	userHandler := user.Handler{Service: &userService}

	playlistRepository := playlist.RepositoryImpl{MongoClient: mongoClient}
	playlistService := playlist.ServiceImpl{Repository: &playlistRepository}
	playlistHandler := playlist.Handler{Service: &playlistService}

	app := fiber.New()
	app.Use(recover.New())
	app.Use(cors.New())

	app.Post("/songs/", songHandler.Create)
	app.Get("/songs/", songHandler.GetAll)
	app.Get("/songs/:id", songHandler.Get)

	app.Post("/users/", userHandler.Create)
	app.Get("/users/", userHandler.GetAll)
	app.Get("/users/:id", userHandler.Get)

	app.Post("/users/:userId/playlists", playlistHandler.Create)
	app.Get("/users/:userId/playlists", playlistHandler.GetAll)
	app.Get("/users/:userId/playlists/:id", playlistHandler.Get)

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
}
