package test

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
	middleware "github.com/thnkrn/comet/api/pkg/api/middleware"
	mlog "github.com/thnkrn/comet/api/pkg/mocks/driver/log"
)

type HTTPRequest struct {
	Method      string
	Path        string
	JsonParam   string
	Description string
}

type errortDependencies struct {
	mockLogger *mlog.Logger
}

func createErrorHandler() (*middleware.ErrorHandler, *errortDependencies) {
	mockLogger := new(mlog.Logger)
	errorHandler := middleware.NewErrorHandler(mockLogger)

	return errorHandler, &errortDependencies{mockLogger}
}

func RequestHandler(urlPattern string, request *http.Request, handler fiber.Handler) *http.Response {
	errorHandler, deps := createErrorHandler()
	deps.mockLogger.On("Error", mock.Anything).Once()

	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler.FiberErrorHandler(),
	})

	switch request.Method {
	case http.MethodPost:
		app.Post(urlPattern, handler)
	case http.MethodPut:
		app.Put(urlPattern, handler)
	case http.MethodDelete:
		app.Delete(urlPattern, handler)
	default:
		app.Get(urlPattern, handler)
	}

	response, _ := app.Test(request, -1)

	return response
}
