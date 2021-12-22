package projectusecase

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/projectdm"
	"github.com/onituka/agile-project-management/project-management/domain/userdm"
	"github.com/onituka/agile-project-management/project-management/usecase/timemanager"
)

type UpdateProjectUsecase interface {
	UpdateProject(ctx context.Context, in *UpdateProjectInput) (*UpdateProjectOutput, error)
}

type updateProjectUsecase struct {
	projectRepository projectdm.ProjectRepository
	timeManager       timemanager.TimeManager
}

func NewUpdateProjectUsecase(UpdateProjectRepository projectdm.ProjectRepository, timeManager timemanager.TimeManager) *updateProjectUsecase {
	return &updateProjectUsecase{
		projectRepository: UpdateProjectRepository,
		timeManager:       timeManager,
	}
}

func (u *updateProjectUsecase) UpdateProject(ctx context.Context, in *UpdateProjectInput) (*UpdateProjectOutput, error) {
	projectIDVo, err := projectdm.NewProjectID(in.ID)
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

	projectDm.MoveToUpdate()

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

	return &UpdateProjectOutput{
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
