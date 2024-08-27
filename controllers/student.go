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
	"time"
)

func StudentList(c *gin.Context) {
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
			consts.RespondWithError(c, -2, "参数异常")
			return
		}

		page = p
	}

	if pageSizeStr != "" {
		pz, err := strconv.Atoi(pageSizeStr)
		if err != nil || pz < 1 {
			consts.RespondWithError(c, -2, "参数异常")
			return
		}

		pageSize = pz
	}

	if sid != "" {
		sid, err := strconv.Atoi(sid)
		if err != nil {
			consts.RespondWithError(c, -2, "参数异常")
			return
		}

		schoolID = sid
	}

	if cid != "" {
		cid, err := strconv.Atoi(cid)
		if err != nil {
			consts.RespondWithError(c, -2, "参数异常")
			return
		}

		classID = cid
	}

	var (
		offset = (page-1) * pageSize
		limit = pageSize
		sids = make([]uint, 0)
		cids = make([]uint, 0)
		result = make([]*models.StudentShow, 0)
	)

	st, err := services.NewStudentService(db).List(offset, limit, uint(schoolID), uint(classID), name)
	if err != nil {
		consts.RespondWithError(c, -2, "内部异常")
		return
	}

	for _, v := range st.Student {
		sids = append(sids, v.SchoolID)
		cids = append(cids, v.ClassID)
	}

	sc, err := services.NewSchoolService(db).FindByID(sids)
	if err != nil {
		consts.RespondWithError(c, -2, "内部异常")
		return
	}

	cl, err := services.NewClassService(db).FindByID(cids)
	if err != nil {
		consts.RespondWithError(c, -2, "内部异常")
		return
	}



	for i := 0; i < len(st.Student); i++ {
		var (
			className string
		)
		cn, ok := cl[st.Student[i].ClassID]
		if ok {
			className = cn.ClassName
		}


		result = append(result, &models.StudentShow{
			ID:          st.Student[i].ID,
			StudentCode: fmt.Sprintf("%d", st.Student[i].ID),
			LoginNumber: st.Student[i].LoginNumber,
			StudentName: st.Student[i].StudentName,
			ParentName:  st.Student[i].ParentName,
			PhoneNumber: st.Student[i].PhoneNumber,
			ClassID:     st.Student[i].ClassID,
			SchoolID:    st.Student[i].SchoolID,
			ClassName:   className,
			SchoolName:  sc[st.Student[i].SchoolID].Name,
			IsDelete:    st.Student[i].IsDelete,
			CreateTime:  st.Student[i].CreateTime,
			UpdateTime:  st.Student[i].UpdateTime,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "获取数据成功",
		"code" : 0,
		"data" : services.StudentListRespShow{
			Student: result,
			Total:   st.Total,
			Page:    st.Page,
		},
	})
}


func StudentUpdate(c *gin.Context) {
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
		st, err := services.NewStudentService(db).Info(req.ID)
		if st == nil || st.ID <= 0 || err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "参数异常,无该学生数据"})
			return
		}
	}

	err := services.NewStudentService(db).Update(&models.Student{
		LoginNumber: req.LoginNumber,
		StudentName: req.StudentName,
		ParentName:  req.ParentName,
		PhoneNumber: req.PhoneNumber,
		ClassID:     req.ClassID,
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

func StudentDelete(c *gin.Context)  {
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
		st, err := services.NewStudentService(db).Info(req.ID)
		if st == nil || st.ID <= 0 || err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "参数异常,无该学生数据"})
			return
		}
	}


	err := services.NewStudentService(db).Update(&models.Student{
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


func StudentAdd(c *gin.Context)  {
	var (
		db = config.GetDB()
	)


	var req models.Student
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.NewStudentService(db).Add(&req)


	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": consts.CodeMsg[-3]})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "学校数据添加成功",
		"code" : 0,
	})
}

func StudentBatchAdd(c *gin.Context) {
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

	// 解析文件内容并处理导入逻辑（假设已经实现）
	n, err := services.NewStudentService(db).ProcessStudentFile(dst, uint(schoolID), uint(classID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"message": "学生批量导入成功",
		"code" : 0,
		"data" : n,
	})
}


func MicroSelf(c *gin.Context) {
	//参数
	var (
		loginNumber = c.Query("login_number")
		db = config.GetDB()

	)


	ln, err := strconv.Atoi(loginNumber)
	if err != nil  {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "传参异常,学校ID不正确"})
		return
	}

	ps, err := services.NewStudentService(db).GetStudentPoints(int64(ln))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "获取数据异常"})
		return
	}


	c.JSON(http.StatusOK, gin.H{
		"message": "获取数据成功",
		"code" : 0,
		"data" : ps,
	})
}

func MicroRank(c *gin.Context) {
	//参数
	var (
		cid = c.Query("class_id")
		db = config.GetDB()

	)


	classID, err := strconv.Atoi(cid)
	if err != nil  {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "传参异常,班级不正确"})
		return
	}

	ps, err := services.NewStudentService(db).GetClassRanking(uint(classID))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "获取数据异常"})
		return
	}


	c.JSON(http.StatusOK, gin.H{
		"message": "获取数据成功",
		"code" : 0,
		"data" : ps,
	})
}

type TaskFinishReq struct {
	StudentID uint `json:"student_id"`
	ResourceID uint `json:"resource_id"`
}

func MicroTask(c *gin.Context) {
	//参数
	var (
		req = TaskFinishReq{}
		db = config.GetDB()
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	today := time.Now().Format("2006-01-02")

	//外部传的loginNumber
	ss, err := services.NewStudentService(db).InfoByLogin(req.StudentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//外部传入的schedule
	sd, err := services.NewScheduleService(db).Info(req.ResourceID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = services.NewStudentService(db).CompleteTask(ss.ID, sd.ResourceID, today)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "任务上报成功",
		"code" : 0,
	})
}