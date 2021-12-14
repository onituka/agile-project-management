package projectusecase

import "time"

type CreateProjectOutput struct {
	ID                string    `json:"id"`
	ProductID         string    `json:"productID"`
	GroupID           string    `json:"groupID"`
	KeyName           string    `json:"keyName"`
	Name              string    `json:"name"`
	LeaderID          string    `json:"leaderID"`
	DefaultAssigneeID string    `json:"defaultAssigneeID"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}

type UpdateProjectOutput struct {
	ID                string    `json:"id"`
	ProductID         string    `json:"productID"`
	GroupID           string    `json:"groupID"`
	KeyName           string    `json:"keyName"`
	Name              string    `json:"name"`
	LeaderID          string    `json:"leaderID"`
	DefaultAssigneeID string    `json:"defaultAssigneeID"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}

type FetchProjectByIDOutput struct {
	ID                string    `json:"id"`
	ProductID         string    `json:"productID"`
	GroupID           string    `json:"groupID"`
	KeyName           string    `json:"keyName"`
	Name              string    `json:"name"`
	LeaderID          string    `json:"leaderID"`
	DefaultAssigneeID string    `json:"defaultAssigneeID"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}

type FetchProjectsOutput []*Project

type Project struct {
	ID                string    `json:"id"`
	ProductID         string    `json:"productID"`
	GroupID           string    `json:"groupID"`
	KeyName           string    `json:"keyName"`
	Name              string    `json:"name"`
	LeaderID          string    `json:"leaderID"`
	DefaultAssigneeID string    `json:"defaultAssigneeID"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}
