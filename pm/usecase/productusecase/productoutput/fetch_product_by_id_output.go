package productoutput

import "time"

type FetchProductByIDOutput struct {
	ID        string    `json:"id"`
	GroupID   string    `json:"groupID"`
	Name      string    `json:"name"`
	LeaderID  string    `json:"leaderID"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
