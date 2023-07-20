package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	usecase "github.com/thnkrn/comet/api/pkg/usecase"
)

type AdminHandler struct {
	adminUsecase usecase.AdminUsecase
}

func NewAdminHandler(usecase usecase.AdminUsecase) *AdminHandler {
	return &AdminHandler{
		adminUsecase: usecase,
	}
}

func (h *AdminHandler) CatchUpWithPrimary(c *fiber.Ctx) error {
	databaseName := c.Params("db")

	err := h.adminUsecase.CatchUpWithPrimary(c.Context(), databaseName)

	if err != nil {
		return err
	} else {
		return c.SendStatus(http.StatusNoContent)
	}
}

func (h *AdminHandler) GetDBProperty(c *fiber.Ctx) error {
	databaseName := c.Params("db")
	property := c.Params("property")

	result, err := h.adminUsecase.GetDBProperty(c.Context(), databaseName, property)

	if err != nil {
		return err
	} else {
		return c.SendString(result)
	}
}

func (h *AdminHandler) CreateCheckPoint(c *fiber.Ctx) error {
	databaseName := c.Params("db")
	directory := c.Params("directory")

	err := h.adminUsecase.CreateCheckPoint(c.Context(), databaseName, directory)

	if err != nil {
		return err
	} else {
		return c.SendStatus(http.StatusNoContent)
	}
}

func (h *AdminHandler) Ingest(c *fiber.Ctx) error {
	databaseName := c.Params("db")
	directory := c.Params("directory")

	err := h.adminUsecase.Ingest(c.Context(), databaseName, directory)

	if err != nil {
		return err
	} else {
		return c.SendStatus(http.StatusNoContent)
	}
}

func (h *AdminHandler) GetLastIngest(c *fiber.Ctx) error {
	databaseName := c.Params("db")

	result, err := h.adminUsecase.GetLastIngest(c.Context(), databaseName)

	if err != nil {
		return err
	} else {
		return c.SendString(string(result))
	}
}
