package repository

import "github.com/MingPV/PostService/internal/entities"

type PostRepository interface {
	Save(post *entities.Post) error
	FindAll() ([]*entities.Post, error)
	FindByID(id int) (*entities.Post, error)
	Patch(id int, post *entities.Post) error
	Delete(id int) error
}
