package test

import (
	"fmt"
	"homeschooledu/utils"
	"testing"
)


func TestAccountLogin(t *testing.T) {
	// POST 请求示例
	postParams := utils.RequestParams{
		Method: "POST",
		URL:    "http://127.0.0.1:9123/api/login",
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: map[string]string{
			"account": "admin",
			"password":     "12345678",
		},
	}

	postResponse, postErr := utils.SendRequest(postParams)
	if postErr != nil {
		fmt.Println("POST request error:", postErr)
	} else {
		fmt.Println(postResponse)
	}
}

func TestMicroLogin(t *testing.T) {
	// POST 请求示例
	postParams := utils.RequestParams{
		Method: "POST",
		URL:    "http://127.0.0.1:9123/api/microLogin",
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: map[string]interface{}{
			"login_number": 222555888,
			"password":     "123456",
		},
	}

	postResponse, postErr := utils.SendRequest(postParams)
	if postErr != nil {
		fmt.Println("POST request error:", postErr)
	} else {
		fmt.Println(postResponse)
	}
}