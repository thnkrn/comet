package handler_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	handler "github.com/thnkrn/comet/api/pkg/api/handler"
	musecase "github.com/thnkrn/comet/api/pkg/mocks/usecase"
	test "github.com/thnkrn/comet/api/pkg/utils/test"
)

type adminDependencies struct {
	mockAdminUsecase *musecase.AdminUsecase
}

func createAdminHandler() (*handler.AdminHandler, *adminDependencies) {
	mockAdminUsecase := new(musecase.AdminUsecase)
	handler := handler.NewAdminHandler(mockAdminUsecase)

	return handler, &adminDependencies{mockAdminUsecase}
}

func TestAdminCatchUpWithPrimary(t *testing.T) {
	urlPattern := "/admin/databases/:db/catch-up-with-primary"

	t.Run("It should return status 204 if catch up with primary successfully", func(t *testing.T) {
		handlerRequest := test.HTTPRequest{
			Method: "POST",
			Path:   "/admin/databases/mp6/catch-up-with-primary",
		}

		request, err := http.NewRequest(handlerRequest.Method, handlerRequest.Path, nil)
		assert.NoError(t, err)

		handler, deps := createAdminHandler()
		deps.mockAdminUsecase.On("CatchUpWithPrimary", mock.Anything, "mp6").Return(nil).Once()

		response := test.RequestHandler(urlPattern, request, handler.CatchUpWithPrimary)

		deps.mockAdminUsecase.AssertExpectations(t)
		assert.Equal(t, 204, response.StatusCode)
	})

	t.Run("It should return status 204 if admin usecase errors", func(t *testing.T) {
		handlerRequest := test.HTTPRequest{
			Method: "POST",
			Path:   "/admin/databases/mp6/catch-up-with-primary",
		}

		request, err := http.NewRequest(handlerRequest.Method, handlerRequest.Path, nil)
		assert.NoError(t, err)

		handler, deps := createAdminHandler()
		deps.mockAdminUsecase.On("CatchUpWithPrimary", mock.Anything, "mp6").Return(errors.New("Unexpected Error")).Once()

		response := test.RequestHandler(urlPattern, request, handler.CatchUpWithPrimary)

		deps.mockAdminUsecase.AssertExpectations(t)
		assert.Equal(t, 500, response.StatusCode)
	})
}

func TestAdminGetDBProperty(t *testing.T) {
	urlPattern := "/admin/databases/:db/properties/:property"

	t.Run("It should return status 200 if get property successfully", func(t *testing.T) {
		handlerRequest := test.HTTPRequest{
			Method: "GET",
			Path:   "/admin/databases/mp6/properties/property",
		}

		request, err := http.NewRequest(handlerRequest.Method, handlerRequest.Path, nil)
		assert.NoError(t, err)

		handler, deps := createAdminHandler()
		deps.mockAdminUsecase.On("GetDBProperty", mock.Anything, "mp6", "property").Return("value", nil).Once()

		response := test.RequestHandler(urlPattern, request, handler.GetDBProperty)

		deps.mockAdminUsecase.AssertExpectations(t)
		assert.Equal(t, 200, response.StatusCode)
	})

	t.Run("It should return status 500 if admin usecase errors", func(t *testing.T) {
		handlerRequest := test.HTTPRequest{
			Method: "GET",
			Path:   "/admin/databases/mp6/properties/property",
		}

		request, err := http.NewRequest(handlerRequest.Method, handlerRequest.Path, nil)
		assert.NoError(t, err)

		handler, deps := createAdminHandler()
		deps.mockAdminUsecase.On("GetDBProperty", mock.Anything, "mp6", "property").Return("", errors.New("Unexpected Error")).Once()

		response := test.RequestHandler(urlPattern, request, handler.GetDBProperty)

		deps.mockAdminUsecase.AssertExpectations(t)
		assert.Equal(t, 500, response.StatusCode)
	})
}

