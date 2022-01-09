package productinput

type CreateProductInput struct {
	GroupID  string `json:"groupID"`
	Name     string `json:"name"`
	LeaderID string `json:"leaderID"`
}
