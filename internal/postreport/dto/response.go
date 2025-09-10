package dto

import (
	"time"

	"github.com/google/uuid"
)

type PostReportResponse struct {
	ID        int       `json:"id"`
	PostId    int       `json:"post_id"`
	Reporter  uuid.UUID `json:"reporter"`
	ReportTo  uuid.UUID `json:"report_to"`
	Detail    string    `json:"detail"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
