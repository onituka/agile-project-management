package persistence

import (
	"context"
	"database/sql"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/groupdm"
	"github.com/onituka/agile-project-management/project-management/domain/productdm"
	"github.com/onituka/agile-project-management/project-management/infrastructure/rdb"
	"github.com/onituka/agile-project-management/project-management/infrastructure/rdb/persistence/datasource"
)

type productRepository struct{}

func NewProductRepository() *productRepository {
	return &productRepository{}
}

func (r *productRepository) CreateProduct(ctx context.Context, product *productdm.Product) error {
	conn, err := rdb.ExecFromCtx(ctx)
	if err != nil {
		return err
	}

	query := `
       INSERT INTO products
       (
         id,
         group_id,
         name,
         leader_id,
         created_at,
         updated_at
        )
       VALUES
         (?, ?, ?, ?, ?, ?)`

	if _, err = conn.ExecContext(
		ctx,
		query,
		product.ID().Value(),
		product.GroupID().Value(),
		product.Name().Value(),
		product.LeaderID().Value(),
		product.CreatedAt(),
		product.UpdatedAt(),
	); err != nil {
		return apperrors.InternalServerError
	}

	return nil
}

func (r *productRepository) UpdateProduct(ctx context.Context, product *productdm.Product) error {
	conn, err := rdb.ExecFromCtx(ctx)
	if err != nil {
		return err
	}

	query := `
        UPDATE
          products
        SET
          name = ?,
          leader_id = ?,
          updated_at = ?
        WHERE
          id = ?`

	if _, err := conn.ExecContext(
		ctx,
		query,
		product.Name().Value(),
		product.LeaderID().Value(),
		product.UpdatedAt(),
		product.ID().Value(),
	); err != nil {
		return apperrors.InternalServerError
	}

	return nil
}

func (r *productRepository) FetchProductByIDForUpdate(ctx context.Context, id productdm.ProductID) (*productdm.Product, error) {
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
           id = ?
         FOR UPDATE`

	var productDto datasource.Product

	if err := conn.QueryRowxContext(ctx, query, id.Value()).StructScan(&productDto); err != nil {
		if apperrors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.NotFound
		}

		return nil, apperrors.InternalServerError
	}

	productDm, err := productdm.Reconstruct(
		productDto.ID,
		productDto.GroupID,
		productDto.Name,
		productDto.LeaderID,
		productDto.CreatedAt,
		productDto.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return productDm, nil
}

func (r *productRepository) FetchProductByID(ctx context.Context, id productdm.ProductID) (*productdm.Product, error) {
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
           id = ?`

	var productDto datasource.Product

	if err := conn.QueryRowxContext(ctx, query, id.Value()).StructScan(&productDto); err != nil {
		if apperrors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.NotFound
		}

		return nil, apperrors.InternalServerError
	}

	productDm, err := productdm.Reconstruct(
		productDto.ID,
		productDto.GroupID,
		productDto.Name,
		productDto.LeaderID,
		productDto.CreatedAt,
		productDto.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return productDm, nil
}

func (r *productRepository) FetchProductByGroupIDAndName(ctx context.Context, groupID groupdm.GroupID, Name productdm.Name) (*productdm.Product, error) {
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
           name = ?`

	var productDto datasource.Product

	if err = conn.QueryRowxContext(ctx, query, groupID.Value(), Name.Value()).StructScan(&productDto); err != nil {
		if apperrors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.NotFound
		}

		return nil, apperrors.InternalServerError
	}

	productDm, err := productdm.Reconstruct(
		productDto.ID,
		productDto.GroupID,
		productDto.Name,
		productDto.LeaderID,
		productDto.CreatedAt,
		productDto.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return productDm, nil
}

func (r *productRepository) FetchProducts(ctx context.Context) ([]*productdm.Product, error) {
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
           products`

	rows, err := conn.QueryxContext(ctx, query)
	if err != nil {
		return nil, apperrors.InternalServerError
	}

	defer rows.Close()

	var productDms []*productdm.Product
	for rows.Next() {
		var productDto datasource.Product
		if err = rows.StructScan(&productDto); err != nil {
			return nil, apperrors.InternalServerError
		}

		productDm, err := productdm.Reconstruct(
			productDto.ID,
			productDto.GroupID,
			productDto.Name,
			productDto.LeaderID,
			productDto.CreatedAt,
			productDto.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		productDms = append(productDms, productDm)
	}

	return productDms, nil
}
