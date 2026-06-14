package delete

import (
	"bidkan-bucket/internal/app/domain/usecase"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service usecase.DeleteBucketService
}

func NewHandler(service usecase.DeleteBucketService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) DeleteFile(c *fiber.Ctx) error {
	filename := c.Params("filename")
	if filename == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Filename is required",
		})
	}

	err := h.service.DeleteFile(c.Context(), filename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "File deleted successfully",
	})
}
