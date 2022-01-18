package productnoteusecase

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/groupdm"
	"github.com/onituka/agile-project-management/project-management/domain/productdm"
	"github.com/onituka/agile-project-management/project-management/domain/productnotedm"
	"github.com/onituka/agile-project-management/project-management/domain/userdm"
	"github.com/onituka/agile-project-management/project-management/usecase/productnoteusecase/productnoteinput"
	"github.com/onituka/agile-project-management/project-management/usecase/productnoteusecase/productnoteoutput"
)

type CreateProductNoteUsecase interface {
	CreateProductNote(ctx context.Context, in *productnoteinput.CreateProductNoteInput) (*productnoteoutput.CreateProductNoteOutput, error)
}

type createProductNoteUsecase struct {
	productNoteRepository productnotedm.ProductNoteRepository
	productRepository     productdm.ProductRepository
}

func NewCreateProductNoteUsecase(ProductNoteRepository productnotedm.ProductNoteRepository, ProductRepository productdm.ProductRepository) *createProductNoteUsecase {
	return &createProductNoteUsecase{
		productNoteRepository: ProductNoteRepository,
		productRepository:     ProductRepository,
	}
}

func (u *createProductNoteUsecase) CreateProductNote(ctx context.Context, in *productnoteinput.CreateProductNoteInput) (*productnoteoutput.CreateProductNoteOutput, error) {
	productIDVo, err := productdm.NewProductID(in.ProductID)
	if err != nil {
		return nil, err
	}

	if _, err := u.productRepository.FetchProductByIDForUpdate(ctx, productIDVo); err != nil {
		return nil, err
	}

	groupIDVo, err := groupdm.NewGroupID(in.GroupID)
	if err != nil {
		return nil, err
	}

	titleVo, err := productnotedm.NewTitle(in.Title)
	if err != nil {
		return nil, err
	}

	contentVo, err := productnotedm.NewContent(in.Content)
	if err != nil {
		return nil, err
	}

	createdBy, err := userdm.NewUserID(in.CreatedBy)
	if err != nil {
		return nil, err
	}

	UpdatedBy, err := userdm.NewUserID(in.UpdatedBy)
	if err != nil {
		return nil, err
	}

	productnoteDm, err := productnotedm.GenProductNoteForCreate(
		productIDVo,
		groupIDVo,
		titleVo,
		contentVo,
		createdBy,
		UpdatedBy,
	)
	if err != nil {
		return nil, err
	}

	productNoteDomainService := productnotedm.NewProductNoteDomainService(u.productNoteRepository)

	exist, err := productNoteDomainService.ExistsProductNoteForCreate(ctx, productnoteDm)
	if err != nil && !apperrors.Is(err, apperrors.NotFound) {
		return nil, err
	} else if exist {
		return nil, apperrors.Conflict
	}

	if err = u.productNoteRepository.CreateProductNote(ctx, productnoteDm); err != nil {
		return nil, err
	}

	return &productnoteoutput.CreateProductNoteOutput{
		ID:        productnoteDm.ID().Value(),
		ProductID: productnoteDm.ProductID().Value(),
		GroupID:   productnoteDm.GroupID().Value(),
		Title:     productnoteDm.Title().Value(),
		Content:   productnoteDm.Content().Value(),
		CreatedBy: productnoteDm.CreatedBy().Value(),
		UpdatedBy: productnoteDm.UpdatedBy().Value(),
		CreatedAt: productnoteDm.CreatedAt(),
		UpdatedAt: productnoteDm.UpdatedAt(),
	}, nil
}
