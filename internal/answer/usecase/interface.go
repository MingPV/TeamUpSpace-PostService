package usecase

import "github.com/MingPV/PostService/internal/entities"

type AnswerUseCase interface {
	FindAllAnswers() ([]*entities.Answer, error)
	CreateAnswer(answer *entities.Answer) error
	DeleteAnswer(id int) error
	FindAnswerByID(id int) (*entities.Answer, error)
	FindAllAnswersByPostID(postId int) ([]*entities.Answer, error)
}