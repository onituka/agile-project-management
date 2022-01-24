package productnoteusecase

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/productdm"
	"github.com/onituka/agile-project-management/project-management/domain/productnotedm"
	"github.com/onituka/agile-project-management/project-management/usecase/productnoteusecase/productnoteinput"
	"github.com/onituka/agile-project-management/project-management/usecase/productnoteusecase/productnoteoutput"
)

type DeleteProductNoteUsecase interface {
	DeleteProductNote(ctx context.Context, in *productnoteinput.DeleteProductNoteInput) (*productnoteoutput.DeleteProductNoteMsg, error)
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

func (u *deleteProductNoteUsecase) DeleteProductNote(ctx context.Context, in *productnoteinput.DeleteProductNoteInput) (*productnoteoutput.DeleteProductNoteMsg, error) {
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

	if err = u.productNoteRepository.DeleteProductNote(ctx, productNoteIDVo, productIDVo); err != nil {
		return nil, err
	}

	return &productnoteoutput.DeleteProductNoteMsg{Message: "プロダクトノートを削除しました。"}, nil
}
