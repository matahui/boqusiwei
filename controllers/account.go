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

func Login(c *gin.Context) {
	var req AccountLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	lc, err, code := services.NewAccountService(config.GetDB()).Login(req.Account, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": consts.CodeMsg[code]})
		return
	}

	//获取token
	token, err := services.GenerateToken(req.Account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	// Return the JWT token in the response
	c.JSON(http.StatusOK, gin.H{
		"message": consts.CodeMsg[code],
		"token":   token,
		"type" : lc.Cate,
		"account" : lc.Account,
	})

}

func StudentLogin(c *gin.Context) {
	var req StudentLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	err:= services.NewStudentService(config.GetDB()).Login(req.LoginNumber, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	//获取token
	ln := strconv.Itoa(int(req.LoginNumber))
	token, err := services.GenerateToken(ln)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	// Return the JWT token in the response
	c.JSON(http.StatusOK, gin.H{
		"message": "登录成功",
		"token":   token,
		"login_number" : req.LoginNumber,
	})

}
