package persistence

import (
	"context"
	"database/sql"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/projectdm"
	"github.com/onituka/agile-project-management/project-management/domain/projectnotedm"
	"github.com/onituka/agile-project-management/project-management/infrastructure/rdb"
	"github.com/onituka/agile-project-management/project-management/infrastructure/rdb/persistence/datasource"
)

type projectNoteRepository struct{}

func NewProjectNoteRepository() *projectNoteRepository {
	return &projectNoteRepository{}
}
func (r *projectNoteRepository) CreateProjectNote(ctx context.Context, projectNote *projectnotedm.ProjectNote) error {
	conn, err := rdb.ExecFromCtx(ctx)
	if err != nil {
		return err
	}

	query := `
       INSERT INTO project_notes
       (
         id,
         product_id,
         project_id,
         group_id,
         title,
         content,
         created_by,
         updated_by,
         created_at,
         updated_at
        )
       VALUES
         (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	if _, err = conn.ExecContext(
		ctx,
		query,
		projectNote.ID().Value(),
		projectNote.ProductID().Value(),
		projectNote.ProjectID().Value(),
		projectNote.GroupID().Value(),
		projectNote.Title().Value(),
		projectNote.Content().Value(),
		projectNote.CreatedBy().Value(),
		projectNote.UpdatedBy().Value(),
		projectNote.CreatedAt(),
		projectNote.UpdatedAt(),
	); err != nil {
		return apperrors.InternalServerError
	}

	return nil
}

func (r *projectNoteRepository) FetchProjectNoteByProjectIDAndTitle(ctx context.Context, projectID projectdm.ProjectID, title projectnotedm.Title) (*projectnotedm.ProjectNote, error) {
	conn, err := rdb.ExecFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	query := `
         SELECT 
           id,
           product_id,
           project_id,
           group_id,
           title,
           content,
           created_by,
           updated_by,
           created_at,
           updated_at
         FROM
           project_notes
         WHERE
           project_id = ?
         AND
           title = ?`

	var projectNoteDto datasource.ProjectNote

	if err = conn.QueryRowxContext(ctx, query, projectID.Value(), title.Value()).StructScan(&projectNoteDto); err != nil {
		if apperrors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.NotFound
		}

		return nil, apperrors.InternalServerError
	}

	projectNoteDm, err := projectnotedm.Reconstruct(
		projectNoteDto.ID,
		projectNoteDto.ProductID,
		projectNoteDto.ProjectID,
		projectNoteDto.GroupID,
		projectNoteDto.Title,
		projectNoteDto.Content,
		projectNoteDto.CreatedBy,
		projectNoteDto.UpdatedBy,
		projectNoteDto.CreatedAt,
		projectNoteDto.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return projectNoteDm, nil
}
