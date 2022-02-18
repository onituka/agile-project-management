package projectnoteusecase

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/productdm"
	"github.com/onituka/agile-project-management/project-management/domain/projectdm"
	"github.com/onituka/agile-project-management/project-management/domain/projectnotedm"
	"github.com/onituka/agile-project-management/project-management/usecase/projectnoteusecase/projectnoteinput"
)

type DeleteProjectNoteUsecase interface {
	DeleteProjectNote(ctx context.Context, in *projectnoteinput.DeleteProjectNoteInput) error
}

type deleteProjectNoteUsecase struct {
	projectNoteRepository projectnotedm.ProjectNoteRepository
	productRepository     productdm.ProductRepository
	projectRepository     projectdm.ProjectRepository
}

func NewDeleteProjectNoteUsecase(ProjectNoteRepository projectnotedm.ProjectNoteRepository, ProductRepository productdm.ProductRepository, ProjectRepository projectdm.ProjectRepository) *deleteProjectNoteUsecase {
	return &deleteProjectNoteUsecase{
		projectNoteRepository: ProjectNoteRepository,
		productRepository:     ProductRepository,
		projectRepository:     ProjectRepository,
	}
}

func (u *deleteProjectNoteUsecase) DeleteProjectNote(ctx context.Context, in *projectnoteinput.DeleteProjectNoteInput) error {
	productIDVo, err := productdm.NewProductID(in.ProductID)
	if err != nil {
		return err
	}

	productDomainService := productdm.NewProductDomainService(u.productRepository)

	if exist, err := productDomainService.ExistsProductByIDForUpdate(ctx, productIDVo); err != nil {
		return err
	} else if !exist {
		return apperrors.NotFound
	}

	projectIDVo, err := projectdm.NewProjectID(in.ProjectID)
	if err != nil {
		return err
	}

	projectDomainService := projectdm.NewProjectDomainService(u.projectRepository)

	if exist, err := projectDomainService.ExistsProjectByIDForUpdate(ctx, projectIDVo, productIDVo); err != nil {
		return err
	} else if !exist {
		return apperrors.NotFound
	}

	projectNoteIDVo, err := projectnotedm.NewProjectNoteID(in.ID)
	if err != nil {
		return err
	}

	projectNoteDomainService := projectnotedm.NewProjectNoteDomainService(u.projectNoteRepository)

	if exist, err := projectNoteDomainService.ExistsProjectNoteByIDForUpdate(ctx, projectNoteIDVo, projectIDVo); err != nil {
		return err
	} else if !exist {
		return apperrors.NotFound
	}

	if err = u.projectNoteRepository.DeleteProjectNote(ctx, projectNoteIDVo, productIDVo, projectIDVo); err != nil {
		return err
	}

	return nil
}
