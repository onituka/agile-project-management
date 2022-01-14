package projectusecase

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/productdm"
	"github.com/onituka/agile-project-management/project-management/domain/projectdm"
	"github.com/onituka/agile-project-management/project-management/domain/userdm"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase/projectinput"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase/projectoutput"
)

type UpdateProjectUsecase interface {
	UpdateProject(ctx context.Context, in *projectinput.UpdateProjectInput) (*projectoutput.UpdateProjectOutput, error)
}

type updateProjectUsecase struct {
	projectRepository projectdm.ProjectRepository
	productRepository productdm.ProductRepository
}

func NewUpdateProjectUsecase(UpdateProjectRepository projectdm.ProjectRepository, productRepository productdm.ProductRepository) *updateProjectUsecase {
	return &updateProjectUsecase{
		projectRepository: UpdateProjectRepository,
		productRepository: productRepository,
	}
}

func (u *updateProjectUsecase) UpdateProject(ctx context.Context, in *projectinput.UpdateProjectInput) (*projectoutput.UpdateProjectOutput, error) {
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
	}

	keyNameVo, err := projectdm.NewKeyName(in.KeyName)
	if err != nil {
		return nil, err
	}

	projectDm.ChangeKeyName(keyNameVo)

	nameVo, err := projectdm.NewName(in.Name)
	if err != nil {
		return nil, err
	}

	projectDm.ChangeName(nameVo)

	leaderIDVo, err := userdm.NewUserID(in.LeaderID)
	if err != nil {
		return nil, err
	}

	projectDm.ChangeLeaderID(leaderIDVo)

	defaultAssigneeID, err := userdm.NewUserID(in.DefaultAssigneeID)
	if err != nil {
		return nil, err
	}

	projectDm.ChangeDefaultAssigneeID(defaultAssigneeID)

	projectDm.ChangeUpdateAt()

	projectDomainService := projectdm.NewProjectDomainService(u.projectRepository)

	exist, err := projectDomainService.ExistUniqueProjectForUpdate(ctx, projectDm)
	if err != nil && !apperrors.Is(err, apperrors.NotFound) {
		return nil, err
	} else if exist {
		return nil, apperrors.Conflict
	}

	if err = u.projectRepository.UpdateProject(ctx, projectDm); err != nil {
		return nil, err
	}

	return &projectoutput.UpdateProjectOutput{
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
