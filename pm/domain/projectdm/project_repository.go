package projectdm

import (
	"github.com/onituka/agile-project-management/project-management/domain/sheredvo"
)

type ProjectRepository interface {
	CreateProject(project *Project) error
	FetchProjectByID(id sheredvo.ProjectID) (*Project, error)
	FetchProjectByGroupIDAndKeyName(groupID sheredvo.GroupID, keyName KeyName) (*Project, error)
	FetchProjectByGroupIDAndName(groupID sheredvo.GroupID, name Name) (*Project, error)
}
