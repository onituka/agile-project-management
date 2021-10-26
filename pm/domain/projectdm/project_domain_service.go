package projectdm

import (
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

func (s *projectDomainService) ExistsUniqueProjectForCreate(projectDm *Project) (bool, error) {
	existingProjectDm, err := s.projectRepository.FetchProjectByGroupIDAndKeyName(projectDm.GroupID(), projectDm.KeyName())
	if err != nil && !apperrors.Is(err, apperrors.NotFound) {
		return false, err
	} else if existingProjectDm != nil {
		return true, nil
	}

	existingProjectDm, err = s.projectRepository.FetchProjectByGroupIDAndName(projectDm.GroupID(), projectDm.Name())
	if existingProjectDm != nil {
		return true, nil
	}

	return false, err
}

func (s *projectDomainService) ExistUniqueProjectForUpdate(projectDm *Project) (bool, error) {
	oldProjectDm, err := s.projectRepository.FetchProjectByID(projectDm.ID())
	if err != nil {
		return false, apperrors.InternalServerError
	}

	if oldProjectDm.KeyName() == projectDm.KeyName() && oldProjectDm.Name() == projectDm.Name() {
		return false, nil
	}

	projectDmByName, errByKeyName := s.projectRepository.FetchProjectByGroupIDAndKeyName(projectDm.GroupID(), projectDm.KeyName())
	if errByKeyName != nil && !apperrors.Is(errByKeyName, apperrors.NotFound) {
		return false, errByKeyName
	}

	projectDmByKeyName, errByName := s.projectRepository.FetchProjectByGroupIDAndName(projectDm.GroupID(), projectDm.Name())
	if errByName != nil && !apperrors.Is(errByName, apperrors.NotFound) {
		return false, errByName
	}

	if apperrors.Is(errByKeyName, apperrors.NotFound) && apperrors.Is(errByName, apperrors.NotFound) {
		return false, apperrors.NotFound
	}

	if projectDmByName != nil {
		if projectDm.EqualKeyName(projectDmByName.KeyName()) && apperrors.Is(errByName, apperrors.NotFound) {
			return false, apperrors.NotFound
		}

		return true, nil
	}

	if projectDmByKeyName != nil {
		if projectDm.EqualName(projectDmByKeyName.Name()) && apperrors.Is(errByKeyName, apperrors.NotFound) {
			return false, apperrors.NotFound
		}

		return true, nil
	}

	return false, apperrors.NotFound
}
