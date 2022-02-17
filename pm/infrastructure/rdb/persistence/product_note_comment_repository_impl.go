package persistence

import (
	"context"

	"github.com/onituka/agile-project-management/project-management/apperrors"
	"github.com/onituka/agile-project-management/project-management/domain/productnotecommentdm"
	"github.com/onituka/agile-project-management/project-management/infrastructure/rdb"
)

type productNoteCommentRepository struct{}

func NewProductNoteCommentRepository() *productNoteCommentRepository {
	return &productNoteCommentRepository{}
}

func (r *productNoteCommentRepository) CreateProductNoteComment(ctx context.Context, productNoteComment *productnotecommentdm.ProductNoteComment) error {
	conn, err := rdb.ExecFromCtx(ctx)
	if err != nil {
		return err
	}

	query := `
	  INSERT INTO product_note_comments
	  (
	    id,
	    product_id,
	    product_note_id,
	    group_id,
	    content,
	    created_by,
	    created_at,
	    updated_at
	   )
	  VALUES
	    (?, ?, ?, ?, ?, ?, ?, ?)`

	if _, err = conn.ExecContext(
		ctx,
		query,
		productNoteComment.ID().Value(),
		productNoteComment.ProductID().Value(),
		productNoteComment.ProductNoteID().Value(),
		productNoteComment.GroupID().Value(),
		productNoteComment.Content().Value(),
		productNoteComment.CreatedBy().Value(),
		productNoteComment.CreatedAt(),
		productNoteComment.UpdatedAt(),
	); err != nil {
		return apperrors.InternalServerError
	}

	query = `
	  INSERT INTO product_note_comment_paths
	  (
	    comment_ancestor_id,
	    comment_descendant_id
	   )
	  VALUES
	    (?, ?)`

	if _, err = conn.ExecContext(
		ctx,
		query,
		productNoteComment.CommentAncestorID().Value(),
		productNoteComment.CommentDescendantID().Value(),
	); err != nil {
		return apperrors.InternalServerError
	}

	return nil
}
