package projectqueryservice

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/domain/productdm"
	"github.com/onituka/agile-project-management/project-management/usecase/projectusecase/projectoutput"
)

type ProjectQueryService interface {
	FetchProjects(ctx context.Context, productID productdm.ProductID, limit int, offset int) ([]*projectoutput.ProjectOutput, error)
	CountProjects(ctx context.Context, productID productdm.ProductID) (totalCount int, err error)
}
