package projectnoteusecase

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/groupdm"
	"github.com/onituka/agile-project-management/project-management/domain/productdm"
	"github.com/onituka/agile-project-management/project-management/domain/projectdm"
	"github.com/onituka/agile-project-management/project-management/domain/projectnotedm"
	"github.com/onituka/agile-project-management/project-management/domain/userdm"
	"github.com/onituka/agile-project-management/project-management/usecase/projectnoteusecase/projectnoteinput"
	"github.com/onituka/agile-project-management/project-management/usecase/projectnoteusecase/projectnoteoutput"
)

type CreateProjectNoteUsecase interface {
	CreateProjectNote(ctx context.Context, in *projectnoteinput.CreateProjectNoteInput) (*projectnoteoutput.CreateProjectNoteOutput, error)
}

type createProjectNoteUsecase struct {
	projectNoteRepository projectnotedm.ProjectNoteRepository
	productRepository     productdm.ProductRepository
	projectRepository     projectdm.ProjectRepository
}

func NewCreateProjectNoteUsecase(ProjectNoteRepository projectnotedm.ProjectNoteRepository, ProductRepository productdm.ProductRepository, ProjectRepository projectdm.ProjectRepository) *createProjectNoteUsecase {
	return &createProjectNoteUsecase{
		projectNoteRepository: ProjectNoteRepository,
		productRepository:     ProductRepository,
		projectRepository:     ProjectRepository,
	}
}

func (u *createProjectNoteUsecase) CreateProjectNote(ctx context.Context, in *projectnoteinput.CreateProjectNoteInput) (*projectnoteoutput.CreateProjectNoteOutput, error) {
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

	groupIDVo, err := groupdm.NewGroupID(in.GroupID)
	if err != nil {
		return nil, err
	}

	titleVo, err := projectnotedm.NewTitle(in.Title)
	if err != nil {
		return nil, err
	}

	contentVo, err := projectnotedm.NewContent(in.Content)
	if err != nil {
		return nil, err
	}

	createdBy, err := userdm.NewUserID(in.CreatedBy)
	if err != nil {
		return nil, err
	}

	updatedBy, err := userdm.NewUserID(in.UpdatedBy)
	if err != nil {
		return nil, err
	}

	projectNoteDm, err := projectnotedm.GenProjectNoteForCreate(
		productIDVo,
		projectIDVo,
		groupIDVo,
		titleVo,
		contentVo,
		createdBy,
		updatedBy,
	)
	if err != nil {
		return nil, err
	}

	projectNoteDomainService := projectnotedm.NewProjectNoteDomainService(u.projectNoteRepository)

	exist, err := projectNoteDomainService.ExistsProjectNoteForCreate(ctx, projectNoteDm)
	if err != nil && !apperrors.Is(err, apperrors.NotFound) {
		return nil, err
	} else if exist {
		return nil, apperrors.Conflict
	}

	if err = u.projectNoteRepository.CreateNoteProject(ctx, projectNoteDm); err != nil {
		return nil, err
	}

	return &projectnoteoutput.CreateProjectNoteOutput{
		ID:        projectNoteDm.ID().Value(),
		ProductID: projectNoteDm.ProductID().Value(),
		ProjectID: projectNoteDm.ProjectID().Value(),
		GroupID:   projectNoteDm.GroupID().Value(),
		Title:     projectNoteDm.Title().Value(),
		Content:   projectNoteDm.Content().Value(),
		CreatedBy: projectNoteDm.CreatedBy().Value(),
		UpdatedBy: projectNoteDm.UpdatedBy().Value(),
		CreatedAt: projectNoteDm.CreatedAt(),
		UpdatedAt: projectNoteDm.UpdatedAt(),
	}, nil
}
