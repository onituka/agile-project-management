package projectdm

import (
	"time"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/groupdm"
	"github.com/onituka/agile-project-management/project-management/domain/productdm"
	"github.com/onituka/agile-project-management/project-management/domain/userdm"
)

type Project struct {
	id                ProjectID
	productID         productdm.ProductID
	groupID           groupdm.GroupID
	keyName           KeyName
	name              Name
	leaderID          userdm.UserID
	defaultAssigneeID userdm.UserID
	trashedAt         *time.Time
	createdAt         time.Time
	updatedAt         time.Time
}

func newProject(
	id ProjectID,
	productID productdm.ProductID,
	groupID groupdm.GroupID,
	keyName KeyName,
	name Name,
	leaderID userdm.UserID,
	defaultAssigneeID userdm.UserID,
	trashedAt *time.Time,
	createdAt time.Time,
	updatedAt time.Time,
) (*Project, error) {
	if createdAt.IsZero() || updatedAt.IsZero() {
		return nil, apperrors.InvalidParameter
	}

	return &Project{
		id:                id,
		productID:         productID,
		groupID:           groupID,
		keyName:           keyName,
		name:              name,
		leaderID:          leaderID,
		defaultAssigneeID: defaultAssigneeID,
		trashedAt:         trashedAt,
		createdAt:         createdAt,
		updatedAt:         updatedAt,
	}, nil
}

func Reconstruct(
	id string,
	productID string,
	groupID string,
	keyName string,
	name string,
	leaderID string,
	defaultAssigneeID string,
	trashedAt *time.Time,
	createdAt time.Time,
	updatedAt time.Time,
) (*Project, error) {
	idVo, err := NewProjectID(id)
	if err != nil {
		return nil, err
	}

	productIDVo, err := productdm.NewProductID(productID)
	if err != nil {
		return nil, err
	}

	groupIDVo, err := groupdm.NewGroupID(groupID)
	if err != nil {
		return nil, err
	}

	keyNameVo, err := NewKeyName(keyName)
	if err != nil {
		return nil, err
	}

	nameVo, err := NewName(name)
	if err != nil {
		return nil, err
	}

	leaderIDVo, err := userdm.NewUserID(leaderID)
	if err != nil {
		return nil, err
	}

	defaultAssigneeIDVo, err := userdm.NewUserID(defaultAssigneeID)
	if err != nil {
		return nil, err
	}

	return newProject(
		idVo,
		productIDVo,
		groupIDVo,
		keyNameVo,
		nameVo,
		leaderIDVo,
		defaultAssigneeIDVo,
		trashedAt,
		createdAt,
		updatedAt,
	)
}

func (p *Project) ID() ProjectID {
	return p.id
}

func (p *Project) ProductID() productdm.ProductID {
	return p.productID
}

func (p *Project) GroupID() groupdm.GroupID {
	return p.groupID
}

func (p *Project) KeyName() KeyName {
	return p.keyName
}

func (p *Project) Name() Name {
	return p.name
}

func (p *Project) LeaderID() userdm.UserID {
	return p.leaderID
}

func (p *Project) DefaultAssigneeID() userdm.UserID {
	return p.defaultAssigneeID
}

func (p *Project) TrashedAt() *time.Time {
	return p.trashedAt
}

func (p *Project) CreatedAt() time.Time {
	return p.createdAt
}

func (p *Project) UpdatedAt() time.Time {
	return p.updatedAt
}

func (p *Project) ChangeKeyName(keyName KeyName) {
	p.keyName = keyName
}

func (p *Project) ChangeName(name Name) {
	p.name = name
}

func (p *Project) ChangeLeaderID(leaderID userdm.UserID) {
	p.leaderID = leaderID
}

func (p *Project) ChangeDefaultAssigneeID(defaultAssigneeID userdm.UserID) {
	p.defaultAssigneeID = defaultAssigneeID
}

func (p *Project) MoveToTrashed() {
	now := time.Now().UTC()
	p.trashedAt = &now
}

func (p *Project) MoveToUpdate() {
	now := time.Now().UTC()
	p.updatedAt = now
}

func (p *Project) IsTrashed() bool {
	return p.trashedAt != nil
}
