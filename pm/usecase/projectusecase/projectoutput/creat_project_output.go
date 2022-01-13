package projectoutput

import "time"

type CreateProjectOutput struct {
	ID                string     `json:"id"`
	ProductID         string     `json:"productID"`
	GroupID           string     `json:"groupID"`
	KeyName           string     `json:"keyName"`
	Name              string     `json:"name"`
	LeaderID          string     `json:"leaderID"`
	DefaultAssigneeID string     `json:"defaultAssigneeID"`
	TrashedAt         *time.Time `json:"trashedAt"`
	CreatedAt         time.Time  `json:"createdAt"`
	UpdatedAt         time.Time  `json:"updatedAt"`
}
