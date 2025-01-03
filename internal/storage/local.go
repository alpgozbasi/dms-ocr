package storage

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

func SaveLocalFile(fileHeader *multipart.FileHeader) (string, error) {
	timestamp := time.Now().UnixNano()
	fileName := fmt.Sprintf("%d_%s", timestamp, fileHeader.Filename)

	uploadDir := "uploads"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return "", err
	}

	filePath := filepath.Join(uploadDir, fileName)
	return filePath, nil
}
