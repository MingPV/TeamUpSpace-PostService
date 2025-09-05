package entities

import (
	"time"

	"github.com/google/uuid"
)

type Answer struct {
	ID 			int 		`gorm:"primaryKey;autoIncrement" json:"id"`
	PostID    	int       	`gorm:"not null" json:"post_id"`
	UserID    	uuid.UUID 	`gorm:"type:uuid;" json:"user_id"`
	Question 	string		`gorm:"size:255;not null" json:"question"`
	Answer 		string		`gorm:"size:255;not null" json:"answer"`
	CreatedAt 	time.Time 	`gorm:"autoCreateTime" json:"created_at"`

	// Relationships
	Post Post `gorm:"foreignKey:PostID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"post"`
}

func (Answer) TableName() string {
	return "answers"
}