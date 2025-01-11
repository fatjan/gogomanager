package handlers

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/fatjan/gogomanager/internal/useCases/file"
)

type FileHandler interface {
	Post(ginCtx *gin.Context)
}

type fileHandler struct {
	fileUseCase file.UseCase
}

func NewFileHandler(fileUseCase file.UseCase) FileHandler {
	return &fileHandler{fileUseCase}
}

func (r *fileHandler) Post(ginCtx *gin.Context) {
	file, header, err := ginCtx.Request.FormFile("file")
	if err != nil {
		ginCtx.JSON(400, gin.H{
			"message": "No file provided",
		})
		return
	}
	defer file.Close()

	if header.Size > (100 * 1024) {
		ginCtx.JSON(400, gin.H{
			"message": "File too large",
		})
		return
	}

	if !isAllowedFileType(header.Header.Get("Content-Type")) {
		ginCtx.JSON(400, gin.H{
			"message": "File type not allowed",
		})
		return
	}

	filename := fmt.Sprintf("%d-%s", time.Now().UnixNano(), header.Filename)
	publicUrl, err := r.fileUseCase.UploadFile(file, filename)
	if err != nil {
		ginCtx.JSON(500, gin.H{
			"message": "Failed to upload file",
		})
		return
	}

	ginCtx.JSON(200, gin.H{
		"uri": publicUrl,
	})
}

func isAllowedFileType(fileType string) bool {
	allowedFileType := []string{"image/jpeg", "image/jpg", "image/png"}
	for _, v := range allowedFileType {
		if fileType == v {
			return true
		}
	}
	return false
}
