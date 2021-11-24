package projectusecase

type CreateProjectInput struct {
	GroupID           string `json:"groupID"`
	KeyName           string `json:"keyName"`
	Name              string `json:"name"`
	LeaderID          string `json:"leaderID"`
	DefaultAssigneeID string `json:"defaultAssigneeID"`
}

type UpdateProjectInput struct {
	ID                string
	GroupID           string `json:"groupID"`
	KeyName           string `json:"keyName"`
	Name              string `json:"name"`
	LeaderID          string `json:"leaderID"`
	DefaultAssigneeID string `json:"defaultAssigneeID"`
}

type FetchProjectByIDInput struct {
	ID string
}
