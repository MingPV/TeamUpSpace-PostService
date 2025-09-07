package usecase

import (
	"github.com/MingPV/PostService/internal/entities"
	"github.com/MingPV/PostService/internal/postlike/repository"
	"github.com/google/uuid"
)

type PostLikeService struct {
	repo repository.PostLikeRepository
}

func NewPostLikeService(repo repository.PostLikeRepository) PostLikeUseCase {
	return &PostLikeService{repo: repo}
}

func (s *PostLikeService) CreatePostLike(postlike *entities.PostLike) error {
	if err := s.repo.Save(postlike); err !=nil {
		return err
	}
	return nil
}

func (s *PostLikeService) FindAllPostLikesByPostID(postId int) ([]*entities.PostLike, error) {
	postlikes, err := s.repo.FindAllByPostID(postId)
	if err != nil {
		return nil, err
	}
	return postlikes, nil
}

func (s *PostLikeService) FindAllPostLikesByUserID(userId uuid.UUID) ([]*entities.PostLike, error) {
	postlikes, err := s.repo.FindAllByUserID(userId)
	if err != nil {
		return nil, err
	}
	return postlikes, nil
}

func (s *PostLikeService) DeletePostLike(postlike *entities.PostLike) error {
	if err := s.repo.Delete(postlike); err != nil {
		return err
	}
	return nil
}