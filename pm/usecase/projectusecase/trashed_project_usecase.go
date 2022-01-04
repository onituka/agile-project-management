package projectusecase

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/projectdm"
)

type TrashedProjectUsecase interface {
	TrashedProject(ctx context.Context, in *TrashedProjectIDInput) (*TrashedProjectOutPut, error)
}

type trashedProjectUsecase struct {
	projectRepository projectdm.ProjectRepository
}

func NewTrashedProjectUsecase(TrashedProjectRepository projectdm.ProjectRepository) *trashedProjectUsecase {
	return &trashedProjectUsecase{
		projectRepository: TrashedProjectRepository,
	}
}

func (u *trashedProjectUsecase) TrashedProject(ctx context.Context, in *TrashedProjectIDInput) (*TrashedProjectOutPut, error) {
	projectIDVo, err := projectdm.NewProjectID(in.ID)
	if err != nil {
		return nil, err
	}

	projectDm, err := u.projectRepository.FetchProjectByIDForUpdate(ctx, projectIDVo)
	if err != nil {
		return nil, err
	} else if projectDm.IsTrashed() {
		return nil, apperrors.Conflict
	}

	projectDm.MoveToTrashed()

	projectDm.ChangeUpdateAt()

	if err = u.projectRepository.UpdateProject(ctx, projectDm); err != nil {
		return nil, err
	}

	return &TrashedProjectOutPut{
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
