package dto

import (
	"github.com/MingPV/PostService/internal/entities"
)

func ToCommentResponse(comment *entities.Comment) *CommentResponse {

    return &CommentResponse{
        ID:            comment.ID,
        PostId:        comment.PostId,
        CommentBy:     comment.CommentBy,
        ParentId:      comment.ParentId,
        Detail:        comment.Detail,
        CreatedAt:     comment.CreatedAt,
        UpdatedAt:     comment.UpdatedAt,
    }
}

func ToCommentResponseList(comments []*entities.Comment) []*CommentResponse {
    result := make([]*CommentResponse, 0, len(comments))
    for _, c := range comments {
        result = append(result, ToCommentResponse(c))
    }   
    return result
}