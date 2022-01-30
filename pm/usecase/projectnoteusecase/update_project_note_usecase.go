package projectnoteusecase

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/productdm"
	"github.com/onituka/agile-project-management/project-management/domain/projectdm"
	"github.com/onituka/agile-project-management/project-management/domain/projectnotedm"
	"github.com/onituka/agile-project-management/project-management/domain/userdm"
	"github.com/onituka/agile-project-management/project-management/usecase/projectnoteusecase/projectnoteinput"
	"github.com/onituka/agile-project-management/project-management/usecase/projectnoteusecase/projectnoteoutput"
)

type UpdateProjectNoteUsecase interface {
	UpdateProjectNote(ctx context.Context, in *projectnoteinput.UpdateProjectNoteInput) (*projectnoteoutput.UpdateProjectNoteOutput, error)
}

type updateProjectNoteUsecase struct {
	projectNoteRepository projectnotedm.ProjectNoteRepository
	productRepository     productdm.ProductRepository
	projectRepository     projectdm.ProjectRepository
}

func NewUpdateProjectNoteUsecase(projectNoteRepository projectnotedm.ProjectNoteRepository, productRepository productdm.ProductRepository, projectRepository projectdm.ProjectRepository) *updateProjectNoteUsecase {
	return &updateProjectNoteUsecase{
		projectNoteRepository: projectNoteRepository,
		productRepository:     productRepository,
		projectRepository:     projectRepository,
	}
}

func (u *updateProjectNoteUsecase) UpdateProjectNote(ctx context.Context, in *projectnoteinput.UpdateProjectNoteInput) (*projectnoteoutput.UpdateProjectNoteOutput, error) {
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

	projectNoteIDVo, err := projectnotedm.NewProjectNoteID(in.ID)
	if err != nil {
		return nil, err
	}

	projectNoteDomainService := projectnotedm.NewProjectNoteDomainService(u.projectNoteRepository)

	if exist, err := projectNoteDomainService.ExistsProjectNoteByIDForUpdate(ctx, projectNoteIDVo, projectIDVo); err != nil {
		return nil, err
	} else if !exist {
		return nil, apperrors.NotFound
	}

	projectNoteDm, err := u.projectNoteRepository.FetchProjectNoteByID(ctx, projectNoteIDVo, projectIDVo)
	if err != nil {
		return nil, err
	}

	titleVo, err := projectnotedm.NewTitle(in.Title)
	if err != nil {
		return nil, err
	}

	projectNoteDm.ChangeTitle(titleVo)

	contentVo, err := projectnotedm.NewContent(in.Content)
	if err != nil {
		return nil, err
	}

	projectNoteDm.ChangeContent(contentVo)

	userIDVo, err := userdm.NewUserID(in.UserID)
	if err != nil {
		return nil, err
	}

	projectNoteDm.ChangeUpdatedBy(userIDVo)

	exist, err := projectNoteDomainService.ExistsProjectNoteForUpdate(ctx, projectNoteDm)
	if err != nil && !apperrors.Is(err, apperrors.NotFound) {
		return nil, err
	} else if exist {
		return nil, apperrors.Conflict
	}

	if err = u.projectNoteRepository.UpdateProjectNote(ctx, projectNoteDm); err != nil {
		return nil, err
	}

	return &projectnoteoutput.UpdateProjectNoteOutput{
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
