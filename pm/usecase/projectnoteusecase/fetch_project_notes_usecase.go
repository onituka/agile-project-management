package projectnoteusecase

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/config"
	"github.com/onituka/agile-project-management/project-management/domain/productdm"
	"github.com/onituka/agile-project-management/project-management/domain/projectdm"
	"github.com/onituka/agile-project-management/project-management/usecase/projectnoteusecase/projectnoteinput"
	"github.com/onituka/agile-project-management/project-management/usecase/projectnoteusecase/projectnoteoutput"
	"github.com/onituka/agile-project-management/project-management/usecase/projectnoteusecase/projectnotequeryservice"
)

type FetchProjectNotesUsecase interface {
	FetchProjectNotes(ctx context.Context, in *projectnoteinput.FetchProjectNotesInput) (*projectnoteoutput.FetchProjectNotesOutput, error)
}

type fetchProjectNotesUsecase struct {
	projectNoteQueryService projectnotequeryservice.ProjectNoteQueryService
	productRepository       productdm.ProductRepository
	projectRepository       projectdm.ProjectRepository
}

func NewFetchProjectNotesUsecase(projectQueryService projectnotequeryservice.ProjectNoteQueryService, productRepository productdm.ProductRepository, projectRepository projectdm.ProjectRepository) *fetchProjectNotesUsecase {
	return &fetchProjectNotesUsecase{
		projectNoteQueryService: projectQueryService,
		productRepository:       productRepository,
		projectRepository:       projectRepository,
	}
}

func (u *fetchProjectNotesUsecase) FetchProjectNotes(ctx context.Context, in *projectnoteinput.FetchProjectNotesInput) (*projectnoteoutput.FetchProjectNotesOutput, error) {
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

	projectIDVo, err := projectdm.NewProjectID(in.ProjectID)
	if err != nil {
		return nil, err
	}

	projectDomainService := projectdm.NewProjectDomainService(u.projectRepository)

	if exist, err := projectDomainService.ExistsProjectByIDForUpdate(ctx, projectIDVo, productIDVo); err != nil {
		return nil, err
	} else if !exist {
		return nil, apperrors.NotFound
	}

	totalCount, err := u.projectNoteQueryService.CountProjectNotesByProductIDAndProjectID(ctx, productIDVo, projectIDVo)
	if err != nil {
		return nil, err
	} else if totalCount == 0 {
		return &projectnoteoutput.FetchProjectNotesOutput{
			TotalCount:   0,
			ProjectNotes: make([]*projectnoteoutput.ProjectNoteOutput, 0),
		}, nil
	}

	offset := in.Page*in.Limit - in.Limit

	projectNotesDto, err := u.projectNoteQueryService.FetchProjectNotes(ctx, productIDVo, projectIDVo, in.Limit, offset)
	if err != nil {
		return nil, err
	}

	return &projectnoteoutput.FetchProjectNotesOutput{
		TotalCount:   totalCount,
		ProjectNotes: projectNotesDto,
	}, nil
}
