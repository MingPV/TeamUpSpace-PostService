package dto

import "github.com/google/uuid"

type CreatePostRequest struct {
    PostBy   uuid.UUID `json:"post_by" binding:"required,uuid"`
    Title    string `json:"title" binding:"required"`
    Detail   string `json:"detail"`
    ImageUrl string `json:"image_url"`
    EventId  int    `json:"event_id"`
    Status   string `json:"status"` // optional, default: "active"
}
