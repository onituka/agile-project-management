package productnotecommentinput

type CreateProductNoteCommentInput struct {
	ProductID     string
	ProductNoteID string
	GroupID       string
	Content       string `json:"content"`
	UserID        string
}
