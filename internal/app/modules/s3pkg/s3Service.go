package s3pkg

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/widcha/openidea-marketplace/internal/pkg"
)

type S3Service struct {
	storage *pkg.Storage
}

func NewS3Service(storage *pkg.Storage) *S3Service {
	return &S3Service{
		storage: storage,
	}
}

func (s *S3Service) UploadFile(filePath string, fileName string) {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	err, _ := storage.S3Uploader.Upload(&s3manager.UploadInput{Bucket: aws.String(storage.S3Config.)})
}
