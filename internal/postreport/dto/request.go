package dto

import "github.com/google/uuid"

type CreatePostReportRequest struct {
    PostId   int       `json:"post_id" binding:"required"`
    Reporter uuid.UUID `json:"reporter" binding:"required,uuid"`
    ReportTo uuid.UUID `json:"report_to" binding:"required,uuid"`
    Detail   string    `json:"detail" binding:"required"`
    Status   string    `json:"status"` // optional, default: "pending"
}