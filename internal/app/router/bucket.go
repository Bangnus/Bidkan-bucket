package router

import (
	"bidkan-bucket/internal/app/usecase/bucket/delete"
	"bidkan-bucket/internal/app/usecase/bucket/upload"

	"github.com/gofiber/fiber/v2"
)

func SetupBucketRouter(app *fiber.App, uploadHandler *upload.Handler, deleteHandler *delete.Handler) {
	api := app.Group("/api")

	// Bucket routes
	api.Post("/upload", uploadHandler.UploadFile)
	api.Delete("/files/:filename", deleteHandler.DeleteFile)
}
