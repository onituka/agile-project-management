package projectnotedm

import (
	"unicode/utf8"

	"github.com/onituka/agile-project-management/project-management/apperrors"
)

type Title string

func NewTitle(title string) (Title, error) {
	if l := utf8.RuneCountInString(title); l < 1 || l > 255 {
		return "", apperrors.InvalidParameter
	}

	return Title(title), nil
}

func (t Title) Value() string {
	return string(t)
}
