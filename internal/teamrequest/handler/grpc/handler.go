package grpc

import (
	"context"

	"github.com/MingPV/PostService/internal/entities"
	"github.com/MingPV/PostService/internal/teamrequest/usecase"
	"github.com/MingPV/PostService/pkg/apperror"
	teamrequestpb "github.com/MingPV/PostService/proto/teamrequest"
	"github.com/google/uuid"
	"google.golang.org/grpc/status"
)

type GrpcTeamRequestHandler struct {
	teamRequestUseCase usecase.TeamRequestUseCase
	teamrequestpb.UnimplementedTeamRequestServiceServer
}

func NewGrpcTeamRequestHandler(uc usecase.TeamRequestUseCase) *GrpcTeamRequestHandler {
	return &GrpcTeamRequestHandler{teamRequestUseCase: uc}
}

// CreateTeamRequest handles creating a new team request
func (h *GrpcTeamRequestHandler) CreateTeamRequest(ctx context.Context, req *teamrequestpb.CreateTeamRequestRequest) (*teamrequestpb.CreateTeamRequestResponse, error) {
	requestBy, err := uuid.Parse(req.RequestBy)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "invalid request_by UUID")
	}
	requestTo, err := uuid.Parse(req.RequestTo)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "invalid request_to UUID")
	}

	tr := &entities.TeamRequest{
		PostId:    int(req.PostId),
		RequestBy: requestBy,
		RequestTo: requestTo,
		Status:    req.Status,
	}

	if err := h.teamRequestUseCase.CreateTeamRequest(tr); err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}

	return &teamrequestpb.CreateTeamRequestResponse{TeamRequest: toProtoTeamRequest(tr)}, nil
}

// FindTeamRequestByID handles retrieving a specific team request
func (h *GrpcTeamRequestHandler) FindTeamRequestByID(ctx context.Context, req *teamrequestpb.FindTeamRequestByIDRequest) (*teamrequestpb.FindTeamRequestByIDResponse, error) {
	requestBy, err := uuid.Parse(req.RequestBy)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "invalid request_by UUID")
	}

	tr, err := h.teamRequestUseCase.FindTeamRequestByID(int(req.PostId), requestBy)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}

	return &teamrequestpb.FindTeamRequestByIDResponse{TeamRequest: toProtoTeamRequest(tr)}, nil
}

// FindAllByPostID returns all team requests for a post
func (h *GrpcTeamRequestHandler) FindAllByPostID(ctx context.Context, req *teamrequestpb.FindAllByPostIDRequest) (*teamrequestpb.FindAllByPostIDResponse, error) {
	teamRequests, err := h.teamRequestUseCase.FindAllByPost(int(req.PostId))
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}

	var protoTRs []*teamrequestpb.TeamRequest
	for _, tr := range teamRequests {
		protoTRs = append(protoTRs, toProtoTeamRequest(tr))
	}

	return &teamrequestpb.FindAllByPostIDResponse{TeamRequests: protoTRs}, nil
}

// FindAllByRequestBy returns all team requests made by a user
func (h *GrpcTeamRequestHandler) FindAllByRequestBy(ctx context.Context, req *teamrequestpb.FindAllByRequestByRequest) (*teamrequestpb.FindAllByRequestByResponse, error) {
	requestBy, err := uuid.Parse(req.RequestBy)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "invalid request_by UUID")
	}

	teamRequests, err := h.teamRequestUseCase.FindAllByRequestBy(requestBy)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}

	var protoTRs []*teamrequestpb.TeamRequest
	for _, tr := range teamRequests {
		protoTRs = append(protoTRs, toProtoTeamRequest(tr))
	}

	return &teamrequestpb.FindAllByRequestByResponse{TeamRequests: protoTRs}, nil
}

// PatchTeamRequest updates a team request
func (h *GrpcTeamRequestHandler) PatchTeamRequest(ctx context.Context, req *teamrequestpb.PatchTeamRequestRequest) (*teamrequestpb.PatchTeamRequestResponse, error) {
	requestBy, err := uuid.Parse(req.RequestBy)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "invalid request_by UUID")
	}

	tr := &entities.TeamRequest{
		Status: req.Status, // Only patching status
	}

	updatedTR, err := h.teamRequestUseCase.PatchTeamRequest(int(req.PostId), requestBy, tr)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}

	return &teamrequestpb.PatchTeamRequestResponse{TeamRequest: toProtoTeamRequest(updatedTR)}, nil
}

// DeleteTeamRequest deletes a team request
func (h *GrpcTeamRequestHandler) DeleteTeamRequest(ctx context.Context, req *teamrequestpb.DeleteTeamRequestRequest) (*teamrequestpb.DeleteTeamRequestResponse, error) {
	requestBy, err := uuid.Parse(req.RequestBy)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "invalid request_by UUID")
	}

	tr := &entities.TeamRequest{
		PostId:    int(req.PostId),
		RequestBy: requestBy,
	}

	if err := h.teamRequestUseCase.DeleteTeamRequest(tr); err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}

	return &teamrequestpb.DeleteTeamRequestResponse{Message: "team request deleted"}, nil
}

// helper converter
func toProtoTeamRequest(tr *entities.TeamRequest) *teamrequestpb.TeamRequest {
	var requestTo string
	if tr.RequestTo != uuid.Nil {
		requestTo = tr.RequestTo.String()
	}

	return &teamrequestpb.TeamRequest{
		PostId:    int32(tr.PostId),
		RequestBy: tr.RequestBy.String(),
		RequestTo: requestTo,
		Status:    tr.Status,
	}
}
