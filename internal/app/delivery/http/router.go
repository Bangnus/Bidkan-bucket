package http

import (
	"bidkan-bucket/internal/app/delivery/http/handler"

	"github.com/gofiber/fiber/v2"
)

// SetupRouter initializes all the routes for the application
func SetupRouter(app *fiber.App, bucketHandler *handler.BucketHandler) {
	api := app.Group("/api")

	// Bucket routes
	api.Post("/upload", bucketHandler.UploadFile)
	api.Delete("/files/:filename", bucketHandler.DeleteFile)
}
