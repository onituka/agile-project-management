package productnoteusecase

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/config"
	"github.com/onituka/agile-project-management/project-management/domain/productdm"
	"github.com/onituka/agile-project-management/project-management/usecase/productnoteusecase/productnoteinput"
	"github.com/onituka/agile-project-management/project-management/usecase/productnoteusecase/productnoteoutput"
	"github.com/onituka/agile-project-management/project-management/usecase/productnoteusecase/productnotequeryservice"
)

type FetchProductNotesUsecase interface {
	FetchProductNotes(ctx context.Context, in *productnoteinput.FetchProductNotesInput) (*productnoteoutput.FetchProductNotesOutput, error)
}

type fetchProductNotesUsecase struct {
	productnoteQueryService productnotequeryservice.ProductNoteQueryService
	productRepository       productdm.ProductRepository
}

func NewFetchProductNotesUsecase(productnoteQueryService productnotequeryservice.ProductNoteQueryService, productRepository productdm.ProductRepository) *fetchProductNotesUsecase {
	return &fetchProductNotesUsecase{
		productnoteQueryService: productnoteQueryService,
		productRepository:       productRepository,
	}
}

func (u *fetchProductNotesUsecase) FetchProductNotes(ctx context.Context, in *productnoteinput.FetchProductNotesInput) (*productnoteoutput.FetchProductNotesOutput, error) {
	if in.Page == 0 || in.Limit == 0 || in.Limit > config.LimitPerPage {
		return nil, apperrors.InvalidParameter
	}

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

	totalCount, err := u.productnoteQueryService.CountProductNotesByProductID(ctx, productIDVo)
	if err != nil {
		return nil, err
	} else if totalCount == 0 {
		return &productnoteoutput.FetchProductNotesOutput{
			TotalCount:   0,
			ProductNotes: make([]*productnoteoutput.ProductNoteOutput, 0),
		}, nil
	}

	offset := in.Page*in.Limit - in.Limit

	productnotesDto, err := u.productnoteQueryService.FetchProductNotes(ctx, productIDVo, in.Limit, offset)
	if err != nil {
		return nil, err
	}

	return &productnoteoutput.FetchProductNotesOutput{
		TotalCount:   totalCount,
		ProductNotes: productnotesDto,
	}, nil
}
