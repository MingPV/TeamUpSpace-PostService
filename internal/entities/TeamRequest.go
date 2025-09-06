package entities

import (
	"time"

	"github.com/google/uuid"
)

type TeamRequest struct {
	ID			int       `gorm:"primaryKey;autoIncrement" json:"id"`
	PostID    	int       `gorm:"not null" json:"post_id"`
	RequestTo	uuid.UUID `gorm:"type:uuid;not null" json:"request_to"`
	RequestBy	uuid.UUID `gorm:"type:uuid;not null" json:"request_by"`
	IsAccept    bool	  `gorm:"default:false" json:"is_accept"`
	CreatedAt 	time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt 	time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// Relationships
	Post Post `gorm:"foreignKey:PostID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"post"`
}

func (TeamRequest) TableName() string {
	return "team_requests"
}