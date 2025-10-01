package repository

import "github.com/MingPV/PostService/internal/entities"

type AnswerRepository interface {
	Save(answer *entities.Answer) error
	FindAll() ([]*entities.Answer, error)
	FindByID(id int) (*entities.Answer, error)
	FindAllByPostID(postId int) ([]*entities.Answer, error)
	FindAllByPostIDAndUserID(postId int, userId string) ([]*entities.Answer, error)
	FindAllByUserID(userId string) ([]*entities.Answer, error)
	Delete(id int) error
}
