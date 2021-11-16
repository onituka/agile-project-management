package productdm

import (
	"unicode/utf8"

	"github.com/onituka/agile-project-management/project-management/apperrors"
)

type Name string

func NewName(name string) (Name, error) {
	if l := utf8.RuneCountInString(name); l < 2 || l > 80 {
		return "", apperrors.InvalidParameter
	}

	return Name(name), nil
}

func (n Name) Value() string {
	return string(n)
}

func (n Name) Equals(name Name) bool {
	return n == name
}
