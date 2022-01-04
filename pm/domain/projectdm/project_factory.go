package projectdm

import (
	"time"

	"github.com/onituka/agile-project-management/project-management/domain/groupdm"
	"github.com/onituka/agile-project-management/project-management/domain/productdm"
	"github.com/onituka/agile-project-management/project-management/domain/userdm"
)

func GenProjectForCreate(
	productID productdm.ProductID,
	groupID groupdm.GroupID,
	keyName KeyName,
	name Name,
	leaderID userdm.UserID,
	defaultAssigneeID userdm.UserID,
) (*Project, error) {
	now := time.Now().UTC()

	return newProject(
		NewProjectIDForCreate(),
		productID,
		groupID,
		keyName,
		name,
		leaderID,
		defaultAssigneeID,
		nil,
		now,
		now,
	)
}
