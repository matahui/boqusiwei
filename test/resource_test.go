package test

import (
	"bytes"
	"fmt"
	"homeschooledu/utils"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"testing"
)


func TestResourceList(t *testing.T) {
	//1分页查询 http://127.0.0.1:9123/api/student/list?page=1&pageSize=10
	//2分页只带学校 http://127.0.0.1:9123/api/student/list?page=1&pageSize=10&school_id=1
	//3分页带学校&班级 http://127.0.0.1:9123/api/student/list?page=1&pageSize=10&school_id=1&class_id=1
	//4带姓名http://127.0.0.1:9123/api/student/list?page=1&pageSize=10&name=郭
	queryParams := utils.RequestParams{
		Method: "GET",
		URL:    "http://127.0.0.1:9123/api/resource/list?page=1&pageSize=10",
		Headers: map[string]string{
			"Authorization" : "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X2lkIjoiYWRtaW4iLCJleHAiOjE3MjQyMjQyNjQsImp0aSI6IjdiNTc1ZmQ3LWI3Y2ItNGYxNS1hY2U5LTVmZjI1YjUwM2M1MSJ9.lUZgFGnwsZz9iXCh7YSSoZSRAuJn0OVlRoCrELL_vHA",
			"account" : "admin",
		},
	}

	postResponse, postErr := utils.SendRequest(queryParams)
	if postErr != nil {
		fmt.Println("GET request error:", postErr)
	} else {
		fmt.Println(postResponse)
	}
}


func TestResourceDelete(t *testing.T) {
	// POST 请求示例
	postParams := utils.RequestParams{
		Method: "POST",
		URL:    "http://127.0.0.1:9123/api/resource/delete",
		Headers: map[string]string{
			"Content-Type": "application/json",
			"Authorization" : "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X2lkIjoiYWRtaW4iLCJleHAiOjE3MjQwNDcxNDIsImp0aSI6IjIzNDZkNjI1LTY2NGEtNGFmNS1iMDFhLWVlZjFhNWZmMTk3ZiJ9.qyoBnn7L2zEqHttBworl1bwwmEO62zjQEf0ny7m47dw",
			"account" : "admin",
		},
		Body: map[string]interface{}{
			"id" : 3,
			"student_code" : "10000003",
			"login_number" : "19900008666",
			"student_name" : "罗德曼",
			"parent_name" : "罗大海",
			"phone_number" : "13489091234",
			"school_id" : 2,
			"class_id" : 11,
		},
	}

	postResponse, postErr := utils.SendRequest(postParams)
	if postErr != nil {
		fmt.Println("POST request error:", postErr)
	} else {
		fmt.Println("POST response:", postResponse)
	}
}


func TestResourceBatchAdd(t *testing.T) {
	// 构建一个缓冲区来存储multipart/form-data数据
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// 添加文件字段
	file, err := os.Open(`D:\company\homeschooledu\test\resources.xlsx`)
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return
	}
	defer file.Close()

	part, err := writer.CreateFormFile("file", "resources.xlsx")
	if err != nil {
		fmt.Println("无法创建文件字段:", err)
		return
	}
	_, err = io.Copy(part, file)
	if err != nil {
		fmt.Println("无法复制文件内容:", err)
		return
	}

	// 关闭multipart writer以结束消息体
	writer.Close()

	// 创建HTTP请求
	req, err := http.NewRequest("POST", "http://127.0.0.1:9123/api/resource/batchAdd", &requestBody)
	if err != nil {
		fmt.Println("无法创建请求:", err)
		return
	}

	// 设置Content-Type为multipart/form-data
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X2lkIjoiYWRtaW4iLCJleHAiOjE3MjQyMjQyNjQsImp0aSI6IjdiNTc1ZmQ3LWI3Y2ItNGYxNS1hY2U5LTVmZjI1YjUwM2M1MSJ9.lUZgFGnwsZz9iXCh7YSSoZSRAuJn0OVlRoCrELL_vHA")

	// 发送HTTP请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("请求失败:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("无法读取响应体:", err)
		return
	}

	fmt.Println("响应状态码:", resp.StatusCode)
	fmt.Println("响应体:", string(body))
}

func TestResourceCate(t *testing.T) {
	//1分页查询 http://127.0.0.1:9123/api/student/list?page=1&pageSize=10
	//2分页只带学校 http://127.0.0.1:9123/api/student/list?page=1&pageSize=10&school_id=1
	//3分页带学校&班级 http://127.0.0.1:9123/api/student/list?page=1&pageSize=10&school_id=1&class_id=1
	//4带姓名http://127.0.0.1:9123/api/student/list?page=1&pageSize=10&name=郭
	queryParams := utils.RequestParams{
		Method: "GET",
		URL:    "http://127.0.0.1:9123/api/resource/cate?level_1=阅读&level_2=闪卡&age_group=小班",
		Headers: map[string]string{
			"Authorization" : "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X2lkIjoiYWRtaW4iLCJleHAiOjE3MjQyMjQyNjQsImp0aSI6IjdiNTc1ZmQ3LWI3Y2ItNGYxNS1hY2U5LTVmZjI1YjUwM2M1MSJ9.lUZgFGnwsZz9iXCh7YSSoZSRAuJn0OVlRoCrELL_vHA",
			"account" : "admin",
		},
	}

	postResponse, postErr := utils.SendRequest(queryParams)
	if postErr != nil {
		fmt.Println("GET request error:", postErr)
	} else {
		fmt.Println(postResponse)
	}
}