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

func ClassList(c *gin.Context) {
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
		result = make([]*models.ClassShow, 0)
	)

	st, err := services.NewClassService(db).List(offset, limit, uint(schoolID), uint(classID), name)
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
		result = append(result, &models.ClassShow{
			ID:         st[i].ID,
			CustomID:   fmt.Sprintf("C%06d", st[i].ID),
			ClassName:  st[i].ClassName,
			SchoolID:   st[i].SchoolID,
			SchoolName: sc[st[i].SchoolID].Name,
			CreateTime: st[i].CreateTime,
			UpdateTime: st[i].UpdateTime,
			IsDelete:   st[i].IsDelete,
		})
	}

	// Return the JWT token in the response
	c.JSON(http.StatusOK, gin.H{
		"message": "获取数据成功",
		"code" : 0,
		"data" : result,
	})
}


func ClassUpdate(c *gin.Context) {
	var (
		req models.Class
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
		st, err := services.NewClassService(db).Info(req.ID)
		if st == nil || st.ID <= 0 || err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "参数异常,无该班级数据"})
			return
		}
	}

	err := services.NewClassService(db).Update(&models.Class{
		ClassName: req.ClassName,
		SchoolID:    req.SchoolID,
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

func ClassDelete(c *gin.Context)  {
	var (
		req models.Class
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
		st, err := services.NewClassService(db).Info(req.ID)
		if st.ID <= 0 || err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "参数异常,无该学生数据"})
			return
		}
	}


	err := services.NewClassService(db).Update(&models.Class{
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


func ClassAdd(c *gin.Context)  {
	var (
		db = config.GetDB()
	)


	var req models.Class
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.NewClassService(db).Add(&req)


	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": consts.CodeMsg[-3]})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "班级数据添加成功",
		"code" : 0,
	})
}

type ClassDetailShow struct {
	Teachers []*models.Teacher `json:"teachers"`
	Students []*models.Student `json:"students"`
}

func ClassDetail(c *gin.Context)  {
	var (
		db = config.GetDB()
		cid = c.Query("class_id")
	)


	classID, err := strconv.Atoi(cid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	teachers, err := services.NewTeacherClassAssignmentService(db).GetClassTeacher(uint(classID))
	if err != nil {

	}

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": consts.CodeMsg[-3]})
		return
	}


	students, err := services.NewStudentService(db).GetClassStudents(uint(classID))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": consts.CodeMsg[-3]})
		return
	}


	c.JSON(http.StatusOK, gin.H{
		"message": "数据更新成功",
		"code" : 0,
		"data" : ClassDetailShow{
			Teachers: teachers,
			Students: students,
		},
	})
}

type BindTeacherReq struct {
	ClassID uint `json:"class_id"`
	BindTeacher []uint `json:"bind_teacher"`
}

func ClassBindTeacher(c *gin.Context)  {
	var (
		db = config.GetDB()
	)


	var req BindTeacherReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tc := make([]*models.TeacherClassAssignment, 0)
	for _, v := range req.BindTeacher {
		tc = append(tc, &models.TeacherClassAssignment{
			TeacherID:  v,
			ClassID:    req.ClassID,
		})
	}


	err := services.NewTeacherClassAssignmentService(db).Create(tc)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": consts.CodeMsg[-3]})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "关联教师成功",
		"code" : 0,
	})
}