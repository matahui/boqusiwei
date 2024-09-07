package test

import (
	"strings"
	"testing"
)

func ConvertWindowsPathToURL(windowsPath string) string {
	// 定义基础 URL
	baseURL := "https://xiaolongrenbq.com/video/"
	// 移除 Windows 路径中的 "D:\幼儿园资源\小龙人\" 前缀
	const prefix = "D:\\幼儿园资源\\小龙人\\"
	if strings.HasPrefix(windowsPath, prefix) {
		windowsPath = windowsPath[len(prefix):]
	}

	// 将反斜杠替换为正斜杠
	urlPath := strings.ReplaceAll(windowsPath, "\\", "/")

	// 拼接完整的 URL
	return baseURL + urlPath
}

func TestConvert(t *testing.T) {
	wp := "D:\\幼儿园资源\\小龙人\\学科\\大班\\a.mp3"
	np := ConvertWindowsPathToURL(wp)
	t.Logf("np:%s\n", np)
}
