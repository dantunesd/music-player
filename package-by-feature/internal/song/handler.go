package song

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type service interface {
	Create(name, artistName, albumName string, number int) (*song, error)
	Get(id string) (*song, error)
	GetAll() ([]*song, error)
}

type handler struct {
	service service
}

func NewHandler(service service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) Create(c *fiber.Ctx) error {
	type Request struct {
		Name       string `json:"name" validate:"required"`
		ArtistName string `json:"artist_name" validate:"required"`
		AlbumName  string `json:"album_name" validate:"required"`
		Number     int    `json:"number" validate:"required"`
	}

	var request Request
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := validator.New().Struct(&request); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	response, err := h.service.Create(request.Name, request.ArtistName, request.AlbumName, request.Number)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.Status(201).JSON(response)
}

func (h *handler) GetAll(c *fiber.Ctx) error {
	response, err := h.service.GetAll()
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.Status(200).JSON(response)
}

func (h *handler) Get(c *fiber.Ctx) error {
	id := c.Params("id")

	response, err := h.service.Get(id)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.Status(200).JSON(response)
}
