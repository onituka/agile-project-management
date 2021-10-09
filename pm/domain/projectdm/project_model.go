package projectdm

import (
	"time"

	"github.com/onituka/agile-project-management/project-management/domain/groupdm"
	"github.com/onituka/agile-project-management/project-management/domain/userdm"
)

type Project struct {
	id                ID
	groupID           groupdm.GroupID
	keyName           KeyName
	name              Name
	leaderID          userdm.UserID
	defaultAssigneeID userdm.UserID
	createdDate       time.Time
	updatedDate       time.Time
}

func NewProjectWithoutDate(id ID, groupID groupdm.GroupID, keyName KeyName, name Name, leaderID userdm.UserID, defaultAssigneeID userdm.UserID) *Project {
	return &Project{
		id:                id,
		groupID:           groupID,
		keyName:           keyName,
		name:              name,
		leaderID:          leaderID,
		defaultAssigneeID: defaultAssigneeID,
	}
}

func NewProject(id ID, groupID groupdm.GroupID, keyName KeyName, name Name, leaderID userdm.UserID, defaultAssigneeID userdm.UserID, createdDate time.Time, updatedDate time.Time) *Project {
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

func (p *Project) ID() ID {
	return p.id
}

func (p *Project) Group() groupdm.GroupID {
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

func (p *Project) CreatedDate() time.Time {
	return p.createdDate
}

func (p *Project) UpdatedDate() time.Time {
	return p.updatedDate
}
