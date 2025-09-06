package handler

import (
	"context"

	answerpb "github.com/MingPV/PostService/proto/answer"

	"github.com/MingPV/PostService/internal/answer/usecase"
	"github.com/MingPV/PostService/internal/entities"
	"github.com/MingPV/PostService/pkg/apperror"
	"github.com/google/uuid"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GrpcAnswerHandler struct {
	answerUseCase usecase.AnswerUseCase
	answerpb.UnimplementedAnswerServiceServer
}

func NewGrpcAnswerHandler(uc usecase.AnswerUseCase) *GrpcAnswerHandler {
	return &GrpcAnswerHandler{answerUseCase: uc}
}

func (h *GrpcAnswerHandler) CreateAnswer(ctx context.Context, req *answerpb.CreateAnswerRequest) (*answerpb.CreateAnswerResponse, error) {
	answerByUUID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	answer := &entities.Answer{
		PostId: int(req.PostId),
		UserId: answerByUUID,
		Question: req.Question,
		Answer: req.Answer,
	}

	if err := h.answerUseCase.CreateAnswer(answer); err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &answerpb.CreateAnswerResponse{Answer: toProtoAnswer(answer)}, nil
}

func (h *GrpcAnswerHandler) FindAnswerByID(ctx context.Context, req *answerpb.FindAnswerByIDRequest) (*answerpb.FindAnswerByIDResponse, error) {
	answer, err := h.answerUseCase.FindAnswerByID(int(req.Id));
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &answerpb.FindAnswerByIDResponse{Answer: toProtoAnswer(answer)}, nil
}

func (h *GrpcAnswerHandler) FindAllAnswersByPostID(ctx context.Context, req *answerpb.FindAllAnswersByPostIDRequest) (*answerpb.FindAllAnswersByPostIDResponse, error) {
	answers, err := h.answerUseCase.FindAllAnswersByPostID(int(req.PostId))
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}

	var protoAnswers []*answerpb.Answer
	for _, o := range answers {
		protoAnswers = append(protoAnswers, toProtoAnswer(o))
	}

	return &answerpb.FindAllAnswersByPostIDResponse{Answers: protoAnswers}, nil
}

func (h *GrpcAnswerHandler) FindAllAnswers(ctx context.Context, req *answerpb.FindAllAnswersRequest) (*answerpb.FindAllAnswersResponse, error) {
	answers, err := h.answerUseCase.FindAllAnswers()
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}

	var protoAnswers []*answerpb.Answer
	for _, o := range answers {
		protoAnswers = append(protoAnswers, toProtoAnswer(o))
	}

	return &answerpb.FindAllAnswersResponse{Answers: protoAnswers}, nil
}

func (h *GrpcAnswerHandler) DeleteAnswer(ctx context.Context, req *answerpb.DeleteAnswerRequest) (*answerpb.DeleteAnswerResponse, error) {
	if err := h.answerUseCase.DeleteAnswer(int(req.Id)); err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &answerpb.DeleteAnswerResponse{Message: "Answer deleted"}, nil
}

func toProtoAnswer(a *entities.Answer) *answerpb.Answer {
	return &answerpb.Answer{
		Id: int32(a.ID),
		PostId: int32(a.PostId),
		UserId: a.UserId.String(),
		Question: a.Question,
		Answer: a.Answer,
		CreatedAt: timestamppb.New(a.CreatedAt),
	}
}