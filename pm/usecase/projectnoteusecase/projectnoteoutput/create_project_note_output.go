package projectnoteoutput

import "time"

type CreateProjectNoteOutput struct {
	ID        string    `json:"id"`
	ProductID string    `json:"productID"`
	ProjectID string    `json:"projectID"`
	GroupID   string    `json:"groupID"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedBy string    `json:"createdBy"`
	UpdatedBy string    `json:"updatedBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
