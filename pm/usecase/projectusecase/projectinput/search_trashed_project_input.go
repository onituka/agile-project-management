package projectinput

type SearchTrashedProjectsInput struct {
	ProductID string
	KeyWord   string
	Page      uint32
	Limit     uint32
}
