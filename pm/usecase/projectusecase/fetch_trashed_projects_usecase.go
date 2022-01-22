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

type FetchTrashedProjectsUsecase interface {
	FetchTrashedProjects(ctx context.Context, in *projectinput.FetchTrashedProjectsInput) (*projectoutput.FetchTrashedProjectsOutput, error)
}

type fetchTrashedProjectsUsecase struct {
	projectQueryService projectqueryservice.ProjectQueryService
	productRepository   productdm.ProductRepository
}

func NewFetchTrashedProjectsUsecase(projectQueryService projectqueryservice.ProjectQueryService, productRepository productdm.ProductRepository) *fetchTrashedProjectsUsecase {
	return &fetchTrashedProjectsUsecase{
		projectQueryService: projectQueryService,
		productRepository:   productRepository,
	}
}

func (u *fetchTrashedProjectsUsecase) FetchTrashedProjects(ctx context.Context, in *projectinput.FetchTrashedProjectsInput) (*projectoutput.FetchTrashedProjectsOutput, error) {
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

	totalCount, err := u.projectQueryService.CountTrashedProjectsByProductID(ctx, productIDVo)
	if err != nil {
		return nil, err
	} else if totalCount == 0 {
		return &projectoutput.FetchTrashedProjectsOutput{
			TotalCount: 0,
			Projects:   make([]*projectoutput.FetchTrashedProjectOutput, 0),
		}, nil
	}

	offset := in.Page*in.Limit - in.Limit

	projectsDto, err := u.projectQueryService.FetchTrashedProjects(ctx, productIDVo, uint32(in.Limit), uint32(offset))
	if err != nil {
		return nil, err
	}

	return &projectoutput.FetchTrashedProjectsOutput{
		TotalCount: totalCount,
		Projects:   projectsDto,
	}, nil
}
