package usecase

import (
	"log"

	"github.com/MingPV/PostService/internal/comment/repository"
	"github.com/MingPV/PostService/internal/entities"
	postrepo "github.com/MingPV/PostService/internal/post/repository"
	"github.com/MingPV/PostService/pkg/mq"
)

type CommentService struct {
	repo     repository.CommentRepository
	postrepo postrepo.PostRepository
	mq       mq.MQPublisher
}

func NewCommentService(repo repository.CommentRepository, postrepo postrepo.PostRepository, mq mq.MQPublisher) CommentUseCase {
	return &CommentService{repo: repo, postrepo: postrepo, mq: mq}
}

func (s *CommentService) CreateComment(comment *entities.Comment) error {
	if err := s.repo.Save(comment); err != nil {
		return err
	}

	// find who is post's owner
	post, err := s.postrepo.FindByID(comment.PostId)
	if err != nil {
		return err
	}

	// check if post's owner is same as comment's user
	if post.PostBy == comment.CommentBy {
		// if same, do not send event
		return nil
	}

	mqevent := entities.CommentCreatedEvent{
		ID:          comment.ID,
		PostId:      comment.PostId,
		PostOwnerId: post.PostBy,
		CommentBy:   comment.CommentBy,
		ParentId:    comment.ParentId,
		Detail:      comment.Detail,
		CreatedAt:   comment.CreatedAt,
		UpdatedAt:   comment.UpdatedAt,
	}

	err = s.mq.Publish("CommentCreated", mqevent)
	if err != nil {
		log.Println("Failed to publish event:", err)
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
