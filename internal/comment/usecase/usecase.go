package usecase

import (
	"github.com/MingPV/PostService/internal/entities"
	"github.com/MingPV/PostService/internal/comment/repository"
)

type CommentService struct {
	repo repository.CommentRepository
}

func NewCommentService(repo repository.CommentRepository) CommentUseCase {
	return &CommentService{repo:repo}
}

func (s *CommentService) CreateComment(comment *entities.Comment) error {
	if err := s.repo.Save(comment); err !=nil {
		return err
	}
	
	return nil
}

func (s *CommentService) FindAllComments() ([]*entities.Comment, error) {
	comments, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (s *CommentService) FindCommentByID(id int) (*entities.Comment, error) {
	comment, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (s *CommentService) DeleteComment(id int) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}
	return nil
}

func (s *CommentService) PatchComment(id int, comment *entities.Comment) (*entities.Comment, error) {
	if err := s.repo.Patch(id, comment); err != nil {
		return nil, err
	}
	updatedComment, _ := s.repo.FindByID(id)
	return updatedComment, nil
}

func (s *CommentService) FindCommentsByPostID(postId int) ([]*entities.Comment, error) {
	comments, err := s.repo.FindByPostID(postId)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (s *CommentService) FindCommentsByUserID(userId string) ([]*entities.Comment, error) {
	comments, err := s.repo.FindByUserID(userId)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (s *CommentService) FindCommentsByParentID(parentId int) ([]*entities.Comment, error) {
	comments, err := s.repo.FindByParentID(parentId)
	if err != nil {
		return nil, err
	}
	return comments, nil
}
