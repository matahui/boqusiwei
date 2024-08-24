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
		pageStr = c.DefaultQuery("page", "1")
		pageSizeStr = c.DefaultQuery("pageSize", "10")
		db = config.GetDB()
	)

	info, err := services.NewAccountService(db).Info(acc)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "account不正确"})
		return
	}

	if info.Cate != consts.AccountCateAdmin {
		//园长账号，只显示园长所在的学校
		if info.Cate == consts.AccountCateDirector {
			sc, err := services.NewSchoolService(db).FindByAccount(acc)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "获取数据异常"})
				return
			}

			if sc != nil {
				sc.CustomId = fmt.Sprintf("Y000%d", sc.ID)
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "获取数据成功",
				"code" : 0,
				"data" : sc,
			})
		}

		return
	}



	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	var (
		offset = (page-1) * pageSize
		limit = pageSize
		sc = make([]*models.School, 0)
	)

	//获取列表数据
	if name == "" {
		//全量查询
		sc, err = services.NewSchoolService(db).List(offset, limit)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "获取数据异常"})
			return
		}
	} else {
		//模糊查询
		sc, err = services.NewSchoolService(db).FuzzySearch(name)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "获取数据异常"})
			return
		}
	}


	for i := 0; i < len(sc); i++ {
		sc[i].CustomId = fmt.Sprintf("Y000%d", sc[i].ID)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "获取数据成功",
		"code" : 0,
		"data" : sc,
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
		c.JSON(http.StatusUnauthorized, gin.H{"error": "account不正确"})
		return
	}

	if info.Cate != consts.AccountCateAdmin {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "非管理员账号没有权限"})
		return
	}

	var req UpdateSchoolRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = services.NewSchoolService(db).Update(&models.School{
		Name:       req.Name,
		Region:     req.Region,
	}, req.ID)


	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": consts.CodeMsg[-3]})
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
		c.JSON(http.StatusUnauthorized, gin.H{"error": "account不正确"})
		return
	}

	if info.Cate != consts.AccountCateAdmin {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "非管理员账号没有权限"})
		return
	}

	var req UpdateSchoolRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = services.NewSchoolService(db).Update(&models.School{
		IsDelete:1,
	}, req.ID)


	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": consts.CodeMsg[-3]})
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
		c.JSON(http.StatusUnauthorized, gin.H{"error": "account不正确"})
		return
	}

	if info.Cate != consts.AccountCateAdmin {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "非管理员账号没有权限"})
		return
	}

	var req AddSchoolRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = services.NewSchoolService(db).Add(&models.School{
		Name:       req.Name,
		Region:     req.Region,
		Account:    acc,
	})


	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": consts.CodeMsg[-3]})
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
		c.JSON(http.StatusUnauthorized, gin.H{"error": "account不正确"})
		return
	}

	if info.Cate != consts.AccountCateAdmin {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "非管理员账号没有权限"})
		return
	}



	re, err := services.NewRegionService(db).List()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "获取校区数据异常"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "获取校区数据成功",
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
		c.JSON(http.StatusUnauthorized, gin.H{"error": "account不正确"})
		return
	}

	if info.Cate != consts.AccountCateAdmin {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "非管理员账号没有权限"})
		return
	}



	var req AddRegionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	re := &models.Region{
		Name:       req.Name,
	}
	err = services.NewRegionService(db).Add(re)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "添加校区失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "添加校区成功",
		"code" : 0,
	})
}