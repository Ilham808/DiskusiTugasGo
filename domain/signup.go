package domain

type SignupStruct struct {
	Name       string `validate:"required" json:"name" form:"name"`
	Email      string `validate:"required,email" json:"email" form:"email"`
	Password   string `validate:"required" json:"password" form:"password"`
	Gender     string `validate:"required" json:"gender" form:"gender"`
	University string `validate:"required" json:"university" form:"university"`
	Avatar     string `json:"avatar" form:"avatar"`
	IsStudent  bool   `json:"is_student" form:"is_student"`
	Status     string `json:"status" form:"status"`
}

type SignupUsecase interface {
	Store(user *User) error
	GetUserByEmail(email string) error
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
}
