package sheredvo

import (
	"github.com/google/uuid"

	"github.com/onituka/agile-project-management/project-management/apperrors"
)

type ProjectID string

func NewProjectIDForCreate() ProjectID {
	return ProjectID(uuid.New().String())
}

func NewProjectID(projectID string) (ProjectID, error) {
	if _, err := uuid.Parse(projectID); err != nil {
		return "", apperrors.InvalidParameter
	}

	return ProjectID(projectID), nil
}

func (i ProjectID) Value() string {
	return string(i)
}
