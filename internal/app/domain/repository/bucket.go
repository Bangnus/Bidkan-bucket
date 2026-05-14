package repository

import (
	"context"
	"io"
)

// BucketRepository defines the interface for interacting with the storage bucket (e.g., MinIO, S3)
type BucketRepository interface {
	UploadFile(ctx context.Context, bucketName, objectName string, reader io.Reader, objectSize int64, contentType string) (string, error)
	DeleteFile(ctx context.Context, bucketName, objectName string) error
}
