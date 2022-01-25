package productnoteusecase

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/productdm"
	"github.com/onituka/agile-project-management/project-management/domain/productnotedm"
	"github.com/onituka/agile-project-management/project-management/usecase/productnoteusecase/productnoteinput"
)

type DeleteProductNoteUsecase interface {
	DeleteProductNote(ctx context.Context, in *productnoteinput.DeleteProductNoteInput) error
}

type deleteProductNoteUsecase struct {
	productNoteRepository productnotedm.ProductNoteRepository
	productRepository     productdm.ProductRepository
}

func NewDeleteProductNoteUsecase(ProductNoteRepository productnotedm.ProductNoteRepository, ProductRepository productdm.ProductRepository) *deleteProductNoteUsecase {
	return &deleteProductNoteUsecase{
		productNoteRepository: ProductNoteRepository,
		productRepository:     ProductRepository,
	}
}

func (u *deleteProductNoteUsecase) DeleteProductNote(ctx context.Context, in *productnoteinput.DeleteProductNoteInput) error {
	productIDVo, err := productdm.NewProductID(in.ProductID)
	if err != nil {
		return err
	}

	productDomainService := productdm.NewProductDomainService(u.productRepository)

	if exist, err := productDomainService.ExistsProductByIDForUpdate(ctx, productIDVo); err != nil {
		return err
	} else if !exist {
		return apperrors.NotFound
	}

	productNoteIDVo, err := productnotedm.NewProductNoteID(in.ID)
	if err != nil {
		return err
	}

	productNoteDomainService := productnotedm.NewProductNoteDomainService(u.productNoteRepository)

	if exist, err := productNoteDomainService.ExistsProductNoteByIDForUpdate(ctx, productNoteIDVo, productIDVo); err != nil {
		return err
	} else if !exist {
		return apperrors.NotFound
	}

	if err = u.productNoteRepository.DeleteProductNote(ctx, productNoteIDVo, productIDVo); err != nil {
		return err
	}

	return nil
}
