package projectinput

type CreateProjectInput struct {
	ProductID         string `json:"productID"`
	GroupID           string `json:"groupID"`
	KeyName           string `json:"keyName"`
	Name              string `json:"name"`
	LeaderID          string `json:"leaderID"`
	DefaultAssigneeID string `json:"defaultAssigneeID"`
}
