package duser

type (
	UserInfo[T any] struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Info     T      `json:"info"`
	}
)
