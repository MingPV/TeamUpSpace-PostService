package entities

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID 			int 		`gorm:"primaryKey;autoIncrement" json:"id"`
	PostId    	int       	`gorm:"not null" json:"post_id"`
	CommentBy	uuid.UUID 	`gorm:"type:uuid;not null" json:"comment_by"`
	ParentId    int      	`gorm:"default:0" json:"parent_id"`
	Detail    	string    	`gorm:"size:255;not null" json:"detail"`
	CreatedAt 	time.Time 	`gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt 	time.Time 	`gorm:"autoUpdateTime" json:"updated_at"`

	// Relationships
	Post Post `gorm:"foreignKey:PostId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"post"`
}

func (Comment) TableName() string {
	return "comments"
}