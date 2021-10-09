package projectdm

import "github.com/onituka/agile-project-management/project-management/domain/groupdm"

type Repository interface {
	CreateProject(project *Project) error
	FetchProjectByID(id ID) (*Project, error)
	FetchProjectByGroupIDAndKeyName(groupID groupdm.GroupID, keyName KeyName) (*Project, error)
	FetchProjectByGroupIDAndName(groupID groupdm.GroupID, name Name) (*Project, error)
}
