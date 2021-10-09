package usecase

import (
	"errors"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/groupdm"
	"github.com/onituka/agile-project-management/project-management/domain/projectdm"
	"github.com/onituka/agile-project-management/project-management/domain/userdm"
	"github.com/onituka/agile-project-management/project-management/usecase/input"
	"github.com/onituka/agile-project-management/project-management/usecase/output"
)

type ProjectUsecase interface {
	CreateProject(in *input.Project) (*output.Project, error)
}

type projectUsecase struct {
	projectRepository projectdm.Repository
}

func NewProjectUsecase(projectRepository projectdm.Repository) *projectUsecase {
	return &projectUsecase{
		projectRepository: projectRepository,
	}
}

func (u *projectUsecase) CreateProject(in *input.Project) (*output.Project, error) {
	idVo := projectdm.NewID()

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

	projectDm, err := u.projectRepository.FetchProjectByGroupIDAndKeyName(groupIDVo, keyNameVo)
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

	projectDm, err = u.projectRepository.FetchProjectByGroupIDAndName(groupIDVo, nameVo)
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

	if err = u.projectRepository.CreateProject(projectDm); err != nil {
		return nil, err
	}

	projectDm, err = u.projectRepository.FetchProjectByID(idVo)
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
