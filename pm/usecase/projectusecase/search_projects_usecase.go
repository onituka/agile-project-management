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

type SearchProjectsUsecase interface {
	SearchProjects(ctx context.Context, in *projectinput.SearchProjectsInput) (*projectoutput.SearchProjectsOutput, error)
}

type searchProjectsUsecase struct {
	projectQueryService projectqueryservice.ProjectQueryService
	productRepository   productdm.ProductRepository
}

func NewSearchProjectsUsecase(projectQueryService projectqueryservice.ProjectQueryService, productRepository productdm.ProductRepository) *searchProjectsUsecase {
	return &searchProjectsUsecase{
		projectQueryService: projectQueryService,
		productRepository:   productRepository,
	}
}

func (u *searchProjectsUsecase) SearchProjects(ctx context.Context, in *projectinput.SearchProjectsInput) (*projectoutput.SearchProjectsOutput, error) {
	if in.Page == 0 || in.Limit == 0 || in.Limit > config.LimitPerPage {
		return nil, apperrors.InvalidParameter
	}

	productIDVo, err := productdm.NewProductID(in.ProductID)
	if err != nil {
		return nil, err
	}

	if _, err := u.productRepository.FetchProductByIDForUpdate(ctx, productIDVo); err != nil {
		return nil, err
	}

	totalCount, err := u.projectQueryService.CountProjectsByKeyNameAndName(ctx, productIDVo, in.KeyWord)
	if err != nil {
		return nil, err
	} else if totalCount == 0 {
		return &projectoutput.SearchProjectsOutput{
			TotalCount: 0,
			Projects:   make([]*projectoutput.SearchProjectOutput, 0),
		}, nil
	}

	offset := in.Page*in.Limit - in.Limit

	projectsDto, err := u.projectQueryService.SearchProjects(ctx, productIDVo, in.KeyWord, in.Limit, offset)
	if err != nil {
		return nil, err
	}

	return &projectoutput.SearchProjectsOutput{
		TotalCount: totalCount,
		Projects:   projectsDto,
	}, nil
}
