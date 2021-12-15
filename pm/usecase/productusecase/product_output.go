package productusecase

import "time"

type CreateProductOutput struct {
	ID        string    `json:"id"`
	GroupID   string    `json:"groupID"`
	Name      string    `json:"name"`
	LeaderID  string    `json:"leaderID"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UpdateProductOutput struct {
	ID        string    `json:"id"`
	GroupID   string    `json:"groupID"`
	Name      string    `json:"name"`
	LeaderID  string    `json:"leaderID"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type FetchProductByIDOutput struct {
	ID        string    `json:"id"`
	GroupID   string    `json:"groupID"`
	Name      string    `json:"name"`
	LeaderID  string    `json:"leaderID"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type FetchProductsOutput []*Product

type Product struct {
	ID        string    `json:"id"`
	GroupID   string    `json:"groupID"`
	Name      string    `json:"name"`
	LeaderID  string    `json:"leaderID"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
