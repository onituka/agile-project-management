package productinput

type FetchProductsInput struct {
	GroupID string
	Page    uint32 `json:"page"`
	Limit   uint32 `json:"limit"`
}
