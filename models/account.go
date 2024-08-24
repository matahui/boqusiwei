package models

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Account struct {
	ID          uint           `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Account     string         `gorm:"type:varchar(255);not null;unique" json:"account"`
	Password    string         `gorm:"type:varchar(255);not null" json:"password"`
	Cate    int8           `gorm:"type:tinyint;not null" json:"cate"`
	CreateTime  time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime  time.Time      `gorm:"not null;default:CURRENT_TIMESTAMP;autoUpdateTime" json:"update_time"`
	IsDelete    int            `gorm:"not null;default:0" json:"is_delete"`
}

func NewAccount() *Account {
	return &Account{}
}

func (A *Account) GetAccountByAccAndPwd(db *gorm.DB, acc, pwd string) (*Account, error, int) {
	var account Account
	result := db.Where("account = ?", acc).Find(&account)
	if result.Error != nil {
		return nil, fmt.Errorf("账号:%s不存在", acc), -1
	}

	if account.Account == "" {
		return nil, fmt.Errorf("账号:%s不存在", acc), -1
	}

	if account.Account == acc && account.Password == pwd {
		return &account, nil, 0
	} else {
		return nil, fmt.Errorf("账号:%s密码不正确", acc), -2
	}

}

func (A *Account) GetAccountByAcc(db *gorm.DB, acc string) (*Account, error) {
	var account Account
	result := db.Where("account = ?", acc).Find(&account)
	if result.Error != nil {
		return nil, fmt.Errorf("账号:%s不存在", acc)
	}

	return &account, nil
}
