package usecase

import (
	"context"
	"mime/multipart"

	"bidkan-bucket/internal/app/domain/entity"
)

// BucketUsecase defines the business logic for file operations
type BucketUsecase interface {
	UploadFile(ctx context.Context, file *multipart.FileHeader) (*entity.FileResponse, error)
	DeleteFile(ctx context.Context, filename string) error
}
