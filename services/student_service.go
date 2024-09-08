package services

import (
	"errors"
	"fmt"
	"github.com/extrame/xls"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"homeschooledu/models"
	"homeschooledu/utils"
	"os"
	"strconv"
)

type StudentService struct {
	DB *gorm.DB
}

func NewStudentService(db *gorm.DB) *StudentService{
	return &StudentService{DB:db}
}
func (s *StudentService) Info(id uint) (*models.Student, error) {
	return models.NewStudent().Info(s.DB, id)
}

func (s *StudentService) InfoByLN(ln int64) (*models.Student, error) {
	var su models.Student
	err := s.DB.Model(&models.Student{}).Where("login_number = ?", ln).Find(&su).Error
	if err != nil {
		return nil, err
	}

	return &su, nil
}

func (s *StudentService) InfoByLogin(ln uint) (*models.Student, error) {
	var st models.Student
	err := s.DB.Model(&models.Student{}).Where("login_number = ?", ln).Find(&st).Error
	if err != nil {
		return nil, err
	}

	return &st, nil
}

type StudentListResp struct {
	Student []*models.Student `json:"student"`
	Total  int64  `json:"total"`
	Page   int64  `json:"page"`
}

type StudentListRespShow struct {
	Student []*models.StudentShow`json:"student"`
	Total  int64  `json:"total"`
	Page   int64  `json:"page"`
}

func (s *StudentService) List(offset, limit int, schoolID, classID uint, name string) (*StudentListResp, error) {
	st, total, page, err := models.NewStudent().List(s.DB, offset, limit, schoolID, classID, name)
	if err != nil {
		return nil, err
	}

	return &StudentListResp{
		Student: st,
		Total:  total,
		Page:   page,
	}, nil
}


func (s *StudentService) Update(st *models.Student, id uint) error  {
	return models.NewStudent().Update(s.DB, id, st)
}

func (s *StudentService) Add(st *models.Student) error  {
	return models.NewStudent().Add(s.DB, st)
}

func (s *StudentService) Delete(id uint) error {
	return models.NewStudent().Del(s.DB, id)
}

//处理excel
func (s *StudentService) ProcessStudentFile(filePath, ext string, schoolID, classID uint) (int, error) {
	switch ext {
	case utils.ExtFileXLSX:
		f, err := excelize.OpenFile(filePath)
		if err != nil {
			return -1, fmt.Errorf("无法打开文件%s 错误提示%v", filePath, err)
		}

		defer func() {
			if err := f.Close(); err != nil {
				fmt.Println(err)
			}
			os.Remove(filePath)  // 处理完后删除文件
		}()


		// 获取第一个工作表
		sheetName := f.GetSheetName(0)
		rows, err := f.GetRows(sheetName)
		if err != nil {
			return -1, fmt.Errorf("无法读取文件工作表%s 错误提示%v", sheetName, err)
		}

		sts := make([]*models.Student, 0)
		// 遍历行数据，跳过标题行
		for i, row := range rows {
			if i == 0 {
				continue // 跳过标题行
			}

			if len(row) < 2 {
				continue // 跳过无效行
			}

			ln, err := strconv.Atoi(row[0])
			if err != nil {
				continue
			}

			student := &models.Student{
				LoginNumber: int64(ln),
				StudentName: row[1],
				SchoolID: schoolID,
				ClassID: classID,
				Password:"123456",
			}

			// 如果有家长姓名
			if len(row) > 2 {
				student.ParentName = row[2]
			}

			// 如果有电话号码
			if len(row) > 3 {
				student.PhoneNumber = row[3]
			}

			sts = append(sts, student)
		}

		if len(sts) <= 0 {
			return -2, fmt.Errorf("请在表格中输入正确的信息")
		}

		return models.NewStudent().BatchInsert(s.DB, sts)

	case utils.ExtFileXLS:
		xlFile, err := xls.Open(filePath, "utf-8")
		if err != nil {
			os.Remove(filePath)  // 处理完后删除文件
			return -1, fmt.Errorf("无法打开文件%s 错误提示%v", filePath, err)
		}

		defer func() {
			os.Remove(filePath)  // 处理完后删除文件
		}()

		sheet := xlFile.GetSheet(0)
		sts := make([]*models.Student, 0)

		if sheet.MaxRow != 0 {
			for i := 0; i < int(sheet.MaxRow); i++ {
				if i == 0 {
					continue
				}

				row := sheet.Row(i)
				if row.LastCol() < 2 {
					continue // 跳过无效行
				}


				ln, err := strconv.Atoi(row.Col(0))
				if err != nil {
					continue
				}

				student := &models.Student{
					LoginNumber: int64(ln),
					StudentName: row.Col(1),
					SchoolID: schoolID,
					ClassID: classID,
				}

				// 如果有家长姓名
				if row.LastCol() > 2 {
					student.ParentName = row.Col(2)
				}

				// 如果有电话号码
				if row.LastCol() > 3 {
					student.PhoneNumber = row.Col(3)
				}

				sts = append(sts, student)
			}

			if len(sts) <= 0 {
				return -2, fmt.Errorf("请在表格中输入正确的信息")
			}

			return models.NewStudent().BatchInsert(s.DB, sts)
		}

	default:
		return -2, fmt.Errorf("暂时不支持其他格式")
	}

	return -2, fmt.Errorf("暂时不支持其他格式")
}

