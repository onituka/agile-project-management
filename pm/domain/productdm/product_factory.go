package productdm

import (
	"time"

	"github.com/onituka/agile-project-management/project-management/domain/groupdm"
	"github.com/onituka/agile-project-management/project-management/domain/userdm"
)

func GenProductForCreate(
	groupID groupdm.GroupID,
	name Name,
	leaderID userdm.UserID,
) (*Product, error) {
	now := time.Now().UTC()

	return newProduct(
		NewProductIDForCreate(),
		groupID,
		name,
		leaderID,
		now,
		now,
	)
}
