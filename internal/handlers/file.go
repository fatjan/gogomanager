package handlers

import (
	"bytes"
	"fmt"
	"image"
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
	fileContent := new(bytes.Buffer)
	if _, err := fileContent.ReadFrom(file); err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file content"})
		return
	}

	if !isAllowedFileType(fileName, fileType, fileContent.Bytes()) {
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

func isAllowedFileType(fileName, fileType string, fileContent []byte) bool {
	allowedMimeTypes := map[string]bool{
		"image/jpeg":               true,
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

	if fileType == "image/jpeg" || fileType == "image/png" {
		if !isValidImage(fileContent) {
			return false
		}
	}

	return true
}

func isValidImage(content []byte) bool {
	_, _, err := image.Decode(bytes.NewReader(content))
	return err == nil
}
