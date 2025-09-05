package entities

import (
	"time"
)

type Question struct {
	ID 			int 		`gorm:"primaryKey;autoIncrement" json:"id"`
	PostID    	int       	`gorm:"not null" json:"post_id"`
	Question 	string		`gorm:"size:255;not null" json:"question"`
	CreatedAt 	time.Time 	`gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt 	time.Time 	`gorm:"autoUpdateTime" json:"updated_at"`

	// Relationships
	Post Post `gorm:"foreignKey:PostID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"post"`
}

func (Question) TableName() string {
	return "questions"
}