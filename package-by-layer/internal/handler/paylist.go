package handler

import (
	"music-player/package-by-layer/internal/domain"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type PlaylistService interface {
	Create(userId, name string, songs []string) (*domain.Playlist, error)
	Get(id string) (*domain.Playlist, error)
	GetAll() ([]*domain.Playlist, error)
}

type Playlist struct {
	Service PlaylistService
}

func (h *Playlist) Create(c *fiber.Ctx) error {
	userId := c.Params("userId")

	type Playlist struct {
		Name  string   `json:"name" validate:"required"`
		Songs []string `json:"Songs" validate:"required"`
	}

	var request Playlist

	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := validator.New().Struct(&request); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	response, err := h.Service.Create(userId, request.Name, request.Songs)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.Status(201).JSON(response)
}

func (h *Playlist) GetAll(c *fiber.Ctx) error {
	response, err := h.Service.GetAll()
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.Status(200).JSON(response)
}

func (h *Playlist) Get(c *fiber.Ctx) error {
	id := c.Params("id")

	response, err := h.Service.Get(id)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.Status(200).JSON(response)
}
