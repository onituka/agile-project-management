package productnotecommentdm

import (
	"github.com/google/uuid"

	"github.com/onituka/agile-project-management/project-management/apperrors"
)

type ProductNoteCommentID string

func NewProductNoteCommentIDForCreate() ProductNoteCommentID {
	return ProductNoteCommentID(uuid.New().String())
}

func NewProductNoteCommentID(productNoteCommentID string) (ProductNoteCommentID, error) {
	if _, err := uuid.Parse(productNoteCommentID); err != nil {
		return "", apperrors.InvalidParameter
	}

	return ProductNoteCommentID(productNoteCommentID), nil
}

func (i ProductNoteCommentID) Value() string {
	return string(i)
}
