package repository

import "github.com/MingPV/PostService/internal/entities"

type PostReportRepository interface {
	Save(report *entities.PostReport) error
	FindAll() ([]*entities.PostReport, error)
	FindByID(id int) (*entities.PostReport, error)
	Patch(id int, report *entities.PostReport) error
	Delete(id int) error
}