package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	config "github.com/thnkrn/comet/puller/pkg/config"
	usecase "github.com/thnkrn/comet/puller/pkg/usecase"
)

type TaskHandler struct {
	taskUsecase usecase.TaskUsecase
	cfg         config.Config
}

func NewTaskHandler(usecase usecase.TaskUsecase) *TaskHandler {
	return &TaskHandler{
		taskUsecase: usecase,
	}
}

func (t *TaskHandler) ManualIngest(c *gin.Context) {

	var (
		uri   TaskIngestURI
		query TaskIngestQuery
	)

	if err := c.ShouldBindUri(&uri); err != nil {
		c.Error(err)
		return
	}

	if err := c.ShouldBindQuery(&query); err != nil {
		c.Error(err)
		return
	}

	err := t.taskUsecase.Perform(c.Request.Context(), uri.DB, query.IgnoreLastIngest)

	if err != nil {
		c.Error(err)
	} else {
		c.Status(http.StatusNoContent)
	}
}
