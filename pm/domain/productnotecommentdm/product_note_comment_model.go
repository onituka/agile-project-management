package productnotecommentdm

import (
	"time"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/groupdm"
	"github.com/onituka/agile-project-management/project-management/domain/productdm"
	"github.com/onituka/agile-project-management/project-management/domain/productnotedm"
	"github.com/onituka/agile-project-management/project-management/domain/userdm"
)

type ProductNoteComment struct {
	id            ProductNoteCommentID
	productID     productdm.ProductID
	productNoteID productnotedm.ProductNoteID
	groupID       groupdm.GroupID
	content       Content
	createdBy     userdm.UserID
	createdAt     time.Time
	updatedAt     time.Time
}

func newProductNoteComment(
	id ProductNoteCommentID,
	productID productdm.ProductID,
	productNoteID productnotedm.ProductNoteID,
	groupID groupdm.GroupID,
	content Content,
	createdBy userdm.UserID,
	createdAt time.Time,
	updatedAt time.Time,
) (*ProductNoteComment, error) {
	if createdAt.IsZero() || updatedAt.IsZero() {
		return nil, apperrors.InvalidParameter
	}

	return &ProductNoteComment{
		id:            id,
		productID:     productID,
		productNoteID: productNoteID,
		groupID:       groupID,
		content:       content,
		createdBy:     createdBy,
		createdAt:     createdAt,
		updatedAt:     updatedAt,
	}, nil
}

func Reconstruct(
	id string,
	productID string,
	productNoteID string,
	groupID string,
	content string,
	createdBy string,
	createdAt time.Time,
	updatedAt time.Time,
) (*ProductNoteComment, error) {
	idVo, err := NewProductNoteCommentID(id)
	if err != nil {
		return nil, err
	}

	productIDVo, err := productdm.NewProductID(productID)
	if err != nil {
		return nil, err
	}

	productNoteIDVo, err := productnotedm.NewProductNoteID(productNoteID)
	if err != nil {
		return nil, err
	}

	groupIDVo, err := groupdm.NewGroupID(groupID)
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

	return newProductNoteComment(
		idVo,
		productIDVo,
		productNoteIDVo,
		groupIDVo,
		contentVo,
		createdByVo,
		createdAt,
		updatedAt,
	)
}

func (p *ProductNoteComment) ID() ProductNoteCommentID {
	return p.id
}

func (p *ProductNoteComment) ProductID() productdm.ProductID {
	return p.productID
}

func (p *ProductNoteComment) ProductNoteID() productnotedm.ProductNoteID {
	return p.productNoteID
}

func (p *ProductNoteComment) GroupID() groupdm.GroupID {
	return p.groupID
}

func (p *ProductNoteComment) Content() Content {
	return p.content
}

func (p *ProductNoteComment) CreatedBy() userdm.UserID {
	return p.createdBy
}

func (p *ProductNoteComment) CreatedAt() time.Time {
	return p.createdAt
}

func (p *ProductNoteComment) UpdatedAt() time.Time {
	return p.updatedAt
}

func (p *ProductNoteComment) CommentAncestorID() ProductNoteCommentID {
	return p.id
}

func (p *ProductNoteComment) CommentDescendantID() ProductNoteCommentID {
	return p.id
}
