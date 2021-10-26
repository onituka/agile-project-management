package persistence

import (
	"database/sql"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/projectdm"
	"github.com/onituka/agile-project-management/project-management/domain/sheredvo"
	"github.com/onituka/agile-project-management/project-management/infrastructure/persistence/datesource"
	"github.com/onituka/agile-project-management/project-management/infrastructure/persistence/rdb"
)

type projectRepository struct {
	*rdb.MySQLHandler
}

func NewProjectRepository(mysqlHandle *rdb.MySQLHandler) *projectRepository {
	return &projectRepository{mysqlHandle}
}

func (r *projectRepository) CreateProject(project *projectdm.Project) error {
	query := `
       INSERT INTO projects
       (
         id,
         group_id,
         key_name,
         name,
         leader_id,
         default_assignee_id
        )
       VALUES
         (?, ?, ?, ?, ?, ?)`

	if _, err := r.Conn.Exec(
		query,
		project.ID().Value(),
		project.GroupID().Value(),
		project.KeyName().Value(),
		project.Name().Value(),
		project.LeaderID().Value(),
		project.DefaultAssigneeID().Value(),
	); err != nil {
		return apperrors.InternalServerError
	}

	return nil
}

func (r *projectRepository) UpdateProject(project *projectdm.Project) error {
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

	if _, err := r.Conn.Exec(
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

func (r *projectRepository) FetchProjectByID(id sheredvo.ProjectID) (*projectdm.Project, error) {
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

	var projectDto datesource.Project
	if err := r.Conn.QueryRowx(query, id.Value()).StructScan(&projectDto); err != nil {
		if apperrors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.NotFound
		}

		return nil, apperrors.InternalServerError
	}

	projectDm := projectdm.Reconstruct(
		projectDto.ID,
		projectDto.GroupID,
		projectDto.KeyName,
		projectDto.Name,
		projectDto.LeaderID,
		projectDto.DefaultAssigneeID,
		projectDto.CreatedDate,
		projectDto.UpdatedDate,
	)

	return projectDm, nil
}

func (r *projectRepository) FetchProjectByGroupIDAndKeyName(groupID sheredvo.GroupID, keyName projectdm.KeyName) (*projectdm.Project, error) {
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

	var projectDto datesource.Project
	if err := r.Conn.QueryRowx(query, groupID.Value(), keyName.Value()).StructScan(&projectDto); err != nil {
		if apperrors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.NotFound
		}

		return nil, apperrors.InternalServerError
	}

	projectDm := projectdm.Reconstruct(
		projectDto.ID,
		projectDto.GroupID,
		projectDto.KeyName,
		projectDto.Name,
		projectDto.LeaderID,
		projectDto.DefaultAssigneeID,
		projectDto.CreatedDate,
		projectDto.UpdatedDate,
	)

	return projectDm, nil
}

func (r *projectRepository) FetchProjectByGroupIDAndName(groupID sheredvo.GroupID, name projectdm.Name) (*projectdm.Project, error) {
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

	var projectDto datesource.Project
	if err := r.Conn.QueryRowx(query, groupID.Value(), name.Value()).StructScan(&projectDto); err != nil {
		if apperrors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.NotFound
		}

		return nil, apperrors.InternalServerError
	}

	projectDm := projectdm.Reconstruct(
		projectDto.ID,
		projectDto.GroupID,
		projectDto.KeyName,
		projectDto.Name,
		projectDto.LeaderID,
		projectDto.DefaultAssigneeID,
		projectDto.CreatedDate,
		projectDto.UpdatedDate,
	)

	return projectDm, nil
}
