package utils

import (
	"io"
	"mime/multipart"
	"os"
	"strings"
)

func UploadFile(file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Destination
	fileName := "product_" + strings.ReplaceAll(file.Filename, " ", "_")
	dst, err := os.Create("./web/static/images/products/" + fileName)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}
	return fileName, nil
}
