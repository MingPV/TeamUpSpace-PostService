package entities

import (
	"time"

	"github.com/google/uuid"
)

type PostReport struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	PostID    int       `gorm:"not null" json:"post_id"`
	Reporter  uuid.UUID `gorm:"type:uuid;not null" json:"reporter"`
	Report_to uuid.UUID `gorm:"type:uuid;not null" json:"report_to"`
	Detail    string    `gorm:"size:255;not null" json:"detail"`
	Status    string    `gorm:"size:50;default:'pending'" json:"status"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// Relationships
	Post Post `gorm:"foreignKey:PostID;" json:"post"`
}

func (PostReport) TableName() string {
	return "post_reports"
}
