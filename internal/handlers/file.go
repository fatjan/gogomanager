package handlers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
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

	fileName := header.Filename
	fileType := header.Header.Get("Content-Type")
	if !isAllowedFileType(fileName, fileType) {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file type"})
		return
	}

	filename := fmt.Sprintf("%d-%s", time.Now().UnixNano(), fileName)
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

func isAllowedFileType(fileName, fileType string) bool {
	allowedMimeTypes := map[string]bool{
		"image/jpeg":               true,
		"image/jpg":                true,
		"image/png":                true,
		"application/octet-stream": true,
	}

	allowedExtensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
	}

	if !allowedMimeTypes[fileType] {
		return false
	}

	if fileType == "application/octet-stream" {
		ext := strings.ToLower(filepath.Ext(fileName))
		if !allowedExtensions[ext] {
			return false
		}
	}

	return true
}
