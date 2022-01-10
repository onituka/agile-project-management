package projectinput

type FetchProjectsInput struct {
	ProductID string `json:"product_id"`
	Page      int    `json:"page"`
	Limit     int    `json:"limit"`
}
