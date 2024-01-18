package internal

import (
	"main/model"

	"github.com/gofiber/fiber/v2"
	store "main/persistent"
)

type Service struct {
	store store.Store
}

type Store interface {
	GetUsers() (*model.User, error)
}

func NewService(store Store) *Service {
	return &Service{store: store}
}

func (s *Service) GetUsers() (*model.User, error) {
	user, err := s.store.GetUsers()
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	return user, nil
}
