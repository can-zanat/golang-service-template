package internal

import (
	"main/model"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service actions
}

type actions interface {
	GetUsers() (model.User, error)
}

func NewHandler(service actions) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(app *fiber.App) {
	app.Get("/users", h.GetUsers)
}

func (h *Handler) GetUsers(ctx *fiber.Ctx) error {
	user, err := h.service.GetUsers()

	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	if user == (model.User{}) {
		return fiber.NewError(fiber.StatusNotFound, "User Not Found")
	}

	return ctx.Status(fiber.StatusOK).JSON(user)
}
