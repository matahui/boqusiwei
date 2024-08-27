package controllers

import (
	"github.com/gin-gonic/gin"
	"homeschooledu/config"
	"homeschooledu/consts"
	"homeschooledu/services"
	"net/http"
	"strconv"
)

type AccountLoginRequest struct {
	Account  string `json:"account" binding:"required,alphanum,min=5"`  // 账号，必须是5位以上的字母或数字
	Password string `json:"password" binding:"required,min=8"`          // 密码，至少8位
}

type StudentLoginRequest struct {
	LoginNumber  int64 `json:"login_number" binding:"required,min=5"`  // 账号，必须是5位以上的字母或数字
	Password string `json:"password" binding:"required,min=6"`          // 密码，至少8位
}

type AccountLoginResponse struct {
	Token string `json:"token"`
	Type int    `json:"type"`
	Account string `json:"account"`
}

type StudentLoginResponse struct {
	Token string `json:"token"`
	LoginNumber int64 `json:"login_number"`
}

func Login(c *gin.Context) {
	var req AccountLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		consts.RespondWithError(c, -1, "参数错误")
		return
	}


	lc, err, code := services.NewAccountService(config.GetDB()).Login(req.Account, req.Password)
	if err != nil {
		consts.RespondWithError(c, -2, "内部处理异常")
		return
	}

	//获取token
	token, err := services.GenerateToken(req.Account)
	if err != nil {
		consts.RespondWithError(c, -2, "内部处理异常")
		return
	}

	// Return the JWT token in the response
	c.JSON(http.StatusOK, gin.H{
		"code" : 0,
		"message": consts.CodeMsg[code],
		"data" : &AccountLoginResponse{
			Token:   token,
			Type:    int(lc.Cate),
			Account: lc.Account,
		},
	})
}

func StudentLogin(c *gin.Context) {
	var req StudentLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		consts.RespondWithError(c, -1, "参数错误")
		return
	}


	err:= services.NewStudentService(config.GetDB()).Login(req.LoginNumber, req.Password)
	if err != nil {
		consts.RespondWithError(c, -2, "内部处理异常")
		return
	}

	//获取token
	ln := strconv.Itoa(int(req.LoginNumber))
	token, err := services.GenerateToken(ln)
	if err != nil {
		consts.RespondWithError(c, -3, "内部处理异常")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : 0,
		"message": consts.CodeMsg[0],
		"data" : &StudentLoginResponse{
			Token:       token,
			LoginNumber: req.LoginNumber,
		},
	})
}
