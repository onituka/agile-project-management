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

type SearchProductNotesUsecase interface {
	SearchProductNotes(ctx context.Context, in *productnoteinput.SearchProductNotesInput) (*productnoteoutput.SearchProductNotesOutput, error)
}

type searchProductNotesUsecase struct {
	productNoteQueryService productnotequeryservice.ProductNoteQueryService
	productRepository       productdm.ProductRepository
}

func NewSearchProductNotesUsecase(productNoteQueryService productnotequeryservice.ProductNoteQueryService, productRepository productdm.ProductRepository) *searchProductNotesUsecase {
	return &searchProductNotesUsecase{
		productNoteQueryService: productNoteQueryService,
		productRepository:       productRepository,
	}
}

func (u *searchProductNotesUsecase) SearchProductNotes(ctx context.Context, in *productnoteinput.SearchProductNotesInput) (*productnoteoutput.SearchProductNotesOutput, error) {
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

	totalCount, err := u.productNoteQueryService.CountProductNotesByTitle(ctx, productIDVo, in.Title)
	if err != nil {
		return nil, err
	} else if totalCount == 0 {
		return &productnoteoutput.SearchProductNotesOutput{
			TotalCount:   0,
			ProductNotes: make([]*productnoteoutput.SearchProductNoteOutput, 0),
		}, nil
	}

	offset := in.Page*in.Limit - in.Limit

	productNotesDto, err := u.productNoteQueryService.SearchProductNotes(ctx, productIDVo, in.Title, in.Limit, offset)
	if err != nil {
		return nil, err
	}

	return &productnoteoutput.SearchProductNotesOutput{
		TotalCount:   totalCount,
		ProductNotes: productNotesDto,
	}, nil
}
