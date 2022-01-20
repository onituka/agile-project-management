package productnoteinput

type FetchProductNotesInput struct {
	ProductID string
	Page      uint32 `json:"page"`
	Limit     uint32 `json:"limit"`
}
