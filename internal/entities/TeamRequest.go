package entities

import (
	"time"

	"github.com/google/uuid"
)

type TeamRequest struct {
	PostId    	int       `gorm:"not null;primaryKey" json:"post_id"`
    RequestBy 	uuid.UUID `gorm:"type:uuid;not null;primaryKey" json:"request_by"`
	RequestTo	uuid.UUID `gorm:"type:uuid;not null" json:"request_to"`
	Status    	string	  `gorm:"size:50;default:'pending'" json:"status"`
	CreatedAt 	time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt 	time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// Relationships
	Post Post `gorm:"foreignKey:PostId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"post"`
}

func (TeamRequest) TableName() string {
	return "team_requests"
}