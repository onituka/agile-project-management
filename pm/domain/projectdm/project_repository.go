package projectdm

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/domain/groupdm"
	"github.com/onituka/agile-project-management/project-management/domain/productdm"
)

type ProjectRepository interface {
	CreateProject(ctx context.Context, project *Project) error
	UpdateProject(ctx context.Context, project *Project) error
	FetchProjectByIDForUpdate(ctx context.Context, id ProjectID, productID productdm.ProductID) (*Project, error)
	FetchProjectByID(ctx context.Context, id ProjectID, productID productdm.ProductID) (*Project, error)
	FetchProjectByGroupIDAndKeyName(ctx context.Context, groupID groupdm.GroupID, keyName KeyName) (*Project, error)
	FetchProjectByGroupIDAndName(ctx context.Context, groupID groupdm.GroupID, name Name) (*Project, error)
}
