package repository

import (
	"github.com/MingPV/PostService/internal/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormPostLikeRepository struct {
	db *gorm.DB
}

func NewGormPostLikeRepository(db *gorm.DB) PostLikeRepository {
	return  &GormPostLikeRepository{db: db}
}

func (r *GormPostLikeRepository) Save(postLike *entities.PostLike) error {
	return r.db.Create(&postLike).Error
}

func (r *GormPostLikeRepository) FindAllByPostID(postId int) ([]*entities.PostLike, error) {
	var postLikeValues []entities.PostLike
	if err := r.db.Where("post_id = ?", postId).Find(&postLikeValues).Error; err != nil {
		return nil, err
	}

	postlikes := make([]*entities.PostLike, len(postLikeValues))
	for i := range postLikeValues {
		postlikes[i] = &postLikeValues[i]
	}
	return postlikes, nil
}

func (r *GormPostLikeRepository) FindAllByUserID(userId uuid.UUID) ([]*entities.PostLike, error) {
	var postLikeValues []entities.PostLike
	if err := r.db.Where("user_id = ?", userId).Find(&postLikeValues).Error; err != nil {
		return nil, err
	}

	postlikes := make([]*entities.PostLike, len(postLikeValues))
	for i := range postLikeValues {
		postlikes[i] = &postLikeValues[i]
	}
	return postlikes, nil
}

func (r *GormPostLikeRepository) Delete(postlike *entities.PostLike) error {
	result := r.db.Where("post_id = ? AND user_id = ?", postlike.PostId, postlike.UserId).Delete(&entities.PostLike{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}



