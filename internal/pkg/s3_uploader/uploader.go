package s3uploader

import (
	"context"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Uploader struct {
	client   *s3.Client
	uploader *manager.Uploader
	cfg      *Config
}

type Config struct {
	BucketName      string
	AccessKeyID     string
	AccessKeySecret string
	Region          string
	PresignDuration time.Duration
	// For testing purpose with Cloudflare R2 only
	// AccountID string
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

func (u *Uploader) UploadFile(file multipart.File, key string) error {
	_, err := u.uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: &u.cfg.BucketName,
		Key:    &key,
		Body:   file,
	})

	return err
}

func (u *Uploader) GetObjectPublicUrls(key string) string {
	publicUrl := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", u.cfg.BucketName, u.cfg.Region, key)

	return publicUrl
}

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
