package persistence

import (
	"database/sql"
	"errors"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/groupdm"
	"github.com/onituka/agile-project-management/project-management/domain/projectdm"
	"github.com/onituka/agile-project-management/project-management/domain/userdm"
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

func (r *ProjectRepository) FetchProjectByID(id projectdm.ID) (*projectdm.Project, error) {
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
		groupdm.GroupID(projectDto.GroupID),
		projectdm.KeyName(projectDto.KeyName),
		projectdm.Name(projectDto.Name),
		userdm.UserID(projectDto.LeaderID),
		userdm.UserID(projectDto.DefaultAssigneeID),
		projectDto.CreatedDate,
		projectDto.UpdatedDate,
	)

	return projectDm, nil
}

func (r *ProjectRepository) FetchProjectByGroupIDAndKeyName(groupID groupdm.GroupID, keyName projectdm.KeyName) (*projectdm.Project, error) {
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
		projectdm.ID(projectDto.ID),
		groupID,
		keyName,
		projectdm.Name(projectDto.Name),
		userdm.UserID(projectDto.LeaderID),
		userdm.UserID(projectDto.DefaultAssigneeID),
		projectDto.CreatedDate,
		projectDto.UpdatedDate,
	)

	return projectDm, nil
}

func (r *ProjectRepository) FetchProjectByGroupIDAndName(groupID groupdm.GroupID, name projectdm.Name) (*projectdm.Project, error) {
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
		projectdm.ID(projectDto.ID),
		groupID,
		projectdm.KeyName(projectDto.KeyName),
		name,
		userdm.UserID(projectDto.LeaderID),
		userdm.UserID(projectDto.DefaultAssigneeID),
		projectDto.CreatedDate,
		projectDto.UpdatedDate,
	)

	return projectDm, nil
}
