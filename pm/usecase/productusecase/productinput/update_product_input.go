package productinput

type UpdateProductInput struct {
	ID       string
	Name     string `json:"name"`
	LeaderID string `json:"leaderID"`
}
