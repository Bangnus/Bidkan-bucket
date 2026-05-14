package handler

import (
	"bidkan-bucket/internal/app/domain/usecase"

	"github.com/gofiber/fiber/v2"
)

type BucketHandler struct {
	bucketUsecase usecase.BucketUsecase
}

// NewBucketHandler creates a new handler for bucket operations
func NewBucketHandler(bucketUsecase usecase.BucketUsecase) *BucketHandler {
	return &BucketHandler{
		bucketUsecase: bucketUsecase,
	}
}

// UploadFile handles multipart file uploads
func (h *BucketHandler) UploadFile(c *fiber.Ctx) error {
	// "file" is the key we expect in the multipart form
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "File is required",
		})
	}

	res, err := h.bucketUsecase.UploadFile(c.Context(), file)
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

// DeleteFile handles file deletion by filename
func (h *BucketHandler) DeleteFile(c *fiber.Ctx) error {
	filename := c.Params("filename")
	if filename == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Filename is required",
		})
	}

	err := h.bucketUsecase.DeleteFile(c.Context(), filename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "File deleted successfully",
	})
}
