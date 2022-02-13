package projectnoteinput

type SearchProjectNotesInput struct {
	ProductID string
	ProjectID string
	Title     string
	Page      uint32
	Limit     uint32
}
