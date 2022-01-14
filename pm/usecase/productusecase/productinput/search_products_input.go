package productinput

type SearchProductsInput struct {
	GroupID     string
	ProductName string `json:"name"`
	Page        uint32 `json:"page"`
	Limit       uint32 `json:"limit"`
}
