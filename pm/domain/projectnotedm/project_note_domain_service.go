package projectnotedm

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/projectdm"
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
	existingProjectNoteDm, err := s.projectNoteRepository.FetchProjectNoteByProjectIDAndTitle(ctx, projectNoteDm.ProjectID(), projectNoteDm.Title())
	if err != nil && !apperrors.Is(err, apperrors.NotFound) {
		return false, err
	} else if existingProjectNoteDm != nil {
		return true, nil
	}

	return false, err
}

func (s *projectNoteDomainService) ExistsProjectNoteByIDForUpdate(ctx context.Context, projectNoteIDVo ProjectNoteID, projectIDVo projectdm.ProjectID) (bool, error) {
	if _, err := s.projectNoteRepository.FetchProjectNoteByIDForUpdate(ctx, projectNoteIDVo, projectIDVo); err != nil {
		return false, err
	}

	return true, nil
}

func (s *projectNoteDomainService) ExistsProjectNoteForUpdate(ctx context.Context, ProjectNoteDm *ProjectNote) (bool, error) {
	oldProjectNote, err := s.projectNoteRepository.FetchProjectNoteByID(ctx, ProjectNoteDm.ID(), ProjectNoteDm.ProjectID())
	if err != nil {
		return false, apperrors.InternalServerError
	}

	projectNoteDmByTitle, err := s.projectNoteRepository.FetchProjectNoteByProjectIDAndTitle(ctx, ProjectNoteDm.ProjectID(), ProjectNoteDm.Title())
	if err != nil && !apperrors.Is(err, apperrors.NotFound) {
		return false, err
	}

	if projectNoteDmByTitle != nil {
		if ProjectNoteDm.Title().Equals(oldProjectNote.Title()) {
			return false, apperrors.NotFound
		}

		return true, nil
	}

	return false, apperrors.NotFound
}
