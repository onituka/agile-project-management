package projectoutput

import "time"

type FetchTrashedProjectsOutput struct {
	TotalCount uint32                       `json:"totalCount"`
	Projects   []*FetchTrashedProjectOutput `json:"projects"`
}

type FetchTrashedProjectOutput struct {
	ID                string     `json:"id"                db:"id"`
	ProductID         string     `json:"productID"         db:"product_id"`
	GroupID           string     `json:"groupID"           db:"group_id"`
	KeyName           string     `json:"keyName"           db:"key_name"`
	Name              string     `json:"name"              db:"name"`
	LeaderID          string     `json:"leaderID"          db:"leader_id"`
	DefaultAssigneeID string     `json:"defaultAssigneeID" db:"default_assignee_id"`
	TrashedAt         *time.Time `json:"trashedAt"         db:"trashed_at"`
	CreatedAt         time.Time  `json:"createdAt"         db:"created_at"`
	UpdatedAt         time.Time  `json:"updatedAt"         db:"updated_at"`
}
