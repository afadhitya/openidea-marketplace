package pkg

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/widcha/openidea-marketplace/configs"
)

type Storage struct {
	S3Client   *s3.S3
	S3Uploader *s3manager.Uploader
	S3Config   *S3Config
}

func NewStorage() *Storage {
	s3Config := S3Config{
		AwsAccessKeyId: configs.Get().AwsAccessKeyId,
		AwsSecretKey:   configs.Get().AwsSecretKey,
		AwsRegion:      configs.Get().AwsRegion,
		BucketName:     configs.Get().S3BucketName,
	}
	s3Client := NewS3Client(s3Config)
	s3Uploader := NewS3Uploader(s3Config)

	return &Storage{
		S3Client:   s3Client,
		S3Uploader: s3Uploader,
		S3Config:   &s3Config,
	}
}
