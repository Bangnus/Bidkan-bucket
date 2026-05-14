package minio

import (
	"context"
	"fmt"
	"io"

	"bidkan-bucket/internal/app/domain/repository"

	"github.com/minio/minio-go/v7"
)

type bucketRepository struct {
	client *minio.Client
}

// NewBucketRepository creates a new instance of BucketRepository using MinIO
func NewBucketRepository(client *minio.Client) repository.BucketRepository {
	return &bucketRepository{
		client: client,
	}
}

func (r *bucketRepository) UploadFile(ctx context.Context, bucketName, objectName string, reader io.Reader, objectSize int64, contentType string) (string, error) {
	_, err := r.client.PutObject(ctx, bucketName, objectName, reader, objectSize, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return "", err
	}

	// Assuming the bucket is public, construct the public URL
	endpoint := r.client.EndpointURL().String()
	url := fmt.Sprintf("%s/%s/%s", endpoint, bucketName, objectName)
	return url, nil
}

func (r *bucketRepository) DeleteFile(ctx context.Context, bucketName, objectName string) error {
	opts := minio.RemoveObjectOptions{
		GovernanceBypass: true,
	}
	return r.client.RemoveObject(ctx, bucketName, objectName, opts)
}
