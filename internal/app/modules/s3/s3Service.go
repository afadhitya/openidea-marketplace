package s3

import "github.com/widcha/openidea-marketplace/internal/pkg"

type S3Service struct {
	storage *pkg.Storage
}

func NewS3Service(storage *pkg.Storage) *S3Service {
	return &S3Service{
		storage: storage,
	}
}

func (s *S3Service) UploadFile() {
	// todo
}
