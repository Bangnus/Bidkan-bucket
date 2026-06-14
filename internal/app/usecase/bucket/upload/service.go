package upload

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
	"bidkan-bucket/internal/config"

	"github.com/google/uuid"
)

type service struct {
	bucketRepo repository.BucketRepository
	cfg        *config.Config
}

// NewService creates a new upload service
func NewService(bucketRepo repository.BucketRepository, cfg *config.Config) usecase.UploadBucketService {
	return &service{
		bucketRepo: bucketRepo,
		cfg:        cfg,
	}
}

func (s *service) UploadFile(ctx context.Context, fileHeader *multipart.FileHeader) (*entity.FileResponse, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	filename := fmt.Sprintf("%d-%s%s", time.Now().Unix(), uuid.New().String(), ext)

	contentType := fileHeader.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	url, err := s.bucketRepo.UploadFile(ctx, s.cfg.MinioBucketName, filename, file, fileHeader.Size, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to upload file: %w", err)
	}

	return &entity.FileResponse{
		URL:      url,
		Filename: filename,
	}, nil
}
