// Package s3uploader provides a simple wrapper around AWS S3 SDK to upload files to S3 bucket.
package s3uploader

import (
	"context"
	"fmt"
	"mime/multipart"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Uploader is a struct that contains the AWS S3 client and uploader
type Uploader struct {
	client   *s3.Client
	uploader *manager.Uploader
	cfg      *Config
}

// Config is a struct that contains the configuration for the S3 uploader
type Config struct {
	BucketName      string
	AccessKeyID     string
	AccessKeySecret string
	Region          string
	PresignDuration time.Duration
	AccountID       string
}

func NewUploader(cfg *Config) (*Uploader, error) {
	s3Config, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(cfg.Region),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				cfg.AccessKeyID,
				cfg.AccessKeySecret,
				"",
			),
		),
	)

	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(s3Config, func(o *s3.Options) {
		// For testing purpose with Cloudflare R2 only
		// o.BaseEndpoint = aws.String(fmt.Sprintf("https://%s.r2.cloudflarestorage.com", cfg.AccountId))
	})

	return &Uploader{
		uploader: manager.NewUploader(client),
		client:   client,
		cfg:      cfg,
	}, nil
}

// UploadFile uploads a file to the S3 bucket
// The key is the path to the file in the bucket
// The file is the multipart file from the request
// It returns an error if the upload failed
func (u *Uploader) UploadFile(file multipart.File, key string) error {
	ext := key[strings.LastIndex(key, ".")+1:]
	contentTypeMap := map[string]string{
		"jpg":  "image/jpeg",
		"jpeg": "image/jpeg",
		"png":  "image/png",
		"gif":  "image/gif",
		"pdf":  "application/pdf",
		"txt":  "text/plain",
	}

	contentType := contentTypeMap[ext]
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	_, err := u.uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:      &u.cfg.BucketName,
		Key:         &key,
		Body:        file,
		ContentType: &contentType,
	})

	return err
}

// GetObjectPublicUrls generates a public URL for the object in the S3 bucket
// The key is the path to the file in the bucket
// It returns the public URL
// The URL is in the format based on this reference:
// https://docs.aws.amazon.com/AmazonS3/latest/userguide/VirtualHosting.html#virtual-hosted-style-access
func (u *Uploader) GetObjectPublicUrls(key string) string {
	publicUrl := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", u.cfg.BucketName, u.cfg.Region, key)

	return publicUrl
}

// GetObjectPresignedUrl generates a presigned URL for the object in the S3 bucket
// The key is the path to the file in the bucket
// It returns the presigned URL
// The URL is in the format based on this reference:
// https://docs.aws.amazon.com/AmazonS3/latest/API/s3_example_s3_Scenario_PresignedUrl_section.html
func (u *Uploader) GetObjectPresignedUrl(key string) (string, error) {
	presign, err := s3.NewPresignClient(u.client).PresignGetObject(context.TODO(),
		&s3.GetObjectInput{
			Bucket: &u.cfg.BucketName,
			Key:    &key,
		},
		s3.WithPresignExpires(u.cfg.PresignDuration),
	)

	if err != nil {
		return "", err
	}

	return presign.URL, nil
}
