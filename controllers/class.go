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
		pageStr = c.Query("page")
		pageSizeStr = c.Query("pageSize")
		db = config.GetDB()
		schoolID int
		classID int
		page int
		pageSize int
	)


	if pageStr != "" {
		p, err := strconv.Atoi(pageStr)
		if err != nil || p < 1 {
			consts.RespondWithError(c, -6, "参数错误")
			return
		}

		page = p
	}

	if pageSizeStr != "" {
		pz, err := strconv.Atoi(pageSizeStr)
		if err != nil || pz < 1 {
			consts.RespondWithError(c, -6, "参数错误")
			return
		}

		pageSize = pz
	}


	if sid != "" {
		si, err := strconv.Atoi(sid)
		if err != nil {
			consts.RespondWithError(c, -6, "参数错误")
			return
		}

		schoolID = si
	}

	if cid != "" {
		ci, err := strconv.Atoi(cid)
		if err != nil {
			consts.RespondWithError(c, -6, "参数错误")
			return
		}

		classID = ci

	}

	var (
		offset = (page-1) * pageSize
		limit = pageSize
		sids = make([]uint, 0)
		result = make([]*models.ClassShow, 0)
	)

	st, err := services.NewClassService(db).List(offset, limit, uint(schoolID), uint(classID), name)
	if err != nil {
		consts.RespondWithError(c, -20, err.Error())
		return
	}

	for _, v := range st.Class {
		sids = append(sids, v.SchoolID)
	}

	sc, err := services.NewSchoolService(db).FindByID(sids)
	if err != nil {
		consts.RespondWithError(c, -20, err.Error())
		return
	}

	if sc == nil {
		consts.RespondWithError(c, -20, err.Error())
		return
	}



	for i := 0; i < len(st.Class); i++ {
		if st != nil && st.Class[i] != nil {
			sn, ok := sc[st.Class[i].SchoolID]
			if !ok || sn == nil 
				continue
			}

			cs := &models.ClassShow{
				ID:         st.Class[i].ID,
				CustomID:   fmt.Sprintf("C%06d", st.Class[i].ID),
				ClassName:  st.Class[i].ClassName,
				SchoolID:   st.Class[i].SchoolID,
				SchoolName: sc[st.Class[i].SchoolID].Name,
				CreateTime: st.Class[i].CreateTime,
				UpdateTime: st.Class[i].UpdateTime,
				IsDelete:   st.Class[i].IsDelete,
			}

			//判断某个班级
			if services.NewClassService(db).CanDel(cs.ID) {
				cs.IsDelete = 1
			}

			result = append(result, cs)
		}

	}

	c.JSON(http.StatusOK, gin.H{
		"message": "成功",
		"code" : 0,
		"data" : services.ClassListRespShow{
			Class: result,
			Total: st.Total,
			Page:  st.Page,
		},
	})
}


func ClassUpdate(c *gin.Context) {
	var (
		req models.Class
		db = config.GetDB()
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		consts.RespondWithError(c, -6, "参数错误")
		return
	}

	if req.ID <= 0 {
		consts.RespondWithError(c, -6, "参数错误")
		return
	} else {
		st, err := services.NewClassService(db).Info(req.ID)
		if st == nil || st.ID <= 0 || err != nil {
			consts.RespondWithError(c, -6, "参数错误")
			return
		}
	}

	sc, err := services.NewClassService(db).FindByName(req.SchoolID, req.ClassName)
	if err != nil {
		consts.RespondWithError(c, -20, err.Error())
		return
	}

	if sc != nil && sc.SchoolID == req.SchoolID && sc.ClassName == req.ClassName {
		consts.RespondWithError(c, -20, "班级名称已存在")
		return
	}

	err = services.NewClassService(db).Update(&models.Class{
		ClassName: req.ClassName,
		SchoolID:    req.SchoolID,
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

func ClassDelete(c *gin.Context)  {
	var (
		req models.Class
		db = config.GetDB()
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		consts.RespondWithError(c, -6, "参数错误")
		return
	}

	if req.ID <= 0 {
		consts.RespondWithError(c, -6, "参数错误")
		return
	} else {
		st, err := services.NewClassService(db).Info(req.ID)
		if st.ID <= 0 || err != nil {
			consts.RespondWithError(c, -6, "参数错误")
			return
		}
	}


	err := services.NewClassService(db).Update(&models.Class{
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


type ClassAddReq struct {
	SchoolID uint `json:"school_id"`
	ClassName []string `json:"class_name"`
}

func ClassAdd(c *gin.Context)  {
	var (
		db = config.GetDB()
	)


	var req ClassAddReq
	if err := c.ShouldBindJSON(&req); err != nil {
		consts.RespondWithError(c, -2, "参数错误")
		return
	}

	if len(req.ClassName) <= 0 {
		consts.RespondWithError(c, -2, "参数错误")
		return
	}

	cla := make([]*models.Class, 0)
	for _, v := range req.ClassName {
		cla = append(cla, &models.Class{
			ClassName:  v,
			SchoolID:   req.SchoolID,
		})
	}

	err := services.NewClassService(db).Add(cla)


	if err != nil {
		consts.RespondWithError(c, -20, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "成功",
		"code" : 0,
		"data" : len(cla),
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
		consts.RespondWithError(c, -2, "参数错误")
		return
	}

	teachers, err := services.NewTeacherClassAssignmentService(db).GetClassTeacher(uint(classID))
	if err != nil {
		consts.RespondWithError(c, -20, err.Error())
		return
	}

	if err != nil {
		consts.RespondWithError(c, -20, err.Error())
		return
	}


	students, err := services.NewStudentService(db).GetClassStudents(uint(classID))
	if err != nil {
		consts.RespondWithError(c, -20, err.Error())
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
		consts.RespondWithError(c, -3, "参数异常")
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
		consts.RespondWithError(c, -20, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "关联教师成功",
		"code" : 0,
	})
}