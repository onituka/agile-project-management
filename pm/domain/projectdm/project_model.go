package projectdm

import (
	"time"

	"github.com/onituka/agile-project-management/project-management/domain/sheredvo"
)

type Project struct {
	id                sheredvo.ProjectID
	groupID           sheredvo.GroupID
	keyName           KeyName
	name              Name
	leaderID          sheredvo.UserID
	defaultAssigneeID sheredvo.UserID
	createdDate       time.Time
	updatedDate       time.Time
}

func newProject(
	id sheredvo.ProjectID,
	groupID sheredvo.GroupID,
	keyName KeyName,
	name Name,
	leaderID sheredvo.UserID,
	defaultAssigneeID sheredvo.UserID,
	createdDate time.Time,
	updatedDate time.Time,
) *Project {
	return &Project{
		id:                id,
		groupID:           groupID,
		keyName:           keyName,
		name:              name,
		leaderID:          leaderID,
		defaultAssigneeID: defaultAssigneeID,
		createdDate:       createdDate,
		updatedDate:       updatedDate,
	}
}

func Reconstruct(
	id string,
	groupID string,
	keyName string,
	name string,
	leaderID string,
	defaultAssigneeID string,
	createdDate time.Time,
	updatedDate time.Time,
) *Project {
	return newProject(
		sheredvo.ProjectID(id),
		sheredvo.GroupID(groupID),
		KeyName(keyName),
		Name(name),
		sheredvo.UserID(leaderID),
		sheredvo.UserID(defaultAssigneeID),
		createdDate,
		updatedDate,
	)
}

func (p *Project) ID() sheredvo.ProjectID {
	return p.id
}

func (p *Project) GroupID() sheredvo.GroupID {
	return p.groupID
}

func (p *Project) KeyName() KeyName {
	return p.keyName
}

func (p *Project) Name() Name {
	return p.name
}

func (p *Project) LeaderID() sheredvo.UserID {
	return p.leaderID
}

func (p *Project) DefaultAssigneeID() sheredvo.UserID {
	return p.defaultAssigneeID
}

func (p *Project) CreatedAt() time.Time {
	return p.createdDate
}

func (p *Project) UpdatedAt() time.Time {
	return p.updatedDate
}

func (p *Project) ChangeKeyName(keyName KeyName) {
	p.keyName = keyName
}

func (p *Project) ChangeName(name Name) {
	p.name = name
}

func (p *Project) ChangeLeaderID(leaderID sheredvo.UserID) {
	p.leaderID = leaderID
}

func (p *Project) ChangeDefaultAssigneeID(defaultAssigneeID sheredvo.UserID) {
	p.defaultAssigneeID = defaultAssigneeID
}
