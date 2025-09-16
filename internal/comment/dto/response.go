package dto

import (
	"time"

	"github.com/google/uuid"
)

type CommentResponse struct {
    ID            int       `json:"id"`
    PostId        int       `json:"post_id"`
    CommentBy     uuid.UUID `json:"comment_by"`
    ParentId      int       `json:"parent_id"`
    Detail        string    `json:"detail"`
    CreatedAt     time.Time `json:"created_at"`
    UpdatedAt     time.Time `json:"updated_at"`
}
