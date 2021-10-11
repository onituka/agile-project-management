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
	projectProjectRepository projectdm.ProjectRepository
}

func NewProjectUsecase(projectProjectRepository projectdm.ProjectRepository) *projectUsecase {
	return &projectUsecase{
		projectProjectRepository: projectProjectRepository,
	}
}

func (u *projectUsecase) CreateProject(in *input.Project) (*output.Project, error) {
	idVo := sheredvo.NewID()

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

	projectDm, err := u.projectProjectRepository.FetchProjectByGroupIDAndKeyName(groupIDVo, keyNameVo)
	if projectDm != nil {
		return nil, apperrors.Conflict
	}
	if errors.Is(err, apperrors.InternalServerError) {
		return nil, err
	}

	projectDm = projectdm.NewProjectWithoutDate(
		idVo,
		groupIDVo,
		keyNameVo,
		nameVo,
		leaderIDVo,
		defaultAssigneeIDVo,
	)

	projectDm, err = u.projectProjectRepository.FetchProjectByGroupIDAndName(groupIDVo, nameVo)
	if projectDm != nil {
		return nil, apperrors.Conflict
	}
	if errors.Is(err, apperrors.InternalServerError) {
		return nil, err
	}

	projectDm = projectdm.NewProjectWithoutDate(
		idVo,
		groupIDVo,
		keyNameVo,
		nameVo,
		leaderIDVo,
		defaultAssigneeIDVo,
	)

	if err = u.projectProjectRepository.CreateProject(projectDm); err != nil {
		return nil, err
	}

	projectDm, err = u.projectProjectRepository.FetchProjectByID(idVo)
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
