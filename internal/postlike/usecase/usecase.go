package usecase

import (
	"log"

	"github.com/MingPV/PostService/internal/entities"
	postrepo "github.com/MingPV/PostService/internal/post/repository"
	"github.com/MingPV/PostService/internal/postlike/repository"
	"github.com/MingPV/PostService/pkg/mq"
	"github.com/google/uuid"
)

type PostLikeService struct {
	repo     repository.PostLikeRepository
	postrepo postrepo.PostRepository
	mq       mq.MQPublisher
}

func NewPostLikeService(repo repository.PostLikeRepository, postrepo postrepo.PostRepository, mq mq.MQPublisher) PostLikeUseCase {
	return &PostLikeService{repo: repo, postrepo: postrepo, mq: mq}
}

func (s *PostLikeService) CreatePostLike(postlike *entities.PostLike) error {
	if err := s.repo.Save(postlike); err != nil {
		return err
	}

	// find who is post's owner
	post, err := s.postrepo.FindByID(postlike.PostId)
	if err != nil {
		return err
	}

	// check if post's owner is same as postlike's user
	if post.PostBy == postlike.UserId {
		// if same, do not send event
		return nil
	}

	// publish event to message broker (RabbitMQ)
	mqevent := entities.PostLikeCreatedEvent{
		PostId:      postlike.PostId,
		PostOwnerId: post.PostBy,
		UserId:      postlike.UserId,
	}

	err = s.mq.Publish("PostLikeCreated", mqevent)
	if err != nil {
		log.Println("Failed to publish event:", err)
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
