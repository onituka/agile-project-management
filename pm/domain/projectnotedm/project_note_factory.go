package projectnotedm

import (
	"time"

	"github.com/onituka/agile-project-management/project-management/domain/groupdm"
	"github.com/onituka/agile-project-management/project-management/domain/productdm"
	"github.com/onituka/agile-project-management/project-management/domain/projectdm"
	"github.com/onituka/agile-project-management/project-management/domain/userdm"
)

func GenProjectNoteForCreate(
	productID productdm.ProductID,
	projectID projectdm.ProjectID,
	groupID groupdm.GroupID,
	title Title,
	content Content,
	createdBy userdm.UserID,
	UpdatedBy userdm.UserID,
) (*ProjectNote, error) {
	now := time.Now().UTC()

	return newProjectNote(
		NewProjectNoteIDForCreate(),
		productID,
		projectID,
		groupID,
		title,
		content,
		createdBy,
		UpdatedBy,
		now,
		now,
	)
}
