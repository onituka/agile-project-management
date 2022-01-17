package productnotedm

import (
	"time"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/groupdm"
	"github.com/onituka/agile-project-management/project-management/domain/productdm"
	"github.com/onituka/agile-project-management/project-management/domain/userdm"
)

type ProductNote struct {
	id        ProductNoteID
	productID productdm.ProductID
	groupID   groupdm.GroupID
	title     Title
	content   Content
	createdBy userdm.UserID
	updatedBy userdm.UserID
	createdAt time.Time
	updatedAt time.Time
}

func newProductNote(
	id ProductNoteID,
	productID productdm.ProductID,
	groupID groupdm.GroupID,
	title Title,
	Content Content,
	createdBy userdm.UserID,
	updatedBy userdm.UserID,
	createdAt time.Time,
	updatedAt time.Time,
) (*ProductNote, error) {
	if createdAt.IsZero() || updatedAt.IsZero() {
		return nil, apperrors.InvalidParameter
	}

	return &ProductNote{
		id:        id,
		productID: productID,
		groupID:   groupID,
		title:     title,
		content:   Content,
		createdBy: createdBy,
		updatedBy: updatedBy,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}, nil
}

func Reconstruct(
	id string,
	productID string,
	groupID string,
	title string,
	content string,
	createdBy string,
	updatedBy string,
	createdAt time.Time,
	updatedAt time.Time,
) (*ProductNote, error) {
	idVo, err := NewProductNoteID(id)
	if err != nil {
		return nil, err
	}

	productIDVo, err := productdm.NewProductID(productID)
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

	return newProductNote(
		idVo,
		productIDVo,
		groupIDVo,
		titleVo,
		contentVo,
		createdByVo,
		updatedByVo,
		createdAt,
		updatedAt,
	)
}

func (p *ProductNote) ID() ProductNoteID {
	return p.id
}

func (p *ProductNote) ProductID() productdm.ProductID {
	return p.productID
}

func (p *ProductNote) GroupID() groupdm.GroupID {
	return p.groupID
}

func (p *ProductNote) Title() Title {
	return p.title
}

func (p *ProductNote) Content() Content {
	return p.content
}

func (p *ProductNote) CreatedBy() userdm.UserID {
	return p.createdBy
}

func (p *ProductNote) UpdatedBy() userdm.UserID {
	return p.updatedBy
}

func (p *ProductNote) CreatedAt() time.Time {
	return p.createdAt
}

func (p *ProductNote) UpdatedAt() time.Time {
	return p.updatedAt
}
