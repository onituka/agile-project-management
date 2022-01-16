package projectusecase

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/productdm"
	"github.com/onituka/agile-project-management/project-management/domain/projectdm"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase/projectinput"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase/projectoutput"
)

type TrashedProjectUsecase interface {
	TrashedProject(ctx context.Context, in *projectinput.TrashedProjectIDInput) (*projectoutput.TrashedProjectOutPut, error)
}

type trashedProjectUsecase struct {
	projectRepository projectdm.ProjectRepository
	productRepository productdm.ProductRepository
}

func NewTrashedProjectUsecase(TrashedProjectRepository projectdm.ProjectRepository, productRepository productdm.ProductRepository) *trashedProjectUsecase {
	return &trashedProjectUsecase{
		projectRepository: TrashedProjectRepository,
		productRepository: productRepository,
	}
}

func (u *trashedProjectUsecase) TrashedProject(ctx context.Context, in *projectinput.TrashedProjectIDInput) (*projectoutput.TrashedProjectOutPut, error) {
	productIDVo, err := productdm.NewProductID(in.ProductID)
	if err != nil {
		return nil, err
	}

	if _, err = u.productRepository.FetchProductByIDForUpdate(ctx, productIDVo); err != nil {
		return nil, err
	}

	projectIDVo, err := projectdm.NewProjectID(in.ID)
	if err != nil {
		return nil, err
	}

	projectDm, err := u.projectRepository.FetchProjectByIDForUpdate(ctx, projectIDVo, productIDVo)
	if err != nil {
		return nil, err
	} else if projectDm.IsTrashed() {
		return nil, apperrors.Conflict
	}

	projectDm.MoveToTrashed()

	projectDm.ChangeUpdateAt()

	if err = u.projectRepository.UpdateProject(ctx, projectDm); err != nil {
		return nil, err
	}

	return &projectoutput.TrashedProjectOutPut{
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
