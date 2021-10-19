package output

import "time"

type CreateProject struct {
	ID                string    `json:"id"`
	GroupID           string    `json:"groupID"`
	KeyName           string    `json:"keyName"`
	Name              string    `json:"name"`
	LeaderID          string    `json:"leaderID"`
	DefaultAssigneeID string    `json:"defaultAssigneeID"`
	CreatedDate       time.Time `json:"createdAt"`
	UpdatedDate       time.Time `json:"updatedAt"`
}

type UpdateProject struct {
	ID                string    `json:"id"`
	GroupID           string    `json:"groupID"`
	KeyName           string    `json:"keyName"`
	Name              string    `json:"name"`
	LeaderID          string    `json:"leaderID"`
	DefaultAssigneeID string    `json:"defaultAssigneeID"`
	CreatedDate       time.Time `json:"createdAt"`
	UpdatedDate       time.Time `json:"updatedAt"`
}
