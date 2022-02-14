package projectnotedm

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/domain/productdm"
	"github.com/onituka/agile-project-management/project-management/domain/projectdm"
)

type ProjectNoteRepository interface {
	CreateProjectNote(ctx context.Context, projectNote *ProjectNote) error
	UpdateProjectNote(ctx context.Context, projectNote *ProjectNote) error
	FetchProjectNoteByProjectIDAndTitle(ctx context.Context, projectID projectdm.ProjectID, title Title) (*ProjectNote, error)
	FetchProjectNoteByIDForUpdate(ctx context.Context, id ProjectNoteID, projectID projectdm.ProjectID) (*ProjectNote, error)
	FetchProjectNoteByID(ctx context.Context, id ProjectNoteID, projectID projectdm.ProjectID) (*ProjectNote, error)
	DeleteProjectNote(ctx context.Context, id ProjectNoteID, productID productdm.ProductID, projectID projectdm.ProjectID) error
}
