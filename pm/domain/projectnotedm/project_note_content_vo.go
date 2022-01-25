package projectnotedm

import (
	"unicode/utf8"

	"github.com/onituka/agile-project-management/project-management/apperrors"
)

type Content string

func NewContent(content string) (Content, error) {
	if l := utf8.RuneCountInString(content); l == 0 || l > 20000 {
		return "", apperrors.InvalidParameter
	}

	return Content(content), nil
}

func (t Content) Value() string {
	return string(t)
}
