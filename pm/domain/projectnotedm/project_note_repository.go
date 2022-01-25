package projectnotedm

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/domain/projectdm"
)

type ProjectNoteRepository interface {
	CreateProjectNote(ctx context.Context, projectNote *ProjectNote) error
	FetchProjectNoteByProjectIDAndTitle(ctx context.Context, projectID projectdm.ProjectID, title Title) (*ProjectNote, error)
}
