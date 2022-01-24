package productnoteinput

type CreateProductNoteInput struct {
	ProductID string
	GroupID   string
	Title     string `json:"title"`
	Content   string `json:"content"`
	UserID    string
}
