package productqueryservice

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/usecase/productusecase/productoutput"
)

type ProductQueryService interface {
	FetchProducts(ctx context.Context, groupID string, limit uint32, offset uint32) ([]*productoutput.ProductOutput, error)
	CountProductsByGroupID(ctx context.Context, groupID string) (totalCount int, err error)
}
