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
		sheredvo.NewProjectID(),
		groupID,
		keyName,
		name,
		leaderID,
		defaultAssigneeID,
		time.Time{},
		time.Time{},
	)
}

func GenProjectForFetch(
	id sheredvo.ProjectID,
	groupID sheredvo.GroupID,
	keyName KeyName,
	name Name,
	leaderID sheredvo.UserID,
	defaultAssigneeID sheredvo.UserID,
	createdDate time.Time,
	updatedDate time.Time,
) *Project {
	return newProject(
		id,
		groupID,
		keyName,
		name,
		leaderID,
		defaultAssigneeID,
		createdDate,
		updatedDate,
	)
}
