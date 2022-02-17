package productnotecommentdm

import "context"

type ProductNoteCommentRepository interface {
	CreateProductNoteComment(ctx context.Context, productNoteComment *ProductNoteComment) error
}
