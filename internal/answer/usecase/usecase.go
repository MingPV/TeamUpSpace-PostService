package usecase

import (
	"github.com/MingPV/PostService/internal/answer/repository"
	"github.com/MingPV/PostService/internal/entities"
)

type AnswerService struct {
	repo repository.AnswerRepository
}

func NewAnswerService(repo repository.AnswerRepository) AnswerUseCase {
	return &AnswerService{repo: repo}
}

func (s *AnswerService) CreateAnswer(answer *entities.Answer) error {
	if err := s.repo.Save(answer); err != nil {
		return err
	}
	return nil
}

func (s *AnswerService) FindAllAnswers() ([]*entities.Answer, error) {
	answers, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return answers, nil
}

func (s *AnswerService) FindAnswerByID(id int) (*entities.Answer, error) {
	answer, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return answer, nil
}

func (s *AnswerService) FindAllAnswersByPostID(postId int) ([]*entities.Answer, error) {
	answers, err := s.repo.FindAllByPostID(postId)
	if err != nil {
		return nil, err
	}
	return answers, nil
}

func (s *AnswerService) FindAllAnswerByPostIDAndUserID(postId int, userId string) ([]*entities.Answer, error) {
	answers, err := s.repo.FindAllByPostIDAndUserID(postId, userId)
	if err != nil {
		return nil, err
	}
	return answers, nil
}

func (s *AnswerService) FindAllAnswerByUserID(userId string) ([]*entities.Answer, error) {
	answers, err := s.repo.FindAllByUserID(userId)
	if err != nil {
		return nil, err
	}
	return answers, nil
}

func (s *AnswerService) DeleteAnswer(id int) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}
	return nil
}
