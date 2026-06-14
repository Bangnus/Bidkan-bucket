package initial

import (
	"bidkan-bucket/internal/app/infrastructure/minio"
	"bidkan-bucket/internal/app/infrastructure/repository"
	"bidkan-bucket/internal/app/router"
	"bidkan-bucket/internal/app/usecase/bucket/delete"
	"bidkan-bucket/internal/app/usecase/bucket/upload"
	"bidkan-bucket/internal/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func InitializeApp(cfg *config.Config) *fiber.App {
	// 1. Initialize Infrastructure
	minioClient := minio.SetupMinioClient(cfg)
	bucketRepo := repository.NewBucketRepository(minioClient)

	// 2. Initialize Services
	uploadService := upload.NewService(bucketRepo, cfg)
	deleteService := delete.NewService(bucketRepo, cfg)

	// 3. Initialize Handlers
	uploadHandler := upload.NewHandler(uploadService)
	deleteHandler := delete.NewHandler(deleteService)

	// 4. Setup Fiber App
	app := fiber.New(fiber.Config{
		BodyLimit: 20 * 1024 * 1024, // 20MB
	})

	app.Use(cors.New())
	app.Use(logger.New())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok", "service": "bidkan-bucket (clean-arch)"})
	})

	// 5. Setup Router
	router.SetupBucketRouter(app, uploadHandler, deleteHandler)

	return app
}
