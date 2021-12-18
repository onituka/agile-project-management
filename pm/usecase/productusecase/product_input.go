package productusecase

type CreateProductInput struct {
	GroupID  string `json:"groupID"`
	Name     string `json:"name"`
	LeaderID string `json:"leaderID"`
}

type UpdateProductInput struct {
	ID       string
	Name     string `json:"name"`
	LeaderID string `json:"leaderID"`
}

type FetchProductByIDInput struct {
	ID string
}

type TrashedProductInput struct {
	ID string
}
