package models

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"homeschooledu/consts"
	"time"
)

type Teacher struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	LoginNumber int64    `gorm:"type:int64;not null;unique" json:"login_number"`
	Password   string   `gorm:"type:varchar(255);not null" json:"password"`
	TeacherName string    `gorm:"type:varchar(255);not null" json:"teacher_name"`
	PhoneNumber string    `gorm:"type:varchar(255);not null" json:"phone_number"`
	SchoolID  uint      `gorm:"not null" json:"school_id"`
	Role      uint      `gorm:"not null" json:"role"`
	IsDelete int        `gorm:"default:0" json:"is_delete"` // 0 表示未删除，1 表示已删除
	CreateTime consts.CustomTime `gorm:"autoCreateTime" json:"create_time"`
	UpdateTime consts.CustomTime `gorm:"autoUpdateTime" json:"update_time"`
}

type TeacherShow struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CustomID  string    `json:"custom_id"`
	Password  string    `json:"password"`
	LoginNumber int64    `gorm:"type:int64;not null;unique" json:"login_number"`
	TeacherName string    `gorm:"type:varchar(255);not null" json:"teacher_name"`
	PhoneNumber string    `gorm:"type:varchar(255);not null" json:"phone_number"`
	SchoolID  uint      `gorm:"not null" json:"school_id"`
	SchoolName string   `json:"school_name"`
	Role      uint      `gorm:"not null" json:"role"`
	RoleName  string    `json:"role_name"`
	TeachingClass []*TeachClass `json:"teaching_class"`
	IsDelete int        `gorm:"default:0" json:"is_delete"` // 0 表示未删除，1 表示已删除
	CreateTime consts.CustomTime `gorm:"autoCreateTime" json:"create_time"`
	UpdateTime consts.CustomTime `gorm:"autoUpdateTime" json:"update_time"`
}

type TeachClass struct {
	ID uint `json:"id"`
	Name string `json:"name"`
}

func NewTeacher() *Teacher {
	return &Teacher{}
}

func (S *Teacher)Info(db *gorm.DB, id uint) (*Teacher, error) {
	var st Teacher
	result := db.Where("id = ?", id).First(&st)
	if result.Error != nil {
		return nil, result.Error
	}

	return &st, nil
}

func (S *Teacher) List(db *gorm.DB, offset, limit int, schoolID, classID uint, name string) ([]*Teacher, int64, int64, error) {
	var (
		sc []*Teacher
		total int64
		page int64
	)

	query := db.Model(&Teacher{}).Where("is_delete = 0 ")

	if schoolID  > 0  {
		query = query.Where("school_id = ?", schoolID)
	}


	if name != "" {
		query = query.Where("teacher_name LIKE ?", "%"+name+"%")
	}

	//总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, 0, err
	}

	query = query.Order("create_time desc")
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

func (S *Teacher) Update(db *gorm.DB, id uint, st *Teacher) error {
	result := db.Model(&Teacher{}).Where("id = ?", id).Updates(st)
	if result.RowsAffected <= 0 {
		return fmt.Errorf("Teacher update error")
	}

	return nil
}

func (S *Teacher) Del(db *gorm.DB, id uint) error {
	return db.Unscoped().Where("id = ?", id).Delete(&Teacher{}).Error
}

func (S *Teacher) Add(db *gorm.DB, st *Teacher) error {
	result := db.Create(st)
	if result.Error != nil {
		if result.Error == gorm.ErrDuplicatedKey {
			return fmt.Errorf("登录账号已存在")
		}
	}

	return nil
}

func (S *Teacher) BatchInsert(db *gorm.DB, sts []*Teacher) (int, error) {
	result := db.Create(sts)
	if result.Error != nil {
		if mysqlErr, ok := result.Error.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			for _, teacher := range sts {
				// 检查单个记录是否已经存在
				var existingTeacher Teacher
				if db.Where("login_number = ?", teacher.LoginNumber).First(&existingTeacher).Error == nil {
					return 0, fmt.Errorf("导入失败，账号%d已存在", teacher.LoginNumber)
				}
			}
		}
	}

	return int(result.RowsAffected), nil
}

func (S *Teacher) FindClassInfoByT(db *gorm.DB, tid uint) ([]*Class, error)  {
	var ci []*Class
	err := db.Table("teacher_class_assignments").
		Select("classes.id as id, classes.class_name as class_name, classes.school_id").
		Joins("join classes on teacher_class_assignments.class_id = classes.id").
		Where("teacher_class_assignments.teacher_id = ? and teacher_class_assignments.is_delete = 0", tid).
		Scan(&ci).Error

	if err != nil {
		return nil, err
	}

	return ci, nil
}


type TeacherClassAssignment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	TeacherID uint    `gorm:"type:int64;not null;unique" json:"teacher_id"`
	ClassID uint    `gorm:"type:int64;not null" json:"class_id"`
	IsDelete int        `gorm:"default:0" json:"is_delete"` // 0 表示未删除，1 表示已删除
	CreateTime time.Time `gorm:"autoCreateTime" json:"create_time"`
	UpdateTime time.Time `gorm:"autoUpdateTime" json:"update_time"`
}

func NewTeacherClassAssignment() *TeacherClassAssignment {
	return &TeacherClassAssignment{}
}

func (S *TeacherClassAssignment) Add(db *gorm.DB, tc []*TeacherClassAssignment) (error) {
	result := db.Create(tc)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (S *TeacherClassAssignment) Info(db *gorm.DB, id uint) (*TeacherClassAssignment, error) {
	var st TeacherClassAssignment
	result := db.Where("id = ?", id).First(&st)
	if result.Error != nil {
		return nil, result.Error
	}

	return &st, nil
}

func (S *TeacherClassAssignment) List(db *gorm.DB, teacherId uint) ([]*TeacherClassAssignment, error) {
	var(
		st []*TeacherClassAssignment
		total int64
	)

	query := db.Model(&Teacher{}).Where("is_delete = ? and teacher_id = ?", 0, teacherId)


	result := query.Count(&total).Find(&st)
	if result.Error != nil {
		return nil, result.Error
	}

	return st, nil
}

