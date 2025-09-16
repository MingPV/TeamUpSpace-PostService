package grpc

import (
	"context"

	commentpb "github.com/MingPV/PostService/proto/comment"

	"github.com/MingPV/PostService/internal/comment/usecase"
	"github.com/MingPV/PostService/internal/entities"
	"github.com/MingPV/PostService/pkg/apperror"
	"github.com/google/uuid"
	"google.golang.org/grpc/status"
)

type GrpcCommentHandler struct {
	commentUseCase usecase.CommentUseCase
	commentpb.UnimplementedCommentServiceServer
}

func NewGrpcCommentHandler(uc usecase.CommentUseCase) *GrpcCommentHandler {
	return &GrpcCommentHandler{commentUseCase: uc}
}

func (h *GrpcCommentHandler) CreateComment(ctx context.Context, req *commentpb.CreateCommentRequest) (*commentpb.CreateCommentResponse, error) {
	commentByUUID, err := uuid.Parse(req.CommentBy)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}

	comment := &entities.Comment{
		PostId:    int(req.PostId),
		CommentBy: commentByUUID,
		ParentId:  int(req.ParentId),
		Detail:    req.Detail,
	}

	if err := h.commentUseCase.CreateComment(comment); err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}

	return &commentpb.CreateCommentResponse{Comment: toProtoComment(comment)},nil
}

func (h *GrpcCommentHandler) FindCommentByID(ctx context.Context, req *commentpb.FindCommentByIDRequest) (*commentpb.FindCommentByIDResponse, error) {
	comment, err := h.commentUseCase.FindCommentByID(int(req.Id))
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &commentpb.FindCommentByIDResponse{Comment: toProtoComment(comment)}, nil
}

func (h *GrpcCommentHandler) FindAllComments(ctx context.Context, req *commentpb.FindAllCommentsRequest) (*commentpb.FindAllCommentsResponse, error) {
	comments, err := h.commentUseCase.FindAllComments()
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}

	var protoComments []*commentpb.Comment
	for _, o := range comments {
		protoComments = append(protoComments, toProtoComment(o))
	}

	return &commentpb.FindAllCommentsResponse{Comments: protoComments}, nil
}

func (h *GrpcCommentHandler) FindCommentsByPostID(ctx context.Context, req *commentpb.FindCommentsByPostIDRequest) (*commentpb.FindCommentsByPostIDResponse, error) {
	comments, err := h.commentUseCase.FindCommentsByPostID(int(req.PostId))
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}

	var protoComments []*commentpb.Comment
	for _, o := range comments {
		protoComments = append(protoComments, toProtoComment(o))
	}

	return &commentpb.FindCommentsByPostIDResponse{Comments: protoComments}, nil
}

func (h *GrpcCommentHandler) FindCommentsByParentID(ctx context.Context, req *commentpb.FindCommentsByParentIDRequest) (*commentpb.FindCommentsByParentIDResponse, error) {
	comments, err := h.commentUseCase.FindCommentsByParentID(int(req.ParentId))
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}

	var protoComments []*commentpb.Comment
	for _, o := range comments {
		protoComments = append(protoComments, toProtoComment(o))
	}

	return &commentpb.FindCommentsByParentIDResponse{Comments: protoComments}, nil
}

func (h *GrpcCommentHandler) FindCommentsByUserID(ctx context.Context, req *commentpb.FindCommentsByUserIDRequest) (*commentpb.FindCommentsByUserIDResponse, error) {
	comments, err := h.commentUseCase.FindCommentsByUserID(req.UserId)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}

	var protoComments []*commentpb.Comment
	for _, o := range comments {
		protoComments = append(protoComments, toProtoComment(o))
	}

	return &commentpb.FindCommentsByUserIDResponse{Comments: protoComments}, nil
}

func (h *GrpcCommentHandler) PatchComment(ctx context.Context, req *commentpb.PatchCommentRequest) (*commentpb.PatchCommentResponse, error) {
	comment := &entities.Comment{
		Detail: req.Detail,
	}
	updatedComment, err := h.commentUseCase.PatchComment(int(req.Id), comment)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &commentpb.PatchCommentResponse{Comment: toProtoComment(updatedComment)}, nil
}

func (h *GrpcCommentHandler) DeleteComment(ctx context.Context, req *commentpb.DeleteCommentRequest) (*commentpb.DeleteCommentResponse, error) {
	if err := h.commentUseCase.DeleteComment(int(req.Id)); err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &commentpb.DeleteCommentResponse{Message: "comment deleted"}, nil
}

func toProtoComment(c *entities.Comment) *commentpb.Comment {

	pbComment := &commentpb.Comment{
        CommentId: int32(c.ID),
        PostId:    int32(c.PostId),
        CommentBy: c.CommentBy.String(),
		ParentId:  int32(c.ParentId),
        Detail:    c.Detail,
        CreatedAt: c.CreatedAt.Format("2006-01-02 15:04:05"),
        UpdatedAt: c.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return pbComment
}