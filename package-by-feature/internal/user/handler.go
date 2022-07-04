package user

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Service interface {
	Create(name string) (*User, error)
	Get(id string) (*User, error)
	GetAll() ([]*User, error)
}

type Handler struct {
	Service Service
}

func (h *Handler) Create(c *fiber.Ctx) error {
	type CreateUserRequest struct {
		Name string `json:"name" validate:"required"`
	}

	var request CreateUserRequest

	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := validator.New().Struct(&request); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	response, err := h.Service.Create(request.Name)
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
