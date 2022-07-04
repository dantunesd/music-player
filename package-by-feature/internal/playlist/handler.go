package playlist

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Service interface {
	Create(userId, name string, songs []string) (*Playlist, error)
	Get(id string) (*Playlist, error)
	GetAll() ([]*Playlist, error)
}

type Handler struct {
	Service Service
}

func (h *Handler) Create(c *fiber.Ctx) error {
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

func (h *Handler) GetAll(c *fiber.Ctx) error {
	response, err := h.Service.GetAll()
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.Status(200).JSON(response)
}

func (h *Handler) Get(c *fiber.Ctx) error {
	id := c.Params("id")

	response, err := h.Service.Get(id)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.Status(200).JSON(response)
}