func (s *StudentService) GetClassStudents(classID uint) ([]*models.Student, error) {
	st, _, _, err := models.NewStudent().List(s.DB, 0, 0, 0, classID, "")
	if err != nil {
		return nil, err
	}

	return st, nil
}

func (s *StudentService) Login(ln int64, pwd string) error {
	var st models.Student
	err := s.DB.Model(&models.Student{}).Where("login_number = ?", ln).Find(&st).Error
	if err != nil {
		return err
	}

	if st.ID < 0 || st.LoginNumber != ln {
		return fmt.Errorf("账号不存在，请找园长注册账号")
	}

	if st.Password != pwd {
		return fmt.Errorf("密码错误")
	}

	return nil
}

func (s *StudentService) GetStudentPoints(ln int64) (*models.StudentPointShow, error) {
	stu, err := models.NewStudent().GetByLoginNumber(s.DB, ln)
	if err != nil {
		return  nil, err
	}


	sp, err := models.NewStudentPoint().Info(s.DB, stu.ID)
	if err != nil {
		return nil, err
	}

	var points int64
	if sp != nil {
		points = sp.Points
	}

	stars := points / 100 % 10
	moons := points / 1000 % 10
	suns := points / 10000 % 10
	apples := points / 100000 % 10


	result :=  &models.StudentPointShow{
		StudentID:stu.ID,
		Name:   stu.StudentName,
		Points: points,
		Stars:  stars,
		Moons:  moons,
		Suns:   suns,
		Apples: apples,
		ClassID:stu.ClassID,
		ClassName: "",
	}

	cla, err := models.NewClass().Info(s.DB, stu.ClassID)
	if cla != nil && err == nil {
		result.ClassName = cla.ClassName
	}

	return result, nil
}

func (s *StudentService) GetStudentSchedule(ln int64) ([]*models.ResourceStudentShow, error) {
	var st []*models.ResourceStudentShow

	err := s.DB.Table("schedules").
		Select("schedules.id AS schedule_id, resources.resource_name, resources.course, resources.age_group, resources.level_1, resources.level_2, resources.path").
		Joins("JOIN resources ON schedules.resource_id = resources.id").
		Joins("JOIN students ON schedules.class_id = students.class_id").
		Where("students.login_number = ? AND CURDATE() BETWEEN DATE(schedules.begin_time) AND DATE(schedules.end_time)", ln).
		Scan(&st).Error

	if err != nil {
		return nil, err
	}

	return st, nil
}

func (s *StudentService) CompleteTask(sid, rid uint, day string) (int, error) {
	var log models.ActivityLog
	err := s.DB.Where("student_id = ? AND resource_id = ? AND activity_date = ?", sid, rid, day).First(&log).Error
	if err == nil && log.ID > 0  {
		return -1, fmt.Errorf("该任务当天已完成")
	}

	points := 10
	if err := s.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "student_id"}, {Name: "resource_id"}, {Name: "activity_date"}},
		DoUpdates: clause.AssignmentColumns([]string{"points_award", "update_time"}),
	}).Create(&models.ActivityLog{
		StudentID:     sid,
		ResourceID:    rid,
		ActivityDate:  day,
		PointsAward: points,
	}).Error; err != nil {
		return -2, fmt.Errorf("当天任务记录异常")
	}

	//更新学生积分
	var studentPoint models.StudentPoint

	// 先检查记录是否存在
	if err := s.DB.Where("student_id = ?", sid).First(&studentPoint).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 如果记录不存在，插入新的记录
			studentPoint = models.StudentPoint{
				StudentID: int64(sid),
				Points:    int64(points),
			}

			if err := s.DB.Create(&studentPoint).Error; err != nil {
				return -3, fmt.Errorf("学生积分插入失败")
			}
		} else {
			// 其他错误处理
			return -4, fmt.Errorf("学生积分查询失败")
		}
	} else {
		// 如果记录存在，更新积分
		if err := s.DB.Model(&studentPoint).Update("points", gorm.Expr("points + ?", points)).Error; err != nil {
			return -5, fmt.Errorf("学生积分更新失败")
		}
	}

	return 0, nil
}

func (s *StudentService) GetClassRanking(cid uint) ([]*models.StudentRanking, error) {
	var ranks []*models.StudentRanking

	err := s.DB.Raw(`SELECT s.id, s.login_number, s.student_name, sp.points, RANK() OVER (ORDER BY sp.points DESC) as ranks FROM students s JOIN student_points sp ON s.id = sp.student_id WHERE s.class_id = ? AND s.is_delete = 0`, cid).Scan(&ranks).Error
	if err != nil {
		return nil, err
	}


	return ranks, nil
}
