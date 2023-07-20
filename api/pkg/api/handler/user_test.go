package handler_test

import (
	"bytes"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	handler "github.com/thnkrn/comet/api/pkg/api/handler"
	musecase "github.com/thnkrn/comet/api/pkg/mocks/usecase"
	test "github.com/thnkrn/comet/api/pkg/utils/test"
)

type userDependencies struct {
	mockUserUsecase *musecase.UserUsecase
}

func createUserHandler() (*handler.UserHandler, *userDependencies) {
	mockUserUsecase := new(musecase.UserUsecase)
	handler := handler.NewUserHandler(mockUserUsecase)

	return handler, &userDependencies{mockUserUsecase}
}

func TestUserGet(t *testing.T) {
	urlPattern := "/databases/:db/keys/:key"

	t.Run("It should return status 200 if get successfully", func(t *testing.T) {
		handlerRequest := test.HTTPRequest{
			Method: "GET",
			Path:   "/databases/mp6/keys/1",
		}

		request, err := http.NewRequest(handlerRequest.Method, handlerRequest.Path, nil)
		assert.NoError(t, err)

		handler, deps := createUserHandler()
		deps.mockUserUsecase.On("Get", mock.Anything, "mp6", "1").Return([]byte("value"), nil).Once()

		response := test.RequestHandler(urlPattern, request, handler.Get)

		deps.mockUserUsecase.AssertExpectations(t)
		assert.Equal(t, 200, response.StatusCode)
	})

	t.Run("It should return status 500 if user usecase errors", func(t *testing.T) {
		handlerRequest := test.HTTPRequest{
			Method: "GET",
			Path:   "/databases/mp6/keys/1",
		}

		request, err := http.NewRequest(handlerRequest.Method, handlerRequest.Path, nil)
		assert.NoError(t, err)

		handler, deps := createUserHandler()
		deps.mockUserUsecase.On("Get", mock.Anything, "mp6", "1").Return(nil, errors.New("Unexpected Error")).Once()

		response := test.RequestHandler(urlPattern, request, handler.Get)

		deps.mockUserUsecase.AssertExpectations(t)
		assert.Equal(t, 500, response.StatusCode)
	})
}

func TestUserCreate(t *testing.T) {
	urlPattern := "/databases/:db/keys/:key"

	t.Run("It should return status 200 if create successfully", func(t *testing.T) {
		mockResult := []byte("value")

		handlerRequest := test.HTTPRequest{
			Method: "PUT",
			Path:   "/databases/mp6/keys/1",
		}

		request, err := http.NewRequest(handlerRequest.Method, handlerRequest.Path, bytes.NewReader(mockResult))
		assert.NoError(t, err)

		handler, deps := createUserHandler()
		deps.mockUserUsecase.On("Create", mock.Anything, "mp6", "1", mockResult).Return(mockResult, nil).Once()

		response := test.RequestHandler(urlPattern, request, handler.Create)

		deps.mockUserUsecase.AssertExpectations(t)
		assert.Equal(t, 200, response.StatusCode)
	})

	t.Run("It should return status 500 if user usecase errors", func(t *testing.T) {
		mockResult := []byte("value")

		handlerRequest := test.HTTPRequest{
			Method: "PUT",
			Path:   "/databases/mp6/keys/1",
		}

		request, err := http.NewRequest(handlerRequest.Method, handlerRequest.Path, bytes.NewReader(mockResult))
		assert.NoError(t, err)

		handler, deps := createUserHandler()
		deps.mockUserUsecase.On("Create", mock.Anything, "mp6", "1", mockResult).Return(nil, errors.New("Unexpected Error")).Once()

		response := test.RequestHandler(urlPattern, request, handler.Create)

		deps.mockUserUsecase.AssertExpectations(t)
		assert.Equal(t, 500, response.StatusCode)
	})
}

