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
	createdAt time.Time,
	updatedAt time.Time,
) *Project {
	return newProject(
		sheredvo.NewProjectIDForCreate(),
		groupID,
		keyName,
		name,
		leaderID,
		defaultAssigneeID,
		createdAt,
		updatedAt,
	)
}
