package query

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/productdm"
	"github.com/onituka/agile-project-management/project-management/domain/projectdm"
	"github.com/onituka/agile-project-management/project-management/infrastructure/rdb"
	"github.com/onituka/agile-project-management/project-management/usecase/projectnoteusecase/projectnoteoutput"
)

type projectNoteQueryServiceImpl struct{}

func NewProjectNoteQueryServiceImpl() *projectNoteQueryServiceImpl {
	return &projectNoteQueryServiceImpl{}
}

func (r *projectNoteQueryServiceImpl) FetchProjectNotes(ctx context.Context, productID productdm.ProductID, projectID projectdm.ProjectID, limit uint32, offset uint32) ([]*projectnoteoutput.ProjectNoteOutput, error) {
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
           product_id = ?
         AND 
           project_id = ?
         ORDER BY
           created_at, title
         LIMIT
           ?
         OFFSET
           ?`

	rows, err := conn.QueryxContext(
		ctx,
		query,
		productID.Value(),
		projectID.Value(),
		limit,
		offset,
	)
	if err != nil {
		return nil, apperrors.InternalServerError
	}

	defer rows.Close()

	var projectnotesDto []*projectnoteoutput.ProjectNoteOutput
	for rows.Next() {
		var projectnoteDto projectnoteoutput.ProjectNoteOutput
		if err = rows.StructScan(&projectnoteDto); err != nil {
			return nil, apperrors.InternalServerError
		}

		projectnotesDto = append(projectnotesDto, &projectnoteDto)
	}

	return projectnotesDto, nil
}

func (r *projectNoteQueryServiceImpl) CountProjectNotesByProductIDAndProjectID(ctx context.Context, productID productdm.ProductID, projectID projectdm.ProjectID) (totalCount uint32, err error) {
	conn, err := rdb.ExecFromCtx(ctx)
	if err != nil {
		return 0, err
	}

	query := `
         SELECT 
           COUNT(*)
         FROM
          project_notes
         WHERE
          product_id = ?
         AND 
          project_id = ?`

	if err = conn.QueryRowxContext(ctx, query, productID, projectID).Scan(&totalCount); err != nil {
		return 0, apperrors.InternalServerError
	}

	return totalCount, nil
}
