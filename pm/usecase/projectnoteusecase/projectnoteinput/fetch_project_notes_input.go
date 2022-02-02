package projectnoteinput

type FetchProjectNotesInput struct {
	ProductID string
	ProjectID string
	Page      uint32 `json:"page"`
	Limit     uint32 `json:"limit"`
}
