package projectdm

import (
	"time"

	"github.com/onituka/agile-project-management/project-management/domain/sharedvo"
)

type Project struct {
	id                ProjectID
	groupID           sharedvo.GroupID
	keyName           KeyName
	name              Name
	leaderID          sharedvo.UserID
	defaultAssigneeID sharedvo.UserID
	createdAt         time.Time
	updatedAt         time.Time
}

func newProject(
	id ProjectID,
	groupID sharedvo.GroupID,
	keyName KeyName,
	name Name,
	leaderID sharedvo.UserID,
	defaultAssigneeID sharedvo.UserID,
	createdAt time.Time,
	updatedAt time.Time,
) *Project {
	return &Project{
		id:                id,
		groupID:           groupID,
		keyName:           keyName,
		name:              name,
		leaderID:          leaderID,
		defaultAssigneeID: defaultAssigneeID,
		createdAt:         createdAt,
		updatedAt:         updatedAt,
	}
}

func Reconstruct(
	id string,
	groupID string,
	keyName string,
	name string,
	leaderID string,
	defaultAssigneeID string,
	createdAt time.Time,
	updatedAt time.Time,
) *Project {
	return newProject(
		ProjectID(id),
		sharedvo.GroupID(groupID),
		KeyName(keyName),
		Name(name),
		sharedvo.UserID(leaderID),
		sharedvo.UserID(defaultAssigneeID),
		createdAt,
		updatedAt,
	)
}

func (p *Project) ID() ProjectID {
	return p.id
}

func (p *Project) GroupID() sharedvo.GroupID {
	return p.groupID
}

func (p *Project) KeyName() KeyName {
	return p.keyName
}

func (p *Project) Name() Name {
	return p.name
}

func (p *Project) LeaderID() sharedvo.UserID {
	return p.leaderID
}

func (p *Project) DefaultAssigneeID() sharedvo.UserID {
	return p.defaultAssigneeID
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

func (p *Project) ChangeLeaderID(leaderID sharedvo.UserID) {
	p.leaderID = leaderID
}

func (p *Project) ChangeDefaultAssigneeID(defaultAssigneeID sharedvo.UserID) {
	p.defaultAssigneeID = defaultAssigneeID
}

func (p *Project) ChangeUpdatedAt(updatedAt time.Time) {
	p.updatedAt = updatedAt
}
