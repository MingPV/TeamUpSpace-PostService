package repository

import (
	"github.com/MingPV/PostService/internal/entities"
	"github.com/google/uuid"
)

type PostLikeRepository interface {
	Save(postLike *entities.PostLike) error
	FindAllByPostID(postId int) ([]*entities.PostLike, error)
	FindAllByUserID(userId uuid.UUID) ([]*entities.PostLike, error)
	Delete(postlike *entities.PostLike) error
}
