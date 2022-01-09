package productqueryservice

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/domain/groupdm"
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase/productoutput"
)

type ProductsQueryService interface {
	FetchProducts(ctx context.Context, groupID groupdm.GroupID, limit int, offset int) ([]*productoutput.ProductOutput, error)
	CountProducts(ctx context.Context, groupID groupdm.GroupID) (totalCount int, err error)
}
