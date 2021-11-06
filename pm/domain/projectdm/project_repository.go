package projectdm

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/domain/sharedvo"
)

type ProjectRepository interface {
	CreateProject(ctx context.Context, project *Project) error
	UpdateProject(ctx context.Context, project *Project) error
	FetchProjectByIDForUpdate(ctx context.Context, id ProjectID) (*Project, error)
	FetchProjectByID(ctx context.Context, id ProjectID) (*Project, error)
	FetchProjectByGroupIDAndKeyName(ctx context.Context, groupID sharedvo.GroupID, keyName KeyName) (*Project, error)
	FetchProjectByGroupIDAndName(ctx context.Context, groupID sharedvo.GroupID, name Name) (*Project, error)
}
