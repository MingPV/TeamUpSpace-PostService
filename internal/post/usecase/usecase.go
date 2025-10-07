package usecase

import (
	"github.com/MingPV/PostService/internal/entities"
	"github.com/MingPV/PostService/internal/post/repository"
	"github.com/MingPV/PostService/pkg/mq"
)

type PostService struct {
	repo repository.PostRepository
	mq   mq.MQPublisher
}

func NewPostService(repo repository.PostRepository, mq mq.MQPublisher) PostUseCase {
	return &PostService{repo: repo, mq: mq}
}

func (s *PostService) CreatePost(post *entities.Post) error {
	if err := s.repo.Save(post); err != nil {
		return err
	}
	return nil
}

func (s *PostService) FindAllPosts() ([]*entities.Post, error) {
	posts, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *PostService) FindPostByID(id int) (*entities.Post, error) {
	post, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s *PostService) FindPostsByUserID(userID string) ([]*entities.Post, error) {
	posts, err := s.repo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *PostService) DeletePost(id int) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}
	return nil
}

func (s *PostService) PatchPost(id int, post *entities.Post) (*entities.Post, error) {
	if err := s.repo.Patch(id, post); err != nil {
		return nil, err
	}
	updatedPost, _ := s.repo.FindByID(id)

	return updatedPost, nil

}
