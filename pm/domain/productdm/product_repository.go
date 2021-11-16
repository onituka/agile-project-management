package productdm

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/domain/groupdm"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *Product) error
	FetchProductByGroupIDAndName(ctx context.Context, groupID groupdm.GroupID, Name Name) (*Product, error)
}
