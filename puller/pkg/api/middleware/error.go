package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	handlerError "github.com/thnkrn/comet/puller/pkg/api/handler/error"
	log "github.com/thnkrn/comet/puller/pkg/driver/log"
	usecaseError "github.com/thnkrn/comet/puller/pkg/usecase/error"
)

type ErrorHandler struct {
	log log.Logger
}

func NewErrorHandler(log log.Logger) *ErrorHandler {
	return &ErrorHandler{log}
}

func (e *ErrorHandler) Handler() gin.HandlerFunc {
	return func(gctx *gin.Context) {
		gctx.Next()
		if len(gctx.Errors) > 0 {
			gerr := gctx.Errors[0].Unwrap()
			e.log.Error(gerr.Error())
			switch e := gerr.(type) {
			case *handlerError.ErrorBadRequest:
				gctx.JSON(http.StatusBadRequest, e.Error())
				return

			case *usecaseError.ErrorBusinessException:
				gctx.JSON(http.StatusBadRequest, e.Error())
				return

			default:
				gctx.JSON(http.StatusInternalServerError, e.Error())
				return
			}
		}
	}
}
