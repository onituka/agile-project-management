package projectnoteinput

type CreateProjectNoteInput struct {
	ProductID string
	ProjectID string
	GroupID   string
	UserID    string
	Title     string `json:"title"`
	Content   string `json:"content"`
}
