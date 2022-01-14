package productusecase

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/groupdm"
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase/productinput"
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase/productoutput"
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase/productqueryservice"
)

type SearchProductsUsecase interface {
	SearchProducts(ctx context.Context, in *productinput.SearchProductsInput) (*productoutput.SearchProductsOutput, error)
}

type searchProductsUsecase struct {
	productQueryService productqueryservice.ProductQueryService
}

func NewSearchProductsUsecase(SearchProductsQueryService productqueryservice.ProductQueryService) *searchProductsUsecase {
	return &searchProductsUsecase{
		productQueryService: SearchProductsQueryService,
	}
}

func (u *searchProductsUsecase) SearchProducts(ctx context.Context, in *productinput.SearchProductsInput) (*productoutput.SearchProductsOutput, error) {
	if _, err := groupdm.NewGroupID(in.GroupID); err != nil {
		return nil, err
	}

	if in.Page == 0 || in.Limit == 0 || 50 < in.Limit {
		return nil, apperrors.InvalidParameter
	}

	totalCount, err := u.productQueryService.CountProductsByName(ctx, in.GroupID, in.ProductName)
	if err != nil {
		return nil, err
	} else if totalCount == 0 {
		return &productoutput.SearchProductsOutput{
			TotalCount: 0,
			Products:   make([]*productoutput.SearchProductOutput, 0),
		}, nil
	}

	offset := in.Page*in.Limit - in.Limit

	productsDto, err := u.productQueryService.SearchProducts(ctx, in.GroupID, in.ProductName, in.Limit, offset)
	if err != nil {
		return nil, err
	}

	return &productoutput.SearchProductsOutput{
		TotalCount: totalCount,
		Products:   productsDto,
	}, nil
}
