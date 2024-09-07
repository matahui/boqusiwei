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
func ValidateFileExtension(file *multipart.FileHeader) (string, error) {
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !allowedExtensions[ext] {
		return "", errors.New("文件格式不支持，只允许 .xlsx 或 .xls")
	}
	return ext, nil
}

const (
	ExtFileXLSX = ".xlsx"
	ExtFileXLS = ".xls"
)


func ConvertWindowsPathToURL(windowsPath string) string {
	// 定义基础 URL
	baseURL := "https://xiaolongrenbq.com/video/"
	// 移除 Windows 路径中的 "D:\幼儿园资源\小龙人\" 前缀
	const prefix = "D:\\幼儿园资源\\小龙人\\"
	if strings.HasPrefix(windowsPath, prefix) {
		windowsPath = windowsPath[len(prefix):]
	}

	newPath := windowsPath[len(prefix)+1:]
	// 将反斜杠替换为正斜杠
	urlPath := strings.ReplaceAll(newPath, `\\`, "/")

	// 拼接完整的 URL
	return baseURL + urlPath
}
