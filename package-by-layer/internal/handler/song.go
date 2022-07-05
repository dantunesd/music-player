package handler

import (
	"music-player/package-by-layer/internal/domain"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type songService interface {
	Create(name, artistName, albumName string, number int) (*domain.Song, error)
	Get(id string) (*domain.Song, error)
	GetAll() ([]*domain.Song, error)
}

type song struct {
	Service songService
}

func NewSong(service songService) *song {
	return &song{
		Service: service,
	}
}

func (h *song) Create(c *fiber.Ctx) error {
	type Song struct {
		Name       string `json:"name" validate:"required"`
		ArtistName string `json:"artist_name" validate:"required"`
		AlbumName  string `json:"album_name" validate:"required"`
		Number     int    `json:"number" validate:"required"`
	}

	var request Song

	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := validator.New().Struct(&request); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	response, err := h.Service.Create(request.Name, request.ArtistName, request.AlbumName, request.Number)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.Status(201).JSON(response)
}

func (h *song) GetAll(c *fiber.Ctx) error {
	response, err := h.Service.GetAll()
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.Status(200).JSON(response)
}

func (h *song) Get(c *fiber.Ctx) error {
	id := c.Params("id")

	response, err := h.Service.Get(id)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.Status(200).JSON(response)
}
