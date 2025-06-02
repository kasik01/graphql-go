package auth

type (
	LoginRequest struct {
		UserName string `json:"username" binding:"required"`
		PassWord string `json:"password" binding:"required"`
	}
	LoginResponse struct {
		AccessToken string `json:"accessToken"`
	}
	TokenPayload struct {
		Id string
	}
)
