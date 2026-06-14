package usecase

import (
	"context"
	"mime/multipart"

	"bidkan-bucket/internal/app/domain/entity"
)

type UploadBucketService interface {
	UploadFile(ctx context.Context, file *multipart.FileHeader) (*entity.FileResponse, error)
}

type DeleteBucketService interface {
	DeleteFile(ctx context.Context, filename string) error
}
