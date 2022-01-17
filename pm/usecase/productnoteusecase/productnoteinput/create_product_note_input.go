package productnoteinput

type CreateProductNoteInput struct {
	ProductID string
	GroupID   string
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedBy string `json:"createdBy"`
	UpdatedBy string `json:"updatedBy"`
}
