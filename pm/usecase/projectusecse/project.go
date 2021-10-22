package projectusecse

import (
	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/projectdm"
	"github.com/onituka/agile-project-management/project-management/domain/sheredvo"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecse/input"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecse/output"
)

type ProjectUsecase interface {
	CreateProject(in *input.CreateProject) (*output.CreateProject, error)
	UpdateProject(in *input.UpdateProject) (*output.UpdateProject, error)
}

type projectUsecase struct {
	projectRepository projectdm.ProjectRepository
}

func NewProjectUsecase(ProjectRepository projectdm.ProjectRepository) *projectUsecase {
	return &projectUsecase{
		projectRepository: ProjectRepository,
	}
}

func (u *projectUsecase) CreateProject(in *input.CreateProject) (*output.CreateProject, error) {
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

	projectDomainService := projectdm.NewProjectDomainService(u.projectRepository)

	exist, err := projectDomainService.ExistsUniqueProjectKeyName(groupIDVo, keyNameVo)
	if !apperrors.Is(err, apperrors.NotFound) {
		return nil, err
	} else if exist {
		return nil, apperrors.Conflict
	}

	exist, err = projectDomainService.ExistsUniqueProjectName(groupIDVo, nameVo)
	if !apperrors.Is(err, apperrors.NotFound) {
		return nil, err
	} else if exist {
		return nil, apperrors.Conflict
	}

	projectDm := projectdm.GenProjectForCreate(
		groupIDVo,
		keyNameVo,
		nameVo,
		leaderIDVo,
		defaultAssigneeIDVo,
	)

	if err = u.projectRepository.CreateProject(projectDm); err != nil {
		return nil, err
	}

	projectDm, err = u.projectRepository.FetchProjectByID(projectDm.ID())
	if err != nil {
		return nil, err
	}

	return &output.CreateProject{
		ID:                projectDm.ID().Value(),
		GroupID:           projectDm.Group().Value(),
		KeyName:           projectDm.KeyName().Value(),
		Name:              projectDm.Name().Value(),
		LeaderID:          projectDm.LeaderID().Value(),
		DefaultAssigneeID: projectDm.DefaultAssigneeID().Value(),
		CreatedDate:       projectDm.CreatedDate(),
		UpdatedDate:       projectDm.UpdatedDate(),
	}, nil

}

func (u *projectUsecase) UpdateProject(in *input.UpdateProject) (*output.UpdateProject, error) {
	projectIDVo, err := sheredvo.NewProjectID(in.ID)
	if err != nil {
		return nil, err
	}

	if _, err = u.projectRepository.FetchProjectByID(projectIDVo); err != nil {
		return nil, err
	}

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

	projectDm := projectdm.GenProjectForUpdate(
		projectIDVo,
		groupIDVo,
		keyNameVo,
		nameVo,
		leaderIDVo,
		defaultAssigneeIDVo,
	)

	projectDomainService := projectdm.NewProjectDomainService(u.projectRepository)

	exist, err := projectDomainService.ExistUniqueProjectForUpdate(projectDm)
	if err != nil && !apperrors.Is(err, apperrors.NotFound) {
		return nil, err
	} else if exist {
		return nil, apperrors.Conflict
	}

	if err = u.projectRepository.UpdateProject(projectDm); err != nil {
		return nil, err
	}

	projectDm, err = u.projectRepository.FetchProjectByID(projectDm.ID())
	if err != nil {
		return nil, err
	}

	return &output.UpdateProject{
		ID:                projectDm.ID().Value(),
		GroupID:           projectDm.Group().Value(),
		KeyName:           projectDm.KeyName().Value(),
		Name:              projectDm.Name().Value(),
		LeaderID:          projectDm.LeaderID().Value(),
		DefaultAssigneeID: projectDm.DefaultAssigneeID().Value(),
		CreatedDate:       projectDm.CreatedDate(),
		UpdatedDate:       projectDm.UpdatedDate(),
	}, nil
}
