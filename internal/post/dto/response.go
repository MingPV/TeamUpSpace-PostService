package dto

import (
	"time"

	"github.com/google/uuid"
)
type PostResponse struct {
    ID            int       `json:"id"`
    PostBy        uuid.UUID    `json:"post_by"`
    Title         string    `json:"title"`
    Detail        string    `json:"detail"`
    ImageUrl      string    `json:"image_url"`
    EventId       int       `json:"event_id"`
    Status        string    `json:"status"`
    CommentsCount int       `json:"comments_count"`
    LikesCount    int       `json:"likes_count"`
    CreatedAt     time.Time `json:"created_at"`
    UpdatedAt     time.Time `json:"updated_at"`
}
