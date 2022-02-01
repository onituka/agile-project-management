package projectnoteinput

type UpdateProjectNoteInput struct {
	ID        string
	ProductID string
	ProjectID string
	UserID    string
	Title     string `json:"title"`
	Content   string `json:"content"`
}
