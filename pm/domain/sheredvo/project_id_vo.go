package sheredvo

import "github.com/google/uuid"

type ProjectID string

func NewProjectID() ProjectID {
	return ProjectID(uuid.New().String())
}

func (i ProjectID) Value() string {
	return string(i)
}
