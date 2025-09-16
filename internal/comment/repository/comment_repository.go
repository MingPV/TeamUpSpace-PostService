package repository

import "github.com/MingPV/PostService/internal/entities"

type CommentRepository interface {
	Save(comment *entities.Comment) error
	FindAll() ([]*entities.Comment, error)
	FindByID(id int) (*entities.Comment, error)
	FindByPostID(postId int) ([]*entities.Comment, error)
	FindByUserID(userId string) ([]*entities.Comment, error)
	FindByParentID(parentId int) ([]*entities.Comment, error)
	Patch(id int, comment *entities.Comment) error
	Delete(id int) error
}
