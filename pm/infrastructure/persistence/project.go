package persistence

import (
	"database/sql"
	"errors"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/projectdm"
	"github.com/onituka/agile-project-management/project-management/domain/sheredvo"
	"github.com/onituka/agile-project-management/project-management/infrastructure/persistence/datesource"
	"github.com/onituka/agile-project-management/project-management/infrastructure/persistence/rdb"
)

type ProjectRepository struct {
	*rdb.MySQLHandler
}

func NewProjectRepository(mysqlHandle *rdb.MySQLHandler) *ProjectRepository {
	return &ProjectRepository{mysqlHandle}
}

func (r *ProjectRepository) CreateProject(project *projectdm.Project) error {
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
		project.Group().Value(),
		project.KeyName().Value(),
		project.Name().Value(),
		project.LeaderID().Value(),
		project.DefaultAssigneeID().Value(),
	); err != nil {
		return apperrors.InternalServerError
	}

	return nil
}

func (r *ProjectRepository) FetchProjectByID(id sheredvo.ID) (*projectdm.Project, error) {
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
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.NotFound
		}

		return nil, apperrors.InternalServerError
	}

	projectDm := projectdm.NewProject(
		id,
		sheredvo.GroupID(projectDto.GroupID),
		projectdm.KeyName(projectDto.KeyName),
		projectdm.Name(projectDto.Name),
		sheredvo.UserID(projectDto.LeaderID),
		sheredvo.UserID(projectDto.DefaultAssigneeID),
		projectDto.CreatedDate,
		projectDto.UpdatedDate,
	)

	return projectDm, nil
}

func (r *ProjectRepository) FetchProjectByGroupIDAndKeyName(groupID sheredvo.GroupID, keyName projectdm.KeyName) (*projectdm.Project, error) {
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
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.NotFound
		}

		return nil, apperrors.InternalServerError
	}

	projectDm := projectdm.NewProject(
		sheredvo.ID(projectDto.ID),
		groupID,
		keyName,
		projectdm.Name(projectDto.Name),
		sheredvo.UserID(projectDto.LeaderID),
		sheredvo.UserID(projectDto.DefaultAssigneeID),
		projectDto.CreatedDate,
		projectDto.UpdatedDate,
	)

	return projectDm, nil
}

func (r *ProjectRepository) FetchProjectByGroupIDAndName(groupID sheredvo.GroupID, name projectdm.Name) (*projectdm.Project, error) {
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
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.NotFound
		}

		return nil, apperrors.InternalServerError
	}

	projectDm := projectdm.NewProject(
		sheredvo.ID(projectDto.ID),
		groupID,
		projectdm.KeyName(projectDto.KeyName),
		name,
		sheredvo.UserID(projectDto.LeaderID),
		sheredvo.UserID(projectDto.DefaultAssigneeID),
		projectDto.CreatedDate,
		projectDto.UpdatedDate,
	)

	return projectDm, nil
}
