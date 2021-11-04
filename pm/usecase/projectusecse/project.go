package projectusecse

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/projectdm"
	"github.com/onituka/agile-project-management/project-management/domain/sheredvo"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecse/input"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecse/output"
)

type ProjectUsecase interface {
	CreateProject(ctx context.Context, in *input.CreateProject) (*output.CreateProject, error)
	UpdateProject(ctx context.Context, in *input.UpdateProject) (*output.UpdateProject, error)
}

type projectUsecase struct {
	projectRepository projectdm.ProjectRepository
}

func NewProjectUsecase(ProjectRepository projectdm.ProjectRepository) *projectUsecase {
	return &projectUsecase{
		projectRepository: ProjectRepository,
	}
}

func (u *projectUsecase) CreateProject(ctx context.Context, in *input.CreateProject) (*output.CreateProject, error) {
	groupIDVo, err := sheredvo.NewGroupID(in.GroupID)
	if err != nil {
		return nil, err
	}

	keyNameVo, err := projectdm.NewKeyName(in.KeyName)
	if err != nil {
		return nil, err
	}

	nameVo, err := projectdm.NewName(in.Name)
	if err != nil {
		return nil, apperrors.InvalidParameter
	}

	leaderIDVo, err := sheredvo.NewUserID(in.DefaultAssigneeID)
	if err != nil {
		return nil, apperrors.InvalidParameter
	}

	defaultAssigneeIDVo, err := sheredvo.NewUserID(in.LeaderID)
	if err != nil {
		return nil, apperrors.InvalidParameter
	}

	projectDm := projectdm.GenProjectForCreate(
		groupIDVo,
		keyNameVo,
		nameVo,
		leaderIDVo,
		defaultAssigneeIDVo,
	)

	projectDomainService := projectdm.NewProjectDomainService(u.projectRepository)

	exist, err := projectDomainService.ExistsUniqueProjectForCreate(ctx, projectDm)
	if err != nil && !apperrors.Is(err, apperrors.NotFound) {
		return nil, err
	} else if exist {
		return nil, apperrors.Conflict
	}

	if err = u.projectRepository.CreateProject(ctx, projectDm); err != nil {
		return nil, err
	}

	projectDm, err = u.projectRepository.FetchProjectByID(ctx, projectDm.ID())
	if err != nil {
		return nil, err
	}

	return &output.CreateProject{
		ID:                projectDm.ID().Value(),
		GroupID:           projectDm.GroupID().Value(),
		KeyName:           projectDm.KeyName().Value(),
		Name:              projectDm.Name().Value(),
		LeaderID:          projectDm.LeaderID().Value(),
		DefaultAssigneeID: projectDm.DefaultAssigneeID().Value(),
		CreatedAt:         projectDm.CreatedAt(),
		UpdatedAt:         projectDm.UpdatedAt(),
	}, nil

}

func (u *projectUsecase) UpdateProject(ctx context.Context, in *input.UpdateProject) (*output.UpdateProject, error) {
	projectIDVo, err := sheredvo.NewProjectID(in.ID)
	if err != nil {
		return nil, err
	}

	projectDm, err := u.projectRepository.FetchProjectByIDForUpdate(ctx, projectIDVo)
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

	leaderIDVo, err := sheredvo.NewUserID(in.LeaderID)
	if err != nil {
		return nil, err
	}

	projectDm.ChangeLeaderID(leaderIDVo)

	defaultAssigneeID, err := sheredvo.NewUserID(in.DefaultAssigneeID)
	if err != nil {
		return nil, err
	}

	projectDm.ChangeDefaultAssigneeID(defaultAssigneeID)

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

	projectDm, err = u.projectRepository.FetchProjectByID(ctx, projectDm.ID())
	if err != nil {
		return nil, err
	}

	return &output.UpdateProject{
		ID:                projectDm.ID().Value(),
		GroupID:           projectDm.GroupID().Value(),
		KeyName:           projectDm.KeyName().Value(),
		Name:              projectDm.Name().Value(),
		LeaderID:          projectDm.LeaderID().Value(),
		DefaultAssigneeID: projectDm.DefaultAssigneeID().Value(),
		CreatedAt:         projectDm.CreatedAt(),
		UpdatedAt:         projectDm.UpdatedAt(),
	}, nil
}
