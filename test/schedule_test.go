package test

import (
	"fmt"
	"homeschooledu/utils"
	"testing"
)


func TestScheduleList(t *testing.T) {
	//1分页查询 http://127.0.0.1:9123/api/student/list?page=1&pageSize=10
	//2分页只带学校 http://127.0.0.1:9123/api/student/list?page=1&pageSize=10&school_id=1
	//3分页带学校&班级 http://127.0.0.1:9123/api/student/list?page=1&pageSize=10&school_id=1&class_id=1
	//4带姓名http://127.0.0.1:9123/api/student/list?page=1&pageSize=10&name=郭
	queryParams := utils.RequestParams{
		Method: "GET",
		URL:    "http://127.0.0.1:9123/api/schedule/list?school_id=1&class_id=4&year_month=202408",
		Headers: map[string]string{
			"Authorization" : "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X2lkIjoiYWRtaW4iLCJleHAiOjE3MjY4MDU3OTQsImp0aSI6ImJlNmExMDhhLTRiMmItNDJhZS04MGU4LWEzNmE4ZDU0ZjlkOSJ9.GUquubPM3g6CXJXtXlFUMghJVrA_TtFlj58epX8kvw8",
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


func TestScheduleAdd(t *testing.T) {
	// POST 请求示例
	postParams := utils.RequestParams{
		Method: "POST",
		URL:    "http://127.0.0.1:9123/api/schedule/add",
		Headers: map[string]string{
			"Content-Type": "application/json",
			"Authorization" : "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X2lkIjoiYWRtaW4iLCJleHAiOjE3MjY4MDU3OTQsImp0aSI6ImJlNmExMDhhLTRiMmItNDJhZS04MGU4LWEzNmE4ZDU0ZjlkOSJ9.GUquubPM3g6CXJXtXlFUMghJVrA_TtFlj58epX8kvw8",
			"account" : "admin",
		},
		Body: map[string]interface{}{
			"resource_id" : 1,
			"school_id" : 1,
			"class_id" : []uint{4, 5, 6},
			"begin_time" : "2024-08-20 15:04:05",
			"end_time" : "2024-09-07 15:04:05",
		},
	}

	postResponse, postErr := utils.SendRequest(postParams)
	if postErr != nil {
		fmt.Println("POST request error:", postErr)
	} else {
		fmt.Println("POST response:", postResponse)
	}
}
