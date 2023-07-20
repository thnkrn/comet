package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	handlerError "github.com/thnkrn/comet/api/pkg/api/handler/error"
	log "github.com/thnkrn/comet/api/pkg/driver/log"
	usecaseError "github.com/thnkrn/comet/api/pkg/usecase/error"
)

type ErrorHandler struct {
	log log.Logger
}

func NewErrorHandler(log log.Logger) *ErrorHandler {
	return &ErrorHandler{log}
}

func (e *ErrorHandler) FiberErrorHandler() func(ctx *fiber.Ctx, err error) error {
	return func(ctx *fiber.Ctx, err error) error {
		ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

		switch e := err.(type) {
		case *handlerError.ErrorBadRequest:
			return ctx.Status(http.StatusBadRequest).SendString(e.Error())

		case *handlerError.ErrorAuthentication:
			return ctx.Status(http.StatusUnauthorized).SendString(e.Error())

		case *usecaseError.ErrorNotFound:
			return ctx.Status(http.StatusNotFound).SendString(e.Error())

		case *usecaseError.ErrorBusinessException:
			return ctx.Status(http.StatusBadRequest).SendString(e.Error())

		default:
			return ctx.Status(http.StatusInternalServerError).SendString(e.Error())
		}
	}
}
