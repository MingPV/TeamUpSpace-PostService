package repository

import (
	"github.com/MingPV/PostService/internal/entities"
	"gorm.io/gorm"
)

type GormPostReportRepository struct {
	db *gorm.DB
}

func NewGormPostReportRepository(db *gorm.DB) PostReportRepository {
	return &GormPostReportRepository{db: db}
}

func (r *GormPostReportRepository) Save(report *entities.PostReport) error {
	return r.db.Create(&report).Error
}

func (r *GormPostReportRepository) FindAll() ([]*entities.PostReport, error) {
	var reportValues []entities.PostReport
	if err := r.db.Find(&reportValues).Error; err != nil {
		return nil, err
	}

	reports := make([]*entities.PostReport, len(reportValues))
	for i := range reportValues {
		reports[i] = &reportValues[i]
	}
	return reports, nil
}

func (r *GormPostReportRepository) FindByID(id int) (*entities.PostReport, error) {
	var report entities.PostReport
	if err := r.db.First(&report, id).Error; err != nil {
		return &entities.PostReport{}, err
	}
	return &report, nil
}

func (r *GormPostReportRepository) Delete(id int) error {
	result := r.db.Delete(&entities.PostReport{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *GormPostReportRepository) Patch(id int, report *entities.PostReport) error {
	result := r.db.Model(&entities.PostReport{}).Where("id = ?", id).Updates(report)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
