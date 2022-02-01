package datasource

import "time"

type ProductNoteComment struct {
	ID            string    `db:"id"`
	ProductID     string    `db:"product_id"`
	ProductNoteID string    `db:"productNoteID"`
	GroupID       string    `db:"group_id"`
	Content       string    `db:"content"`
	UserID        string    `db:"userID"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}
