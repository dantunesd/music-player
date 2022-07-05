package handler

import (
	component "music-player/package-by-component/internal/playlist"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type playlist struct {
	service component.Component
}

func NewPlaylist(service component.Component) *playlist {
	return &playlist{
		service: service,
	}
}

func (h *playlist) Create(c *fiber.Ctx) error {
	type Request struct {
		Name  string   `json:"name" validate:"required"`
		Songs []string `json:"Songs" validate:"required"`
	}

	var request Request
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := validator.New().Struct(&request); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	userId := c.Params("userId")
	response, err := h.service.Create(userId, request.Name, request.Songs)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.Status(201).JSON(response)
}

func (h *playlist) GetAll(c *fiber.Ctx) error {
	response, err := h.service.GetAll()
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.Status(200).JSON(response)
}

func (h *playlist) Get(c *fiber.Ctx) error {
	id := c.Params("id")

	response, err := h.service.Get(id)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.Status(200).JSON(response)
}
