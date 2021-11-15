package productdm

import (
	"github.com/google/uuid"

	"github.com/onituka/agile-project-management/project-management/apperrors"
)

type ProductID string

func NewProductIDForCreate() ProductID {
	return ProductID(uuid.New().String())
}

func NewProductID(productID string) (ProductID, error) {
	if _, err := uuid.Parse(productID); err != nil {
		return "", apperrors.InvalidParameter
	}

	return ProductID(productID), nil
}

func (i ProductID) Value() string {
	return string(i)
}
