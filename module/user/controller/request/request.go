package request

type (
	CreateUserRequest struct {
		Email    string `json:"email" form:"email"`
		Name     string `json:"name" form:"name"`
		Password string `json:"password" form:"password"`
	}
)
