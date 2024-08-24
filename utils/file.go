package utils

import (
	"errors"
	"mime/multipart"
	"path/filepath"
	"strings"
)


// 允许的文件扩展名
var allowedExtensions = map[string]bool{
	".xlsx": true,
	".xls":  true,
}

// 验证文件扩展名
func ValidateFileExtension(file *multipart.FileHeader) error {
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !allowedExtensions[ext] {
		return errors.New("文件格式不支持，只允许 .xlsx 或 .xls")
	}
	return nil
}

