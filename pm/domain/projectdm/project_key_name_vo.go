package projectdm

import (
	"regexp"

	"github.com/onituka/agile-project-management/project-management/apperrors"
)

type KeyName string

var keyNameRegex = regexp.MustCompile(`^[A-Z][A-Z0-9]{1,9}$`)

func NewKeyName(keyName string) (KeyName, error) {
	if !keyNameRegex.MatchString(keyName) {
		return "", apperrors.InvalidParameter
	}

	return KeyName(keyName), nil
}

func (k KeyName) Value() string {
	return string(k)
}

func (k KeyName) Equals(keyName KeyName) bool {
	return k == keyName
}
