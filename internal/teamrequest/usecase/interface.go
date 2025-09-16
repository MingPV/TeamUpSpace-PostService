package usecase

import (
	"github.com/MingPV/PostService/internal/entities"
	"github.com/google/uuid"
)

type TeamRequestUseCase interface {
	CreateTeamRequest(teamRequest *entities.TeamRequest) error
	FindAllByPost(postId int) ([]*entities.TeamRequest, error)
	FindAllByRequestBy(userId uuid.UUID) ([]*entities.TeamRequest, error)
	FindTeamRequestByID(postId int, requestBy uuid.UUID) (*entities.TeamRequest, error)
	PatchTeamRequest(postId int, requestBy uuid.UUID, update *entities.TeamRequest) (*entities.TeamRequest, error)
	DeleteTeamRequest(teamRequest *entities.TeamRequest) error
}
