package datasource

import "time"

type ProjectNote struct {
	ID        string    `db:"id"`
	ProductID string    `db:"product_id"`
	ProjectID string    `db:"project_id"`
	GroupID   string    `db:"group_id"`
	Title     string    `db:"title"`
	Content   string    `db:"content"`
	CreatedBy string    `db:"created_by"`
	UpdatedBy string    `db:"updated_by"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
