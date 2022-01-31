package productnoteinput

type SearchProductNotesInput struct {
	ProductID string
	Title     string
	Page      uint32
	Limit     uint32
}
