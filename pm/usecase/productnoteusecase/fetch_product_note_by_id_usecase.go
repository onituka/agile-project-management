package productnoteusecase

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/productdm"
	"github.com/onituka/agile-project-management/project-management/domain/productnotedm"
	"github.com/onituka/agile-project-management/project-management/usecase/productnoteusecase/productnoteinput"
	"github.com/onituka/agile-project-management/project-management/usecase/productnoteusecase/productnoteoutput"
)

type FetchProductNoteByIDUsecase interface {
	FetchProductNoteByID(ctx context.Context, in *productnoteinput.FetchProductNoteByIDInput) (*productnoteoutput.FetchProductNoteByIDOutput, error)
}

type fetchProductNoteByIDUsecase struct {
	productNoteRepository productnotedm.ProductNoteRepository
	productRepository     productdm.ProductRepository
}

func NewFetchProductNoteByIDUsecase(FetchProductNoteByIDRepository productnotedm.ProductNoteRepository, productRepository productdm.ProductRepository) *fetchProductNoteByIDUsecase {
	return &fetchProductNoteByIDUsecase{
		productNoteRepository: FetchProductNoteByIDRepository,
		productRepository:     productRepository,
	}
}

func (u *fetchProductNoteByIDUsecase) FetchProductNoteByID(ctx context.Context, in *productnoteinput.FetchProductNoteByIDInput) (*productnoteoutput.FetchProductNoteByIDOutput, error) {
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

	productNoteIDVo, err := productnotedm.NewProductNoteID(in.ID)
	if err != nil {
		return nil, err
	}

	productNoteDomainService := productnotedm.NewProductNoteDomainService(u.productNoteRepository)

	if exist, err := productNoteDomainService.ExistsProductNoteByIDForUpdate(ctx, productNoteIDVo, productIDVo); err != nil {
		return nil, err
	} else if !exist {
		return nil, apperrors.NotFound
	}

	productNoteDm, err := u.productNoteRepository.FetchProductNoteByID(ctx, productNoteIDVo, productIDVo)
	if err != nil {
		return nil, err
	}

	return &productnoteoutput.FetchProductNoteByIDOutput{
		ID:        productNoteDm.ID().Value(),
		ProductID: productNoteDm.ProductID().Value(),
		GroupID:   productNoteDm.GroupID().Value(),
		Title:     productNoteDm.Title().Value(),
		Content:   productNoteDm.Content().Value(),
		CreatedBy: productNoteDm.CreatedBy().Value(),
		UpdatedBy: productNoteDm.UpdatedBy().Value(),
		CreatedAt: productNoteDm.CreatedAt(),
		UpdatedAt: productNoteDm.UpdatedAt(),
	}, nil
}
