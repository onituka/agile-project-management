package productoutput

import "time"

type FetchProductsOutput struct {
	TotalCount int              `json:"totalCount"`
	Products   []*ProductOutput `json:"products"`
}

type ProductOutput struct {
	ID        string    `json:"id"        db:"id"`
	GroupID   string    `json:"groupID"   db:"group_id"`
	Name      string    `json:"name"      db:"name"`
	LeaderID  string    `json:"leaderID"  db:"leader_id"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}
