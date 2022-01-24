package projectnotedm

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/apperrors"
)

type projectNoteDomainService struct {
	projectNoteRepository ProjectNoteRepository
}

func NewProjectNoteDomainService(projectNoteRepository ProjectNoteRepository) *projectNoteDomainService {
	return &projectNoteDomainService{
		projectNoteRepository: projectNoteRepository,
	}
}

func (s *projectNoteDomainService) ExistsProjectNoteForCreate(ctx context.Context, projectNoteDm *ProjectNote) (bool, error) {
	existingProjectNoteDm, err := s.projectNoteRepository.FetchProjectNoteProjectIDAndTitle(ctx, projectNoteDm.ProjectID(), projectNoteDm.Title())
	if err != nil && !apperrors.Is(err, apperrors.NotFound) {
		return false, err
	} else if existingProjectNoteDm != nil {
		return true, err
	}

	return false, err
}
