package dto

import "github.com/google/uuid"

type CreatePostRequest struct {
    PostBy   uuid.UUID `json:"post_by" binding:"required,uuid"`
    Title    string `json:"title" binding:"required"`
    Detail   string `json:"detail"`
    ImageURL string `json:"image_url"`
    EventID  int    `json:"event_id"`
    Status   string `json:"status"` // optional, default: "active"
}
