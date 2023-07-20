package test

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	middleware "github.com/thnkrn/comet/puller/pkg/api/middleware"
	mlog "github.com/thnkrn/comet/puller/pkg/mocks/driver/log"
)

type HTTPRequest struct {
	Method      string
	Path        string
	JsonParam   string
	Description string
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type errortDependencies struct {
	mockLogger *mlog.Logger
}

func createErrorHandler() (*middleware.ErrorHandler, *errortDependencies) {
	mockLogger := new(mlog.Logger)
	errorHandler := middleware.NewErrorHandler(mockLogger)

	return errorHandler, &errortDependencies{mockLogger}
}

func RequestHandler(urlPattern string, request *http.Request, handler gin.HandlerFunc) (*httptest.ResponseRecorder, *gin.Context) {
	response := httptest.NewRecorder()
	gctx, engine := gin.CreateTestContext(response)
	gctx.Request = request

	errorHandler, deps := createErrorHandler()
	deps.mockLogger.On("Error", mock.Anything).Once()

	engine.Use(errorHandler.Handler())

	switch request.Method {
	case http.MethodPost:
		engine.POST(urlPattern, handler)
	case http.MethodPut:
		engine.PUT(urlPattern, handler)
	case http.MethodDelete:
		engine.DELETE(urlPattern, handler)
	default:
		engine.GET(urlPattern, handler)
	}

	engine.ServeHTTP(response, request)

	return response, gctx
}
