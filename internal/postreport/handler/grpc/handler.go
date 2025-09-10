package grpc

import (
	"context"

	postreportpb "github.com/MingPV/PostService/proto/postreport"

	"github.com/MingPV/PostService/internal/entities"
	"github.com/MingPV/PostService/internal/postreport/usecase"
	"github.com/MingPV/PostService/pkg/apperror"
	"github.com/google/uuid"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GrpcPostReportHandler struct {
	postReportUseCase usecase.PostReportUseCase
	postreportpb.UnimplementedPostReportServiceServer
}

func NewGrpcPostReportHandler(uc usecase.PostReportUseCase) *GrpcPostReportHandler {
	return &GrpcPostReportHandler{postReportUseCase: uc}
}

func (h *GrpcPostReportHandler) CreatePostReport(ctx context.Context, req *postreportpb.CreatePostReportRequest) (*postreportpb.CreatePostReportResponse, error) {
	reporterUUID, err := uuid.Parse(req.Reporter)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	reportToUUID, err := uuid.Parse(req.ReportTo)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}

	report := &entities.PostReport{
		PostId:    int(req.PostId),
		Reporter:  reporterUUID,
		Report_to: reportToUUID,
		Detail:    req.Detail,
		Status:    req.Status,
	}

	if err := h.postReportUseCase.CreatePostReport(report); err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}

	return &postreportpb.CreatePostReportResponse{PostReport: toProtoPostReport(report)}, nil
}

func (h *GrpcPostReportHandler) FindPostReportByID(ctx context.Context, req *postreportpb.FindPostReportByIDRequest) (*postreportpb.FindPostReportByIDResponse, error) {
	report, err := h.postReportUseCase.FindPostReportByID(int(req.Id))
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &postreportpb.FindPostReportByIDResponse{PostReport: toProtoPostReport(report)}, nil
}

func (h *GrpcPostReportHandler) FindAllPostReports(ctx context.Context, req *postreportpb.FindAllPostReportsRequest) (*postreportpb.FindAllPostReportsResponse, error) {
	reports, err := h.postReportUseCase.FindAllPostReports()
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}

	var protoReports []*postreportpb.PostReport
	for _, r := range reports {
		protoReports = append(protoReports, toProtoPostReport(r))
	}

	return &postreportpb.FindAllPostReportsResponse{PostReports: protoReports}, nil
}

func (h *GrpcPostReportHandler) PatchPostReport(ctx context.Context, req *postreportpb.PatchPostReportRequest) (*postreportpb.PatchPostReportResponse, error) {
	reporterUUID, err := uuid.Parse(req.Reporter)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	reportToUUID, err := uuid.Parse(req.ReportTo)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}

	report := &entities.PostReport{
		PostId:    int(req.PostId),
		Reporter:  reporterUUID,
		Report_to: reportToUUID,
		Detail:    req.Detail,
		Status:    req.Status,
	}

	updatedReport, err := h.postReportUseCase.PatchPostReport(int(req.Id), report)
	if err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}

	return &postreportpb.PatchPostReportResponse{PostReport: toProtoPostReport(updatedReport)}, nil
}

func (h *GrpcPostReportHandler) DeletePostReport(ctx context.Context, req *postreportpb.DeletePostReportRequest) (*postreportpb.DeletePostReportResponse, error) {
	if err := h.postReportUseCase.DeletePostReport(int(req.Id)); err != nil {
		return nil, status.Errorf(apperror.GRPCCode(err), "%s", err.Error())
	}
	return &postreportpb.DeletePostReportResponse{Message: "post report deleted"}, nil
}

func toProtoPostReport(r *entities.PostReport) *postreportpb.PostReport {
	return &postreportpb.PostReport{
		Id:        int32(r.ID),
		PostId:    int32(r.PostId),
		Reporter:  r.Reporter.String(),
		ReportTo:  r.Report_to.String(),
		Detail:    r.Detail,
		Status:    r.Status,
		CreatedAt: timestamppb.New(r.CreatedAt),
		UpdatedAt: timestamppb.New(r.UpdatedAt),
	}
}
