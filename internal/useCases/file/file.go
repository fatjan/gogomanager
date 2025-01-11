package file

import (
	"log"
	"mime/multipart"

	"github.com/fatjan/gogomanager/internal/config"
	s3uploader "github.com/fatjan/gogomanager/internal/pkg/s3_uploader"
)

type useCase struct {
	config config.Config
	s3     *s3uploader.Uploader
}

func NewUseCase(config config.Config) UseCase {
	s3Config := &s3uploader.Config{
		BucketName:      config.Aws.BucketName,
		AccessKeyID:     config.Aws.AccessKeyID,
		AccessKeySecret: config.Aws.SecretAccessKey,
		Region:          config.Aws.Region,
		AccountID:       config.Aws.AccountID,
	}

	s3, _ := s3uploader.NewUploader(s3Config)

	return &useCase{config, s3}
}

func (uc *useCase) UploadFile(file multipart.File, fileName string) (string, error) {
	err := uc.s3.UploadFile(file, fileName)

	if err != nil {
		log.Printf("Failed to upload file to S3: %v", err)
		return "", err
	}

	publicUrl := uc.s3.GetObjectPublicUrls(fileName)
	return publicUrl, nil
}
