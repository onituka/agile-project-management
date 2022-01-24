package projectnotedm

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/domain/projectdm"
)

type ProjectNoteRepository interface {
	CreateNoteProject(ctx context.Context, projectNote *ProjectNote) error
	FetchProjectNoteProjectIDAndTitle(ctx context.Context, projectID projectdm.ProjectID, title Title) (*ProjectNote, error)
}
