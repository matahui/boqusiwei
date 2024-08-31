package models

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"homeschooledu/consts"
)

type School struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Region    string    `gorm:"type:varchar(100);not null" json:"region"`
	Account   string    `gorm:"type:varchar(100);not null" json:"account"`
	CustomId  string    `gorm:"type:varchar(100);not null" json:"custom_id"`
	IsDelete int        `gorm:"default:0" json:"is_delete"` // 0 表示未删除，1 表示已删除
	CreateTime consts.CustomTime `gorm:"autoCreateTime" json:"create_time"`
	UpdateTime consts.CustomTime `gorm:"autoUpdateTime" json:"update_time"`
}

func NewSchool() *School {
	return &School{}
}

func (S *School) List(db *gorm.DB, offset, limit int, name string) ([]*School, int64, int64, error) {
	var (
		sc []*School
		total int64
		page int64
	)

	query := db.Model(&School{}).Where("is_delete = 0 ")
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
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

func (S *School) QueryByName(db *gorm.DB, name string) ([]*School, error)  {
	var sc []*School

	result := db.Where("name LIKE ? and is_delete = 0 ", "%"+name+"%").Find(&sc)
	if result.Error != nil {
		return nil , result.Error
	}

	return sc, nil
}

func (S *School) Update(db *gorm.DB, id uint, su *School) error {
		result := db.Model(&School{}).Where("id = ?", id).Updates(su)
		if result.RowsAffected <= 0 {
		return fmt.Errorf("school update error")
	}

	return nil
}

func (S *School) Add(db *gorm.DB, su *School) error {
	result := db.Create(su)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (S *School) FindByID(db *gorm.DB, ids []uint) ([]*School, error) {
	var sc []*School
	result := db.Where("id in ?", ids).Find(&sc)
	if result.Error != nil {
		return nil, result.Error
	}

	return sc, nil
}

func (S *School) FindByAccount(db *gorm.DB, acc string) (*School, error) {
	var sc School
	result := db.Where("account = ?", acc).First(&sc)
	if result.Error != nil {
		return nil, result.Error
	}

	return &sc, nil
}


type Region struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	IsDelete int        `gorm:"default:0" json:"is_delete"` // 0 表示未删除，1 表示已删除
	CreateTime consts.CustomTime `gorm:"autoCreateTime" json:"create_time"`
	UpdateTime consts.CustomTime `gorm:"autoUpdateTime" json:"update_time"`
}

func NewRegion() *Region {
	return &Region{}
}

func (r *Region) List(db *gorm.DB) ([]*Region, error) {
	var re []*Region
	result := db.Where("is_delete = 0 ").Find(&re)
	if result.Error != nil {
		return nil , result.Error
	}

	return re, nil
}

func (r *Region) Add(db *gorm.DB, re *Region) error {
	result := db.Create(re)
	if result.Error != nil {
		if mysqlErr, ok := result.Error.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			return fmt.Errorf("校区名称已存在")
		}

		return result.Error
	}

	return nil
}