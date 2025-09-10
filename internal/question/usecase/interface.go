package usecase

import "github.com/MingPV/PostService/internal/entities"

type QuestionUseCase interface {
	FindAllQuestions() ([]*entities.Question, error)
	CreateQuestion(question *entities.Question) error
	PatchQuestion(id int, question *entities.Question) (*entities.Question, error)
	DeleteQuestion(id int) error
	FindQuestionByID(id int) (*entities.Question, error)
	FindAllQuestionsByPostID(postId int) ([]*entities.Question, error)
}