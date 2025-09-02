package entities

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID 			  int 		`gorm:"primaryKey;autoIncrement" json:"id"`
	PostBy        uuid.UUID `gorm:"type:uuid;not null" json:"post_by"`
	Title         string    `gorm:"size:255;not null" json:"title"`
	Detail        string    `gorm:"type:text" json:"detail"`
	ImageURL      string    `gorm:"size:500" json:"image_url"`
	EventID       int       `json:"event_id"`
	Status        string    `gorm:"size:50;default:'active'" json:"status"`
	CommentsCount int       `gorm:"default:0" json:"comments_count"`
	LikesCount    int       `gorm:"default:0" json:"likes_count"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (Post) TableName() string {
	return "posts"
}