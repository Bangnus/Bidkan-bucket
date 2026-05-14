package bucket

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"bidkan-bucket/internal/app/domain/entity"
	"bidkan-bucket/internal/app/domain/repository"
	"bidkan-bucket/internal/app/domain/usecase"
	"bidkan-bucket/internal/pkg/config"

	"github.com/google/uuid"
)

type bucketUsecase struct {
	bucketRepo repository.BucketRepository
	cfg        *config.Config
}

// NewBucketUsecase creates a new instance of BucketUsecase
func NewBucketUsecase(bucketRepo repository.BucketRepository, cfg *config.Config) usecase.BucketUsecase {
	return &bucketUsecase{
		bucketRepo: bucketRepo,
		cfg:        cfg,
	}
}

func (u *bucketUsecase) UploadFile(ctx context.Context, fileHeader *multipart.FileHeader) (*entity.FileResponse, error) {
	// Open the file
	file, err := fileHeader.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Generate a unique filename using timestamp and UUID
	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	filename := fmt.Sprintf("%d-%s%s", time.Now().Unix(), uuid.New().String(), ext)

	// Determine content type
	contentType := fileHeader.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	// Upload using repository
	url, err := u.bucketRepo.UploadFile(ctx, u.cfg.MinioBucketName, filename, file, fileHeader.Size, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to upload file: %w", err)
	}

	return &entity.FileResponse{
		URL:      url,
		Filename: filename,
	}, nil
}

func (u *bucketUsecase) DeleteFile(ctx context.Context, filename string) error {
	return u.bucketRepo.DeleteFile(ctx, u.cfg.MinioBucketName, filename)
}
