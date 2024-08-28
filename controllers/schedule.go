package controllers

import (
	"github.com/gin-gonic/gin"
	"homeschooledu/config"
	"homeschooledu/consts"
	"homeschooledu/models"
	"homeschooledu/services"
	"homeschooledu/utils"
	"net/http"
	"sort"
	"strconv"
	"time"
)

func ScheduleList(c *gin.Context) {
	//参数
	var (
		//acc = c.GetHeader("account")
		sid = c.Query("school_id")
		cid = c.Query("class_id")
		ym = c.Query("year_month")
		db = config.GetDB()
		schoolID int
		classID int
	)

	if sid != "" {
		schoolID, _ = strconv.Atoi(sid)
	} else {
		consts.RespondWithError(c, -6, "参数异常")
		return
	}

	if cid != "" {
		classID, _ = strconv.Atoi(cid)
	} else {
		consts.RespondWithError(c, -6, "参数异常")
		return
	}


	year, month, err := utils.ParseYearMonth(ym)
	if err != nil {
		consts.RespondWithError(c, -6, "参数异常")
		return
	}

	// 获取当月的起始日期和结束日期
	beginOfMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	endOfMonth := beginOfMonth.AddDate(0, 1, -1).Add(+time.Second)


	st, err := services.NewScheduleService(db).List(uint(schoolID), uint(classID), beginOfMonth, endOfMonth)
	if err != nil {
		consts.RespondWithError(c, -20, err.Error())
		return
	}

	if st == nil || len(st) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "暂无数据",
			"code" : 0,
		})
		return
	}


	//获取资源id
	re := services.NewResourceService(db).ExtractResourceIDs(st)
	if len(re) <= 0 {
		consts.RespondWithError(c, -6, "暂无数据")
		return
	}

	resource, err := services.NewResourceService(db).GetByID(re)
	if err != nil {
		consts.RespondWithError(c, -6, err.Error())
		return
	}

	calendar := generateCalendar(st, resource, beginOfMonth, endOfMonth, ym)
	sort.Slice(calendar, func(i, j int) bool {
		return calendar[i].Date < calendar[j].Date
	})

	c.JSON(http.StatusOK, gin.H{
		"message": "获取数据成功",
		"code" : 0,
		"data" : calendar,
	})
}

type CalendarEntry struct {
	Date      string     `json:"date"`
	Resources []*models.Resource `json:"resources"`
}

func generateCalendar(schedules []*models.Schedule, resources []*models.Resource, beginOfMonth, endOfMonth time.Time, ym string) []CalendarEntry {
	calendarMap := make(map[string][]*models.Resource)

	// 遍历每条排课记录
	for _, schedule := range schedules {
		// 获取资源信息
		resource := findResourceByID(resources, schedule.ResourceID)

		// 计算出 schedule 涉及的每一天
		date := schedule.BeginTime
		for {
			if date.Before(beginOfMonth) {
				continue
			}

			dayStr := date.Format("20060102")
			if dayStr[:6] != ym {
				break
			}

			if date.After(schedule.EndTime) {
				break
			}

			calendarMap[dayStr] = append(calendarMap[dayStr], resource)
			date = date.AddDate(0, 0, 1)
		}
	}

	var calendar []CalendarEntry
	for date, resources := range calendarMap {
		calendar = append(calendar, CalendarEntry{
			Date:      date,
			Resources: resources,
		})
	}


	return calendar
}

func findResourceByID(resources []*models.Resource, resourceID uint) *models.Resource {
	for _, resource := range resources {
		if resource.ID == resourceID {
			return resource
		}
	}
	return nil
}

type ScheduleAddReq struct {
	ResourceID uint `json:"resource_id"`
	SchoolID   uint  `json:"school_id"`
	ClassID    []uint `json:"class_id"`
	BeginTime  string `json:"begin_time"`
	EndTime string `json:"end_time"`
}

func ScheduleAdd(c *gin.Context) {
	var (
		req ScheduleAddReq
		db = config.GetDB()
		sd = make([]*models.Schedule, 0)
		beginTime, endTime time.Time
		err error
	)

	 err = c.ShouldBindJSON(&req)
	 if err != nil {
		 consts.RespondWithError(c, -6, "参数异常")
		 return
	}

	beginTime, err = time.Parse(consts.TimeFormatLayout, req.BeginTime)
	if err != nil {
		consts.RespondWithError(c, -6, "begin_time格式错误，需要2006-01-02 15:04:05")
		return
	}

	endTime, err = time.Parse(consts.TimeFormatLayout, req.EndTime)
	if err != nil {
		consts.RespondWithError(c, -6, "end_time格式错误，需要2006-01-02 15:04:05")
		return
	}


	for _, v := range req.ClassID {
		sd = append(sd, &models.Schedule{
			ResourceID: req.ResourceID,
			SchoolID:  req.SchoolID,
			ClassID:   v,
			BeginTime: beginTime,
			EndTime:   endTime,
		})
	}

	n, err := services.NewScheduleService(db).BatchAdd(sd)
	if err != nil {
		consts.RespondWithError(c, -20, err.Error())
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"message": "成功",
		"code" : 0,
		"data" : n,
	})
}
