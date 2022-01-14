package projectusecase

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/groupdm"
	"github.com/onituka/agile-project-management/project-management/domain/productdm"
	"github.com/onituka/agile-project-management/project-management/domain/projectdm"
	"github.com/onituka/agile-project-management/project-management/domain/userdm"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase/projectinput"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase/projectoutput"
)

type CreateProjectUsecase interface {
	CreateProject(ctx context.Context, in *projectinput.CreateProjectInput) (*projectoutput.CreateProjectOutput, error)
}

type createProjectUsecase struct {
	projectRepository projectdm.ProjectRepository
	productRepository productdm.ProductRepository
}

func NewCreateProjectUsecase(createProjectRepository projectdm.ProjectRepository, productRepository productdm.ProductRepository) *createProjectUsecase {
	return &createProjectUsecase{
		projectRepository: createProjectRepository,
		productRepository: productRepository,
	}
}

func (u *createProjectUsecase) CreateProject(ctx context.Context, in *projectinput.CreateProjectInput) (*projectoutput.CreateProjectOutput, error) {
	productIDVo, err := productdm.NewProductID(in.ProductID)
	if err != nil {
		return nil, err
	}

	if _, err = u.productRepository.FetchProductByIDForUpdate(ctx, productIDVo); err != nil {
		return nil, err
	}

	groupIDVo, err := groupdm.NewGroupID(in.GroupID)
	if err != nil {
		return nil, err
	}

	keyNameVo, err := projectdm.NewKeyName(in.KeyName)
	if err != nil {
		return nil, err
	}

	nameVo, err := projectdm.NewName(in.Name)
	if err != nil {
		return nil, err
	}

	leaderIDVo, err := userdm.NewUserID(in.LeaderID)
	if err != nil {
		return nil, err
	}

	defaultAssigneeIDVo, err := userdm.NewUserID(in.DefaultAssigneeID)
	if err != nil {
		return nil, err
	}

	projectDm, err := projectdm.GenProjectForCreate(
		productIDVo,
		groupIDVo,
		keyNameVo,
		nameVo,
		leaderIDVo,
		defaultAssigneeIDVo,
	)
	if err != nil {
		return nil, err
	}

	projectDomainService := projectdm.NewProjectDomainService(u.projectRepository)

	exist, err := projectDomainService.ExistsProjectForCreate(ctx, projectDm)
	if err != nil && !apperrors.Is(err, apperrors.NotFound) {
		return nil, err
	} else if exist {
		return nil, apperrors.Conflict
	}

	if err = u.projectRepository.CreateProject(ctx, projectDm); err != nil {
		return nil, err
	}

	return &projectoutput.CreateProjectOutput{
		ID:                projectDm.ID().Value(),
		ProductID:         projectDm.ProductID().Value(),
		GroupID:           projectDm.GroupID().Value(),
		KeyName:           projectDm.KeyName().Value(),
		Name:              projectDm.Name().Value(),
		LeaderID:          projectDm.LeaderID().Value(),
		DefaultAssigneeID: projectDm.DefaultAssigneeID().Value(),
		CreatedAt:         projectDm.CreatedAt(),
		UpdatedAt:         projectDm.UpdatedAt(),
	}, nil
}
