package productnoteoutput

import "time"

type FetchProductNoteByIDOutput struct {
	ID        string    `json:"id"`
	ProductID string    `json:"productID"`
	GroupID   string    `json:"groupID"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedBy string    `json:"createdBy"`
	UpdatedBy string    `json:"updatedBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
