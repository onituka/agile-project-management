package projectdm

import (
	"time"

	"github.com/onituka/agile-project-management/project-management/domain/sharedvo"
)

func GenProjectForCreate(
	groupID sharedvo.GroupID,
	keyName KeyName,
	name Name,
	leaderID sharedvo.UserID,
	defaultAssigneeID sharedvo.UserID,
	createdAt time.Time,
	updatedAt time.Time,
) *Project {
	return newProject(
		NewProjectIDForCreate(),
		groupID,
		keyName,
		name,
		leaderID,
		defaultAssigneeID,
		createdAt,
		updatedAt,
	)
}
