package projectusecase

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/groupdm"
	"github.com/onituka/agile-project-management/project-management/domain/projectdm"
	"github.com/onituka/agile-project-management/project-management/domain/userdm"
	"github.com/onituka/agile-project-management/project-management/usecase/timemanager"
)

type CreateProjectUsecase interface {
	CreateProject(ctx context.Context, in *CreateProjectInput) (*CreateProjectOutput, error)
}

type createProjectUsecase struct {
	projectRepository projectdm.ProjectRepository
	timeManager       timemanager.TimeManager
}

func NewCreateProjectUsecase(CreateProjectRepository projectdm.ProjectRepository, timeManager timemanager.TimeManager) *createProjectUsecase {
	return &createProjectUsecase{
		projectRepository: CreateProjectRepository,
		timeManager:       timeManager,
	}
}

func (u *createProjectUsecase) CreateProject(ctx context.Context, in *CreateProjectInput) (*CreateProjectOutput, error) {
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

	now := u.timeManager.Now()

	projectDm, err := projectdm.GenProjectForCreate(
		groupIDVo,
		keyNameVo,
		nameVo,
		leaderIDVo,
		defaultAssigneeIDVo,
		now,
		now,
	)

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

	return &CreateProjectOutput{
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
