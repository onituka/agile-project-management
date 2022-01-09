package productinput

type FetchProductsInput struct {
	GroupID string `json:"groupID"`
	Page    int    `json:"page"`
	Limit   int    `json:"limit"`
}
