package usecase

import (
	"github.com/MingPV/PostService/internal/entities"
	"github.com/MingPV/PostService/internal/question/repository"
)

type QuestionService struct {
	repo repository.QuestionRepository
}

func NewQuestionService(repo repository.QuestionRepository) QuestionUseCase {
	return &QuestionService{repo: repo}
}

func (s *QuestionService) CreateQuestion(question *entities.Question) error {
	if err := s.repo.Save(question); err != nil {
		return err
	}
	return nil
}

func (s *QuestionService) FindAllQuestions() ([]*entities.Question, error){
	questions, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return questions, nil
}

func (s *QuestionService) FindQuestionByID(id int) (*entities.Question, error) {
	question, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return question, nil
}

func (s *QuestionService) FindAllQuestionsByPostID(postId int) ([]*entities.Question, error) {
	questions, err := s.repo.FindAllByPostID(postId)
	if err != nil {
		return nil, err
	}
	return questions, nil
}

func (s *QuestionService) DeleteQuestion(id int) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}
	return nil
}

func (s *QuestionService) PatchQuestion(id int, question *entities.Question) (*entities.Question, error) {
	if err := s.repo.Patch(id, question); err != nil {
		return nil, err
	}
	updatedQuestion, _ := s.repo.FindByID(id)

	return updatedQuestion, nil
	
}




