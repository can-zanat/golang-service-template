package internal

import (
	"main/model"

	store "main/persistent"

	"github.com/gofiber/fiber/v2"
)

type Service struct {
	store store.Store
}

type Store interface {
	GetUsers() (model.User, error)
}

func NewService(s Store) *Service {
	return &Service{store: s}
}

func (s *Service) GetUsers() (model.User, error) {
	user, err := s.store.GetUsers()

	if err != nil {
		return model.User{}, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}

	return user, nil
}
