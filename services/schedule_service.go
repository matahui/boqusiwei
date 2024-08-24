package services

import (
	"gorm.io/gorm"
	"homeschooledu/models"
	"time"
)

type ScheduleService struct {
	DB *gorm.DB
}

func NewScheduleService(db *gorm.DB) *ScheduleService{
	return &ScheduleService{DB:db}
}
func (s *ScheduleService) Info(id uint) (*models.Schedule, error) {
	return models.NewSchedule().Info(s.DB, id)
}
func (s *ScheduleService) List(sid, cid uint, bt, et time.Time) ([]*models.Schedule, error) {
	var sc []*models.Schedule
	if err := s.DB.Where("school_id = ? AND class_id = ? AND ((begin_time <= ? AND end_time >= ?) OR (begin_time BETWEEN ? AND ?) OR (end_time BETWEEN ? AND ?))",
		sid, cid, et, bt, bt, et, bt, et).Find(&sc).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return sc, nil
}


func (s *ScheduleService) Update(st *models.Schedule, id uint) error  {
	return models.NewSchedule().Update(s.DB, id, st)
}

func (s *ScheduleService) Add(st *models.Schedule) error  {
	return models.NewSchedule().Add(s.DB, st)
}

func (s *ScheduleService) BatchAdd(sd []*models.Schedule) (int, error) {
	return models.NewSchedule().BatchInsert(s.DB, sd)
}