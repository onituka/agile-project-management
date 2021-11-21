package productusecse

type CreateProductInput struct {
	GroupID  string `json:"groupID"`
	Name     string `json:"name"`
	LeaderID string `json:"leaderID"`
}

type UpdateProductInput struct {
	ID       string `json:"ID"`
	GroupID  string `json:"groupID"`
	Name     string `json:"name"`
	LeaderID string `json:"leaderID"`
}
