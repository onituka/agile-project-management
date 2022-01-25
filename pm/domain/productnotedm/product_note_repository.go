package productnotedm

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/domain/productdm"
)

type ProductNoteRepository interface {
	CreateProductNote(ctx context.Context, productNote *ProductNote) error
	UpdateProductNote(ctx context.Context, productNote *ProductNote) error
	FetchProductNoteByProductIDAndTitle(ctx context.Context, productID productdm.ProductID, Title Title) (*ProductNote, error)
	FetchProductNoteByIDForUpdate(ctx context.Context, id ProductNoteID, productID productdm.ProductID) (*ProductNote, error)
	FetchProductNoteByID(ctx context.Context, id ProductNoteID, productID productdm.ProductID) (*ProductNote, error)
	DeleteProductNote(ctx context.Context, id ProductNoteID, productID productdm.ProductID) error
}
