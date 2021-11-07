package projectdm

import (
	"time"

	"github.com/onituka/agile-project-management/project-management/domain/groupdm"
	"github.com/onituka/agile-project-management/project-management/domain/userdm"
)

func GenProjectForCreate(
	groupID groupdm.GroupID,
	keyName KeyName,
	name Name,
	leaderID userdm.UserID,
	defaultAssigneeID userdm.UserID,
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
