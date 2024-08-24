package models

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Schedule struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ResourceID uint     `gorm:"not null" json:"resource_id"`
	SchoolID uint       `gorm:"not null" json:"school_id"`
	ClassID   uint      `gorm:"not null" json:"class_id"`
	BeginTime time.Time `gorm:"not null" json:"begin_time"`
	EndTime   time.Time `gorm:"not null" json:"end_time"`
	IsDelete int        `gorm:"default:0" json:"is_delete"` // 0 表示未删除，1 表示已删除
	CreateTime time.Time `gorm:"autoCreateTime" json:"create_time"`
	UpdateTime time.Time `gorm:"autoUpdateTime" json:"update_time"`
}


func NewSchedule() *Schedule {
	return &Schedule{}
}

func (S *Schedule)Info(db *gorm.DB, id uint) (*Schedule, error) {
	var st Schedule
	result := db.Where("id = ?", id).First(&st)
	if result.Error != nil {
		return nil, result.Error
	}

	return &st, nil
}


func (S *Schedule) Update(db *gorm.DB, id uint, st *Schedule) error {
	result := db.Model(&Schedule{}).Where("id = ?", id).Updates(st)
	if result.RowsAffected <= 0 {
		return fmt.Errorf("Schedule update error")
	}

	return nil
}

func (S *Schedule) Add(db *gorm.DB, st *Schedule) error {
	result := db.Create(st)
	if result.Error != nil {
		return fmt.Errorf("Schedule add error")
	}

	return nil
}

func (S *Schedule) BatchInsert(db *gorm.DB, sts []*Schedule) (int, error) {
	result := db.Create(sts)
	if result.Error != nil {
		return 0, result.Error
	}

	return int(result.RowsAffected), nil
}

