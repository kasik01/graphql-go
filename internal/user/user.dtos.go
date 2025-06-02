package user

type (
	RegisterUserRequest struct {
		UserName string `json:"username" binding:"required"`
		PassWord string `json:"password" binding:"required"`
	}
	RegisterUserResponse struct {
		UserName string `json:"username"`
	}
)
