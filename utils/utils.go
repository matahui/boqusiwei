package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// 定义一个结构体来表示请求的参数
type RequestParams struct {
	Method  string            // 请求方法 GET 或 POST
	URL     string            // 请求的 URL
	Headers map[string]string // 请求头
	Body    interface{}       // 请求体，可以是 string 或者其他类型
}

// 发送 HTTP 请求的通用函数
func SendRequest(params RequestParams) (string, error) {
	var bodyBytes []byte
	if params.Body != nil {
		// 如果请求体是 string 类型，直接转换为 []byte
		if body, ok := params.Body.(string); ok {
			bodyBytes = []byte(body)
		} else {
			// 否则尝试使用 json 编码请求体
			var err error
			bodyBytes, err = json.Marshal(params.Body)
			if err != nil {
				return "", fmt.Errorf("error marshaling request body: %v", err)
			}
		}
	}

	// 创建请求
	req, err := http.NewRequest(params.Method, params.URL, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	// 设置请求头
	if params.Headers != nil {
		for key, value := range params.Headers {
			req.Header.Set(key, value)
		}
	}

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	respBodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	return string(respBodyBytes), nil
}

func Difference(a, b []uint) []uint {
	m := make(map[uint]bool)
	for _, item := range b {
		m[item] = true
	}

	var diff []uint
	for _, item := range a {
		if !m[item] {
			diff = append(diff, item)
		}
	}
	return diff
}

func Contains(slice []uint, item uint) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func ParseYearMonth(yearMonth string) (int, int, error) {
	if len(yearMonth) != 6 {
		return 0, 0, fmt.Errorf("Invalid format")
	}
	year := yearMonth[:4]
	month := yearMonth[4:]

	yy, _ := strconv.Atoi(year)
	mm, _ := strconv.Atoi(month)
	return yy, mm , nil
}