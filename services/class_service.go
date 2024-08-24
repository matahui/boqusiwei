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
func (s *ClassService) List(offset, limit int, schoolID, classID uint, name string) ([]*models.Class, error) {
	return  models.NewClass().List(s.DB, offset, limit, schoolID, classID, name)
}


func (s *ClassService) Update(st *models.Class, id uint) error  {
	return models.NewClass().Update(s.DB, id, st)
}

func (s *ClassService) Add(st *models.Class) error  {
	return models.NewClass().Add(s.DB, st)
}