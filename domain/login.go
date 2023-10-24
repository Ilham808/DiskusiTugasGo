package domain

type LoginRequest struct {
	Email    string `validate:"required,email" json:"email" form:"email"`
	Password string `validate:"required" json:"password" form:"password"`
}

type LoginUsecase interface {
	GetUserByEmail(email string) (User, error)
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
}
