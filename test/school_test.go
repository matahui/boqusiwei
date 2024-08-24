package test

import (
	"fmt"
	"homeschooledu/utils"
	"testing"
)


func TestSchoolList(t *testing.T) {
	queryParams := utils.RequestParams{
		Method: "GET",
		URL:    "http://127.0.0.1:9123/api/school/list?page=1&pageSize=10",
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


func TestSchoolUpdate(t *testing.T) {
	// POST 请求示例
	postParams := utils.RequestParams{
		Method: "POST",
		URL:    "http://127.0.0.1:9123/api/school/update",
		Headers: map[string]string{
			"Content-Type": "application/json",
			"Authorization" : "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X2lkIjoiYWRtaW4iLCJleHAiOjE3MjM5NTc1NTksImp0aSI6ImVjNzhiYzJiLTc2YWEtNGIxNi05ZmRmLWFmMmVkNGZkYWEzYSJ9.C2xmFNNAybrJ4SvApK7BNIR2Q2LUO_u19IJ_o1o9Kds",
			"account" : "admin",
		},
		Body: map[string]interface{}{
			"id": 4,
			"region":     "深圳市",
			"name" : "深圳市华侨小学",
		},
	}

	postResponse, postErr := utils.SendRequest(postParams)
	if postErr != nil {
		fmt.Println("POST request error:", postErr)
	} else {
		fmt.Println("POST response:", postResponse)
	}
}

func TestSchoolDelete(t *testing.T) {
	// POST 请求示例
	postParams := utils.RequestParams{
		Method: "POST",
		URL:    "http://127.0.0.1:9123/api/school/delete",
		Headers: map[string]string{
			"Content-Type": "application/json",
			"Authorization" : "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X2lkIjoiYWRtaW4iLCJleHAiOjE3MjM5NTc1NTksImp0aSI6ImVjNzhiYzJiLTc2YWEtNGIxNi05ZmRmLWFmMmVkNGZkYWEzYSJ9.C2xmFNNAybrJ4SvApK7BNIR2Q2LUO_u19IJ_o1o9Kds",
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


func TestSchoolAdd(t *testing.T) {
	// POST 请求示例
	postParams := utils.RequestParams{
		Method: "POST",
		URL:    "http://127.0.0.1:9123/api/school/add",
		Headers: map[string]string{
			"Content-Type": "application/json",
			"Authorization" : "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X2lkIjoiYWRtaW4iLCJleHAiOjE3MjQwNDcxNDIsImp0aSI6IjIzNDZkNjI1LTY2NGEtNGFmNS1iMDFhLWVlZjFhNWZmMTk3ZiJ9.qyoBnn7L2zEqHttBworl1bwwmEO62zjQEf0ny7m47dw",
			"account" : "admin",
		},
		Body: map[string]interface{}{
			"region": "深圳",
			"name" : "深圳市袋鼠小学",
		},
	}

	postResponse, postErr := utils.SendRequest(postParams)
	if postErr != nil {
		fmt.Println("POST request error:", postErr)
	} else {
		fmt.Println("POST response:", postResponse)
	}
}


func TestRegionList(t *testing.T) {
	queryParams := utils.RequestParams{
		Method: "GET",
		URL:    "http://127.0.0.1:9123/api/school/regionList",
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

func TestRegionAdd(t *testing.T) {
	// POST 请求示例
	postParams := utils.RequestParams{
		Method: "POST",
		URL:    "http://127.0.0.1:9123/api/school/regionAdd",
		Headers: map[string]string{
			"Content-Type": "application/json",
			"Authorization" : "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50X2lkIjoiYWRtaW4iLCJleHAiOjE3MjQwNDcxNDIsImp0aSI6IjIzNDZkNjI1LTY2NGEtNGFmNS1iMDFhLWVlZjFhNWZmMTk3ZiJ9.qyoBnn7L2zEqHttBworl1bwwmEO62zjQEf0ny7m47dw",
			"account" : "admin",
		},
		Body: map[string]interface{}{
			"region": "澳大利亚",
			"name" : "澳大利亚袋鼠小学",
		},
	}

	postResponse, postErr := utils.SendRequest(postParams)
	if postErr != nil {
		fmt.Println("POST request error:", postErr)
	} else {
		fmt.Println("POST response:", postResponse)
	}
}