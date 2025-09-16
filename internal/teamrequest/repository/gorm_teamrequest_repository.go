package repository

import (
	"github.com/MingPV/PostService/internal/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormTeamRequestRepository struct {
	db *gorm.DB
}

func NewGormTeamRequestRepository(db *gorm.DB) TeamRequestRepository {
	return &GormTeamRequestRepository{db: db}
}

func (r *GormTeamRequestRepository) Save(teamRequest *entities.TeamRequest) error {
	return r.db.Create(teamRequest).Error
}

func (r *GormTeamRequestRepository) FindAllByPostID(postId int) ([]*entities.TeamRequest, error) {
	var values []entities.TeamRequest
	if err := r.db.Where("post_id = ?", postId).Find(&values).Error; err != nil {
		return nil, err
	}
	requests := make([]*entities.TeamRequest, len(values))
	for i := range values {
		requests[i] = &values[i]
	}
	return requests, nil
}

func (r *GormTeamRequestRepository) FindAllByRequestBy(userId uuid.UUID) ([]*entities.TeamRequest, error) {
	var values []entities.TeamRequest
	if err := r.db.Where("request_by = ?", userId).Find(&values).Error; err != nil {
		return nil, err
	}
	requests := make([]*entities.TeamRequest, len(values))
	for i := range values {
		requests[i] = &values[i]
	}
	return requests, nil
}

func (r *GormTeamRequestRepository) FindByID(postId int, requestBy uuid.UUID) (*entities.TeamRequest, error) {
	var teamRequest entities.TeamRequest
	if err := r.db.Where("post_id = ? AND request_by = ?", postId, requestBy).First(&teamRequest).Error; err != nil {
		return nil, err
	}
	return &teamRequest, nil
}

func (r *GormTeamRequestRepository) Patch(postId int, requestBy uuid.UUID, update *entities.TeamRequest) error {
	result := r.db.Model(&entities.TeamRequest{}).
		Where("post_id = ? AND request_by = ?", postId, requestBy).
		Updates(update)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *GormTeamRequestRepository) Delete(postId int, requestBy uuid.UUID) error {
	result := r.db.Where("post_id = ? AND request_by = ?", postId, requestBy).Delete(&entities.TeamRequest{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
