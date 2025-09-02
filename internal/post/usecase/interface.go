package usecase

import "github.com/MingPV/PostService/internal/entities"

type PostUseCase interface {
	FindAllPosts() ([]*entities.Post, error)
	CreatePost(post *entities.Post) error
	PatchPost(id int, post *entities.Post) (*entities.Post, error)
	DeletePost(id int) error
	FindPostByID(id int) (*entities.Post, error)
}