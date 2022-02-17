package productnotecommentdm

import (
	"time"

	"github.com/onituka/agile-project-management/project-management/domain/groupdm"
	"github.com/onituka/agile-project-management/project-management/domain/productdm"
	"github.com/onituka/agile-project-management/project-management/domain/productnotedm"
	"github.com/onituka/agile-project-management/project-management/domain/userdm"
)

func GenProductNoteCommentForCreate(
	productID productdm.ProductID,
	productNoteID productnotedm.ProductNoteID,
	groupID groupdm.GroupID,
	content Content,
	userID userdm.UserID,
) (*ProductNoteComment, error) {
	now := time.Now().UTC()

	return newProductNoteComment(
		NewProductNoteCommentIDForCreate(),
		productID,
		productNoteID,
		groupID,
		content,
		userID,
		now,
		now,
	)
}
