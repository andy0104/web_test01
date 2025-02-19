package uploader

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client struct {
	client *s3.Client
}

// func init() {
// 	cfg, err := config.LoadDefaultConfig(context.TODO())
// 	if err != nil {
// 		return err
// 	}
// 	client := s3.NewFromConfig(cfg)
// 	s3Client := S3Client{client: client}

// 	s3Uploader := manager.NewUploader(client)
// }
