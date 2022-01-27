package query

import (
	"context"
	"fmt"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/productdm"
	"github.com/onituka/agile-project-management/project-management/infrastructure/rdb"
	"github.com/onituka/agile-project-management/project-management/usecase/productnoteusecase/productnoteoutput"
)

type productNoteQueryServiceImpl struct{}

func NewProductNoteQueryServiceImpl() *productNoteQueryServiceImpl {
	return &productNoteQueryServiceImpl{}
}

func (r *productNoteQueryServiceImpl) FetchProductNotes(ctx context.Context, productID productdm.ProductID, limit uint32, offset uint32) ([]*productnoteoutput.ProductNoteOutput, error) {
	conn, err := rdb.ExecFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	query := `
         SELECT 
           id,
           product_id,
           group_id,
           title,
           content,
           created_by,
           updated_by,
           created_at,
           updated_at
         FROM
           product_notes
         WHERE
           product_id = ?
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
		limit,
		offset,
	)
	if err != nil {
		return nil, apperrors.InternalServerError
	}

	defer rows.Close()

	var productnotesDto []*productnoteoutput.ProductNoteOutput
	for rows.Next() {
		var productnoteDto productnoteoutput.ProductNoteOutput
		if err = rows.StructScan(&productnoteDto); err != nil {
			return nil, apperrors.InternalServerError
		}

		productnotesDto = append(productnotesDto, &productnoteDto)
	}

	return productnotesDto, nil
}

func (r *productNoteQueryServiceImpl) CountProductNotesByProductID(ctx context.Context, productID productdm.ProductID) (totalCount uint32, err error) {
	conn, err := rdb.ExecFromCtx(ctx)
	if err != nil {
		return 0, err
	}

	query := `
         SELECT 
           COUNT(*)
         FROM
          product_notes
         WHERE
          product_id = ?`

	if err = conn.QueryRowxContext(ctx, query, productID).Scan(&totalCount); err != nil {
		return 0, apperrors.InternalServerError
	}

	return totalCount, nil
}

func (r *productNoteQueryServiceImpl) SearchProductNotes(ctx context.Context, productID productdm.ProductID, title string, limit uint32, offset uint32) ([]*productnoteoutput.SearchProductNoteOutput, error) {
	conn, err := rdb.ExecFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	query := `
        SELECT
          id,
          product_id,
          group_id,
          title,
          content,
          created_by,
          updated_by,
          created_at,
          updated_at
        FROM
          product_notes 
        WHERE
          product_id = ?
        AND
          title LIKE ?
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
		fmt.Sprintf("%%%s%%", title),
		limit,
		offset,
	)
	if err != nil {
		return nil, apperrors.InternalServerError
	}
	defer rows.Close()

	var productNotesDto []*productnoteoutput.SearchProductNoteOutput
	for rows.Next() {
		var productNoteDto productnoteoutput.SearchProductNoteOutput
		if err = rows.StructScan(&productNoteDto); err != nil {
			return nil, apperrors.InternalServerError
		}

		productNotesDto = append(productNotesDto, &productNoteDto)
	}

	return productNotesDto, nil
}

func (r *productNoteQueryServiceImpl) CountProductNotesByTitle(ctx context.Context, productID productdm.ProductID, title string) (totalCount uint32, err error) {
	conn, err := rdb.ExecFromCtx(ctx)
	if err != nil {
		return 0, err
	}

	query := `
        SELECT
          COUNT(*)
        FROM
          product_notes
        WHERE
          product_id = ?
        AND
          title LIKE ?`

	if err = conn.QueryRowxContext(ctx, query, productID.Value(), fmt.Sprintf("%%%s%%", title)).Scan(&totalCount); err != nil {
		return 0, apperrors.InternalServerError
	}

	return totalCount, nil
}
