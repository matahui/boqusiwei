package services

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
	"homeschooledu/models"
	"os"
)

type ResourceService struct {
	DB *gorm.DB
}

func NewResourceService(db *gorm.DB) *ResourceService{
	return &ResourceService{DB:db}
}
func (s *ResourceService) Info(id uint) (*models.Resource, error) {
	return models.NewResource().Info(s.DB, id)
}
func (s *ResourceService) List(offset, limit int, lv1, lv2, name string) ([]*models.Resource, error) {
	return  models.NewResource().List(s.DB, offset, limit, lv1, lv2, name)
}


func (s *ResourceService) Update(st *models.Resource, id uint) error  {
	return models.NewResource().Update(s.DB, id, st)
}

func (s *ResourceService) Add(st *models.Resource) error  {
	return models.NewResource().Add(s.DB, st)
}

//处理excel
func (s *ResourceService) ProcessSourceFile(filePath string) (int, error) {
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

	sts := make([]*models.Resource, 0)
	// 遍历行数据，跳过标题行
	for i, row := range rows {
		if i == 0 {
			continue // 跳过标题行
		}

		if len(row) < 6 {
			err = fmt.Errorf("文件格式有误")
			return 0, err
		}


		so := &models.Resource{
			ResourceName: row[1],
			AgeGroup:     row[2],
			Course:       row[3],
			Level1:       row[4],
			Level2:       row[5],
		}

		sts = append(sts, so)
	}

	return models.NewResource().BatchInsert(s.DB, sts)
}

//获取分类
func (s *ResourceService) GetLevel(lv1, lv2, age string) ([]string, []string, []string, []*models.Resource, error) {
	var (
		r1 []string
		r2 []string
		ageGroup []string
		st []*models.Resource
	)

	if lv1 == "" {
		//全部lv1和lv2数据
		err := s.DB.Model(&models.Resource{}).
			Distinct("level_1").
			Where("is_delete = ?", 0). // 仅包含未删除的记录
			Pluck("level_1", &r1).Error
		if err != nil {
			return nil, nil, nil, nil, err
		}

		err = s.DB.Model(&models.Resource{}).
			Distinct("level_2").
			Where("is_delete = ?", 0). // 仅包含未删除的记录
			Pluck("level_2", &r2).Error
		if err != nil {
			return nil, nil, nil, nil, err
		}

		return r1, r2, nil, nil, err
	}

	//lv1下的所有lv2
	if lv1 != "" && lv2 == "" {
		err := s.DB.Model(&models.Resource{}).
			Distinct("level_2").
			Where("level_1 = ? AND is_delete = ?", lv1, 0).
			Pluck("level_2", &r2).Error
		if err != nil {
			return nil, nil, nil, nil, err
		}

		return r1, r2, nil, nil, err
	}

	//lv1,lv2下的适合年龄
	if lv1 != "" && lv2 != "" && age == "" {
		err := s.DB.Model(&models.Resource{}).Where("level_1 = ? AND level_2 = ? AND is_delete = ?", lv1, lv2, 0).Pluck("age_group", &ageGroup).Error
		if err != nil {
			return nil, nil, nil, nil, err
		}

		return r1, r2, ageGroup, nil, err
	}

	//资源
	if lv1 != "" && lv2 != "" && age != "" {
		err := s.DB.Model(&models.Resource{}).Where("level_1 = ? AND level_2 = ? AND age_group = ? AND is_delete = ?", lv1, lv2, age, 0).Find(&st).Error
		if err != nil {
			return nil, nil, nil, nil, err
		}

		return r1, r2, ageGroup, st, nil
	}

	return nil, nil, nil, nil, fmt.Errorf("参数异常")
}

func (s *ResourceService) ExtractResourceIDs(schedules []*models.Schedule) []uint {
	resourceIDs := make([]uint, 0)
	seen := make(map[uint]bool)
	for _, schedule := range schedules {
		if !seen[schedule.ResourceID] {
			resourceIDs = append(resourceIDs, schedule.ResourceID)
			seen[schedule.ResourceID] = true
		}
	}

	return resourceIDs
}

func (s *ResourceService) GetByID(ids []uint) ([]*models.Resource, error) {
	return models.NewResource().FindByIDS(s.DB, ids)
}