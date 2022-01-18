package productnotedm

import (
	"github.com/google/uuid"

	"github.com/onituka/agile-project-management/project-management/apperrors"
)

type ProductNoteID string

func NewProductNoteIDForCreate() ProductNoteID {
	return ProductNoteID(uuid.New().String())
}

func NewProductNoteID(productNoteID string) (ProductNoteID, error) {
	if _, err := uuid.Parse(productNoteID); err != nil {
		return "", apperrors.InvalidParameter
	}

	return ProductNoteID(productNoteID), nil
}

func (i ProductNoteID) Value() string {
	return string(i)
}
