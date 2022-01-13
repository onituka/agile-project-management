package query

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/infrastructure/rdb"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase/projectoutput"
)

type projectQuery struct{}

func NewProjectQuery() *projectQuery {
	return &projectQuery{}
}

func (r *projectQuery) FetchProjects(ctx context.Context, productID string, limit uint32, offset uint32) ([]*projectoutput.ProjectOutput, error) {
	conn, err := rdb.ExecFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	query := `
         SELECT 
           id,
           product_id,
           group_id,
           key_name,
           name,
           leader_id,
           default_assignee_id,
           trashed_at,
           created_at,
           updated_at
         FROM
           projects
         WHERE
           product_id = ?
         ORDER BY
           created_at, name
         LIMIT
           ?
         OFFSET
           ?`

	rows, err := conn.QueryxContext(
		ctx,
		query,
		productID,
		limit,
		offset,
	)
	if err != nil {
		return nil, apperrors.InternalServerError
	}

	defer rows.Close()

	var projectsDto []*projectoutput.ProjectOutput
	for rows.Next() {
		var projectDto projectoutput.ProjectOutput
		if err = rows.StructScan(&projectDto); err != nil {
			return nil, apperrors.InternalServerError
		}

		projectsDto = append(projectsDto, &projectDto)
	}

	return projectsDto, nil
}

func (r *projectQuery) CountProjectsByProductID(ctx context.Context, productID string) (totalCount int, err error) {
	conn, err := rdb.ExecFromCtx(ctx)
	if err != nil {
		return 0, err
	}

	query := `
         SELECT 
           COUNT(*)
         FROM
           projects
         WHERE
           product_id = ?
         AND 
           trashed_at IS NULL`

	if err = conn.QueryRowxContext(ctx, query, productID).Scan(&totalCount); err != nil {
		return 0, apperrors.InternalServerError
	}

	return totalCount, nil
}
