package productinput

type CreateProductInput struct {
	GroupID  string
	Name     string `json:"name"`
	LeaderID string `json:"leaderID"`
}
