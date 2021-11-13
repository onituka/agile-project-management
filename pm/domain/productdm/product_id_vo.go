package productdm

import (
	"github.com/google/uuid"
)

type ProductID string

func NewProductIDForCreate() ProductID {
	return ProductID(uuid.New().String())
}

func (i ProductID) Value() string {
	return string(i)
}
