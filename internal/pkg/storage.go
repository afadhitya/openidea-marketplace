package pkg

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/widcha/openidea-marketplace/configs"
)

type Storage struct {
	S3 *s3.S3
}

func NewStorage() *Storage {
	s3Client := NewS3Client(S3Config{
		AwsAccessKeyId: configs.Get().AwsAccessKeyId,
		AwsSecretKey:   configs.Get().AwsSecretKey,
		AwsRegion:      configs.Get().AwsRegion,
		BucketName:     configs.Get().S3BucketName,
	})

	return &Storage{
		S3: s3Client,
	}
}
