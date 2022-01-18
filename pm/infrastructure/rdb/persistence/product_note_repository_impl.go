package persistence

import (
	"context"
	"database/sql"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/productdm"
	"github.com/onituka/agile-project-management/project-management/domain/productnotedm"
	"github.com/onituka/agile-project-management/project-management/infrastructure/rdb"
	"github.com/onituka/agile-project-management/project-management/infrastructure/rdb/persistence/datasource"
)

type productNoteRepository struct{}

func NewProductNoteRepository() *productNoteRepository {
	return &productNoteRepository{}
}

func (r *productNoteRepository) CreateProductNote(ctx context.Context, productNote *productnotedm.ProductNote) error {
	conn, err := rdb.ExecFromCtx(ctx)
	if err != nil {
		return err
	}

	query := `
       INSERT INTO product_notes
       (
         id,
         product_id,
         group_id,
         title,
         content,
         created_by,
         updated_by,
         created_at,
         updated_at
        )
       VALUES
         (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	if _, err = conn.ExecContext(
		ctx,
		query,
		productNote.ID().Value(),
		productNote.ProductID().Value(),
		productNote.GroupID().Value(),
		productNote.Title().Value(),
		productNote.Content().Value(),
		productNote.CreatedBy().Value(),
		productNote.UpdatedBy().Value(),
		productNote.CreatedAt(),
		productNote.UpdatedAt(),
	); err != nil {
		return apperrors.InternalServerError
	}

	return nil
}

func (r *productNoteRepository) UpdateProductNote(ctx context.Context, productNote *productnotedm.ProductNote) error {
	conn, err := rdb.ExecFromCtx(ctx)
	if err != nil {
		return err
	}

	query := `
        UPDATE
          product_notes
        SET
          title = ?,
          content = ?,
          created_by = ?,
          updated_by = ?
        WHERE
          id = ?`

	if _, err := conn.ExecContext(
		ctx,
		query,
		productNote.Title().Value(),
		productNote.Content().Value(),
		productNote.CreatedBy().Value(),
		productNote.UpdatedBy().Value(),
		productNote.ID().Value(),
	); err != nil {
		return apperrors.InternalServerError
	}

	return nil
}

func (r *productNoteRepository) FetchProductNoteByProductIDAndTitle(ctx context.Context, productID productdm.ProductID, Title productnotedm.Title) (*productnotedm.ProductNote, error) {
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
           title = ?`

	var productNoteDto datasource.ProductNote

	if err = conn.QueryRowxContext(ctx, query, productID.Value(), Title.Value()).StructScan(&productNoteDto); err != nil {
		if apperrors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.NotFound
		}

		return nil, apperrors.InternalServerError
	}

	productNoteDm, err := productnotedm.Reconstruct(
		productNoteDto.ID,
		productNoteDto.ProductID,
		productNoteDto.GroupID,
		productNoteDto.Title,
		productNoteDto.Content,
		productNoteDto.CreatedBy,
		productNoteDto.UpdatedBy,
		productNoteDto.CreatedAt,
		productNoteDto.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return productNoteDm, nil
}

func (r *productNoteRepository) FetchProductNoteByIDForUpdate(ctx context.Context, id productnotedm.ProductNoteID, productID productdm.ProductID) (*productnotedm.ProductNote, error) {
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
           id = ?
         AND
           product_id = ?
         FOR UPDATE`

	var productNoteDto datasource.ProductNote

	if err := conn.QueryRowxContext(ctx, query, id.Value(), productID.Value()).StructScan(&productNoteDto); err != nil {
		if apperrors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.NotFound
		}

		return nil, apperrors.InternalServerError
	}

	productNoteDm, err := productnotedm.Reconstruct(
		productNoteDto.ID,
		productNoteDto.ProductID,
		productNoteDto.GroupID,
		productNoteDto.Title,
		productNoteDto.Content,
		productNoteDto.CreatedBy,
		productNoteDto.UpdatedBy,
		productNoteDto.CreatedAt,
		productNoteDto.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return productNoteDm, nil
}

func (r *productNoteRepository) FetchProductNoteByID(ctx context.Context, id productnotedm.ProductNoteID, productID productdm.ProductID) (*productnotedm.ProductNote, error) {
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
           id = ?
         AND
           product_id = ?`

	var productNoteDto datasource.ProductNote

	if err := conn.QueryRowxContext(ctx, query, id.Value(), productID.Value()).StructScan(&productNoteDto); err != nil {
		if apperrors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.NotFound
		}

		return nil, apperrors.InternalServerError
	}

	productNoteDm, err := productnotedm.Reconstruct(
		productNoteDto.ID,
		productNoteDto.ProductID,
		productNoteDto.GroupID,
		productNoteDto.Title,
		productNoteDto.Content,
		productNoteDto.CreatedBy,
		productNoteDto.UpdatedBy,
		productNoteDto.CreatedAt,
		productNoteDto.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return productNoteDm, nil
}
