package projectusecse

import (
	"errors"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/projectdm"
	"github.com/onituka/agile-project-management/project-management/domain/sheredvo"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecse/input"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecse/output"
)

type ProjectUsecase interface {
	CreateProject(in *input.Project) (*output.Project, error)
}

type projectUsecase struct {
	ProjectRepository projectdm.ProjectRepository
}

func NewProjectUsecase(ProjectRepository projectdm.ProjectRepository) *projectUsecase {
	return &projectUsecase{
		ProjectRepository: ProjectRepository,
	}
}

func (u *projectUsecase) CreateProject(in *input.Project) (*output.Project, error) {
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

	validProjectDm, err := u.ProjectRepository.FetchProjectByGroupIDAndKeyName(groupIDVo, keyNameVo)
	if validProjectDm != nil {
		return nil, apperrors.Conflict
	}
	if errors.Is(err, apperrors.InternalServerError) {
		return nil, err
	}

	validProjectDm, err = u.ProjectRepository.FetchProjectByGroupIDAndName(groupIDVo, nameVo)
	if validProjectDm != nil {
		return nil, apperrors.Conflict
	}
	if errors.Is(err, apperrors.InternalServerError) {
		return nil, err
	}

	projectDm := projectdm.GenProjectForCreate(
		groupIDVo,
		keyNameVo,
		nameVo,
		leaderIDVo,
		defaultAssigneeIDVo,
	)

	if err = u.ProjectRepository.CreateProject(projectDm); err != nil {
		return nil, err
	}

	projectDm, err = u.ProjectRepository.FetchProjectByID(projectDm.ID())
	if err != nil {
		return nil, err
	}

	return &output.Project{
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
