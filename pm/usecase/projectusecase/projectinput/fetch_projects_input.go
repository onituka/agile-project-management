package projectinput

type FetchProjectsInput struct {
	ProductID string `json:"product_id"`
	Page      uint32 `json:"page"`
	Limit     uint32 `json:"limit"`
}
