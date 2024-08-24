package models

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Class struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ClassName      string    `gorm:"type:varchar(100);not null" json:"class_name"`
	SchoolID  uint      `gorm:"not null" json:"school_id"`
	CreateTime  time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime  time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP;autoUpdateTime" json:"update_time"`
	IsDelete    int            `gorm:"not null;default:0" json:"is_delete"`
}


type ClassShow struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CustomID  string    `json:"custom_id"`
	ClassName      string    `gorm:"type:varchar(100);not null" json:"class_name"`
	SchoolID  uint      `gorm:"not null" json:"school_id"`
	SchoolName string `json:"school_name"`
	CreateTime  time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime  time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP;autoUpdateTime" json:"update_time"`
	IsDelete    int            `gorm:"not null;default:0" json:"is_delete"`
}


func NewClass() *Class {
	return &Class{}
}

func (S *Class)Info(db *gorm.DB, id uint) (*Class, error) {
	var st Class
	result := db.Where("id = ?", id).First(&st)
	if result.Error != nil {
		return nil, result.Error
	}

	return &st, nil
}

func (S *Class) List(db *gorm.DB, offset, limit int, schoolID, classID uint, name string) ([]*Class, error) {
	var(
		st []*Class
		total int64
	)

	query := db.Model(&Class{}).Where("is_delete = ?", 0)

	if schoolID  > 0  {
		query = query.Where("school_id = ?", schoolID)
	}

	if classID > 0 {
		query = query.Where("id = ?", classID)
	}
	if name != "" {
		query = query.Where("class_name LIKE ?", "%"+name+"%")
	}

	result := query.Count(&total).Offset(offset).Limit(limit).Find(&st)
	if result.Error != nil {
		return nil, result.Error
	}

	return st, nil
}

func (S *Class) Update(db *gorm.DB, id uint, st *Class) error {
	result := db.Model(&Class{}).Where("id = ?", id).Updates(st)
	if result.RowsAffected <= 0 {
		return fmt.Errorf("Class update error")
	}

	return nil
}

func (S *Class) Add(db *gorm.DB, st *Class) error {
	result := db.Create(st)
	if result.Error != nil {
		return fmt.Errorf("Class add error")
	}

	return nil
}

func (S *Class) BatchInsert(db *gorm.DB, sts []*Class) (int, error) {
	result := db.Create(sts)
	if result.Error != nil {
		return 0, result.Error
	}

	return int(result.RowsAffected), nil
}