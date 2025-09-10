package usecase

import (
	"github.com/MingPV/PostService/internal/entities"
	"github.com/MingPV/PostService/internal/postreport/repository"
)

type PostReportService struct {
	repo repository.PostReportRepository
}

func NewPostReportService(repo repository.PostReportRepository) PostReportUseCase {
	return &PostReportService{repo: repo}
}

func (s *PostReportService) CreatePostReport(report *entities.PostReport) error {
	if err := s.repo.Save(report); err != nil {
		return err
	}
	return nil
}

func (s *PostReportService) FindAllPostReports() ([]*entities.PostReport, error) {
	reports, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return reports, nil
}

func (s *PostReportService) FindPostReportByID(id int) (*entities.PostReport, error) {
	report, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return report, nil
}

func (s *PostReportService) DeletePostReport(id int) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}
	return nil
}

func (s *PostReportService) PatchPostReport(id int, report *entities.PostReport) (*entities.PostReport, error) {
	if err := s.repo.Patch(id, report); err != nil {
		return nil, err
	}
	updatedReport, _ := s.repo.FindByID(id)
	return updatedReport, nil
}