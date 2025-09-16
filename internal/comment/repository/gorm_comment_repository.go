package repository

import (
	"github.com/MingPV/PostService/internal/entities"
	"gorm.io/gorm"
)

type GormCommentRepository struct {
	db *gorm.DB
}

func NewGormCommentRepository(db *gorm.DB) CommentRepository {
	return &GormCommentRepository{db:db}
}

func (r *GormCommentRepository) Save(comment *entities.Comment) error {
	return r.db.Create(&comment).Error
}

func (r *GormCommentRepository) FindAll() ([]*entities.Comment, error) {
	var commentValues []entities.Comment
	if err := r.db.Find(&commentValues).Error; err != nil {
		return nil, err
	}

	comments := make([]*entities.Comment, len(commentValues))
	for i := range commentValues {
		comments[i] = &commentValues[i]
	}
	return comments, nil
}

func (r *GormCommentRepository) FindByID(id int) (*entities.Comment, error) {
	var comment entities.Comment
	if err := r.db.First(&comment, id).Error; err != nil {
		return &entities.Comment{}, err
	}
	return &comment, nil
}

func (r *GormCommentRepository) FindByPostID(postId int) ([]*entities.Comment, error) {
	var commentValues []entities.Comment
	if err := r.db.Where("post_id = ?", postId).Find(&commentValues).Error; err != nil {
		return nil, err
	}

	comments := make([]*entities.Comment, len(commentValues))
	for i := range commentValues {
		comments[i] = &commentValues[i]
	}
	return comments, nil
}

func (r *GormCommentRepository) FindByUserID(userId string) ([]*entities.Comment, error) {
	var commentValues []entities.Comment
	if err := r.db.Where("comment_by = ?", userId).Find(&commentValues).Error; err != nil {
		return nil, err
	}

	comments := make([]*entities.Comment, len(commentValues))
	for i := range commentValues {
		comments[i] = &commentValues[i]
	}
	return comments, nil
}

func (r *GormCommentRepository) FindByParentID(parentId int) ([]*entities.Comment, error) {
	var commentValues []entities.Comment
	if err := r.db.Where("parent_id = ?", parentId).Find(&commentValues).Error; err != nil {
		return nil, err
	}

	comments := make([]*entities.Comment, len(commentValues))
	for i := range commentValues {
		comments[i] = &commentValues[i]
	}
	return comments, nil
}

func (r *GormCommentRepository) Delete(id int) error {
	result := r.db.Delete(&entities.Comment{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *GormCommentRepository) Patch(id int, comment *entities.Comment) error {
	result := r.db.Model(&entities.Comment{}).Where("id = ?", id).Updates(comment)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

