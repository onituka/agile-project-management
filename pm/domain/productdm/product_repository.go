package productdm

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/domain/groupdm"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *Product) error
	UpdateProduct(ctx context.Context, product *Product) error
	FetchProductByGroupIDAndName(ctx context.Context, groupID groupdm.GroupID, Name Name) (*Product, error)
	FetchProductByIDForUpdate(ctx context.Context, id ProductID) (*Product, error)
	FetchProductByID(ctx context.Context, id ProductID) (*Product, error)
	FetchProducts(ctx context.Context) ([]*Product, error)
}
