package internal

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/gofiber/fiber/v2"
)

func TestHandler_GetUsers(t *testing.T) {
	t.Run("should return users properly", func(t *testing.T) {
		mockService, mockServiceController := createMockService(t)
		defer mockServiceController.Finish()

		app := createTestApp()

		mockService.
			EXPECT().
			GetUsers().
			Return(testUser, nil).
			Times(1)

		handler := NewHandler(mockService)
		handler.RegisterRoutes(app)

		req := NewUsersRequest(http.MethodGet, "/users")
		res, err := app.Test(req)
		defer res.Body.Close()

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode)
	})

	t.Run("should return error", func(t *testing.T) {
		mockService, mockServiceController := createMockService(t)
		defer mockServiceController.Finish()

		app := createTestApp()

		mockService.
			EXPECT().
			GetUsers().
			Return(nil, fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")).
			Times(1)

		handler := NewHandler(mockService)
		handler.RegisterRoutes(app)

		req := NewUsersRequest(http.MethodGet, "/users")
		res, err := app.Test(req)
		defer res.Body.Close()

		assert.Nil(t, err)
		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	})

	t.Run("should return error", func(t *testing.T) {
		mockService, mockServiceController := createMockService(t)
		defer mockServiceController.Finish()

		app := createTestApp()

		mockService.
			EXPECT().
			GetUsers().
			Return(nil, nil).
			Times(1)

		handler := NewHandler(mockService)
		handler.RegisterRoutes(app)

		req := NewUsersRequest(http.MethodGet, "/users")
		res, err := app.Test(req)
		defer res.Body.Close()

		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, res.StatusCode)
	})
}

func createMockService(t *testing.T) (*Mockactions, *gomock.Controller) {
	t.Helper()

	mockServiceController := gomock.NewController(t)
	mockService := NewMockactions(mockServiceController)

	return mockService, mockServiceController
}

func createTestApp() *fiber.App {
	return fiber.New(fiber.Config{})
}

func NewUsersRequest(method, url string) *http.Request {
	return httptest.NewRequest(method, url, http.NoBody)
}
