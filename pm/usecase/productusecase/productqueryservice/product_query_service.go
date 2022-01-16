package productqueryservice

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/domain/groupdm"
	"github.com/onituka/agile-project-management/project-management/usecase/productusecase/productoutput"
)

type ProductQueryService interface {
	FetchProducts(ctx context.Context, groupID string, limit uint32, offset uint32) ([]*productoutput.ProductOutput, error)
	CountProductsByGroupID(ctx context.Context, groupID string) (totalCount uint32, err error)
	SearchProducts(ctx context.Context, groupID groupdm.GroupID, name string, limit uint32, offset uint32) ([]*productoutput.SearchProductOutput, error)
	CountProductsByName(ctx context.Context, groupID groupdm.GroupID, name string) (totalCount uint32, err error)
}
