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


func TestStudentList(t *testing.T) {
	//1分页查询 http://127.0.0.1:9123/api/student/list?page=1&pageSize=10
	//2分页只带学校 http://127.0.0.1:9123/api/student/list?page=1&pageSize=10&school_id=1
	//3分页带学校&班级 http://127.0.0.1:9123/api/student/list?page=1&pageSize=10&school_id=1&class_id=1
	//4带姓名http://127.0.0.1:9123/api/student/list?page=1&pageSize=10&name=郭
	queryParams := utils.RequestParams{
		Method: "GET",
		URL:    "http://127.0.0.1:9123/api/student/list?page=1&pageSize=10",
		Headers: map[string]string{
			"Authorization" : "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X2lkIjoiYWRtaW4iLCJleHAiOjE3MjQwNDcxNDIsImp0aSI6IjIzNDZkNjI1LTY2NGEtNGFmNS1iMDFhLWVlZjFhNWZmMTk3ZiJ9.qyoBnn7L2zEqHttBworl1bwwmEO62zjQEf0ny7m47dw",
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


func TestStudentUpdate(t *testing.T) {
	// POST 请求示例
	postParams := utils.RequestParams{
		Method: "POST",
		URL:    "http://127.0.0.1:9123/api/student/update",
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

func TestStudentDelete(t *testing.T) {
	// POST 请求示例
	postParams := utils.RequestParams{
		Method: "POST",
		URL:    "http://127.0.0.1:9123/api/student/delete",
		Headers: map[string]string{
			"Content-Type": "application/json",
			"Authorization" : "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X2lkIjoiYWRtaW4iLCJleHAiOjE3MjQwNDcxNDIsImp0aSI6IjIzNDZkNjI1LTY2NGEtNGFmNS1iMDFhLWVlZjFhNWZmMTk3ZiJ9.qyoBnn7L2zEqHttBworl1bwwmEO62zjQEf0ny7m47dw",
			"account" : "admin",
		},
		Body: map[string]interface{}{
			"id": 4,
		},
	}

	postResponse, postErr := utils.SendRequest(postParams)
	if postErr != nil {
		fmt.Println("POST request error:", postErr)
	} else {
		fmt.Println("POST response:", postResponse)
	}
}


func TestStudentAdd(t *testing.T) {
	// POST 请求示例
	postParams := utils.RequestParams{
		Method: "POST",
		URL:    "http://127.0.0.1:9123/api/student/add",
		Headers: map[string]string{
			"Content-Type": "application/json",
			"Authorization" : "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X2lkIjoiYWRtaW4iLCJleHAiOjE3MjQwNDcxNDIsImp0aSI6IjIzNDZkNjI1LTY2NGEtNGFmNS1iMDFhLWVlZjFhNWZmMTk3ZiJ9.qyoBnn7L2zEqHttBworl1bwwmEO62zjQEf0ny7m47dw",
			"account" : "admin",
		},
		Body: map[string]interface{}{
			"student_code" : "10000011",
			"login_number" : "15611224567",
			"student_name" : "罗菲菲",
			"parent_name" : "罗永胜",
			"phone_number" : "13652348890",
			"school_id" : 1,
			"class_id" : 3,
		},
	}

	postResponse, postErr := utils.SendRequest(postParams)
	if postErr != nil {
		fmt.Println("POST request error:", postErr)
	} else {
		fmt.Println("POST response:", postResponse)
	}
}

func TestStudentBatchAdd(t *testing.T) {
	// 构建一个缓冲区来存储multipart/form-data数据
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// 添加 school_id 字段
	_ = writer.WriteField("school_id", "1")

	// 添加 class_id 字段
	_ = writer.WriteField("class_id", "123")

	// 添加文件字段
	file, err := os.Open(`D:\company\homeschooledu\test\students.xlsx`)
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return
	}
	defer file.Close()

	part, err := writer.CreateFormFile("file", "students.xlsx")
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
	req, err := http.NewRequest("POST", "http://127.0.0.1:9123/api/student/batchAdd", &requestBody)
	if err != nil {
		fmt.Println("无法创建请求:", err)
		return
	}

	// 设置Content-Type为multipart/form-data
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X2lkIjoiYWRtaW4iLCJleHAiOjE3MjQwNDcxNDIsImp0aSI6IjIzNDZkNjI1LTY2NGEtNGFmNS1iMDFhLWVlZjFhNWZmMTk3ZiJ9.qyoBnn7L2zEqHttBworl1bwwmEO62zjQEf0ny7m47dw")

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

func TestMicroSelf(t *testing.T) {
	// POST 请求示例
	postParams := utils.RequestParams{
		Method: "GET",
		URL:    "http://127.0.0.1:9123/api/micro/self?login_number=222555888",
		Headers: map[string]string{
			"Content-Type": "application/json",
			"Authorization" : "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X2lkIjoiMjIyNTU1ODg4IiwiZXhwIjoxNzI2ODE5NzMyLCJqdGkiOiIzMjhiN2E0Mi01NGVjLTRkZjEtODFjNi00NzAzMjE2M2I2NWUifQ.P-kiblLCreO1CSqGoiFymQgZ_rUcl6BGiyeiknUTYrM",
		},

	}

	postResponse, postErr := utils.SendRequest(postParams)
	if postErr != nil {
		fmt.Println("POST request error:", postErr)
	} else {
		fmt.Println("POST response:", postResponse)
	}
}

func TestMicroHome(t *testing.T) {
	// POST 请求示例
	url := "http://121.37.191.233:9123/api/micro/home?login_number=222555888"
	postParams := utils.RequestParams{
		Method: "GET",
		//URL:    "http://127.0.0.1:9123/api/micro/home?login_number=222555888",
		URL : url,
		Headers: map[string]string{
			"Content-Type": "application/json",
			"Authorization" : "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X2lkIjoiMjIyNTU1ODg4IiwiZXhwIjoxNzI2ODE5NzMyLCJqdGkiOiIzMjhiN2E0Mi01NGVjLTRkZjEtODFjNi00NzAzMjE2M2I2NWUifQ.P-kiblLCreO1CSqGoiFymQgZ_rUcl6BGiyeiknUTYrM",
		},

	}

	postResponse, postErr := utils.SendRequest(postParams)
	if postErr != nil {
		fmt.Println("POST request error:", postErr)
	} else {
		fmt.Println(postResponse)
	}
}

func TestMicroTask(t *testing.T) {
	// POST 请求示例
	postParams := utils.RequestParams{
		Method: "POST",
		URL:    "http://127.0.0.1:9123/api/micro/task",
		Headers: map[string]string{
			"Content-Type": "application/json",
			"Authorization" : "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X2lkIjoiMjIyNTU1ODg4IiwiZXhwIjoxNzI2ODE5NzMyLCJqdGkiOiIzMjhiN2E0Mi01NGVjLTRkZjEtODFjNi00NzAzMjE2M2I2NWUifQ.P-kiblLCreO1CSqGoiFymQgZ_rUcl6BGiyeiknUTYrM",
		},
		Body: map[string]interface{}{
			"student_id": 222555888,
			"resource_id" : 9,
		},
	}

	postResponse, postErr := utils.SendRequest(postParams)
	if postErr != nil {
		fmt.Println("POST request error:", postErr)
	} else {
		fmt.Println(postResponse)
	}
}


func TestMicroRank(t *testing.T) {
	// POST 请求示例
	postParams := utils.RequestParams{
		Method: "GET",
		URL:    "http://127.0.0.1:9123/api/micro/rank?class_id=123",
		Headers: map[string]string{
			"Content-Type": "application/json",
			"Authorization" : "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X2lkIjoiMjIyNTU1ODg4IiwiZXhwIjoxNzI2ODE5NzMyLCJqdGkiOiIzMjhiN2E0Mi01NGVjLTRkZjEtODFjNi00NzAzMjE2M2I2NWUifQ.P-kiblLCreO1CSqGoiFymQgZ_rUcl6BGiyeiknUTYrM",
		},

	}

	postResponse, postErr := utils.SendRequest(postParams)
	if postErr != nil {
		fmt.Println("POST request error:", postErr)
	} else {
		fmt.Println(postResponse)
	}
}