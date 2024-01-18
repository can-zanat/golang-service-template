package internal

import (
	"main/model"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var testUser = &model.User{
	ID:       12345678,
	Username: "Coniboy",
	FullName: "Muharrem Can Zanat",
	Email:    "can.zanat@gmail.com",
}

func TestService_GetUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("should return users properly", func(t *testing.T) {
		mockRepository := NewMockStore(ctrl)

		//TODO: gomock.Any() is not the best practise but it is ok for now
		mockRepository.
			EXPECT().
			GetUsers().
			Return(testUser, nil).
			Times(1)

		service := NewService(mockRepository)

		users, _ := service.GetUsers()
		assert.Equal(t, testUser, users)
	})

	t.Run("return error", func(t *testing.T) {
		mockRepository := NewMockStore(ctrl)

		var err error
		expectedError := fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")

		mockRepository.
			EXPECT().
			GetUsers().
			Return(nil, expectedError).
			Times(1)

		service := NewService(mockRepository)

		_, err = service.GetUsers()
		assert.Equal(t, expectedError, err)
	})
}
