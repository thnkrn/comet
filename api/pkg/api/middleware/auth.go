package middleware

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
	handlerError "github.com/thnkrn/comet/api/pkg/api/handler/error"
)

const (
	DEV_ROLE    = "dev"
	DEV_TOKEN   = "devtoken"
	ADMIN_ROLE  = "admin"
	ADMIN_TOKEN = "admintoken"
)

var ErrAuthentication = errors.New("unable to successfully authenticate your request")

type Authentication struct {
}

func NewAuthentication() *Authentication {
	return &Authentication{}
}

func validateToken(token, role string) error {
	if (role == DEV_ROLE && token == DEV_TOKEN) || (role == ADMIN_ROLE && token == ADMIN_TOKEN) {
		return nil
	}
	return handlerError.NewErrorAuthentication(ErrAuthentication)
}

func (a *Authentication) Authentication(role string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		s := ctx.Get("Authorization")
		token := strings.TrimPrefix(s, "Bearer ")

		if err := validateToken(token, role); err != nil {
			return err
		}
		return ctx.Next()
	}
}
