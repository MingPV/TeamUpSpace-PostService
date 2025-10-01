package grpc

import (
	"context"

	postpb "github.com/MingPV/PostService/proto/post"

	"github.com/MingPV/PostService/internal/entities"
	"github.com/MingPV/PostService/internal/post/usecase"
	"github.com/MingPV/PostService/pkg/apperror"
	"github.com/google/uuid"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GrpcPostHandler struct {
	postUseCase usecase.PostUseCase
	postpb.UnimplementedPostServiceServer
}

func NewGrpcPostHandler(uc usecase.PostUseCase) *GrpcPostHandler {
	return &GrpcPostHandler{postUseCase: uc}
}

func (h *GrpcPostHandler) CreatePost(ctx context.Context, req *postpb.CreatePostRequest) (*postpb.CreatePostResponse, error) {
	postByUUID, err := uuid.Parse(req.PostBy)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	post := &entities.Post{
		PostBy:   postByUUID,
		Title:    req.Title,
		Detail:   req.Detail,
		ImageUrl: req.ImageUrl,
		EventId:  int(req.EventId),
		Status:   req.Status,
	}

	if err := h.postUseCase.CreatePost(post); err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &postpb.CreatePostResponse{Post: toProtoPost(post)}, nil
}

func (h *GrpcPostHandler) FindPostByID(ctx context.Context, req *postpb.FindPostByIDRequest) (*postpb.FindPostByIDResponse, error) {
	post, err := h.postUseCase.FindPostByID((int(req.Id)))
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &postpb.FindPostByIDResponse{Post: toProtoPost(post)}, nil
}

func (h *GrpcPostHandler) FindAllPosts(ctx context.Context, req *postpb.FindAllPostsRequest) (*postpb.FindAllPostsResponse, error) {
	posts, err := h.postUseCase.FindAllPosts()
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}

	var protoPosts []*postpb.Post
	for _, o := range posts {
		protoPosts = append(protoPosts, toProtoPost(o))
	}

	return &postpb.FindAllPostsResponse{Posts: protoPosts}, nil
}

func (h *GrpcPostHandler) PatchPost(ctx context.Context, req *postpb.PatchPostRequest) (*postpb.PatchPostResponse, error) {
	// postByUUID, err := uuid.Parse(req.PostBy)
	// if err != nil {
	// 	return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	// }
	post := &entities.Post{
		PostBy:        uuid.Nil,
		Title:         req.Title,
		Detail:        req.Detail,
		ImageUrl:      req.ImageUrl,
		EventId:       int(req.EventId),
		Status:        req.Status,
		CommentsCount: int(req.CommentsCount),
		LikesCount:    int(req.LikesCount),
	}
	updatedpost, err := h.postUseCase.PatchPost(int(req.Id), post)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &postpb.PatchPostResponse{Post: toProtoPost(updatedpost)}, nil
}

func (h *GrpcPostHandler) DeletePost(ctx context.Context, req *postpb.DeletePostRequest) (*postpb.DeletePostResponse, error) {
	if err := h.postUseCase.DeletePost(int(req.Id)); err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &postpb.DeletePostResponse{Message: "post deleted"}, nil
}

func toProtoPost(p *entities.Post) *postpb.Post {
	return &postpb.Post{
		Id:            int32(p.ID),
		PostBy:        p.PostBy.String(),
		Title:         p.Title,
		Detail:        p.Detail,
		ImageUrl:      p.ImageUrl,
		EventId:       int32(p.EventId),
		Status:        p.Status,
		CommentsCount: int32(p.CommentsCount),
		LikesCount:    int32(p.LikesCount),
		CreatedAt:     timestamppb.New(p.CreatedAt),
		UpdatedAt:     timestamppb.New(p.UpdatedAt),
	}
}
