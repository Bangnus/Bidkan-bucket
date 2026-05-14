package minio

import (
	"context"
	"log"

	"bidkan-bucket/internal/pkg/config"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// SetupMinioClient initializes and returns a MinIO client
func SetupMinioClient(cfg *config.Config) *minio.Client {
	minioClient, err := minio.New(cfg.MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.MinioAccessKey, cfg.MinioSecretKey, ""),
		Secure: cfg.MinioUseSSL,
	})
	if err != nil {
		log.Fatalf("Failed to initialize MinIO client: %v", err)
	}

	// Verify connection
	_, err = minioClient.ListBuckets(context.Background())
	if err != nil {
		log.Fatalf("Failed to connect to MinIO server: %v", err)
	}

	log.Println("✅ Successfully connected to MinIO")
	return minioClient
}
