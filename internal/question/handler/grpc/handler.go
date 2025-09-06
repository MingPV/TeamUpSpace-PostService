package grpc

import (
	"context"

	questionpb "github.com/MingPV/PostService/proto/question"

	"github.com/MingPV/PostService/internal/entities"
	"github.com/MingPV/PostService/internal/question/usecase"
	"github.com/MingPV/PostService/pkg/apperror"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GrpcQuestionHandler struct {
	questionUseCase usecase.QuestionUseCase
	questionpb.UnimplementedQuestionServiceServer
}

func NewGrpcQuestionHandler(uc usecase.QuestionUseCase) *GrpcQuestionHandler {
	return &GrpcQuestionHandler{questionUseCase: uc}
}

func (h *GrpcQuestionHandler) CreateQuestion(ctx context.Context, req *questionpb.CreateQuestionRequest) (*questionpb.CreateQuestionResponse, error) {
	question := &entities.Question{
		PostId: int(req.PostId),
		Question: req.Question,
	}

	if err := h.questionUseCase.CreateQuestion(question); err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &questionpb.CreateQuestionResponse{Question: toProtoQuestion(question)}, nil
}

func (h *GrpcQuestionHandler) FindQuestionByID(ctx context.Context, req *questionpb.FindQuestionByIDRequest) (*questionpb.FindQuestionByIDResponse, error) {
	question, err := h.questionUseCase.FindQuestionByID(int(req.Id));
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &questionpb.FindQuestionByIDResponse{Question: toProtoQuestion(question)}, nil
}

func (h *GrpcQuestionHandler) FindAllQuestionsByPostID(ctx context.Context, req *questionpb.FindAllQuestionsByPostIDRequest) (*questionpb.FindAllQuestionsByPostIDResponse, error) {
	questions, err := h.questionUseCase.FindAllQuestionsByPostID(int(req.PostId))
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}

	var protoQuestions []*questionpb.Question
	for _, o := range questions {
		protoQuestions = append(protoQuestions, toProtoQuestion(o))
	}

	return &questionpb.FindAllQuestionsByPostIDResponse{Questions: protoQuestions}, nil
}

func (h *GrpcQuestionHandler) FindAllQuestions(ctx context.Context, req *questionpb.FindAllQuestionsRequest) (*questionpb.FindAllQuestionsResponse, error) {
	questions, err := h.questionUseCase.FindAllQuestions()
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}

	var protoQuestions []*questionpb.Question
	for _, o := range questions {
		protoQuestions = append(protoQuestions, toProtoQuestion(o))
	}

	return &questionpb.FindAllQuestionsResponse{Questions: protoQuestions}, nil
}

func (h *GrpcQuestionHandler) PatchQuestion(ctx context.Context, req *questionpb.PatchQuestionRequest) (*questionpb.PatchQuestionResponse, error) {
	question := &entities.Question{
		PostId: int(req.PostId),
		Question: req.Question,
	}

	updatedQuestion, err := h.questionUseCase.PatchQuestion(int(req.Id), question)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &questionpb.PatchQuestionResponse{Question: toProtoQuestion(updatedQuestion)}, nil
}

func (h *GrpcQuestionHandler) DeleteQuestion(ctx context.Context, req *questionpb.DeleteQuestionRequest) (*questionpb.DeleteQuestionResponse, error) {
	if err := h.questionUseCase.DeleteQuestion(int(req.Id)); err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &questionpb.DeleteQuestionResponse{Message: "question deleted"}, nil
}

func toProtoQuestion(q *entities.Question) *questionpb.Question {
	return &questionpb.Question{
		Id: int32(q.ID),
		PostId: int32(q.PostId),
		Question: q.Question,
		CreatedAt: timestamppb.New(q.CreatedAt),
		UpdatedAt: timestamppb.New(q.UpdatedAt),
	}
}