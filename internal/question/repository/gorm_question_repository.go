package repository

import (
	"github.com/MingPV/PostService/internal/entities"
	"gorm.io/gorm"
)

type GormQuestionRepository struct {
	db *gorm.DB
}

func NewGormQuestionRepository(db *gorm.DB) QuestionRepository {
	return &GormQuestionRepository{db: db}
}

func (r *GormQuestionRepository) Save(question *entities.Question) error {
	return r.db.Create(&question).Error
}

func (r *GormQuestionRepository) FindAll() ([]*entities.Question, error) {
	var questionValues []entities.Question
	if err := r.db.Find(&questionValues).Error; err != nil {
		return nil, err
	}

	questions := make([]*entities.Question, len(questionValues))
	for i := range questionValues {
		questions[i] = &questionValues[i]
	}
	return questions, nil
}

func (r *GormQuestionRepository) FindByID(id int) (*entities.Question, error) {
	var question entities.Question
	if err := r.db.First(&question, id).Error; err != nil {
		return &entities.Question{}, err
	}
	return &question, nil
}

func (r *GormQuestionRepository) FindAllByPostID(postId int) ([]*entities.Question, error) {
	var questionValues []entities.Question
	if err := r.db.Where("post_id = ?", postId).Find(&questionValues).Error; err != nil {
		return nil, err
	}

	questions := make([]*entities.Question, len(questionValues))
	for i := range questionValues {
		questions[i] = &questionValues[i]
	}
	return questions, nil
}

func (r *GormQuestionRepository) Delete(id int) error {
	result := r.db.Delete(&entities.Question{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *GormQuestionRepository) Patch(id int, question *entities.Question) error {
	result := r.db.Model(&entities.Question{}).Where("id = ?", id).Updates(question)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

