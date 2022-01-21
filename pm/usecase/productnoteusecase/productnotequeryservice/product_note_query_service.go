package productnotequeryservice

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/domain/productdm"
	"github.com/onituka/agile-project-management/project-management/usecase/productnoteusecase/productnoteoutput"
)

type ProductNoteQueryService interface {
	FetchProductNotes(ctx context.Context, productID productdm.ProductID, limit uint32, offset uint32) ([]*productnoteoutput.ProductNoteOutput, error)
	CountProductNotesByProductID(ctx context.Context, productID productdm.ProductID) (totalCount uint32, err error)
}
