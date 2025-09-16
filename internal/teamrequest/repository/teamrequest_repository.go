package repository

import (
	"github.com/MingPV/PostService/internal/entities"
	"github.com/google/uuid"
)

type TeamRequestRepository interface {
	Save(teamRequest *entities.TeamRequest) error
	FindAllByPostID(postId int) ([]*entities.TeamRequest, error)
	FindAllByRequestBy(userId uuid.UUID) ([]*entities.TeamRequest, error)
	FindByID(postId int, requestBy uuid.UUID) (*entities.TeamRequest, error)
	Patch(postId int, requestBy uuid.UUID, update *entities.TeamRequest) error
	Delete(postId int, requestBy uuid.UUID) error
}
