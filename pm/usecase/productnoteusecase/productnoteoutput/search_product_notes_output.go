package productnoteoutput

import "time"

type SearchProductNotesOutput struct {
	TotalCount   uint32                     `json:"totalCount"`
	ProductNotes []*SearchProductNoteOutput `json:"productNotes"`
}

type SearchProductNoteOutput struct {
	ID        string    `json:"id"        db:"id"`
	ProductID string    `json:"productID" db:"product_id"`
	GroupID   string    `json:"groupID"   db:"group_id"`
	Title     string    `json:"title"     db:"title"`
	Content   string    `json:"content"   db:"content"`
	CreatedBy string    `json:"createdBy" db:"created_by"`
	UpdatedBy string    `json:"updatedBy" db:"updated_by"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}
