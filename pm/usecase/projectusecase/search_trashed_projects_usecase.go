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

type SearchTrashedProjectsUsecase interface {
	SearchTrashedProjects(ctx context.Context, in *projectinput.SearchTrashedProjectsInput) (*projectoutput.SearchTrashedProjectsOutput, error)
}

type searchTrashedProjectsUsecase struct {
	projectQueryService projectqueryservice.ProjectQueryService
	productRepository   productdm.ProductRepository
}

func NewSearchTrashedProjectsUsecase(projectQueryService projectqueryservice.ProjectQueryService, productRepository productdm.ProductRepository) *searchTrashedProjectsUsecase {
	return &searchTrashedProjectsUsecase{
		projectQueryService: projectQueryService,
		productRepository:   productRepository,
	}
}

func (u *searchTrashedProjectsUsecase) SearchTrashedProjects(ctx context.Context, in *projectinput.SearchTrashedProjectsInput) (*projectoutput.SearchTrashedProjectsOutput, error) {
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

	totalCount, err := u.projectQueryService.CountTrashedProjectsByKeyNameAndName(ctx, productIDVo, in.KeyWord)
	if err != nil {
		return nil, err
	} else if totalCount == 0 {
		return &projectoutput.SearchTrashedProjectsOutput{
			TotalCount: 0,
			Projects:   make([]*projectoutput.SearchTrashedProjectOutput, 0),
		}, nil
	}

	offset := in.Page*in.Limit - in.Limit

	projectsDto, err := u.projectQueryService.SearchTrashedProjects(ctx, productIDVo, in.KeyWord, in.Limit, offset)
	if err != nil {
		return nil, err
	}

	return &projectoutput.SearchTrashedProjectsOutput{
		TotalCount: totalCount,
		Projects:   projectsDto,
	}, nil
}
