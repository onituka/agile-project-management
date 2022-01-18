package productnotedm

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/domain/productdm"
)

type ProductNoteRepository interface {
	CreateProductNote(ctx context.Context, productNote *ProductNote) error
	FetchProductNoteByProductIDAndTitle(ctx context.Context, productID productdm.ProductID, Title Title) (*ProductNote, error)
}
