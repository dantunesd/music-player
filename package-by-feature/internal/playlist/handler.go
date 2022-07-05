package playlist

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type service interface {
	Create(userId, name string, songs []string) (*playlist, error)
	Get(id string) (*playlist, error)
	GetAll() ([]*playlist, error)
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
