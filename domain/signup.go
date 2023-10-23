package domain

type SignupUsecase interface {
	Create(user *User) error
	GetUserByEmail(email string) (User, error)
}
