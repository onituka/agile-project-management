package persistence

import (
	"context"
	"database/sql"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/groupdm"
	"github.com/onituka/agile-project-management/project-management/domain/projectdm"
	"github.com/onituka/agile-project-management/project-management/infrastructure/persistence/datasource"
)

type projectRepository struct{}

func NewProjectRepository() *projectRepository {
	return &projectRepository{}
}

func (r *projectRepository) CreateProject(ctx context.Context, project *projectdm.Project) error {
	conn, err := execFromCtx(ctx)
	if err != nil {
		return err
	}

	query := `
       INSERT INTO projects
       (
         id,
         group_id,
         key_name,
         name,
         leader_id,
         default_assignee_id,
         created_at,
         updated_at
        )
       VALUES
         (?, ?, ?, ?, ?, ?, ?, ?)`

	if _, err = conn.ExecContext(
		ctx,
		query,
		project.ID().Value(),
		project.GroupID().Value(),
		project.KeyName().Value(),
		project.Name().Value(),
		project.LeaderID().Value(),
		project.DefaultAssigneeID().Value(),
		project.CreatedAt(),
		project.UpdatedAt(),
	); err != nil {
		return apperrors.InternalServerError
	}

	return nil
}

func (r *projectRepository) UpdateProject(ctx context.Context, project *projectdm.Project) error {
	conn, err := execFromCtx(ctx)
	if err != nil {
		return err
	}

	query := `
        UPDATE 
          projects
        SET 
          key_name = ?,
          name = ?,
          leader_id = ?,
          default_assignee_id = ?
        WHERE
          id = ?`

	if _, err := conn.ExecContext(
		ctx,
		query,
		project.KeyName().Value(),
		project.Name().Value(),
		project.LeaderID().Value(),
		project.DefaultAssigneeID().Value(),
		project.ID().Value(),
	); err != nil {
		return apperrors.InternalServerError
	}

	return nil
}

func (r *projectRepository) FetchProjectByIDForUpdate(ctx context.Context, id projectdm.ProjectID) (*projectdm.Project, error) {
	conn, err := execFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	query := `
         SELECT 
           id,
           group_id,
           key_name,
           name,
           leader_id,
           default_assignee_id,
           created_at,
           updated_at
         FROM
           projects
         WHERE
           id = ?
         FOR UPDATE`

	var projectDto datasource.Project

	if err := conn.QueryRowxContext(ctx, query, id.Value()).StructScan(&projectDto); err != nil {
		if apperrors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.NotFound
		}

		return nil, apperrors.InternalServerError
	}

	projectDm, err := projectdm.Reconstruct(
		projectDto.ID,
		projectDto.GroupID,
		projectDto.KeyName,
		projectDto.Name,
		projectDto.LeaderID,
		projectDto.DefaultAssigneeID,
		projectDto.CreatedAt,
		projectDto.UpdatedAt,
	)

	return projectDm, nil
}

func (r *projectRepository) FetchProjectByID(ctx context.Context, id projectdm.ProjectID) (*projectdm.Project, error) {
	conn, err := execFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	query := `
         SELECT 
           id,
           group_id,
           key_name,
           name,
           leader_id,
           default_assignee_id,
           created_at,
           updated_at
         FROM
           projects
         WHERE
           id = ?`

	var projectDto datasource.Project

	if err := conn.QueryRowxContext(ctx, query, id.Value()).StructScan(&projectDto); err != nil {
		if apperrors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.NotFound
		}

		return nil, apperrors.InternalServerError
	}

	projectDm, err := projectdm.Reconstruct(
		projectDto.ID,
		projectDto.GroupID,
		projectDto.KeyName,
		projectDto.Name,
		projectDto.LeaderID,
		projectDto.DefaultAssigneeID,
		projectDto.CreatedAt,
		projectDto.UpdatedAt,
	)

	return projectDm, nil
}

func (r *projectRepository) FetchProjectByGroupIDAndKeyName(ctx context.Context, groupID groupdm.GroupID, keyName projectdm.KeyName) (*projectdm.Project, error) {
	conn, err := execFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	query := `
         SELECT 
           id,
           group_id,
           key_name,
           name,
           leader_id,
           default_assignee_id,
           created_at,
           updated_at
         FROM
           projects
         WHERE
           group_id = ?
         AND
           key_name = ?`

	var projectDto datasource.Project

	if err = conn.QueryRowxContext(ctx, query, groupID.Value(), keyName.Value()).StructScan(&projectDto); err != nil {
		if apperrors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.NotFound
		}

		return nil, apperrors.InternalServerError
	}

	projectDm, err := projectdm.Reconstruct(
		projectDto.ID,
		projectDto.GroupID,
		projectDto.KeyName,
		projectDto.Name,
		projectDto.LeaderID,
		projectDto.DefaultAssigneeID,
		projectDto.CreatedAt,
		projectDto.UpdatedAt,
	)

	return projectDm, nil
}

func (r *projectRepository) FetchProjectByGroupIDAndName(ctx context.Context, groupID groupdm.GroupID, name projectdm.Name) (*projectdm.Project, error) {
	conn, err := execFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	query := `
         SELECT 
           id,
           group_id,
           key_name,
           name,
           leader_id,
           default_assignee_id,
           created_at,
           updated_at
         FROM
           projects
         WHERE
           group_id = ?
         AND
           name = ?`

	var projectDto datasource.Project

	if err := conn.QueryRowxContext(ctx, query, groupID.Value(), name.Value()).StructScan(&projectDto); err != nil {
		if apperrors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.NotFound
		}

		return nil, apperrors.InternalServerError
	}

	projectDm, err := projectdm.Reconstruct(
		projectDto.ID,
		projectDto.GroupID,
		projectDto.KeyName,
		projectDto.Name,
		projectDto.LeaderID,
		projectDto.DefaultAssigneeID,
		projectDto.CreatedAt,
		projectDto.UpdatedAt,
	)

	return projectDm, nil
}

func (r *projectRepository) FetchProjects(ctx context.Context) ([]*projectdm.Project, error) {
	conn, err := execFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	query := `
         SELECT 
           id,
           group_id,
           key_name,
           name,
           leader_id,
           default_assignee_id,
           created_at,
           updated_at
         FROM
           projects`

	rows, err := conn.QueryxContext(ctx, query)
	if err != nil {
		return nil, apperrors.InternalServerError
	}

	defer rows.Close()

	var projectsDto []datasource.Project
	for rows.Next() {
		var projectDto datasource.Project
		if err := rows.StructScan(&projectDto); err != nil {
			return nil, apperrors.InternalServerError
		}

		projectsDto = append(projectsDto, projectDto)

	}

	projectDms := make([]*projectdm.Project, len(projectsDto))
	for i, projectDto := range projectsDto {
		projectDms[i], err = projectdm.Reconstruct(
			projectDto.ID,
			projectDto.GroupID,
			projectDto.KeyName,
			projectDto.Name,
			projectDto.LeaderID,
			projectDto.DefaultAssigneeID,
			projectDto.CreatedAt,
			projectDto.UpdatedAt,
		)
	}

	return projectDms, nil
}
