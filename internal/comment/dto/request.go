package dto

import (
    "github.com/google/uuid"
)

type CreateCommentRequest struct {
    PostId    int       `json:"post_id" binding:"required"`
    CommentBy uuid.UUID `json:"comment_by" binding:"required,uuid"`
    ParentId  int       `json:"parent_id" binding:"omitempty,gt=0"`
    Detail    string    `json:"detail" binding:"required"`
}
