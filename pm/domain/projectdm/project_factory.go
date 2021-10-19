package projectdm

import (
	"time"

	"github.com/onituka/agile-project-management/project-management/domain/sheredvo"
)

func GenProjectForCreate(
	groupID sheredvo.GroupID,
	keyName KeyName,
	name Name,
	leaderID sheredvo.UserID,
	defaultAssigneeID sheredvo.UserID,
) *Project {
	return newProject(
		sheredvo.NewProjectIDForCreate(),
		groupID,
		keyName,
		name,
		leaderID,
		defaultAssigneeID,
		time.Time{},
		time.Time{},
	)
}

func GenProjectForUpdate(
	id sheredvo.ProjectID,
	groupID sheredvo.GroupID,
	keyName KeyName,
	name Name,
	leaderID sheredvo.UserID,
	defaultAssigneeID sheredvo.UserID,
) *Project {
	return newProject(
		id,
		groupID,
		keyName,
		name,
		leaderID,
		defaultAssigneeID,
		time.Time{},
		time.Time{},
	)
}
