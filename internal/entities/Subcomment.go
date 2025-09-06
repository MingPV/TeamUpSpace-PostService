package entities

import (
	"time"

	"github.com/google/uuid"
)

type Subcomment struct {
	ID			int     	`gorm:"primaryKey;autoIncrement" json:"id"`
	CommentID	int 		`gorm:"not null" json:"comment_id"`
	CommentBy	uuid.UUID 	`gorm:"type:uuid;not null" json:"comment_by"`
	Detail    	string    	`gorm:"size:255;not null" json:"detail"`
	CreatedAt 	time.Time 	`gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt 	time.Time 	`gorm:"autoUpdateTime" json:"updated_at"`

	// Relationships
	Comment Comment 	`gorm:"foreignKey:CommentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"comment"`
}

func (Subcomment) TableName() string {
	return "subcomments"
}