package services

import (
	"gorm.io/gorm"
	"homeschooledu/models"
)

type SchoolService struct {
	DB *gorm.DB
}

func NewSchoolService(db *gorm.DB) *SchoolService{
	return &SchoolService{DB:db}
}

type SchoolListResp struct {
	School []*models.School `json:"school"`
	Total  int64  `json:"total"`
	Page   int64  `json:"page"`
}

func (s *SchoolService) List(offset, limit int, name string) (*SchoolListResp, error) {
	list, total, page, err := models.NewSchool().List(s.DB, offset, limit, name)
	if err != nil {
		return nil, err
	}

	return &SchoolListResp{
		School: list,
		Total:  total,
		Page:   page,
	}, nil
}

func (s *SchoolService) FuzzySearch(name string) ([]*models.School, error)  {
	return models.NewSchool().QueryByName(s.DB, name)
}

func (s *SchoolService) Update(su *models.School, id uint) error  {
	return models.NewSchool().Update(s.DB, id, su)
}

func (s *SchoolService) Add(su *models.School) error  {
	return models.NewSchool().Add(s.DB, su)
}

func (s *SchoolService) FindByID(ids []uint) (map[uint]*models.School, error)  {
	sc, err := models.NewSchool().FindByID(s.DB, ids)
	if err != nil {
		return nil, err
	}

	r := make(map[uint]*models.School)
	for _, v := range sc {
		r[v.ID] = v
	}

	return r, nil
}

func (s *SchoolService) FindByAccount(acc string) (*SchoolListResp, error) {
	sc, err := models.NewSchool().FindByAccount(s.DB, acc)
	if err != nil {
		return nil, err
	}

	return &SchoolListResp{
		School: []*models.School{sc},
		Total:  0,
		Page:   0,
	}, nil
}

func (s *SchoolService) FindByName(region, name string) (*models.School, error) {
	var sc models.School
	err := s.DB.Model(&models.School{}).Where("region = ? and name = ? and is_delete = 0", region, name).Find(&sc).Error
	if err != nil {
		return nil, err
	}

	return &sc, nil
}

type RegionService struct {
	DB *gorm.DB
}

func NewRegionService(db *gorm.DB) *RegionService{
	return &RegionService{DB:db}
}

func (r *RegionService) List() ([]models.Region, error)  {
	var re []models.Region
	err := r.DB.Model(&models.Region{}).Where("is_delete = ? and name != ?", 0, "").Find(&re).Error

	if err != nil {
		return nil, err
	}

	return re, nil
}

func (r *RegionService) Add(re *models.Region) error  {
	return models.NewRegion().Add(r.DB, re)
}