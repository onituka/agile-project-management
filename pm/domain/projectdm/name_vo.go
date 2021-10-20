package projectdm

import "github.com/onituka/agile-project-management/project-management/apperrors"

type Name string

func NewName(name string) (Name, error) {
	if l := len(name); l < 2 || l > 80 {
		return "", apperrors.InvalidParameter
	}

	return Name(name), nil
}

func (n Name) Value() string {
	return string(n)
}

func (p *Project) EqualName(name Name) bool {
	return p.Name() == name
}
