package controllers

import (
	"fmt"
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

func TeacherList(c *gin.Context) {
	//参数
	var (
		//acc = c.GetHeader("account")
		name = c.Query("name")
		sid = c.Query("school_id")
		cid = c.Query("class_id")
		pageStr = c.DefaultQuery("page", "1")
		pageSizeStr = c.DefaultQuery("pageSize", "10")
		db = config.GetDB()
		schoolID int
		classID int
	)


	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	if sid != "" {
		schoolID, err = strconv.Atoi(sid)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "传参异常,学校ID不正确"})
			return
		}
	}

	if cid != "" {
		classID, err = strconv.Atoi(cid)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "传参异常,班级ID不正确"})
			return
		}

	}

	var (
		offset = (page-1) * pageSize
		limit = pageSize
		sids = make([]uint, 0)
		result = make([]*models.TeacherShow, 0)
	)

	st, err := services.NewTeacherService(db).List(offset, limit, uint(schoolID), uint(classID), name)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "获取数据异常"})
		return
	}

	for _, v := range st {
		sids = append(sids, v.SchoolID)
	}

	sc, err := services.NewSchoolService(db).FindByID(sids)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "获取数据异常"})
		return
	}


	for i := 0; i < len(st); i++ {
		ts := &models.TeacherShow{
			ID:            st[i].ID,
			CustomID:      fmt.Sprintf("L%07d", st[i].ID),
			LoginNumber:   st[i].LoginNumber,
			TeacherName:   st[i].TeacherName,
			PhoneNumber:   st[i].PhoneNumber,
			SchoolID:      st[i].SchoolID,
			SchoolName:    sc[st[i].SchoolID].Name,
			Role:          st[i].Role,
			RoleName:      consts.TeacherRole[st[i].Role],
			TeachingClass: make([]*models.TeachClass, 0),
			IsDelete:      st[i].IsDelete,
			CreateTime:    st[i].CreateTime,
			UpdateTime:    st[i].UpdateTime,
		}


		//确定老师教授班级
		ci, _ := services.NewTeacherService(db).FindClassInfoByT(st[i].ID)
		for _, v := range ci {
			ts.TeachingClass = append(ts.TeachingClass, &models.TeachClass{
				ID:   v.ID,
				Name: v.ClassName,
			})
		}

		result = append(result, ts)
	}

	// Return the JWT token in the response
	c.JSON(http.StatusOK, gin.H{
		"message": "获取数据成功",
		"code" : 0,
		"data" : result,
	})
}


func TeacherUpdate(c *gin.Context) {
	var (
		req TeacherAddReq
		db = config.GetDB()
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.ID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数异常,id不正确"})
		return
	} else {
		st, err := services.NewTeacherService(db).Info(req.ID)
		if st == nil || st.ID <= 0 || err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "参数异常,无该老师数据"})
			return
		}
	}

	err := services.NewTeacherService(db).Update(&models.Teacher{
		LoginNumber: req.LoginNumber,
		TeacherName: req.Name,
		PhoneNumber: req.PhoneNumber,
		SchoolID:    req.SchoolID,
		Role:        req.Role,
	}, req.ID)


	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": consts.CodeMsg[-3]})
		return
	}

	//处理关系变更
	err = services.NewTeacherClassAssignmentService(db).Update(req.ID, req.ClassID)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "数据更新成功",
		"code" : 0,
	})
}

func TeacherDelete(c *gin.Context)  {
	var (
		req models.Student
		db = config.GetDB()
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.ID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数异常,id不正确"})
		return
	} else {
		st, err := services.NewTeacherService(db).Info(req.ID)
		if st.ID <= 0 || err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "参数异常,无该学生数据"})
			return
		}
	}


	err := services.NewTeacherService(db).Update(&models.Teacher{
		IsDelete:1,
	}, req.ID)


	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": consts.CodeMsg[-3]})
		return
	}

	//关联授课信息
	n, err := services.NewTeacherClassAssignmentService(db).DeleteTeacher(req.ID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "老师授课信息删除失败"})
		return
	}


	c.JSON(http.StatusOK, gin.H{
		"message": "数据删除成功",
		"code" : 0,
		"data" : n,
	})
}


type TeacherAddReq struct {
	ID          uint    `json:"id"`
	LoginNumber int64    `json:"login_number"`
	Password    string   `json:"password"`
	Name string    `json:"name"`
	PhoneNumber string    `json:"phone_number"`
	ClassID     []uint      `json:"class_id"`
	SchoolID    uint     `json:"school_id"`
	Role        uint     `json:"role"`
}

func TeacherAdd(c *gin.Context)  {
	var (
		db = config.GetDB()
	)


	var req TeacherAddReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.NewTeacherService(db).Add(&models.Teacher{
		LoginNumber: req.LoginNumber,
		TeacherName: req.Name,
		PhoneNumber: req.PhoneNumber,
		SchoolID:    req.SchoolID,
		Role:        req.Role,
	}, req.ClassID)


	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": consts.CodeMsg[-3]})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "学校数据添加成功",
		"code" : 0,
	})
}

func TeacherBatchAdd(c *gin.Context) {
	var (
		db = config.GetDB()
		schoolID int
		classID int
	)

	sid := c.PostForm("school_id")
	if sid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "未选择学校"})
		return
	} else {
		schoolID, _ = strconv.Atoi(sid)
	}

	// 校验是否选择了班级
	cid := c.PostForm("class_id")
	if cid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "未选择班级"})
		return
	} else {
		classID, _ = strconv.Atoi(cid)
	}

	// 处理文件上传
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请先上传文件"})
		return
	}

	// 验证文件扩展名
	if err := utils.ValidateFileExtension(file); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 保存文件到服务器本地
	dst := filepath.Join("uploads", file.Filename)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "文件保存失败"})
		return
	}

	// 解析文件内容并处理导入逻辑
	n, err := services.NewTeacherService(db).ProcessTeacherFile(dst, uint(schoolID), uint(classID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"message": "老师批量导入成功",
		"code" : 0,
		"data" : n,
	})
}