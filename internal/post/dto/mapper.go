package dto

import "github.com/MingPV/PostService/internal/entities"

func ToPostResponse(post *entities.Post) *PostResponse {
    return &PostResponse{
        ID:            post.ID,
        PostBy:        post.PostBy,
        Title:         post.Title,
        Detail:        post.Detail,
        ImageUrl:      post.ImageUrl,
        EventId:       post.EventId,
        Status:        post.Status,
        CommentsCount: post.CommentsCount,
        LikesCount:    post.LikesCount,
        CreatedAt:     post.CreatedAt,
        UpdatedAt:     post.UpdatedAt,
    }
}

func ToPostResponseList(posts []*entities.Post) []*PostResponse {
    result := make([]*PostResponse, 0, len(posts))
	for _, p := range posts {
		result = append(result, ToPostResponse(p))
	}
	return result
}