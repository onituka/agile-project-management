package projectnotedm

import (
	"github.com/google/uuid"

	"github.com/onituka/agile-project-management/project-management/apperrors"
)

type ProjectNoteID string

func NewProjectNoteIDForCreate() ProjectNoteID {
	return ProjectNoteID(uuid.New().String())
}

func NewProjectNoteID(projectNoteID string) (ProjectNoteID, error) {
	if _, err := uuid.Parse(projectNoteID); err != nil {
		return "", apperrors.InvalidParameter
	}

	return ProjectNoteID(projectNoteID), nil
}

func (i ProjectNoteID) Value() string {
	return string(i)
}
