package main

import (
	"log"

	"bidkan-bucket/internal/app/delivery/http"
	"bidkan-bucket/internal/app/delivery/http/handler"
	"bidkan-bucket/internal/app/repository/minio"
	"bidkan-bucket/internal/app/usecase/bucket"
	"bidkan-bucket/internal/pkg/config"
	pkgminio "bidkan-bucket/internal/pkg/minio"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// 1. Load Configuration
	cfg := config.LoadConfig()

	// 2. Initialize MinIO Client
	minioClient := pkgminio.SetupMinioClient(cfg)

	// 3. Initialize Repository layer
	bucketRepo := minio.NewBucketRepository(minioClient)

	// 4. Initialize Usecase layer
	bucketUsecase := bucket.NewBucketUsecase(bucketRepo, cfg)

	// 5. Initialize Delivery layer (HTTP Handlers)
	bucketHandler := handler.NewBucketHandler(bucketUsecase)

	// 6. Setup Fiber App
	app := fiber.New(fiber.Config{
		BodyLimit: 20 * 1024 * 1024, // Limit file upload size to 20MB
	})

	// Middlewares
	app.Use(cors.New())
	app.Use(logger.New())

	// Health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok", "service": "bidkan-bucket"})
	})

	// 7. Setup Routes
	http.SetupRouter(app, bucketHandler)

	// 8. Start Server
	log.Printf("Server starting on port %s", cfg.ServerPort)
	if err := app.Listen(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
