package query

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/groupdm"
	"github.com/onituka/agile-project-management/project-management/infrastructure/rdb/persistence"
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase/productoutput"
)

type productsQueryServiceImpl struct{}

func NewProductsQueryServiceImpl() *productsQueryServiceImpl {
	return &productsQueryServiceImpl{}
}

func (r *productsQueryServiceImpl) FetchProducts(ctx context.Context, groupID groupdm.GroupID, limit int, offset int) ([]*productoutput.ProductOutput, error) {
	conn, err := persistence.ExecFromCtx(ctx)
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

func (r *productsQueryServiceImpl) CountProducts(ctx context.Context, groupID groupdm.GroupID) (totalCount int, err error) {
	conn, err := persistence.ExecFromCtx(ctx)
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

	if err = conn.QueryRowxContext(ctx, query, groupID.Value()).Scan(&totalCount); err != nil {
		return 0, apperrors.InternalServerError
	}

	return totalCount, nil
}
