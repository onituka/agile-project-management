package productnoteusecase

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/productdm"
	"github.com/onituka/agile-project-management/project-management/domain/productnotedm"
	"github.com/onituka/agile-project-management/project-management/domain/userdm"
	"github.com/onituka/agile-project-management/project-management/usecase/productnoteusecase/productnoteinput"
	"github.com/onituka/agile-project-management/project-management/usecase/productnoteusecase/productnoteoutput"
)

type UpdateProductNoteUsecase interface {
	UpdateProductNote(ctx context.Context, in *productnoteinput.UpdateProductNoteInput) (*productnoteoutput.UpdateProductNoteOutput, error)
}

type updateProductNoteUsecase struct {
	productNoteRepository productnotedm.ProductNoteRepository
	productRepository     productdm.ProductRepository
}

func NewUpdateProductNoteUsecase(ProductNoteRepository productnotedm.ProductNoteRepository, ProductRepository productdm.ProductRepository) *updateProductNoteUsecase {
	return &updateProductNoteUsecase{
		productNoteRepository: ProductNoteRepository,
		productRepository:     ProductRepository,
	}
}

func (u *updateProductNoteUsecase) UpdateProductNote(ctx context.Context, in *productnoteinput.UpdateProductNoteInput) (*productnoteoutput.UpdateProductNoteOutput, error) {
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

	titleVo, err := productnotedm.NewTitle(in.Title)
	if err != nil {
		return nil, err
	}

	productNoteDm.ChangeTitle(titleVo)

	contentVo, err := productnotedm.NewContent(in.Content)
	if err != nil {
		return nil, err
	}

	productNoteDm.ChangeContent(contentVo)

	UserIDVo, err := userdm.NewUserID(in.UserID)
	if err != nil {
		return nil, err
	}

	productNoteDm.ChangeUpdatedBy(UserIDVo)

	productNoteDm.ChangeUpdateAt()

	exist, err := productNoteDomainService.ExistsProductNoteForUpdate(ctx, productNoteDm)
	if err != nil && !apperrors.Is(err, apperrors.NotFound) {
		return nil, err
	} else if exist {
		return nil, apperrors.Conflict
	}

	if err = u.productNoteRepository.UpdateProductNote(ctx, productNoteDm); err != nil {
		return nil, err
	}

	return &productnoteoutput.UpdateProductNoteOutput{
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
