package usecase

import (
	"github.com/MingPV/PostService/internal/entities"
	"github.com/google/uuid"
)
type PostLikeUseCase interface {
	CreatePostLike(postlike *entities.PostLike) error
	FindAllPostLikesByPostID(postId int) ([]*entities.PostLike, error)
	FindAllPostLikesByUserID(userId uuid.UUID) ([]*entities.PostLike, error)
	DeletePostLike(postlike *entities.PostLike) error
}