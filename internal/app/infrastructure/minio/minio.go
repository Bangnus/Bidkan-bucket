package minio

import (
	"context"
	"log"

	"bidkan-bucket/internal/config"

	minioClient "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// SetupMinioClient initializes and returns a MinIO client
func SetupMinioClient(cfg *config.Config) *minioClient.Client {
	client, err := minioClient.New(cfg.MinioEndpoint, &minioClient.Options{
		Creds:  credentials.NewStaticV4(cfg.MinioAccessKey, cfg.MinioSecretKey, ""),
		Secure: cfg.MinioUseSSL,
	})
	if err != nil {
		log.Fatalf("Failed to initialize MinIO client: %v", err)
	}

	// Verify connection
	_, err = client.ListBuckets(context.Background())
	if err != nil {
		log.Fatalf("Failed to connect to MinIO server: %v", err)
	}

	log.Println("✅ Successfully connected to MinIO")
	return client
}