func TestAdminCreateCheckPoint(t *testing.T) {
	urlPattern := "/admin/databases/:db/checkpoint/:directory"

	t.Run("It should return status 204 if create directory successfully", func(t *testing.T) {
		handlerRequest := test.HTTPRequest{
			Method: "PUT",
			Path:   "/admin/databases/mp6/checkpoint/directory",
		}

		request, err := http.NewRequest(handlerRequest.Method, handlerRequest.Path, nil)
		assert.NoError(t, err)

		handler, deps := createAdminHandler()
		deps.mockAdminUsecase.On("CreateCheckPoint", mock.Anything, "mp6", "directory").Return(nil).Once()

		response := test.RequestHandler(urlPattern, request, handler.CreateCheckPoint)

		deps.mockAdminUsecase.AssertExpectations(t)
		assert.Equal(t, 204, response.StatusCode)
	})

	t.Run("It should return status 500 if admin usecase errors", func(t *testing.T) {
		handlerRequest := test.HTTPRequest{
			Method: "PUT",
			Path:   "/admin/databases/mp6/checkpoint/directory",
		}

		request, err := http.NewRequest(handlerRequest.Method, handlerRequest.Path, nil)
		assert.NoError(t, err)

		handler, deps := createAdminHandler()
		deps.mockAdminUsecase.On("CreateCheckPoint", mock.Anything, "mp6", "directory").Return(errors.New("Unexpected Error")).Once()

		response := test.RequestHandler(urlPattern, request, handler.CreateCheckPoint)

		deps.mockAdminUsecase.AssertExpectations(t)
		assert.Equal(t, 500, response.StatusCode)
	})
}

func TestAdminIngest(t *testing.T) {
	urlPattern := "/admin/databases/:db/ingests/:directory"

	t.Run("It should return status 204 if ingest successfully", func(t *testing.T) {
		handlerRequest := test.HTTPRequest{
			Method: "PUT",
			Path:   "/admin/databases/mp6/ingests/directory",
		}

		request, err := http.NewRequest(handlerRequest.Method, handlerRequest.Path, nil)
		assert.NoError(t, err)

		handler, deps := createAdminHandler()
		deps.mockAdminUsecase.On("Ingest", mock.Anything, "mp6", "directory").Return(nil).Once()

		response := test.RequestHandler(urlPattern, request, handler.Ingest)

		deps.mockAdminUsecase.AssertExpectations(t)
		assert.Equal(t, 204, response.StatusCode)
	})

	t.Run("It should return status 500 if admin usecase errors", func(t *testing.T) {
		handlerRequest := test.HTTPRequest{
			Method: "PUT",
			Path:   "/admin/databases/mp6/ingests/directory",
		}

		request, err := http.NewRequest(handlerRequest.Method, handlerRequest.Path, nil)
		assert.NoError(t, err)

		handler, deps := createAdminHandler()
		deps.mockAdminUsecase.On("Ingest", mock.Anything, "mp6", "directory").Return(errors.New("Unexpected Error")).Once()

		response := test.RequestHandler(urlPattern, request, handler.Ingest)

		deps.mockAdminUsecase.AssertExpectations(t)
		assert.Equal(t, 500, response.StatusCode)
	})
}

func TestAdminGetLastIngest(t *testing.T) {
	urlPattern := "/admin/databases/:db/ingests/last"

	t.Run("It should return status 200 if get property successfully", func(t *testing.T) {
		handlerRequest := test.HTTPRequest{
			Method: "GET",
			Path:   "/admin/databases/mp6/ingests/last",
		}

		request, err := http.NewRequest(handlerRequest.Method, handlerRequest.Path, nil)
		assert.NoError(t, err)

		handler, deps := createAdminHandler()
		deps.mockAdminUsecase.On("GetLastIngest", mock.Anything, "mp6").Return([]byte("value"), nil).Once()

		response := test.RequestHandler(urlPattern, request, handler.GetLastIngest)

		deps.mockAdminUsecase.AssertExpectations(t)
		assert.Equal(t, 200, response.StatusCode)
	})

	t.Run("It should return status 500 if admin usecase errors", func(t *testing.T) {
		handlerRequest := test.HTTPRequest{
			Method: "GET",
			Path:   "/admin/databases/mp6/ingests/last",
		}

		request, err := http.NewRequest(handlerRequest.Method, handlerRequest.Path, nil)
		assert.NoError(t, err)

		handler, deps := createAdminHandler()
		deps.mockAdminUsecase.On("GetLastIngest", mock.Anything, "mp6").Return(nil, errors.New("Unexpected Error")).Once()

		response := test.RequestHandler(urlPattern, request, handler.GetLastIngest)

		deps.mockAdminUsecase.AssertExpectations(t)
		assert.Equal(t, 500, response.StatusCode)
	})
}
