package projectinput

type UpdateProjectInput struct {
	ID                string
	ProductID         string
	KeyName           string `json:"keyName"`
	Name              string `json:"name"`
	LeaderID          string `json:"leaderID"`
	DefaultAssigneeID string `json:"defaultAssigneeID"`
}
