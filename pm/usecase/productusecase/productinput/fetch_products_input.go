package productinput

type FetchProductsInput struct {
	GroupID string `json:"groupID"`
	Page    uint32 `json:"page"`
	Limit   uint32 `json:"limit"`
}
