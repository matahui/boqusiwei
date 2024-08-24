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

func (s *SchoolService) List(offset, limit int) ([]*models.School, error) {
	return  models.NewSchool().List(s.DB, offset, limit)
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

func (s *SchoolService) FindByAccount(acc string) (*models.School, error) {
	return models.NewSchool().FindByAccount(s.DB, acc)
}
type RegionService struct {
	DB *gorm.DB
}

func NewRegionService(db *gorm.DB) *RegionService{
	return &RegionService{DB:db}
}

func (r *RegionService) List() ([]string, error)  {
	var regions []string
	err := r.DB.Model(&models.School{}).
		Distinct("region").
		Where("is_delete = ?", 0).
		Pluck("region", &regions).Error

	if err != nil {
		return nil, err
	}

	return regions, nil
}

func (r *RegionService) Add(re *models.Region) error  {
	return models.NewRegion().Add(r.DB, re)
}