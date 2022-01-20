package productnoteinput

type UpdateProductNoteInput struct {
	ID        string
	ProductID string
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedBy string `json:"createdBy"`
	UpdatedBy string `json:"updatedBy"`
}
