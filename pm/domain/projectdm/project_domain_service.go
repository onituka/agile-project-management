package projectdm

import (
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

func (s *projectDomainService) ExistsProjectKeyByUnique(groupID sheredvo.GroupID, keyName KeyName) (bool, error) {
	projectDm, err := s.projectRepository.FetchProjectByGroupIDAndKeyName(groupID, keyName)
	if projectDm != nil {
		return true, nil
	}

	return false, err
}

func (s *projectDomainService) ExistsProjectNameByUnique(groupID sheredvo.GroupID, name Name) (bool, error) {
	projectDm, err := s.projectRepository.FetchProjectByGroupIDAndName(groupID, name)
	if projectDm != nil {
		return true, nil
	}

	return false, err
}
