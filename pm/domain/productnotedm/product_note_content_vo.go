package productnotedm

import (
	"github.com/onituka/agile-project-management/project-management/apperrors"
)

type Content string

func NewContent(content string) (Content, error) {
	if l := len(content); l == 0 || l > 65535 {
		return "", apperrors.InvalidParameter
	}

	return Content(content), nil
}

func (t Content) Value() string {
	return string(t)
}
