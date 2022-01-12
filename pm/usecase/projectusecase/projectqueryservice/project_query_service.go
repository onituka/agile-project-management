package projectqueryservice

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/domain/productdm"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase/projectoutput"
)

type ProjectQueryService interface {
	FetchProjects(ctx context.Context, productID productdm.ProductID, limit uint32, offset uint32) ([]*projectoutput.ProjectOutput, error)
	CountProjectsByProductID(ctx context.Context, productID productdm.ProductID) (totalCount int, err error)
}
