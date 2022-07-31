package request

type (
	DeleteUserRequest struct {
		ID []int `json:"user_ids"`
	}
)
