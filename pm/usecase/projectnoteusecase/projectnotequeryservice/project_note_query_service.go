package projectnotequeryservice

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/domain/productdm"
	"github.com/onituka/agile-project-management/project-management/domain/projectdm"
	"github.com/onituka/agile-project-management/project-management/usecase/projectnoteusecase/projectnoteoutput"
)

type ProjectNoteQueryService interface {
	FetchProjectNotes(ctx context.Context, productID productdm.ProductID, projectID projectdm.ProjectID, limit uint32, offset uint32) ([]*projectnoteoutput.ProjectNoteOutput, error)
	CountProjectNotesByProductIDAndProjectID(ctx context.Context, productID productdm.ProductID, projectID projectdm.ProjectID) (totalCount uint32, err error)
}
