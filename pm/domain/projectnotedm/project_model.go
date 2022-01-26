package projectnotedm

import (
	"time"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/groupdm"
	"github.com/onituka/agile-project-management/project-management/domain/productdm"
	"github.com/onituka/agile-project-management/project-management/domain/projectdm"
	"github.com/onituka/agile-project-management/project-management/domain/userdm"
)

type ProjectNote struct {
	id        ProjectNoteID
	productID productdm.ProductID
	projectID projectdm.ProjectID
	groupID   groupdm.GroupID
	title     Title
	content   Content
	createdBy userdm.UserID
	updatedBy userdm.UserID
	createdAt time.Time
	updatedAt time.Time
}

func newProjectNote(
	id ProjectNoteID,
	productID productdm.ProductID,
	projectID projectdm.ProjectID,
	groupID groupdm.GroupID,
	title Title,
	content Content,
	createdBy userdm.UserID,
	updatedBy userdm.UserID,
	createdAt time.Time,
	updatedAt time.Time,
) (*ProjectNote, error) {
	if createdAt.IsZero() || updatedAt.IsZero() {
		return nil, apperrors.InvalidParameter
	}

	return &ProjectNote{
		id:        id,
		productID: productID,
		projectID: projectID,
		groupID:   groupID,
		title:     title,
		content:   content,
		createdBy: createdBy,
		updatedBy: updatedBy,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}, nil
}

func Reconstruct(
	id string,
	productID string,
	projectID string,
	groupID string,
	title string,
	content string,
	createdBy string,
	updatedBy string,
	createdAt time.Time,
	updatedAt time.Time,
) (*ProjectNote, error) {
	idVo, err := NewProjectNoteID(id)
	if err != nil {
		return nil, err
	}

	productIDVo, err := productdm.NewProductID(productID)
	if err != nil {
		return nil, err
	}

	projectIDVo, err := projectdm.NewProjectID(projectID)
	if err != nil {
		return nil, err
	}

	groupIDVo, err := groupdm.NewGroupID(groupID)
	if err != nil {
		return nil, err
	}

	titleVo, err := NewTitle(title)
	if err != nil {
		return nil, err
	}

	contentVo, err := NewContent(content)
	if err != nil {
		return nil, err
	}

	createdByVo, err := userdm.NewUserID(createdBy)
	if err != nil {
		return nil, err
	}

	updatedByVo, err := userdm.NewUserID(updatedBy)
	if err != nil {
		return nil, err
	}

	return newProjectNote(
		idVo,
		productIDVo,
		projectIDVo,
		groupIDVo,
		titleVo,
		contentVo,
		createdByVo,
		updatedByVo,
		createdAt,
		updatedAt,
	)
}

func (p *ProjectNote) ID() ProjectNoteID {
	return p.id
}

func (p *ProjectNote) ProductID() productdm.ProductID {
	return p.productID
}

func (p *ProjectNote) ProjectID() projectdm.ProjectID {
	return p.projectID
}

func (p *ProjectNote) GroupID() groupdm.GroupID {
	return p.groupID
}

func (p *ProjectNote) Title() Title {
	return p.title
}

func (p *ProjectNote) Content() Content {
	return p.content
}

func (p *ProjectNote) CreatedBy() userdm.UserID {
	return p.createdBy
}

func (p *ProjectNote) UpdatedBy() userdm.UserID {
	return p.updatedBy
}

func (p *ProjectNote) CreatedAt() time.Time {
	return p.createdAt
}

func (p *ProjectNote) UpdatedAt() time.Time {
	return p.updatedAt
}

func (p *ProjectNote) ChangeTitle(title Title) {
	p.title = title
}

func (p *ProjectNote) ChangeContent(content Content) {
	p.content = content
}

func (p *ProjectNote) ChangeUpdatedBy(updatedBy userdm.UserID) {
	p.updatedBy = updatedBy
}

func (p *ProjectNote) ChangeUpdateAt() {
	now := time.Now().UTC()
	p.updatedAt = now
}
