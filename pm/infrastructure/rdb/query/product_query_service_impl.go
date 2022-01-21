package query

import (
	"context"
	"fmt"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/groupdm"
	"github.com/onituka/agile-project-management/project-management/infrastructure/rdb"
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase/productoutput"
)

type productQueryServiceImpl struct{}

func NewProductQueryServiceImpl() *productQueryServiceImpl {
	return &productQueryServiceImpl{}
}

func (r *productQueryServiceImpl) FetchProducts(ctx context.Context, groupID groupdm.GroupID, limit uint32, offset uint32) ([]*productoutput.ProductOutput, error) {
	conn, err := rdb.ExecFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	query := `
         SELECT 
           id,
           group_id,
           name,
           leader_id,
           created_at,
           updated_at
         FROM
           products
         WHERE
           group_id = ?
         ORDER BY
           created_at, name
         LIMIT
           ?
         OFFSET
           ?`

	rows, err := conn.QueryxContext(
		ctx,
		query,
		groupID.Value(),
		limit,
		offset,
	)
	if err != nil {
		return nil, apperrors.InternalServerError
	}

	defer rows.Close()

	var productsDto []*productoutput.ProductOutput
	for rows.Next() {
		var productDto productoutput.ProductOutput
		if err = rows.StructScan(&productDto); err != nil {
			return nil, apperrors.InternalServerError
		}

		productsDto = append(productsDto, &productDto)
	}

	return productsDto, nil
}

func (r *productQueryServiceImpl) CountProductsByGroupID(ctx context.Context, groupID groupdm.GroupID) (totalCount uint32, err error) {
	conn, err := rdb.ExecFromCtx(ctx)
	if err != nil {
		return 0, err
	}

	query := `
         SELECT 
           COUNT(*)
         FROM
           products
         WHERE
           group_id = ?`

	if err = conn.QueryRowxContext(ctx, query, groupID).Scan(&totalCount); err != nil {
		return 0, apperrors.InternalServerError
	}

	return totalCount, nil
}

func (r *productQueryServiceImpl) SearchProducts(ctx context.Context, groupID groupdm.GroupID, name string, limit uint32, offset uint32) ([]*productoutput.SearchProductOutput, error) {
	conn, err := rdb.ExecFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	query := `
        SELECT
          id,
          group_id,
          name,
          leader_id,
          created_at,
          updated_at
        FROM
          products
        WHERE
          group_id = ?
        AND
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
		groupID.Value(),
		fmt.Sprintf("%%%s%%", name),
		limit,
		offset,
	)
	if err != nil {
		return nil, apperrors.InternalServerError
	}

	defer rows.Close()

	var productsDto []*productoutput.SearchProductOutput
	for rows.Next() {
		var productDto productoutput.SearchProductOutput
		if err = rows.StructScan(&productDto); err != nil {
			return nil, apperrors.InternalServerError
		}

		productsDto = append(productsDto, &productDto)
	}

	return productsDto, nil
}

func (r *productQueryServiceImpl) CountProductsByName(ctx context.Context, groupID groupdm.GroupID, name string) (totalCount uint32, err error) {
	conn, err := rdb.ExecFromCtx(ctx)
	if err != nil {
		return 0, err
	}

	query := `
        SELECT
          COUNT(*)
        FROM
          products
        WHERE
          group_id = ?
        AND
          name LIKE ?`

	if err = conn.QueryRowxContext(ctx, query, groupID.Value(), fmt.Sprintf("%%%s%%", name)).Scan(&totalCount); err != nil {
		return 0, apperrors.InternalServerError
	}

	return totalCount, nil
}
