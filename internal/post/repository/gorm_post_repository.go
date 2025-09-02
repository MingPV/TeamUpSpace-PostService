package repository

import (
	"github.com/MingPV/PostService/internal/entities"
	"gorm.io/gorm"
)

type GormPostRepository struct {
	db *gorm.DB
}

func NewGormPostRepository(db *gorm.DB) PostRepository {
	return &GormPostRepository{db:db}
}

func (r *GormPostRepository) Save(post *entities.Post) error {
	return r.db.Create(&post).Error
}

func (r *GormPostRepository) FindAll() ([]*entities.Post, error) {
	var postValues []entities.Post
	if err := r.db.Find(&postValues).Error; err != nil {
		return nil, err
	}

	posts := make([]*entities.Post, len(postValues))
	for i := range postValues {
		posts[i] = &postValues[i]
	}
	return posts, nil
}

func (r *GormPostRepository) FindByID(id int) (*entities.Post, error) {
	var post entities.Post
	if err := r.db.First(&post, id).Error; err != nil {
		return &entities.Post{}, err
	}
	return &post, nil
}

func (r *GormPostRepository) Delete(id int) error {
	result := r.db.Delete(&entities.Post{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *GormPostRepository) Patch(id int, post *entities.Post) error {
	result := r.db.Model(&entities.Post{}).Where("id = ?", id).Updates(post)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

