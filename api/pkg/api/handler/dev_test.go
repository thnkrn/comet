package handler_test

import (
	"bytes"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	handler "github.com/thnkrn/comet/api/pkg/api/handler"
	domain "github.com/thnkrn/comet/api/pkg/domain"
	musecase "github.com/thnkrn/comet/api/pkg/mocks/usecase"
	test "github.com/thnkrn/comet/api/pkg/utils/test"
)

type devDependencies struct {
	mockDevUsecase *musecase.DevUsecase
}

func createDevHandler() (*handler.DevHandler, *devDependencies) {
	mockDevUsecase := new(musecase.DevUsecase)
	handler := handler.NewDevHandler(mockDevUsecase)

	return handler, &devDependencies{mockDevUsecase}
}

func TestDevAddValueToSSTFile(t *testing.T) {
	urlPattern := "/dev/:fileName/sst/:key"

	t.Run("It should return status 200 if add value to SST file successfully", func(t *testing.T) {
		mockResult := []byte("value")

		handlerRequest := test.HTTPRequest{
			Method: "PUT",
			Path:   "/dev/fileName/sst/key",
		}

		request, err := http.NewRequest(handlerRequest.Method, handlerRequest.Path, bytes.NewReader(mockResult))
		assert.NoError(t, err)

		handler, deps := createDevHandler()
		deps.mockDevUsecase.On("AddValueToSSTFile", mock.Anything, "fileName", "key", mockResult).Return(mockResult, nil).Once()

		response := test.RequestHandler(urlPattern, request, handler.AddValueToSSTFile)

		deps.mockDevUsecase.AssertExpectations(t)
		assert.Equal(t, 200, response.StatusCode)
	})

	t.Run("It should return status 500 if dev usecase errors", func(t *testing.T) {
		mockResult := []byte("value")

		handlerRequest := test.HTTPRequest{
			Method: "PUT",
			Path:   "/dev/fileName/sst/key",
		}

		request, err := http.NewRequest(handlerRequest.Method, handlerRequest.Path, bytes.NewReader(mockResult))
		assert.NoError(t, err)

		handler, deps := createDevHandler()
		deps.mockDevUsecase.On("AddValueToSSTFile", mock.Anything, "fileName", "key", mockResult).Return(nil, errors.New("Unexpected Error")).Once()

		response := test.RequestHandler(urlPattern, request, handler.AddValueToSSTFile)

		deps.mockDevUsecase.AssertExpectations(t)
		assert.Equal(t, 500, response.StatusCode)
	})
}

func TestDevPullFile(t *testing.T) {
	urlPattern := "/dev/sst/:fileName/ingest/:source/:ingestFolder"

	t.Run("It should return status 204 if pull successfully", func(t *testing.T) {
		handlerRequest := test.HTTPRequest{
			Method: "PUT",
			Path:   "/dev/sst/fileName/ingest/source/ingestFolder",
		}

		request, err := http.NewRequest(handlerRequest.Method, handlerRequest.Path, nil)
		assert.NoError(t, err)

		handler, deps := createDevHandler()
		deps.mockDevUsecase.On("PullFile", mock.Anything, "fileName", "source", "ingestFolder").Return(nil).Once()

		response := test.RequestHandler(urlPattern, request, handler.PullFile)

		deps.mockDevUsecase.AssertExpectations(t)
		assert.Equal(t, 204, response.StatusCode)
	})

	t.Run("It should return status 500 if dev usercase errors", func(t *testing.T) {
		handlerRequest := test.HTTPRequest{
			Method: "PUT",
			Path:   "/dev/sst/fileName/ingest/source/ingestFolder",
		}

		request, err := http.NewRequest(handlerRequest.Method, handlerRequest.Path, nil)
		assert.NoError(t, err)

		handler, deps := createDevHandler()
		deps.mockDevUsecase.On("PullFile", mock.Anything, "fileName", "source", "ingestFolder").Return(errors.New("Unexpected Error")).Once()

		response := test.RequestHandler(urlPattern, request, handler.PullFile)

		deps.mockDevUsecase.AssertExpectations(t)
		assert.Equal(t, 500, response.StatusCode)
	})
}

func TestDevListDB(t *testing.T) {
	urlPattern := "/dev/db"

	t.Run("It should return status 200 if get DB list successfully", func(t *testing.T) {
		mockResponse := []domain.DB{
			{
				Name: "db1",
				Mode: "read",
			},
			{
				Name: "db1",
				Mode: "read-write",
			},
		}

		handlerRequest := test.HTTPRequest{
			Method: "GET",
			Path:   "/dev/db",
		}

		request, err := http.NewRequest(handlerRequest.Method, handlerRequest.Path, nil)
		assert.NoError(t, err)

		handler, deps := createDevHandler()
		deps.mockDevUsecase.On("ListDB", mock.Anything).Return(mockResponse).Once()

		response := test.RequestHandler(urlPattern, request, handler.ListDB)

		deps.mockDevUsecase.AssertExpectations(t)
		assert.Equal(t, 200, response.StatusCode)
	})
}
