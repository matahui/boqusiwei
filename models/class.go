package models

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"homeschooledu/consts"
)

type Class struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ClassName      string    `gorm:"type:varchar(100);not null" json:"class_name"`
	SchoolID  uint      `gorm:"not null" json:"school_id"`
	CreateTime  consts.CustomTime      `gorm:"not null;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime  consts.CustomTime      `gorm:"not null;default:CURRENT_TIMESTAMP;autoUpdateTime" json:"update_time"`
	IsDelete    int            `gorm:"not null;default:0" json:"is_delete"`
}


type ClassShow struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CustomID  string    `json:"custom_id"`
	ClassName      string    `gorm:"type:varchar(100);not null" json:"class_name"`
	SchoolID  uint      `gorm:"not null" json:"school_id"`
	SchoolName string `json:"school_name"`
	CreateTime  consts.CustomTime      `gorm:"not null;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime  consts.CustomTime      `gorm:"not null;default:CURRENT_TIMESTAMP;autoUpdateTime" json:"update_time"`
	IsDelete    int            `gorm:"not null;default:0" json:"is_delete"` // 0:不可删除 1:可删除
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

func (S *Class) List(db *gorm.DB, offset, limit int, schoolID, classID uint, name string) ([]*Class, int64, int64, error) {
	var(
		st []*Class
		total int64
		page int64
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

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, 0, err
	}

	query = query.Order("create_time desc")

	if limit > 0 {
		err := query.Limit(limit).Offset(offset).Find(&st).Error
		if err != nil {
			return nil, 0, 0, err
		}

		page = (total + int64(limit) - 1) / int64(limit)

		return st, total, page ,nil
	}


	err := query.Find(&st).Error
	if err != nil {
		return nil, 0, 0, err
	}

	return st, total, page ,nil
}

func (S *Class) Update(db *gorm.DB, id uint, st *Class) error {
	result := db.Model(&Class{}).Where("id = ?", id).Updates(st)
	if result.RowsAffected <= 0 {
		return fmt.Errorf("Class update error")
	}

	return nil
}


func (S *Class) Del(db *gorm.DB, id uint) error {
	return db.Unscoped().Where("id = ?", id).Delete(&Class{}).Error
}

func (S *Class) Add(db *gorm.DB, st []*Class) error {
	error := db.Create(st).Error
	if error != nil {
		if mysqlErr, ok := error.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			for _, cla := range st {
				// 检查单个记录是否已经存在
				var existingClass Class
				if db.Where("school_id = ? and class_name = ?", cla.SchoolID, cla.ClassName).First(&existingClass).Error == nil {
					return fmt.Errorf("班级名称已存在:%s", cla.ClassName)
				}
			}

			return error
		}
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

func (S *Class) FindByID(db *gorm.DB, ids []uint) ([]*Class, error) {
	var sc []*Class
	result := db.Where("id in ?", ids).Find(&sc)
	if result.Error != nil {
		return nil, result.Error
	}

	return sc, nil
}