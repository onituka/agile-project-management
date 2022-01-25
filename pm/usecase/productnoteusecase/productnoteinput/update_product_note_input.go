package productnoteinput

type UpdateProductNoteInput struct {
	ID        string
	ProductID string
	Title     string `json:"title"`
	Content   string `json:"content"`
	UserID    string
}
