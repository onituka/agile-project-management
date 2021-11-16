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
	createdAt time.Time,
	updatedAt time.Time,
) (*Product, error) {
	return newProduct(
		NewProductIDForCreate(),
		groupID,
		name,
		leaderID,
		createdAt,
		updatedAt,
	)
}
