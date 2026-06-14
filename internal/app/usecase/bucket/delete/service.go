package delete

import (
	"context"

	"bidkan-bucket/internal/app/domain/repository"
	"bidkan-bucket/internal/app/domain/usecase"
	"bidkan-bucket/internal/config"
)

type service struct {
	bucketRepo repository.BucketRepository
	cfg        *config.Config
}

// NewService creates a new delete service
func NewService(bucketRepo repository.BucketRepository, cfg *config.Config) usecase.DeleteBucketService {
	return &service{
		bucketRepo: bucketRepo,
		cfg:        cfg,
	}
}

func (s *service) DeleteFile(ctx context.Context, filename string) error {
	return s.bucketRepo.DeleteFile(ctx, s.cfg.MinioBucketName, filename)
}
