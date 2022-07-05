package main

import (
	"context"
	"log"
	"music-player/package-by-layer/infrastructure"
	"music-player/package-by-layer/internal/handler"
	"music-player/package-by-layer/internal/repository"
	"music-player/package-by-layer/internal/service"

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
	songClient := infrastructure.NewMongoDBAdapter(mongoClient, "music-player", "song")
	playlistClient := infrastructure.NewMongoDBAdapter(mongoClient, "music-player", "playlist")
	userClient := infrastructure.NewMongoDBAdapter(mongoClient, "music-player", "user")

	songRepository := repository.NewSong(songClient)
	playlistRepository := repository.NewPlaylist(playlistClient)
	userRepository := repository.NewUser(userClient)

	songService := service.NewSong(songRepository)
	userService := service.NewUser(userRepository)
	playlistService := service.NewPlaylist(playlistRepository)

	songHandler := handler.NewSong(songService)
	userHandler := handler.NewUser(userService)
	playlistHandler := handler.NewPlaylist(playlistService)

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
