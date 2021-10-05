package datesource

import "time"

type Project struct {
	ID                string    `db:"id"`
	GroupID           string    `db:"group_id"`
	KeyName           string    `db:"key_name"`
	Name              string    `db:"name"`
	LeaderID          string    `db:"leader_id"`
	DefaultAssigneeID string    `db:"default_assignee_id"`
	CreatedDate       time.Time `db:"created_at"`
	UpdatedDate       time.Time `db:"updated_at"`
}
