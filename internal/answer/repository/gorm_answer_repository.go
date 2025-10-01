package repository

import (
	"github.com/MingPV/PostService/internal/entities"
	"gorm.io/gorm"
)

type GormAnswerRepository struct {
	db *gorm.DB
}

func NewGormAnswerRepository(db *gorm.DB) AnswerRepository {
	return &GormAnswerRepository{db: db}
}

func (r *GormAnswerRepository) Save(answer *entities.Answer) error {
	return r.db.Create(&answer).Error
}

func (r *GormAnswerRepository) FindAll() ([]*entities.Answer, error) {
	var answerValues []entities.Answer
	if err := r.db.Find(&answerValues).Error; err != nil {
		return nil, err
	}

	answers := make([]*entities.Answer, len(answerValues))
	for i := range answerValues {
		answers[i] = &answerValues[i]
	}
	return answers, nil
}

func (r *GormAnswerRepository) FindByID(id int) (*entities.Answer, error) {
	var answer entities.Answer
	if err := r.db.First(&answer, id).Error; err != nil {
		return &entities.Answer{}, err
	}
	return &answer, nil
}

func (r *GormAnswerRepository) FindAllByPostID(postId int) ([]*entities.Answer, error) {
	var answerValues []entities.Answer
	if err := r.db.Where("post_id = ?", postId).Find(&answerValues).Error; err != nil {
		return nil, err
	}

	answers := make([]*entities.Answer, len(answerValues))
	for i := range answerValues {
		answers[i] = &answerValues[i]
	}
	return answers, nil
}

func (r *GormAnswerRepository) FindAllByPostIDAndUserID(postId int, userId string) ([]*entities.Answer, error) {
	var answerValues []entities.Answer
	if err := r.db.Where("post_id = ? AND user_id = ?", postId, userId).Find(&answerValues).Error; err != nil {
		return nil, err
	}

	answers := make([]*entities.Answer, len(answerValues))
	for i := range answerValues {
		answers[i] = &answerValues[i]
	}
	return answers, nil
}

func (r *GormAnswerRepository) FindAllByUserID(userId string) ([]*entities.Answer, error) {
	var answerValues []entities.Answer
	if err := r.db.Where("user_id = ?", userId).Find(&answerValues).Error; err != nil {
		return nil, err
	}

	answers := make([]*entities.Answer, len(answerValues))
	for i := range answerValues {
		answers[i] = &answerValues[i]
	}
	return answers, nil
}

func (r *GormAnswerRepository) Delete(id int) error {
	result := r.db.Delete(&entities.Answer{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
