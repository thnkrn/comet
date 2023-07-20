package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	usecase "github.com/thnkrn/comet/api/pkg/usecase"
)

type DevHandler struct {
	devUsecase usecase.DevUsecase
}

func NewDevHandler(usecase usecase.DevUsecase) *DevHandler {
	return &DevHandler{
		devUsecase: usecase,
	}
}

func (h *DevHandler) AddValueToSSTFile(c *fiber.Ctx) error {
	fileName := c.Params("fileName")
	key := c.Params("key")
	value := c.Body()

	result, err := h.devUsecase.AddValueToSSTFile(c.Context(), fileName, key, value)

	if err != nil {
		return err
	} else {
		c.Set("Content-Type", c.Get("Content-Type"))
		return c.Send(result)
	}
}

func (h *DevHandler) PullFile(c *fiber.Ctx) error {
	fileName := c.Params("fileName")
	source := c.Params("source")
	ingestFolder := c.Params("ingestFolder")

	err := h.devUsecase.PullFile(c.Context(), fileName, source, ingestFolder)

	if err != nil {
		return err
	} else {
		return c.SendStatus(http.StatusNoContent)
	}
}

func (h *DevHandler) ListDB(c *fiber.Ctx) error {
	result := h.devUsecase.ListDB(c.Context())

	response := make([]DBListReponse, len(result))

	for i, v := range result {
		response[i] = DBListReponse{
			DBName:   v.Name,
			OpenMode: v.Mode,
		}
	}

	return c.JSON(response)
}
