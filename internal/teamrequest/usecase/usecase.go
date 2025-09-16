package usecase

import (
	"github.com/MingPV/PostService/internal/entities"
	"github.com/MingPV/PostService/internal/teamrequest/repository"
	"github.com/google/uuid"
)

type TeamRequestService struct {
	repo repository.TeamRequestRepository
}

func NewTeamRequestService(repo repository.TeamRequestRepository) TeamRequestUseCase {
	return &TeamRequestService{repo: repo}
}

func (s *TeamRequestService) CreateTeamRequest(teamRequest *entities.TeamRequest) error {
	return s.repo.Save(teamRequest)
}

func (s *TeamRequestService) FindAllByPost(postId int) ([]*entities.TeamRequest, error) {
	return s.repo.FindAllByPostID(postId)
}

func (s *TeamRequestService) FindAllByRequestBy(userId uuid.UUID) ([]*entities.TeamRequest, error) {
	return s.repo.FindAllByRequestBy(userId)
}

func (s *TeamRequestService) FindTeamRequestByID(postId int, requestBy uuid.UUID) (*entities.TeamRequest, error) {
	return s.repo.FindByID(postId, requestBy)
}

func (s *TeamRequestService) PatchTeamRequest(postId int, requestBy uuid.UUID, update *entities.TeamRequest) (*entities.TeamRequest, error) {
	if err := s.repo.Patch(postId, requestBy, update); err != nil {
		return nil, err
	}

	updatedRequest, err := s.repo.FindByID(postId, requestBy)
	if err != nil {
		return nil, err
	}

	return updatedRequest, nil
}

func (s *TeamRequestService) DeleteTeamRequest(teamRequest *entities.TeamRequest) error {
	return s.repo.Delete(teamRequest.PostId, teamRequest.RequestBy)
}