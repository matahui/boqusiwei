package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"homeschooledu/config"
	"homeschooledu/consts"
	"homeschooledu/models"
	"homeschooledu/services"
	"net/http"
	"strconv"
)

func SchoolList(c *gin.Context) {
	//参数
	var (
		acc = c.GetHeader("account")
		name = c.Query("name")
		pageStr = c.Query("page")
		pageSizeStr = c.Query("pageSize")
		db = config.GetDB()
		page int
		pageSize int
	)

	info, err := services.NewAccountService(db).Info(acc)
	if err != nil {
		consts.RespondWithError(c, -6, "参数异常,账号不存在")
		return
	}

	if info.Cate != consts.AccountCateAdmin {
		//园长账号，只显示园长所在的学校
		if info.Cate == consts.AccountCateDirector {
			sc, err := services.NewSchoolService(db).FindByAccount(acc)
			if err != nil {
				consts.RespondWithError(c, -20, "内部异常")
				return
			}

			if sc != nil && sc.School != nil {
				sc.School[0].CustomId = fmt.Sprintf("Y000%d", sc.School[0].ID)
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "获取数据成功",
				"code" : 0,
				"data" : sc,
			})
		}

		return
	}



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


	resp, err := services.NewSchoolService(db).List(offset, limit, name)
	if err != nil {
		consts.RespondWithError(c, -20, err.Error())
		return
	}


	for i := 0; i < len(resp.School); i++ {
		resp.School[i].CustomId = fmt.Sprintf("Y000%d", resp.School[i].ID)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "成功",
		"code" : 0,
		"data" : resp,
	})
}

type UpdateSchoolRequest struct {
	ID uint `json:"id"`
	Region string `json:"region"`
	Name string `json:"name"`
}

func SchoolUpdate(c *gin.Context) {
	var (
		acc = c.GetHeader("account")
		db = config.GetDB()
	)

	info, err := services.NewAccountService(db).Info(acc)
	if err != nil {
		consts.RespondWithError(c, -6, "参数异常,账号不存在")
		return
	}

	if info.Cate != consts.AccountCateAdmin {
		consts.RespondWithError(c, -6, "非管理员账号没有权限")
		return
	}

	var req UpdateSchoolRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		consts.RespondWithError(c, -6, "参数异常")
		return
	}

	err = services.NewSchoolService(db).Update(&models.School{
		Name:       req.Name,
		Region:     req.Region,
	}, req.ID)


	if err != nil {
		consts.RespondWithError(c, -20, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "数据更新成功",
		"code" : 0,
	})
}

func SchoolDelete(c *gin.Context)  {
	var (
		acc = c.GetHeader("account")
		db = config.GetDB()
	)

	info, err := services.NewAccountService(db).Info(acc)
	if err != nil {
		consts.RespondWithError(c, -6, "参数异常,账号不存在")
		return
	}

	if info.Cate != consts.AccountCateAdmin {
		consts.RespondWithError(c, -6, "非管理员账号没有权限")
		return
	}

	var req UpdateSchoolRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		consts.RespondWithError(c, -6, "参数异常")
		return
	}

	err = services.NewSchoolService(db).Update(&models.School{
		IsDelete:1,
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


type AddSchoolRequest struct {
	Region string `json:"region"`
	Name string `json:"name"`
}

func SchoolAdd(c *gin.Context)  {
	var (
		acc = c.GetHeader("account")
		db = config.GetDB()
	)

	info, err := services.NewAccountService(db).Info(acc)
	if err != nil {
		consts.RespondWithError(c, -6, "参数异常,账号不存在")
		return
	}

	if info.Cate != consts.AccountCateAdmin {
		consts.RespondWithError(c, -6, "非管理员账号没有权限")
		return
	}

	var req AddSchoolRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		consts.RespondWithError(c, -6, "参数异常")
		return
	}

	if req.Name == "" {
		consts.RespondWithError(c, -6, "参数异常")
		return
	}


	sc, err := services.NewSchoolService(db).FindByName(req.Region, req.Name)
	if err != nil {
		consts.RespondWithError(c, -20, err.Error())
		return
	}

	if sc != nil && sc.Region == req.Region && req.Name == sc.Name {
		consts.RespondWithError(c, -20, "学校名称已存在")
		return
	}


	err = services.NewSchoolService(db).Add(&models.School{
		Name:       req.Name,
		Region:     req.Region,
		Account:    acc,
	})


	if err != nil {
		consts.RespondWithError(c, -20, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "学校数据添加成功",
		"code" : 0,
	})
}

func RegionList(c *gin.Context)  {
	var (
		acc = c.GetHeader("account")
		db = config.GetDB()
	)

	info, err := services.NewAccountService(db).Info(acc)
	if err != nil {
		consts.RespondWithError(c, -6, "参数异常,账号不存在")
		return
	}

	if info.Cate != consts.AccountCateAdmin {
		consts.RespondWithError(c, -6, "非管理员账号没有权限")
		return
	}


	re, err := services.NewRegionService(db).List()
	if err != nil {
		consts.RespondWithError(c, -20, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "成功",
		"code" : 0,
		"data" : re,
	})
}


type AddRegionRequest struct {
	Name string `json:"name"`
}

func RegionAdd(c *gin.Context)  {
	var (
		acc = c.GetHeader("account")
		db = config.GetDB()
	)

	info, err := services.NewAccountService(db).Info(acc)
	if err != nil {
		consts.RespondWithError(c, -6, "参数异常,账号不存在")
		return
	}

	if info.Cate != consts.AccountCateAdmin {
		consts.RespondWithError(c, -6, "非管理员账号没有权限")
		return
	}



	var req AddRegionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		consts.RespondWithError(c, -6, "参数异常")
		return
	}

	re := &models.Region{
		Name:       req.Name,
	}

	err = services.NewRegionService(db).Add(re)
	if err != nil {
		consts.RespondWithError(c, -20, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "成功",
		"code" : 0,
	})
}


func SchoolClass(c *gin.Context) {
	//参数
	var (
		school = c.Query("school_id")
		db = config.GetDB()
	)

	if school == "" {
		consts.RespondWithError(c, -6, "参数错误")
		return
	}

	schoolId, err := strconv.Atoi(school)
	if err != nil {
		consts.RespondWithError(c, -6, "参数错误")
		return
	}

	resp, err := services.NewClassService(db).List(0, 0, uint(schoolId), 0, "")
	if err != nil {
		consts.RespondWithError(c, -20, err.Error())
		return
	}



	c.JSON(http.StatusOK, gin.H{
		"message": "成功",
		"code" : 0,
		"data" : resp,
	})
}