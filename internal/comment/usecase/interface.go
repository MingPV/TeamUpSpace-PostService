package usecase

import "github.com/MingPV/PostService/internal/entities"

type CommentUseCase interface {
	FindAllComments() ([]*entities.Comment, error)
	CreateComment(comment *entities.Comment) error
	PatchComment(id int, comment *entities.Comment) (*entities.Comment, error)
	DeleteComment(id int) error
	FindCommentByID(id int) (*entities.Comment, error)
	FindCommentsByPostID(postId int) ([]*entities.Comment, error)
	FindCommentsByUserID(userId string) ([]*entities.Comment, error)
	FindCommentsByParentID(parentId int) ([]*entities.Comment, error)
}