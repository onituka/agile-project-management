package projectusecase

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/domain/projectdm"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase/projectinput"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase/projectoutput"
)

type FetchProjectByIDUsecase interface {
	FetchProjectByID(ctx context.Context, in *projectinput.FetchProjectByIDInput) (*projectoutput.FetchProjectByIDOutput, error)
}

type fetchProjectByIDUsecase struct {
	projectRepository projectdm.ProjectRepository
}

func NewFetchProjectByIDUsecase(FetchProjectByIDRepository projectdm.ProjectRepository) *fetchProjectByIDUsecase {
	return &fetchProjectByIDUsecase{
		projectRepository: FetchProjectByIDRepository,
	}
}

func (u *fetchProjectByIDUsecase) FetchProjectByID(ctx context.Context, in *projectinput.FetchProjectByIDInput) (*projectoutput.FetchProjectByIDOutput, error) {
	projectIDVo, err := projectdm.NewProjectID(in.ID)
	if err != nil {
		return nil, err
	}

	projectDm, err := u.projectRepository.FetchProjectByID(ctx, projectIDVo)
	if err != nil {
		return nil, err
	}

	return &projectoutput.FetchProjectByIDOutput{
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
