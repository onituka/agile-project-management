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

type SearchProjectNotesUsecase interface {
	SearchProjectNotes(ctx context.Context, in *projectnoteinput.SearchProjectNotesInput) (*projectnoteoutput.SearchProjectNotesOutput, error)
}

type searchProjectNotesUsecase struct {
	projectNoteQueryService projectnotequeryservice.ProjectNoteQueryService
	productRepository       productdm.ProductRepository
	projectRepository       projectdm.ProjectRepository
}

func NewSearchProjectNotesUsecase(projectNoteQueryService projectnotequeryservice.ProjectNoteQueryService, productRepository productdm.ProductRepository, projectRepository projectdm.ProjectRepository) *searchProjectNotesUsecase {
	return &searchProjectNotesUsecase{
		projectNoteQueryService: projectNoteQueryService,
		productRepository:       productRepository,
		projectRepository:       projectRepository,
	}
}

func (u *searchProjectNotesUsecase) SearchProjectNotes(ctx context.Context, in *projectnoteinput.SearchProjectNotesInput) (*projectnoteoutput.SearchProjectNotesOutput, error) {
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

	totalCount, err := u.projectNoteQueryService.CountProjectNotesByTitle(ctx, productIDVo, projectIDVo, in.Title)
	if err != nil {
		return nil, err
	} else if totalCount == 0 {
		return &projectnoteoutput.SearchProjectNotesOutput{
			TotalCount:   0,
			ProjectNotes: make([]*projectnoteoutput.SearchProjectNoteOutPut, 0),
		}, nil
	}

	offset := in.Page*in.Limit - in.Limit

	projectNotesDto, err := u.projectNoteQueryService.SearchProjectNotes(ctx, productIDVo, projectIDVo, in.Title, in.Limit, offset)
	if err != nil {
		return nil, err
	}

	return &projectnoteoutput.SearchProjectNotesOutput{
		TotalCount:   totalCount,
		ProjectNotes: projectNotesDto,
	}, nil
}
