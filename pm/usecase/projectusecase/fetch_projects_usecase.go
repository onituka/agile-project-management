package projectusecase

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/apperrors"
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
}

func NewFetchProjectsUsecase(projectQueryService projectqueryservice.ProjectQueryService) *fetchProjectsUsecase {
	return &fetchProjectsUsecase{
		projectQueryService: projectQueryService,
	}
}

func (u *fetchProjectsUsecase) FetchProjects(ctx context.Context, in *projectinput.FetchProjectsInput) (*projectoutput.FetchProjectsOutput, error) {
	productIDVo, err := productdm.NewProductID(in.ProductID)
	if err != nil {
		return nil, err
	}

	if in.Page <= 0 || in.Limit <= 0 || in.Limit > 50 {
		return nil, apperrors.InvalidParameter
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
