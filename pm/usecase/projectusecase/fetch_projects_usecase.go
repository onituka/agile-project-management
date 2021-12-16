package projectusecase

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/domain/projectdm"
)

type FetchProjectsUsecase interface {
	FetchProjects(ctx context.Context) (FetchProjectsOutput, error)
}

type fetchProjectsUsecase struct {
	projectRepository projectdm.ProjectRepository
}

func NewFetchProjectsUsecase(FetchProjectsRepository projectdm.ProjectRepository) *fetchProjectsUsecase {
	return &fetchProjectsUsecase{
		projectRepository: FetchProjectsRepository,
	}
}

func (u *fetchProjectsUsecase) FetchProjects(ctx context.Context) (FetchProjectsOutput, error) {
	projectsDm, err := u.projectRepository.FetchProjects(ctx)
	if err != nil {
		return nil, err
	}

	projectsDto := make(FetchProjectsOutput, len(projectsDm))
	for i, projectDm := range projectsDm {
		projectsDto[i] = &Project{
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
		}
	}

	return projectsDto, nil
}
