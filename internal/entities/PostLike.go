package entities

import (
	"time"

	"github.com/google/uuid"
)

type PostLike struct {
	PostID    int       `gorm:"primaryKey" json:"post_id"`
	UserID    uuid.UUID `gorm:"type:uuid;primaryKey" json:"user_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	// Relationships
	Post Post `gorm:"foreignKey:PostID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"post"`
}

func (PostLike) TableName() string {
	return "post_likes"
}
