package productdm

import (
	"time"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/groupdm"
	"github.com/onituka/agile-project-management/project-management/domain/userdm"
)

type Product struct {
	id        ProductID
	groupID   groupdm.GroupID
	name      Name
	leaderID  userdm.UserID
	trashedAt *time.Time
	createdAt time.Time
	updatedAt time.Time
}

func newProduct(
	id ProductID,
	groupID groupdm.GroupID,
	name Name,
	leaderID userdm.UserID,
	trashedAt *time.Time,
	createdAt time.Time,
	updatedAt time.Time,
) (*Product, error) {
	if createdAt.IsZero() || updatedAt.IsZero() {
		return nil, apperrors.InvalidParameter
	}

	return &Product{
		id:        id,
		groupID:   groupID,
		name:      name,
		leaderID:  leaderID,
		trashedAt: trashedAt,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}, nil
}

func Reconstruct(
	id string,
	groupID string,
	name string,
	leaderID string,
	trashedAt *time.Time,
	createdAt time.Time,
	updatedAt time.Time,
) (*Product, error) {
	idVo, err := NewProductID(id)
	if err != nil {
		return nil, err
	}

	groupIDVo, err := groupdm.NewGroupID(groupID)
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

	return newProduct(
		idVo,
		groupIDVo,
		nameVo,
		leaderIDVo,
		trashedAt,
		createdAt,
		updatedAt,
	)
}

func (p *Product) ID() ProductID {
	return p.id
}

func (p *Product) GroupID() groupdm.GroupID {
	return p.groupID
}

func (p *Product) Name() Name {
	return p.name
}

func (p *Product) LeaderID() userdm.UserID {
	return p.leaderID
}

func (p *Product) TrashedAt() *time.Time {
	return p.trashedAt
}

func (p *Product) CreatedAt() time.Time {
	return p.createdAt
}

func (p *Product) UpdatedAt() time.Time {
	return p.updatedAt
}

func (p *Product) ChangeName(name Name) {
	p.name = name
}

func (p *Product) ChangeLeaderID(leaderID userdm.UserID) {
	p.leaderID = leaderID
}

func (p *Product) ChangeUpdatedAt(updatedAt time.Time) {
	p.updatedAt = updatedAt
}

func (p *Product) IsTrashed() bool {
	return p.trashedAt != nil
}

func (p *Product) MoveTrash() {
	now := time.Now().UTC()
	p.trashedAt = &now
}
