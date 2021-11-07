package output

import "time"

type CreateProject struct {
	ID                string    `json:"id"`
	GroupID           string    `json:"groupID"`
	KeyName           string    `json:"keyName"`
	Name              string    `json:"name"`
	LeaderID          string    `json:"leaderID"`
	DefaultAssigneeID string    `json:"defaultAssigneeID"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}

type UpdateProject struct {
	ID                string    `json:"id"`
	GroupID           string    `json:"groupID"`
	KeyName           string    `json:"keyName"`
	Name              string    `json:"name"`
	LeaderID          string    `json:"leaderID"`
	DefaultAssigneeID string    `json:"defaultAssigneeID"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}

type FetchProjectByID struct {
	ID                string    `json:"id"`
	GroupID           string    `json:"groupID"`
	KeyName           string    `json:"keyName"`
	Name              string    `json:"name"`
	LeaderID          string    `json:"leaderID"`
	DefaultAssigneeID string    `json:"defaultAssigneeID"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}

type FetchProjects struct {
	ID                string    `json:"id"`
	GroupID           string    `json:"groupID"`
	KeyName           string    `json:"keyName"`
	Name              string    `json:"name"`
	LeaderID          string    `json:"leaderID"`
	DefaultAssigneeID string    `json:"defaultAssigneeID"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}
