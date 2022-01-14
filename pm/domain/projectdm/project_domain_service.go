package projectdm

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/apperrors"
)

type projectDomainService struct {
	projectRepository ProjectRepository
}

func NewProjectDomainService(projectRepository ProjectRepository) *projectDomainService {
	return &projectDomainService{
		projectRepository: projectRepository,
	}
}

func (s *projectDomainService) ExistsProjectForCreate(ctx context.Context, projectDm *Project) (bool, error) {
	existingProjectDm, err := s.projectRepository.FetchProjectByGroupIDAndKeyName(ctx, projectDm.GroupID(), projectDm.KeyName())
	if err != nil && !apperrors.Is(err, apperrors.NotFound) {
		return false, err
	} else if existingProjectDm != nil {
		return true, nil
	}

	existingProjectDm, err = s.projectRepository.FetchProjectByGroupIDAndName(ctx, projectDm.GroupID(), projectDm.Name())
	if existingProjectDm != nil {
		return true, nil
	}

	return false, err
}

// TODO: シンプルなロジックにリファクタ
func (s *projectDomainService) ExistUniqueProjectForUpdate(ctx context.Context, projectDm *Project) (bool, error) {
	oldProjectDm, err := s.projectRepository.FetchProjectByID(ctx, projectDm.ID(), projectDm.productID)
	if err != nil {
		return false, apperrors.InternalServerError
	}

	if projectDm.KeyName().Equals(oldProjectDm.KeyName()) && projectDm.Name().Equals(oldProjectDm.Name()) {
		return false, nil
	}

	projectDmByKeyName, errByKeyName := s.projectRepository.FetchProjectByGroupIDAndKeyName(ctx, projectDm.GroupID(), projectDm.KeyName())
	if errByKeyName != nil && !apperrors.Is(errByKeyName, apperrors.NotFound) {
		return false, errByKeyName
	}

	projectDmByName, errByName := s.projectRepository.FetchProjectByGroupIDAndName(ctx, projectDm.GroupID(), projectDm.Name())
	if errByName != nil && !apperrors.Is(errByName, apperrors.NotFound) {
		return false, errByName
	}

	if apperrors.Is(errByKeyName, apperrors.NotFound) && apperrors.Is(errByName, apperrors.NotFound) {
		return false, apperrors.NotFound
	}

	if projectDmByKeyName != nil {
		if projectDm.KeyName().Equals(oldProjectDm.KeyName()) && apperrors.Is(errByName, apperrors.NotFound) {
			return false, apperrors.NotFound
		}

		return true, nil
	}

	if projectDmByName != nil {
		if projectDm.Name().Equals(oldProjectDm.Name()) && apperrors.Is(errByKeyName, apperrors.NotFound) {
			return false, apperrors.NotFound
		}

		return true, nil
	}

	return false, apperrors.NotFound
}
