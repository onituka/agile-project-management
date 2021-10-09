package groupdm

import (
	"github.com/google/uuid"

	"github.com/onituka/agile-project-management/project-management/apperrors"
)

type GroupID string

func NewGroupID(groupID string) (GroupID, error) {
	if _, err := uuid.Parse(groupID); err != nil {
		return "", apperrors.InvalidParameter
	}

	return GroupID(groupID), nil
}

func (i GroupID) Value() string {
	return string(i)
}
