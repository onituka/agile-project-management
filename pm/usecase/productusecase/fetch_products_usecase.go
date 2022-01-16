package productusecase

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/groupdm"
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase/productinput"
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase/productoutput"
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase/productqueryservice"
)

type FetchProductsUsecase interface {
	FetchProducts(ctx context.Context, in *productinput.FetchProductsInput) (*productoutput.FetchProductsOutput, error)
}

type fetchProductsUsecase struct {
	productQueryService productqueryservice.ProductQueryService
}

func NewFetchProductsUsecase(FetchProductsQueryService productqueryservice.ProductQueryService) *fetchProductsUsecase {
	return &fetchProductsUsecase{
		productQueryService: FetchProductsQueryService,
	}
}

func (u *fetchProductsUsecase) FetchProducts(ctx context.Context, in *productinput.FetchProductsInput) (*productoutput.FetchProductsOutput, error) {
	groupIDVo, err := groupdm.NewGroupID(in.GroupID)
	if err != nil {
		return nil, err
	}

	if in.Page == 0 || in.Limit == 0 || 50 < in.Limit {
		return nil, apperrors.InvalidParameter
	}

	totalCount, err := u.productQueryService.CountProductsByGroupID(ctx, groupIDVo)
	if err != nil {
		return nil, err
	} else if totalCount == 0 {
		return &productoutput.FetchProductsOutput{
			TotalCount: 0,
			Products:   make([]*productoutput.ProductOutput, 0),
		}, nil
	}

	offset := in.Page*in.Limit - in.Limit

	productsDto, err := u.productQueryService.FetchProducts(ctx, groupIDVo, in.Limit, offset)
	if err != nil {
		return nil, err
	}

	return &productoutput.FetchProductsOutput{
		TotalCount: totalCount,
		Products:   productsDto,
	}, nil
}
