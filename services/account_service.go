package services

import (
	"gorm.io/gorm"
	"homeschooledu/models"
)

type AccountService struct {
	DB *gorm.DB
}

func NewAccountService(db *gorm.DB) *AccountService{
	return &AccountService{DB:db}
}

func (s *AccountService) Login(acc, pwd string) (*models.Account, error, int) {
	return models.NewAccount().GetAccountByAccAndPwd(s.DB, acc, pwd)
}

func (s *AccountService) Info(acc string) (*models.Account, error) {
	return models.NewAccount().GetAccountByAcc(s.DB, acc)
}