package usecase

import "github.com/MingPV/PostService/internal/entities"

type PostReportUseCase interface {
	FindAllPostReports() ([]*entities.PostReport, error)
	CreatePostReport(report *entities.PostReport) error
	PatchPostReport(id int, report *entities.PostReport) (*entities.PostReport, error)
	DeletePostReport(id int) error
	FindPostReportByID(id int) (*entities.PostReport, error)
}
