package projectusecase

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/productdm"
	"github.com/onituka/agile-project-management/project-management/domain/projectdm"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase/projectinput"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase/projectoutput"
)

type RestoreFromTrashProjectUsecase interface {
	RestoreFromTrashProject(ctx context.Context, in *projectinput.RestoreFromTrashProjectIDInput) (*projectoutput.RestoreFromTrashProjectIDOutPut, error)
}

type restoreFromTrashProjectUsecase struct {
	projectRepository projectdm.ProjectRepository
	productRepository productdm.ProductRepository
}

func NewRestoreFromTrashProjectUsecase(restoreFromTrashProjectHandler projectdm.ProjectRepository, productRepository productdm.ProductRepository) *restoreFromTrashProjectUsecase {
	return &restoreFromTrashProjectUsecase{
		projectRepository: restoreFromTrashProjectHandler,
		productRepository: productRepository,
	}
}

func (u *restoreFromTrashProjectUsecase) RestoreFromTrashProject(ctx context.Context, in *projectinput.RestoreFromTrashProjectIDInput) (*projectoutput.RestoreFromTrashProjectIDOutPut, error) {
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

	projectIDVo, err := projectdm.NewProjectID(in.ID)
	if err != nil {
		return nil, err
	}

	projectDomainService := projectdm.NewProjectDomainService(u.projectRepository)

	if exist, err := projectDomainService.ExistsProjectByIDForUpdate(ctx, projectIDVo, productIDVo); err != nil {
		return nil, err
	} else if !exist {
		return nil, apperrors.NotFound
	}

	projectDm, err := u.projectRepository.FetchProjectByIDForUpdate(ctx, projectIDVo, productIDVo)
	if err != nil {
		return nil, err
	} else if projectDm.IsNotTrashed() {
		return nil, apperrors.Conflict
	}

	projectDm.MoveFromTrashed()

	projectDm.ChangeUpdateAt()

	if err = u.projectRepository.UpdateProject(ctx, projectDm); err != nil {
		return nil, err
	}

	return &projectoutput.RestoreFromTrashProjectIDOutPut{
		ID:                projectDm.ID().Value(),
		ProductID:         projectDm.ProductID().Value(),
		GroupID:           projectDm.GroupID().Value(),
		KeyName:           projectDm.KeyName().Value(),
		Name:              projectDm.Name().Value(),
		LeaderID:          projectDm.LeaderID().Value(),
		DefaultAssigneeID: projectDm.DefaultAssigneeID().Value(),
		TrashedAt:         projectDm.TrashedAt(),
		CreatedAt:         projectDm.CreatedAt(),
		UpdatedAt:         projectDm.UpdatedAt(),
	}, nil
}
