package services

import (
	"fmt"
	"github.com/extrame/xls"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
	"homeschooledu/models"
	"homeschooledu/utils"
	"os"
	"strconv"
)

type TeacherService struct {
	DB *gorm.DB
}

func NewTeacherService(db *gorm.DB) *TeacherService{
	return &TeacherService{DB:db}
}
func (s *TeacherService) Info(id uint) (*models.Teacher, error) {
	return models.NewTeacher().Info(s.DB, id)
}

func (s *TeacherService) FindByLN(ln int64) (*models.Teacher, error) {
	var tea models.Teacher
	result := s.DB.Model(&models.Teacher{}).Where("login_number = ?", ln).Find(&tea)
	if result.Error != nil {
		return nil, result.Error
	}

	return &tea, nil
}
type TeacherListResp struct {
	Teacher []*models.Teacher `json:"teacher"`
	Total  int64  `json:"total"`
	Page   int64  `json:"page"`
}

type TeacherListRespShow struct {
	Teacher []*models.TeacherShow`json:"teacher"`
	Total  int64  `json:"total"`
	Page   int64  `json:"page"`
}


func (s *TeacherService) List(offset, limit int, schoolID, classID uint, name string) (*TeacherListResp, error) {
	 st, total, page, err := models.NewTeacher().List(s.DB, offset, limit, schoolID, classID, name)
	 if err != nil {
	 	return nil, err
	 }

	 return &TeacherListResp{
		 Teacher: st,
		 Total:   total,
		 Page:    page,
	 }, nil
}


func (s *TeacherService) Update(st *models.Teacher, id uint) error  {
	return models.NewTeacher().Update(s.DB, id, st)
}

func (s *TeacherService) Add(st *models.Teacher, cids []uint) error  {
	err := models.NewTeacher().Add(s.DB, st)
	if err != nil {
		return err
	}

	tc := make([]*models.TeacherClassAssignment, 0)
	for _, v := range cids {
		tc = append(tc, &models.TeacherClassAssignment{
			TeacherID:  st.ID,
			ClassID:    v,
		})
	}

	if len(tc) > 0 {
		return models.NewTeacherClassAssignment().Add(s.DB, tc)
	}

	return nil
}

func (s *TeacherService) Delete(id uint) error {
	return models.NewTeacher().Del(s.DB, id)
}

//处理excel
func (s *TeacherService) ProcessTeacherFile(filePath, ext string, schoolID uint) (int, error) {
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

		sts := make([]*models.Teacher, 0)
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

			teacher := &models.Teacher{
				LoginNumber: int64(ln),
				Password:"123456",
				TeacherName: row[1],
				SchoolID:    schoolID,
				Role:        2,
			}

			if len(row) > 2 {
				teacher.PhoneNumber = row[2]
			}


			sts = append(sts, teacher)
		}

		if len(sts) <= 0 {
			return -2, fmt.Errorf("请在表格中输入正确的信息")
		}

		return models.NewTeacher().BatchInsert(s.DB, sts)

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
		sts := make([]*models.Teacher, 0)
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

				teacher := &models.Teacher{
					LoginNumber: int64(ln),
					Password:"123456",
					TeacherName: row.Col(1),
					SchoolID:    schoolID,
					Role:        2,
				}

				if row.LastCol() > 2 {
					teacher.PhoneNumber = row.Col(2)
				}

				sts = append(sts, teacher)
			}

			if len(sts) <= 0 {
				return -2, fmt.Errorf("请在表格中输入正确的信息")
			}

			return models.NewTeacher().BatchInsert(s.DB, sts)
		}
	default:
		return -2, fmt.Errorf("暂时不支持其他格式")
	}

	return -2, fmt.Errorf("暂时不支持其他格式")
}

func (s *TeacherService) FindClassInfoByT(tid uint) ([]*models.Class, error) {
	return models.NewTeacher().FindClassInfoByT(s.DB, tid)
}

type TeacherClassAssignmentService struct {
	DB *gorm.DB
}

func NewTeacherClassAssignmentService(db *gorm.DB) *TeacherClassAssignmentService{
	return &TeacherClassAssignmentService{DB:db}
}

func (s *TeacherClassAssignmentService) Create(tc []*models.TeacherClassAssignment) error {
	return models.NewTeacherClassAssignment().Add(s.DB, tc)
}
func (s *TeacherClassAssignmentService) Info(id uint) (*models.TeacherClassAssignment, error) {
	return models.NewTeacherClassAssignment().Info(s.DB, id)
}


func (s *TeacherClassAssignmentService) List(tid uint) ([]*models.TeacherClassAssignment, error) {
	return models.NewTeacherClassAssignment().List(s.DB, tid)
}

//更新教师授课班级
func (s *TeacherClassAssignmentService) Update(tid uint, cid []uint) error {
	var currentAssignments []models.TeacherClassAssignment

	result := s.DB.Model(&models.TeacherClassAssignment{}).Where("teacher_id = ? and is_delete = 0", tid).Find(&currentAssignments)
	if result.Error != nil {
		return result.Error
	}

	currentClassMap := make(map[uint]bool)
	for _, assignment := range currentAssignments {
		currentClassMap[assignment.ClassID] = true
	}

	// Step 2: 找出需要删除的记录
	for _, assignment := range currentAssignments {
		if !utils.Contains(cid, assignment.ClassID) {
			// 软删除不再存在的关联记录
			s.DB.Model(&assignment).Update("is_delete", 1)
		}
	}

	for _, classID := range cid {
		if !currentClassMap[classID] {
			// 如果是新记录，插入到数据库中
			newAssignment := &models.TeacherClassAssignment{
				TeacherID: tid,
				ClassID:   classID,
			}

			if err := s.DB.Create(&newAssignment).Error; err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *TeacherClassAssignmentService) DeleteTeacher(tid uint) (int64, error) {
	result := s.DB.Model(&models.TeacherClassAssignment{}).Where("teacher_id = ? and is_delete = 0", tid).Update("is_delete", 1)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}

func (s *TeacherClassAssignmentService) GetClassTeacher(cid uint) ([]*models.Teacher, error) {
	var teacherIDs []int
	var teachers []*models.Teacher

	// Step 1: 从 teacher_class_assignments 表中获取与 class_id 相关的 teacher_id
	if err := s.DB.Table("teacher_class_assignments").
		Where("class_id = ? AND is_delete = ?", cid, 0).
		Pluck("teacher_id", &teacherIDs).Error; err != nil {
		return nil, err
	}

	if len(teacherIDs) == 0 {
		return teachers, nil // 没有找到相关的 teacher_id
	}

	// Step 2: 从 teachers 表中获取老师信息
	if err := s.DB.Where("id IN (?)", teacherIDs).Find(&teachers).Error; err != nil {
		return nil, err
	}

	return teachers, nil
}