package productqueryservice

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/domain/groupdm"
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase/productoutput"
)

type ProductQueryService interface {
	FetchProducts(ctx context.Context, groupID groupdm.GroupID, limit uint32, offset uint32) ([]*productoutput.ProductOutput, error)
	CountProducts(ctx context.Context, groupID groupdm.GroupID) (totalCount int, err error)
}
