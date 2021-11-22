package productusecse

type CreateProductInput struct {
	GroupID  string `json:"groupID"`
	Name     string `json:"name"`
	LeaderID string `json:"leaderID"`
}

type UpdateProductInput struct {
	ID       string
	GroupID  string `json:"groupID"`
	Name     string `json:"name"`
	LeaderID string `json:"leaderID"`
}
