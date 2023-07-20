package handler_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	handler "github.com/thnkrn/comet/puller/pkg/api/handler"
	musecase "github.com/thnkrn/comet/puller/pkg/mocks/usecase"
	test "github.com/thnkrn/comet/puller/pkg/utils/test"
)

type taskDependencies struct {
	mockTaskUsecase *musecase.TaskUsecase
}

func createTaskHandler() (*handler.TaskHandler, *taskDependencies) {
	mockTaskUsecase := new(musecase.TaskUsecase)
	handler := handler.NewTaskHandler(mockTaskUsecase)

	return handler, &taskDependencies{mockTaskUsecase}
}

func TestManualIngest(t *testing.T) {
	urlPattern := "/manual/:db"

	t.Run("It should return status 204 if get successfully", func(t *testing.T) {
		handlerRequest := test.HTTPRequest{
			Method: "POST",
			Path:   "/manual/db1",
		}

		request, err := http.NewRequest(handlerRequest.Method, handlerRequest.Path, nil)
		assert.NoError(t, err)

		handler, deps := createTaskHandler()
		deps.mockTaskUsecase.On("Perform", context.TODO(), mock.Anything, mock.Anything).Return(nil).Once()

		response, _ := test.RequestHandler(urlPattern, request, handler.ManualIngest)

		assert.Equal(t, 204, response.Code)
		deps.mockTaskUsecase.AssertExpectations(t)

	})

	t.Run("It should return status 500 if failed", func(t *testing.T) {
		handlerRequest := test.HTTPRequest{
			Method: "POST",
			Path:   "/manual/db1",
		}

		request, err := http.NewRequest(handlerRequest.Method, handlerRequest.Path, nil)
		assert.NoError(t, err)

		handler, deps := createTaskHandler()
		deps.mockTaskUsecase.On("Perform", context.TODO(), mock.Anything, mock.Anything).Return(errors.New("couldn't perform manual ingest")).Once()

		response, _ := test.RequestHandler(urlPattern, request, handler.ManualIngest)

		assert.Equal(t, 500, response.Code)
		deps.mockTaskUsecase.AssertExpectations(t)

	})
}
