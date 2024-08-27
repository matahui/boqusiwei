package models

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Resource struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	ResourceName string   `gorm:"type:varchar(255)" json:"resource_name"`
	AgeGroup    string    `gorm:"type:varchar(255)" json:"age_group"`
	Course    string    `gorm:"type:varchar(255)" json:"course"`
	Level1    string    `gorm:"column:level_1;type:varchar(255)" json:"level_1"`
	Level2    string    `gorm:"column:level_2;type:varchar(255)" json:"level_2"`
	IsDelete int        `gorm:"default:0" json:"is_delete"` // 0 表示未删除，1 表示已删除
	CreateTime time.Time `gorm:"autoCreateTime" json:"create_time"`
	UpdateTime time.Time `gorm:"autoUpdateTime" json:"update_time"`
}

type ResourceStudentShow struct {
	ScheduleID   uint      `json:"schedule_id"`
	ResourceName string   `gorm:"type:varchar(255)" json:"resource_name"`
	AgeGroup    string    `gorm:"type:varchar(255)" json:"age_group"`
	Course    string    `gorm:"type:varchar(255)" json:"course"`
	Level1    string    `gorm:"column:level_1;type:varchar(255)" json:"level_1"`
	Level2    string    `gorm:"column:level_2;type:varchar(255)" json:"level_2"`
}

func NewResource() *Resource {
	return &Resource{}
}

func (S *Resource)Info(db *gorm.DB, id uint) (*Resource, error) {
	var st Resource
	result := db.Where("id = ?", id).First(&st)
	if result.Error != nil {
		return nil, result.Error
	}

	return &st, nil
}

func (S *Resource) List(db *gorm.DB, offset, limit int, lv1, lv2, name string) ([]*Resource, int64, int64, error) {
	var (
		sc []*Resource
		total int64
		page int64
	)

	query := db.Model(&Resource{}).Where("is_delete = 0 ")

	if lv1  != ""   {
		query = query.Where("level_1 = ?", lv1)
	}

	if lv2  != ""   {
		query = query.Where("level_2 = ?", lv2)
	}

	if name != "" {
		query = query.Where("resource_name LIKE ?", "%"+name+"%")
	}

	//总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, 0, err
	}

	if limit > 0 {
		err := query.Limit(limit).Offset(offset).Find(&sc).Error
		if err != nil {
			return nil, 0, 0, err
		}

		page = (total + int64(limit) - 1) / int64(limit)

		return sc, total, page ,nil
	}


	err := query.Find(&sc).Error
	if err != nil {
		return nil, 0, 0, err
	}

	return sc, total, page ,nil
}

func (S *Resource) Update(db *gorm.DB, id uint, st *Resource) error {
	result := db.Model(&Resource{}).Where("id = ?", id).Updates(st)
	if result.RowsAffected <= 0 {
		return fmt.Errorf("Resource update error")
	}

	return nil
}

func (S *Resource) Add(db *gorm.DB, st *Resource) error {
	result := db.Create(st)
	if result.Error != nil {
		return fmt.Errorf("Resource add error")
	}

	return nil
}

func (S *Resource) BatchInsert(db *gorm.DB, sts []*Resource) (int, error) {
	result := db.Create(sts)
	if result.Error != nil {
		return 0, result.Error
	}

	return int(result.RowsAffected), nil
}


func (S *Resource) FindByIDS(db *gorm.DB, ids []uint) ([]*Resource, error) {
	var re []*Resource
	result := db.Model(&Resource{}).Where("id in ?", ids).Find(&re)
	if result.Error != nil {
		return nil, result.Error
	}

	return re, nil
}