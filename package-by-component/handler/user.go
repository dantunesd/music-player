package handler

import (
	component "music-player/package-by-component/internal/user"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type user struct {
	service component.Component
}

func NewUser(service component.Component) *user {
	return &user{
		service: service,
	}
}

func (h *user) Create(c *fiber.Ctx) error {
	type Request struct {
		Name string `json:"name" validate:"required"`
	}

	var request Request
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := validator.New().Struct(&request); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	response, err := h.service.Create(request.Name)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.Status(201).JSON(response)
}

func (h *user) GetAll(c *fiber.Ctx) error {
	response, err := h.service.GetAll()
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.Status(200).JSON(response)
}

func (h *user) Get(c *fiber.Ctx) error {
	id := c.Params("id")

	response, err := h.service.Get(id)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	return c.Status(200).JSON(response)
}
