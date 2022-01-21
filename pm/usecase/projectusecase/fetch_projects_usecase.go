package projectusecase

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/config"
	"github.com/onituka/agile-project-management/project-management/domain/productdm"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase/projectinput"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase/projectoutput"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase/projectqueryservice"
)

type FetchProjectsUsecase interface {
	FetchProjects(ctx context.Context, in *projectinput.FetchProjectsInput) (*projectoutput.FetchProjectsOutput, error)
}

type fetchProjectsUsecase struct {
	projectQueryService projectqueryservice.ProjectQueryService
	productRepository   productdm.ProductRepository
}

func NewFetchProjectsUsecase(projectQueryService projectqueryservice.ProjectQueryService, productRepository productdm.ProductRepository) *fetchProjectsUsecase {
	return &fetchProjectsUsecase{
		projectQueryService: projectQueryService,
		productRepository:   productRepository,
	}
}

func (u *fetchProjectsUsecase) FetchProjects(ctx context.Context, in *projectinput.FetchProjectsInput) (*projectoutput.FetchProjectsOutput, error) {
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

	totalCount, err := u.projectQueryService.CountProjectsByProductID(ctx, productIDVo)
	if err != nil {
		return nil, err
	} else if totalCount == 0 {
		return &projectoutput.FetchProjectsOutput{
			TotalCount: 0,
			Projects:   make([]*projectoutput.ProjectOutput, 0),
		}, nil
	}

	offset := in.Page*in.Limit - in.Limit

	projectsDto, err := u.projectQueryService.FetchProjects(ctx, productIDVo, uint32(in.Limit), uint32(offset))
	if err != nil {
		return nil, err
	}

	return &projectoutput.FetchProjectsOutput{
		TotalCount: totalCount,
		Projects:   projectsDto,
	}, nil
}
