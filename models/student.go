package models

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"homeschooledu/consts"
)


type Student struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	LoginNumber int64    `gorm:"type:int64;not null;unique" json:"login_number"`
	Password    string   `gorm:"type:varchar(255);not null" json:"password"`
	StudentName string    `gorm:"type:varchar(255);not null" json:"student_name"`
	ParentName  string    `gorm:"type:varchar(255)" json:"parent_name"`
	PhoneNumber string    `gorm:"type:varchar(20)" json:"phone_number"`
	ClassID     uint      `gorm:"not null" json:"class_id"`
	SchoolID    uint     `gorm:"not null" json:"school_id"`
	IsDelete int        `gorm:"default:0" json:"is_delete"` // 0 表示未删除，1 表示已删除
	CreateTime consts.CustomTime `gorm:"autoCreateTime" json:"create_time"`
	UpdateTime consts.CustomTime `gorm:"autoUpdateTime" json:"update_time"`
}

type StudentShow struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	StudentCode string    `gorm:"type:varchar(20);not null;unique" json:"student_code"`
	LoginNumber int64    `gorm:"type:int64;not null;unique" json:"login_number"`
	Password    string   `gorm:"type:varchar(255);not null" json:"password"`
	StudentName string    `gorm:"type:varchar(255);not null" json:"student_name"`
	ParentName  string    `gorm:"type:varchar(255)" json:"parent_name"`
	PhoneNumber string    `gorm:"type:varchar(20)" json:"phone_number"`
	ClassID     uint      `gorm:"not null" json:"class_id"`
	SchoolID    uint     `gorm:"not null" json:"school_id"`
	ClassName   string   `json:"class_name"`
	SchoolName  string   `json:"school_name"`
	IsDelete int        `gorm:"default:0" json:"is_delete"` // 0 表示未删除，1 表示已删除
	CreateTime consts.CustomTime `gorm:"autoCreateTime" json:"create_time"`
	UpdateTime consts.CustomTime `gorm:"autoUpdateTime" json:"update_time"`
}

type StudentRanking struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	LoginNumber int64    `gorm:"type:int64;not null;unique" json:"login_number"`
	StudentName string `json:"student_name"`
	Avatar      string `json:"avatar"`
	Points      int    `json:"points"`
	Ranks        int    `json:"ranks"`
}

func NewStudent() *Student {
	return &Student{}
}

func (S *Student)Info(db *gorm.DB, id uint) (*Student, error) {
	var st Student
	result := db.Where("id = ?", id).First(&st)
	if result.Error != nil {
		return nil, result.Error
	}

	return &st, nil
}

func (S *Student)GetByLoginNumber(db *gorm.DB, ln int64) (*Student, error) {
	var st Student
	result := db.Where("login_number = ?", ln).First(&st)
	if result.Error != nil {
		return nil, result.Error
	}

	return &st, nil
}

func (S *Student) List(db *gorm.DB, offset, limit int, schoolID, classID uint, name string) ([]*Student, int64, int64, error) {
	var (
		sc []*Student
		total int64
		page int64
	)

	query := db.Model(&Student{}).Where("is_delete = 0 ")

	if schoolID  > 0  {
		query = query.Where("school_id = ?", schoolID)
	}

	if classID > 0 {
		query = query.Where("class_id = ?", classID)
	}

	if name != "" {
		query = query.Where("student_name LIKE ?", "%"+name+"%")
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

func (S *Student) Update(db *gorm.DB, id uint, st *Student) error {
	result := db.Model(&Student{}).Where("id = ?", id).Updates(st)
	if result.RowsAffected <= 0 {
		return fmt.Errorf("student update error")
	}

	return nil
}

func (S *Student) Add(db *gorm.DB, st *Student) error {
	result := db.Create(st)
	if result.Error != nil {
		if result.Error == gorm.ErrDuplicatedKey {
			return fmt.Errorf("登录账号已存在login_number:%d", st.LoginNumber)
		}
		return fmt.Errorf("student add error")
	}

	return nil
}


func (S *Student) Del(db *gorm.DB, id uint) error {
	return db.Unscoped().Where("id = ?", id).Delete(&Student{}).Error
}

func (S *Student) BatchInsert(db *gorm.DB, sts []*Student) (int, error) {
	result := db.Create(sts)
	if result.Error != nil {
		if mysqlErr, ok := result.Error.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			for _, student := range sts {
				// 检查单个记录是否已经存在
				var existingStudent Student
				if db.Where("login_number = ?", student.LoginNumber).First(&existingStudent).Error == nil {
					return 0, fmt.Errorf("插入失败，存在重复的 login_number:%d", student.LoginNumber)
				}
			}
		}

		return 0, result.Error
	}

	return int(result.RowsAffected), nil
}



type StudentPoint struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	StudentID int64    `gorm:"type:int64;not null;unique" json:"student_id"`
	Points    int64     `gorm:"not null" json:"points"`
	IsDelete int        `gorm:"default:0" json:"is_delete"` // 0 表示未删除，1 表示已删除
	CreateTime consts.CustomTime `gorm:"autoCreateTime" json:"create_time"`
	UpdateTime consts.CustomTime `gorm:"autoUpdateTime" json:"update_time"`
}

type StudentPointShow struct {
	StudentID uint    `gorm:"type:int64;not null;unique" json:"student_id"`
	Name      string   `json:"name"`
	Points    int64     `gorm:"not null" json:"points"`
	Stars    int64    `json:"stars"`
	Moons    int64    `json:"moons"`
	Suns     int64    `json:"suns"`
	Apples   int64    `json:"apples"`
	ClassID  uint   `json:"class_id"`
	ClassName string `json:"class_name"`
}


func NewStudentPoint() *StudentPoint {
	return &StudentPoint{}
}

func (S *StudentPoint)Info(db *gorm.DB, studentID uint) (*StudentPoint, error) {
	var st StudentPoint
	result := db.Model(&StudentPoint{}).Where("student_id = ?", studentID).Find(&st)
	if result.Error != nil {
		return nil, result.Error
	}

	return &st, nil
}


type ActivityLog struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	StudentID uint    `gorm:"not null" json:"student_id"`
	ResourceID    uint     `gorm:"not null" json:"resource_id"`
	ActivityDate string `gorm:"not null" json:"activity_date"`
	PointsAward  int    `gorm:"not null" json:"points_award"`
	IsDelete int        `gorm:"default:0" json:"is_delete"` // 0 表示未删除，1 表示已删除
	CreateTime consts.CustomTime `gorm:"autoCreateTime" json:"create_time"`
	UpdateTime consts.CustomTime `gorm:"autoUpdateTime" json:"update_time"`
}

