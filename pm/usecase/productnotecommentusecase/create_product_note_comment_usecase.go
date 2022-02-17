package productnotecommentusecase

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/groupdm"
	"github.com/onituka/agile-project-management/project-management/domain/productdm"
	"github.com/onituka/agile-project-management/project-management/domain/productnotecommentdm"
	"github.com/onituka/agile-project-management/project-management/domain/productnotedm"
	"github.com/onituka/agile-project-management/project-management/domain/userdm"
	"github.com/onituka/agile-project-management/project-management/usecase/productnotecommentusecase/productnotecommentinput"
	"github.com/onituka/agile-project-management/project-management/usecase/productnotecommentusecase/productnotecommentoutput"
)

type CreateProductNoteCommentUsecase interface {
	CreateProductNoteComment(ctx context.Context, in *productnotecommentinput.CreateProductNoteCommentInput) (*productnotecommentoutput.CreateProductNoteCommentOutput, error)
}

type createProductNoteCommentUsecase struct {
	productNoteCommentRepository productnotecommentdm.ProductNoteCommentRepository
	productNoteRepository        productnotedm.ProductNoteRepository
	productRepository            productdm.ProductRepository
}

func NewCreateProductNoteCommentUsecase(ProductNoteCommentRepository productnotecommentdm.ProductNoteCommentRepository, ProductNoteRepository productnotedm.ProductNoteRepository, ProductRepository productdm.ProductRepository) *createProductNoteCommentUsecase {
	return &createProductNoteCommentUsecase{
		productNoteCommentRepository: ProductNoteCommentRepository,
		productNoteRepository:        ProductNoteRepository,
		productRepository:            ProductRepository,
	}
}

func (u *createProductNoteCommentUsecase) CreateProductNoteComment(ctx context.Context, in *productnotecommentinput.CreateProductNoteCommentInput) (*productnotecommentoutput.CreateProductNoteCommentOutput, error) {
	productIDVo, err := productdm.NewProductID(in.ProductID)
	if err != nil {
		return nil, err
	}

	productDomainService := productdm.NewProductDomainService(u.productRepository)

	if exist, err := productDomainService.ExistsProductByIDForUpdate(ctx, productIDVo); err != nil {
		return nil, err
	} else if !exist {
		return nil, apperrors.NotFound
	}

	productNoteIDVo, err := productnotedm.NewProductNoteID(in.ProductNoteID)
	if err != nil {
		return nil, err
	}

	productNoteDomainService := productnotedm.NewProductNoteDomainService(u.productNoteRepository)

	if exist, err := productNoteDomainService.ExistsProductNoteByIDForUpdate(ctx, productNoteIDVo, productIDVo); err != nil {
		return nil, err
	} else if !exist {
		return nil, apperrors.NotFound
	}

	groupIDVo, err := groupdm.NewGroupID(in.GroupID)
	if err != nil {
		return nil, err
	}

	contentVo, err := productnotecommentdm.NewContent(in.Content)
	if err != nil {
		return nil, err
	}

	userIDVo, err := userdm.NewUserID(in.UserID)
	if err != nil {
		return nil, err
	}

	productNoteCommentDm, err := productnotecommentdm.GenProductNoteCommentForCreate(
		productIDVo,
		productNoteIDVo,
		groupIDVo,
		contentVo,
		userIDVo,
	)
	if err != nil {
		return nil, err
	}

	if err = u.productNoteCommentRepository.CreateProductNoteComment(ctx, productNoteCommentDm); err != nil {
		return nil, err
	}

	return &productnotecommentoutput.CreateProductNoteCommentOutput{
		ID:            productNoteCommentDm.ID().Value(),
		ProductID:     productNoteCommentDm.ProductID().Value(),
		ProductNoteID: productNoteCommentDm.ProductNoteID().Value(),
		GroupID:       productNoteCommentDm.GroupID().Value(),
		Content:       productNoteCommentDm.Content().Value(),
		CreatedBy:     productNoteCommentDm.CreatedBy().Value(),
		CreatedAt:     productNoteCommentDm.CreatedAt(),
		UpdatedAt:     productNoteCommentDm.UpdatedAt(),
	}, nil
}
