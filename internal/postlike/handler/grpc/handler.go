package grpc

import (
	"context"

	postlikepb "github.com/MingPV/PostService/proto/postlike"

	"github.com/MingPV/PostService/internal/entities"
	"github.com/MingPV/PostService/internal/postlike/usecase"
	"github.com/MingPV/PostService/pkg/apperror"
	"github.com/google/uuid"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GrpcPostLikeHandler struct {
	postlikeUseCase usecase.PostLikeUseCase
	postlikepb.UnimplementedPostLikeServiceServer
}

func NewGprcPostLikeHandler(uc usecase.PostLikeUseCase) *GrpcPostLikeHandler {
	return &GrpcPostLikeHandler{postlikeUseCase: uc}
}

func (h *GrpcPostLikeHandler) CreatePostLike(ctx context.Context, req *postlikepb.CreatePostLikeRequest) (*postlikepb.CreatePostLikeResponse, error) {
	likeByUUID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}

	postlike := &entities.PostLike{
		PostId: int(req.PostId),
		UserId: likeByUUID,
	}

	if err := h.postlikeUseCase.CreatePostLike(postlike); err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &postlikepb.CreatePostLikeResponse{Postlike: toProtoPostLike(postlike)}, nil
}

func (h *GrpcPostLikeHandler) FindAllPostLikesByPostID(ctx context.Context, req *postlikepb.FindAllPostLikesByPostIDRequest) (*postlikepb.FindAllPostLikesByPostIDResponse, error) {
	postlikes, err := h.postlikeUseCase.FindAllPostLikesByPostID(int(req.PostId))
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}

	var protoPostLikes []*postlikepb.PostLike
	for _,p := range postlikes {
		protoPostLikes = append(protoPostLikes, toProtoPostLike(p))
	}

	return &postlikepb.FindAllPostLikesByPostIDResponse{Postlikes: protoPostLikes}, nil
} 

func (h *GrpcPostLikeHandler) FindAllPostLikesByUserID(ctx context.Context, req *postlikepb.FindAllPostLikesByUserIDRequest) (*postlikepb.FindAllPostLikesByUserIDResponse, error) {
	likeByUUID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	postlikes, err := h.postlikeUseCase.FindAllPostLikesByUserID(likeByUUID)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}

	var protoPostLikes []*postlikepb.PostLike
	for _,p := range postlikes {
		protoPostLikes = append(protoPostLikes, toProtoPostLike(p))
	}

	return &postlikepb.FindAllPostLikesByUserIDResponse{Postlikes: protoPostLikes}, nil
} 

func (h *GrpcPostLikeHandler) DeletePostLike(ctx context.Context, req *postlikepb.DeletePostLikeRequest) (*postlikepb.DeletePostLikeResponse, error) {
	likeByUUID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}

	postlike := &entities.PostLike{
		PostId: int(req.PostId),
		UserId: likeByUUID,
	}
	if err := h.postlikeUseCase.DeletePostLike(postlike); err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &postlikepb.DeletePostLikeResponse{Message: "postlike deleted"}, nil
}


func toProtoPostLike(pl *entities.PostLike) *postlikepb.PostLike {
	return &postlikepb.PostLike{
		PostId:	int32(pl.PostId),
		UserId: pl.UserId.String(),
		CreatedAt: timestamppb.New(pl.CreatedAt),
	}
}