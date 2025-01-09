package file

import (
	"mime/multipart"
)

type UseCase interface {
	UploadFile(multipart.File, string) (string, error)
}
