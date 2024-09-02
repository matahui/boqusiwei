package services

import (
	"gorm.io/gorm"
	"homeschooledu/models"
)

type ClassService struct {
	DB *gorm.DB
}

func NewClassService(db *gorm.DB) *ClassService{
	return &ClassService{DB:db}
}
func (s *ClassService) Info(id uint) (*models.Class, error) {
	return models.NewClass().Info(s.DB, id)
}

type ClassListResp struct {
	Class []*models.Class `json:"class"`
	Total  int64  `json:"total"`
	Page   int64  `json:"page"`
}


type ClassListRespShow struct {
	Class []*models.ClassShow `json:"class"`
	Total  int64  `json:"total"`
	Page   int64  `json:"page"`
}

func (s *ClassService) List(offset, limit int, schoolID, classID uint, name string) (*ClassListResp, error) {
	  st, total, page, err := models.NewClass().List(s.DB, offset, limit, schoolID, classID, name)
	  if err != nil {
	  	return nil, err
	  }

	  return &ClassListResp{
		  Class: st,
		  Total: total,
		  Page:  page,
	  }, nil
}


func (s *ClassService) Update(st *models.Class, id uint) error  {
	return models.NewClass().Update(s.DB, id, st)
}

func (s *ClassService) Add(st []*models.Class) error  {
	return models.NewClass().Add(s.DB, st)
}

func (s *ClassService) FindByID(ids []uint) (map[uint]*models.Class, error)  {
	sc, err := models.NewClass().FindByID(s.DB, ids)
	if err != nil {
		return nil, err
	}

	r := make(map[uint]*models.Class)
	for _, v := range sc {
		r[v.ID] = v
	}

	return r, nil
}

func (s *ClassService) FindByName(sid uint, name string) (*models.Class, error) {
	var cl models.Class
	err := s.DB.Model(&models.Class{}).Where("school_id = ? and class_name = ? and is_delete = 0", sid, name).Find(&cl).Error
	if err != nil {
		return nil, err
	}

	return &cl, nil
}

func (s *ClassService) CanDel(cid uint) bool {
	//班级能否被删除
	var sn int64
	err := s.DB.Model(&models.Student{}).Where("class_id = ? and is_delete = 0", cid).Count(&sn).Error
	if err != nil {
		return false
	}

	if sn > 0 {
		//有学生还在，不能删除
		return false
	}

	//有授课
	err = s.DB.Model(&models.TeacherClassAssignment{}).Where("class_id = ? and is_delete = 0", cid).Count(&sn).Error
	if err != nil {
		return false
	}

	if sn > 0 {
		return false
	}

	return true
}