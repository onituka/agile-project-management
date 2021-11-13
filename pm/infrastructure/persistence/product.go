package persistence

import (
	"context"
	"database/sql"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/groupdm"
	"github.com/onituka/agile-project-management/project-management/domain/productdm"
	"github.com/onituka/agile-project-management/project-management/infrastructure/persistence/datesource"
)

type productRepository struct{}

func NewProductRepository() *productRepository {
	return &productRepository{}
}

func (r *productRepository) CreateProduct(ctx context.Context, product *productdm.Product) error {
	conn, err := execFromCtx(ctx)
	if err != nil {
		return err
	}

	query := `
       INSERT INTO products
       (
         id,
         group_id,
         name,
         leader_id
        )
       VALUES
         (?, ?, ?, ?)`

	if _, err = conn.ExecContext(
		ctx,
		query,
		product.ID().Value(),
		product.GroupID().Value(),
		product.Name().Value(),
		product.LeaderID().Value(),
	); err != nil {
		return apperrors.InternalServerError
	}

	return nil
}

func (r *productRepository) FetchProductByGroupIDAndName(ctx context.Context, groupID groupdm.GroupID, Name productdm.Name) (*productdm.Product, error) {
	conn, err := execFromCtx(ctx)
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
           name = ?`

	var productDto datesource.Product

	if err = conn.QueryRowxContext(ctx, query, groupID.Value(), Name.Value()).StructScan(&productDto); err != nil {
		if apperrors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.NotFound
		}

		return nil, apperrors.InternalServerError
	}

	productDm := productdm.Reconstruct(
		productDto.ID,
		productDto.GroupID,
		productDto.Name,
		productDto.LeaderID,
		productDto.CreatedAt,
		productDto.UpdatedAt,
	)

	return productDm, nil
}
