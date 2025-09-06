package repository

import "github.com/MingPV/PostService/internal/entities"

type QuestionRepository interface {
	Save(question *entities.Question) error
	FindAll() ([]*entities.Question, error)
	FindByID(id int) (*entities.Question, error)
	FindAllByPostID(postId int) ([]*entities.Question, error)
	Patch(id int, question *entities.Question) error
	Delete(id int) error

}