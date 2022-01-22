package query

import (
	"context"
	"fmt"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/productdm"
	"github.com/onituka/agile-project-management/project-management/infrastructure/rdb"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase/projectoutput"
)

type projectQuery struct{}

func NewProjectQuery() *projectQuery {
	return &projectQuery{}
}

func (r *projectQuery) FetchProjects(ctx context.Context, productID productdm.ProductID, limit uint32, offset uint32) ([]*projectoutput.ProjectOutput, error) {
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
         AND 
           trashed_at IS NULL 
         ORDER BY
           created_at, name
         LIMIT
           ?
         OFFSET
           ?`

	rows, err := conn.QueryxContext(
		ctx,
		query,
		productID.Value(),
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

func (r *projectQuery) CountProjectsByProductID(ctx context.Context, productID productdm.ProductID) (totalCount uint32, err error) {
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

	if err = conn.QueryRowxContext(ctx, query, productID.Value()).Scan(&totalCount); err != nil {
		return 0, apperrors.InternalServerError
	}

	return totalCount, nil
}

func (r *projectQuery) SearchProjects(ctx context.Context, productID productdm.ProductID, keyword string, limit uint32, offset uint32) ([]*projectoutput.SearchProjectOutput, error) {
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
         AND 
           trashed_at IS NULL 
         AND
           key_name LIKE ?
         OR
           name LIKE ?
         ORDER BY
           created_at, name
         LIMIT
           ?
         OFFSET
           ?`

	rows, err := conn.QueryxContext(
		ctx,
		query,
		productID.Value(),
		fmt.Sprintf("%%%s%%", keyword),
		fmt.Sprintf("%%%s%%", keyword),
		limit,
		offset,
	)
	if err != nil {
		return nil, apperrors.InternalServerError
	}

	defer rows.Close()

	var projectsDto []*projectoutput.SearchProjectOutput
	for rows.Next() {
		var projectDto projectoutput.SearchProjectOutput
		if err = rows.StructScan(&projectDto); err != nil {
			return nil, apperrors.InternalServerError
		}

		projectsDto = append(projectsDto, &projectDto)
	}

	return projectsDto, nil
}

func (r *projectQuery) CountProjectsByKeyNameAndName(ctx context.Context, productID productdm.ProductID, keyword string) (totalCount uint32, err error) {
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
           trashed_at IS NULL
         AND 
           key_name LIKE ?
         OR  
           name LIKE ?`

	if err = conn.QueryRowxContext(ctx, query, productID.Value(), fmt.Sprintf("%%%s%%", keyword), fmt.Sprintf("%%%s%%", keyword)).Scan(&totalCount); err != nil {
		return 0, apperrors.InternalServerError
	}

	return totalCount, nil
}

func (r *projectQuery) FetchTrashedProjects(ctx context.Context, productID productdm.ProductID, limit uint32, offset uint32) ([]*projectoutput.FetchTrashedProjectOutput, error) {
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
         AND 
           trashed_at IS NOT NULL 
         ORDER BY
           created_at, name
         LIMIT
           ?
         OFFSET
           ?`

	rows, err := conn.QueryxContext(
		ctx,
		query,
		productID.Value(),
		limit,
		offset,
	)
	if err != nil {
		return nil, apperrors.InternalServerError
	}

	defer rows.Close()

	var projectsDto []*projectoutput.FetchTrashedProjectOutput
	for rows.Next() {
		var projectDto projectoutput.FetchTrashedProjectOutput
		if err = rows.StructScan(&projectDto); err != nil {
			return nil, apperrors.InternalServerError
		}

		projectsDto = append(projectsDto, &projectDto)
	}

	return projectsDto, nil
}

func (r *projectQuery) CountTrashedProjectsByProductID(ctx context.Context, productID productdm.ProductID) (totalCount uint32, err error) {
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
           trashed_at IS NOT NULL`

	if err = conn.QueryRowxContext(ctx, query, productID.Value()).Scan(&totalCount); err != nil {
		return 0, apperrors.InternalServerError
	}

	return totalCount, nil
}
