package projectinput

type CreateProjectInput struct {
	ProductID         string
	GroupID           string
	KeyName           string `json:"keyName"`
	Name              string `json:"name"`
	LeaderID          string `json:"leaderID"`
	DefaultAssigneeID string `json:"defaultAssigneeID"`
}
