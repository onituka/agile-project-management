package productnotecommentoutput

import "time"

type CreateProductNoteCommentOutput struct {
	ID            string    `json:"id"`
	ProductID     string    `json:"productID"`
	ProductNoteID string    `json:"productNoteID"`
	GroupID       string    `json:"groupID"`
	Content       string    `json:"content"`
	CreatedBy     string    `json:"createdBy"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}
