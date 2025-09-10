package dto

import "github.com/MingPV/PostService/internal/entities"

func ToPostReportResponse(report *entities.PostReport) *PostReportResponse {
	return &PostReportResponse{
		ID:        report.ID,
		PostId:    report.PostId,
		Reporter:  report.Reporter,
		ReportTo:  report.Report_to,
		Detail:    report.Detail,
		Status:    report.Status,
		CreatedAt: report.CreatedAt,
		UpdatedAt: report.UpdatedAt,
	}
}

func ToPostReportResponseList(reports []*entities.PostReport) []*PostReportResponse {
	result := make([]*PostReportResponse, 0, len(reports))
	for _, r := range reports {
		result = append(result, ToPostReportResponse(r))
	}
	return result
}
