package sheredvo

import "github.com/google/uuid"

type ID string

func NewID() ID {
	return ID(uuid.New().String())
}

func (i ID) Value() string {
	return string(i)
}
