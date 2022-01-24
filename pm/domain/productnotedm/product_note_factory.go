package productnotedm

import (
	"time"

	"github.com/onituka/agile-project-management/project-management/domain/groupdm"
	"github.com/onituka/agile-project-management/project-management/domain/productdm"
	"github.com/onituka/agile-project-management/project-management/domain/userdm"
)

func GenProductNoteForCreate(
	productID productdm.ProductID,
	groupID groupdm.GroupID,
	title Title,
	content Content,
	UserID userdm.UserID,
) (*ProductNote, error) {
	now := time.Now().UTC()

	return newProductNote(
		NewProductNoteIDForCreate(),
		productID,
		groupID,
		title,
		content,
		UserID,
		UserID,
		now,
		now,
	)
}
