package projectusecse

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/groupdm"
	"github.com/onituka/agile-project-management/project-management/domain/projectdm"
	"github.com/onituka/agile-project-management/project-management/domain/userdm"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecse/input"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecse/output"
	"github.com/onituka/agile-project-management/project-management/usecase/timemanager"
)

type ProjectUsecase interface {
	CreateProject(ctx context.Context, in *input.CreateProject) (*output.CreateProject, error)
	UpdateProject(ctx context.Context, in *input.UpdateProject) (*output.UpdateProject, error)
	FetchProjectByID(ctx context.Context, in *input.FetchProjectByID) (*output.FetchProjectByID, error)
	FetchProjects(ctx context.Context) ([]*output.FetchProjects, error)
}

type projectUsecase struct {
	projectRepository projectdm.ProjectRepository
	timeManager       timemanager.TimeManager
}

func NewProjectUsecase(ProjectRepository projectdm.ProjectRepository, timeManager timemanager.TimeManager) *projectUsecase {
	return &projectUsecase{
		projectRepository: ProjectRepository,
		timeManager:       timeManager,
	}
}

func (u *projectUsecase) CreateProject(ctx context.Context, in *input.CreateProject) (*output.CreateProject, error) {
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
		return nil, apperrors.InvalidParameter
	}

	leaderIDVo, err := userdm.NewUserID(in.DefaultAssigneeID)
	if err != nil {
		return nil, apperrors.InvalidParameter
	}

	defaultAssigneeIDVo, err := userdm.NewUserID(in.LeaderID)
	if err != nil {
		return nil, apperrors.InvalidParameter
	}

	now := u.timeManager.Now()

	projectDm := projectdm.GenProjectForCreate(
		groupIDVo,
		keyNameVo,
		nameVo,
		leaderIDVo,
		defaultAssigneeIDVo,
		now,
		now,
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

	projectDm.ChangeUpdatedAt(u.timeManager.Now())

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

func (u *projectUsecase) FetchProjectByID(ctx context.Context, in *input.FetchProjectByID) (*output.FetchProjectByID, error) {
	projectIDVo, err := projectdm.NewProjectID(in.ID)
	if err != nil {
		return nil, err
	}

	projectDm, err := u.projectRepository.FetchProjectByID(ctx, projectIDVo)
	if err != nil {
		return nil, err
	}

	return &output.FetchProjectByID{
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

func (u *projectUsecase) FetchProjects(ctx context.Context) ([]*output.FetchProjects, error) {

	projectsDm, err := u.projectRepository.FetchProjects(ctx)
	if err != nil {
		return nil, err
	}
	projectsDto := make([]*output.FetchProjects, len(projectsDm))
	for i, projectDm := range projectsDm {
		projectsDto[i] = &output.FetchProjects{
			ID:                projectDm.ID().Value(),
			GroupID:           projectDm.GroupID().Value(),
			KeyName:           projectDm.KeyName().Value(),
			Name:              projectDm.Name().Value(),
			LeaderID:          projectDm.LeaderID().Value(),
			DefaultAssigneeID: projectDm.DefaultAssigneeID().Value(),
			CreatedAt:         projectDm.CreatedAt(),
			UpdatedAt:         projectDm.UpdatedAt(),
		}
	}

	return projectsDto, nil
}
