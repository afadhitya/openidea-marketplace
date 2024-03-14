package pkg

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Config struct {
	AwsAccessKeyId string
	AwsSecretKey   string
	AwsRegion      string
	BucketName     string
}

func NewS3Client(config S3Config) *s3.S3 {
	sess, err := getNewSession(config)

	if err != nil {
		fmt.Println("Failed to create AWS session:", err)
		return nil
	}

	s3Client := s3.New(sess)
	fmt.Println("S3 session & client initialized")

	return s3Client
}

func getNewSession(config S3Config) (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(config.AwsRegion),
		Credentials: credentials.NewStaticCredentials(
			config.AwsAccessKeyId,
			config.AwsSecretKey,
			"",
		),
	})

	if err != nil {
		return nil, err
	}

	return sess, nil
}
