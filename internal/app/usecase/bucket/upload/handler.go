package upload

import (
	"bidkan-bucket/internal/app/domain/usecase"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service usecase.UploadBucketService
}

func NewHandler(service usecase.UploadBucketService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) UploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "File is required",
		})
	}

	res, err := h.service.UploadFile(c.Context(), file)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "File uploaded successfully",
		"data":    res,
	})
}
