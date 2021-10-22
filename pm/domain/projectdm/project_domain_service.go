package projectdm

import (
	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/sheredvo"
)

type projectDomainService struct {
	projectRepository ProjectRepository
}

func NewProjectDomainService(projectRepository ProjectRepository) *projectDomainService {
	return &projectDomainService{
		projectRepository: projectRepository,
	}
}

func (s *projectDomainService) ExistsUniqueProjectKeyName(groupID sheredvo.GroupID, keyName KeyName) (bool, error) {
	projectDm, err := s.projectRepository.FetchProjectByGroupIDAndKeyName(groupID, keyName)
	if projectDm != nil {
		return true, nil
	}

	return false, err
}

func (s *projectDomainService) ExistsUniqueProjectName(groupID sheredvo.GroupID, name Name) (bool, error) {
	projectDm, err := s.projectRepository.FetchProjectByGroupIDAndName(groupID, name)
	if projectDm != nil {
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

	projectDmByName, errByKeyName := s.projectRepository.FetchProjectByGroupIDAndKeyName(projectDm.Group(), projectDm.KeyName())
	if errByKeyName != nil && !apperrors.Is(errByKeyName, apperrors.NotFound) {
		return false, errByKeyName
	}

	projectDmByKeyName, errByName := s.projectRepository.FetchProjectByGroupIDAndName(projectDm.groupID, projectDm.Name())
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