func TestUserDelete(t *testing.T) {
	urlPattern := "/databases/:db/keys/:key"

	t.Run("It should return status 204 if delete successfully", func(t *testing.T) {
		handlerRequest := test.HTTPRequest{
			Method: "DELETE",
			Path:   "/databases/mp6/keys/1",
		}

		request, err := http.NewRequest(handlerRequest.Method, handlerRequest.Path, nil)
		assert.NoError(t, err)

		handler, deps := createUserHandler()
		deps.mockUserUsecase.On("Delete", mock.Anything, "mp6", "1").Return(nil).Once()

		response := test.RequestHandler(urlPattern, request, handler.Delete)

		deps.mockUserUsecase.AssertExpectations(t)
		assert.Equal(t, 204, response.StatusCode)
	})

	t.Run("It should return status 500 if user usercase errors", func(t *testing.T) {
		handlerRequest := test.HTTPRequest{
			Method: "DELETE",
			Path:   "/databases/mp6/keys/1",
		}

		request, err := http.NewRequest(handlerRequest.Method, handlerRequest.Path, nil)
		assert.NoError(t, err)

		handler, deps := createUserHandler()
		deps.mockUserUsecase.On("Delete", mock.Anything, "mp6", "1").Return(errors.New("Unexpected Error")).Once()

		response := test.RequestHandler(urlPattern, request, handler.Delete)

		deps.mockUserUsecase.AssertExpectations(t)
		assert.Equal(t, 500, response.StatusCode)
	})
}

func TestUserCount(t *testing.T) {
	urlPattern := "/databases/:db"

	t.Run("It should return status 200 if count successfully", func(t *testing.T) {
		handlerRequest := test.HTTPRequest{
			Method: "GET",
			Path:   "/databases/mp6",
		}

		request, err := http.NewRequest(handlerRequest.Method, handlerRequest.Path, nil)
		assert.NoError(t, err)

		handler, deps := createUserHandler()
		deps.mockUserUsecase.On("Count", mock.Anything, "mp6").Return("1", nil).Once()

		response := test.RequestHandler(urlPattern, request, handler.Count)

		deps.mockUserUsecase.AssertExpectations(t)
		assert.Equal(t, 200, response.StatusCode)
	})

	t.Run("It should return status 500 if user usecase errors", func(t *testing.T) {
		handlerRequest := test.HTTPRequest{
			Method: "GET",
			Path:   "/databases/mp6",
		}

		request, err := http.NewRequest(handlerRequest.Method, handlerRequest.Path, nil)
		assert.NoError(t, err)

		handler, deps := createUserHandler()
		deps.mockUserUsecase.On("Count", mock.Anything, "mp6").Return("", errors.New("Unexpected Error")).Once()

		response := test.RequestHandler(urlPattern, request, handler.Count)

		deps.mockUserUsecase.AssertExpectations(t)
		assert.Equal(t, 500, response.StatusCode)
	})
}

func TestUserMultiGet(t *testing.T) {
	urlPattern := "/databases/:db/keys"

	t.Run("It should return status 200 if get successfully", func(t *testing.T) {
		handlerRequest := test.HTTPRequest{
			Method: "GET",
			Path:   "/databases/mp6/keys?keys=1&keys=2",
		}

		request, err := http.NewRequest(handlerRequest.Method, handlerRequest.Path, nil)
		assert.NoError(t, err)

		handler, deps := createUserHandler()
		deps.mockUserUsecase.On("MultiGet", mock.Anything, "mp6", []string{"1", "2"}).Return([]string{"val1", "val2"}, nil).Once()

		response := test.RequestHandler(urlPattern, request, handler.MultiGet)

		deps.mockUserUsecase.AssertExpectations(t)
		assert.Equal(t, 200, response.StatusCode)
	})

	t.Run("It should return status 500 if user usecase errors", func(t *testing.T) {
		handlerRequest := test.HTTPRequest{
			Method: "GET",
			Path:   "/databases/mp6/keys?keys=1&keys=2",
		}

		request, err := http.NewRequest(handlerRequest.Method, handlerRequest.Path, nil)
		assert.NoError(t, err)

		handler, deps := createUserHandler()
		deps.mockUserUsecase.On("MultiGet", mock.Anything, "mp6", []string{"1", "2"}).Return(nil, errors.New("Unexpected Error")).Once()

		response := test.RequestHandler(urlPattern, request, handler.MultiGet)

		deps.mockUserUsecase.AssertExpectations(t)
		assert.Equal(t, 500, response.StatusCode)
	})

	t.Run("It should return status 400 if query invalid", func(t *testing.T) {
		handlerRequest := test.HTTPRequest{
			Method: "GET",
			Path:   "/databases/mp6/keys?x=1",
		}

		request, err := http.NewRequest(handlerRequest.Method, handlerRequest.Path, nil)
		assert.NoError(t, err)

		handler, deps := createUserHandler()

		response := test.RequestHandler(urlPattern, request, handler.MultiGet)

		deps.mockUserUsecase.AssertExpectations(t)
		assert.Equal(t, 400, response.StatusCode)
	})
}
