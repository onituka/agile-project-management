package datasource

import "time"

type Product struct {
	ID        string     `db:"id"`
	GroupID   string     `db:"group_id"`
	Name      string     `db:"name"`
	LeaderID  string     `db:"leader_id"`
	TrashedAt *time.Time `db:"trashed_at"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
}
