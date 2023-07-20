package handler

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"

	handlerError "github.com/thnkrn/comet/api/pkg/api/handler/error"
	usecase "github.com/thnkrn/comet/api/pkg/usecase"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(usecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: usecase,
	}
}

func (h *UserHandler) Get(c *fiber.Ctx) error {
	databaseName := c.Params("db")
	key := c.Params("key")

	result, err := h.userUsecase.Get(c.Context(), databaseName, key)

	if err != nil {
		return err
	} else {
		c.Set("Content-Type", c.Get("Accept", "application/octet-stream"))
		return c.Send(result)
	}
}

func (h *UserHandler) Create(c *fiber.Ctx) error {
	databaseName := c.Params("db")
	key := c.Params("key")

	value := c.Body()

	result, err := h.userUsecase.Create(c.Context(), databaseName, key, value)

	if err != nil {
		return err
	} else {
		c.Set("Content-Type", c.Get("Content-Type"))
		return c.Send(result)
	}
}

func (h *UserHandler) Delete(c *fiber.Ctx) error {
	databaseName := c.Params("db")
	key := c.Params("key")

	err := h.userUsecase.Delete(c.Context(), databaseName, key)

	if err != nil {
		return err
	} else {
		return c.SendStatus(http.StatusNoContent)
	}
}

func (h *UserHandler) Count(c *fiber.Ctx) error {
	databaseName := c.Params("db")

	result, err := h.userUsecase.Count(c.Context(), databaseName)

	if err != nil {
		return err
	} else {
		return c.SendString(result)
	}
}

func (h *UserHandler) MultiGet(c *fiber.Ctx) error {
	databaseName := c.Params("db")
	query := new(UserMultiGetQuery)

	if err := c.QueryParser(query); err != nil {
		return err
	}

	if len(query.Keys) <= 0 {
		return handlerError.NewErrorBadRequest(errors.New("query not valid"))
	}

	result, err := h.userUsecase.MultiGet(c.Context(), databaseName, query.Keys)

	if err != nil {
		return err
	} else {
		return c.JSON(result)
	}
}
