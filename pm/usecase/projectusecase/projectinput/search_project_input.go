package projectinput

type SearchProjectsInput struct {
	ProductID string
	KeyWord   string
	Page      uint32
	Limit     uint32
}
