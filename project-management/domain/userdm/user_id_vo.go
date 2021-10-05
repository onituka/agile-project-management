package userdm

import (
	"github.com/google/uuid"

	"github.com/onituka/agile-project-management/project-management/apperrors"
)

type UserID string

func NewUserID(userID string) (UserID, error) {
	if _, err := uuid.Parse(userID); err != nil {
		return "", apperrors.InvalidParameter
	}

	return UserID(userID), nil
}

func (i UserID) Value() string {
	return string(i)
}
