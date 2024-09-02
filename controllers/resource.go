package controllers

import (
	"github.com/gin-gonic/gin"
	"homeschooledu/config"
	"homeschooledu/consts"
	"homeschooledu/models"
	"homeschooledu/services"
	"homeschooledu/utils"
	"net/http"
	"path/filepath"
	"strconv"
)

func ResourceList(c *gin.Context) {
	//参数
	var (
		//acc = c.GetHeader("account")
		name = c.Query("name")
		lv1 = c.Query("level_1")
		lv2 = c.Query("level_2")
		pageStr = c.Query("page")
		pageSizeStr = c.Query("pageSize")
		db = config.GetDB()
		page int
		pageSize int
	)


	if pageStr != "" {
		p, err := strconv.Atoi(pageStr)
		if err != nil || p < 1 {
			consts.RespondWithError(c, -6, "参数异常")
			return
		}

		page = p
	}

	if pageSizeStr != "" {
		pz, err := strconv.Atoi(pageSizeStr)
		if err != nil || pz < 1 {
			consts.RespondWithError(c, -6, "参数异常")
			return
		}

		pageSize = pz
	}


	var (
		offset = (page-1) * pageSize
		limit = pageSize
	)

	st, err := services.NewResourceService(db).List(offset, limit, lv1, lv2, name)
	if err != nil {
		consts.RespondWithError(c, -20, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "获取数据成功",
		"code" : 0,
		"data" : st,
	})
}



func ResourceDelete(c *gin.Context)  {
	var (
		req models.Resource
		db = config.GetDB()
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		consts.RespondWithError(c, -6, "参数异常")
		return
	}

	if req.ID <= 0 {
		consts.RespondWithError(c, -6, "参数异常")
		return
	}


	err := services.NewResourceService(db).Update(&models.Resource{
		IsDelete:     1,
	}, req.ID)


	if err != nil {
		consts.RespondWithError(c, -20, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "数据删除成功",
		"code" : 0,
	})
}


func ResourceBatchAdd(c *gin.Context) {
	var (
		db = config.GetDB()
	)

	// 处理文件上传
	file, err := c.FormFile("file")
	if err != nil {
		consts.RespondWithError(c, -6, "参数异常,请上传文件file")
		return
	}

	// 验证文件扩展名
	ext, err := utils.ValidateFileExtension(file)
	if err != nil {
		consts.RespondWithError(c, -6, "参数异常,文件格式")
		return
	}

	// 保存文件到服务器本地
	dst := filepath.Join("uploads", file.Filename)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		consts.RespondWithError(c, -20, err.Error())
		return
	}

	// 解析文件内容并处理导入逻辑
	n, err := services.NewResourceService(db).ProcessSourceFile(dst, ext)
	if err != nil {
		consts.RespondWithError(c, -20, err.Error())
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"message": "资源导入成功",
		"code" : 0,
		"data" : n,
	})
}

func ResourceCate(c *gin.Context) {
	var (
		lv1 = c.Query("level_1")
		lv2 = c.Query("level_2")
		age = c.Query("age_group")
		db = config.GetDB()
	)

	l1, l2, ageGroup, re, err := services.NewResourceService(db).GetLevel(lv1, lv2, age)
	if err != nil {
		consts.RespondWithError(c, -6, "参数异常")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "数据获取成功",
		"code" : 0,
		"data" : struct {
			Level1 []string `json:"level_1"`
			Level2 []string `json:"level_2"`
			AgeGroup []string `json:"age_group"`
			Resources []*models.Resource `json:"resources"`
		}{
			Level1:l1,
			Level2:l2,
			AgeGroup:ageGroup,
			Resources:re,
		},
	})
}

func MicroHome(c *gin.Context) {
	//参数
	var (
		loginNumber = c.Query("login_number")
		db = config.GetDB()

	)


	ln, err := strconv.Atoi(loginNumber)
	if err != nil  {
		consts.RespondWithError(c, -6, "参数异常")
		return
	}

	ps, err := services.NewStudentService(db).GetStudentSchedule(int64(ln))
	if err != nil {
		consts.RespondWithError(c, -20, err.Error())
		return
	}


	c.JSON(http.StatusOK, gin.H{
		"message": "成功",
		"code" : 0,
		"data" : ps,
	})
}