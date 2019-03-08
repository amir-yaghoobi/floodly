package user

type UserRepository interface {
	Create(i *User) error
}
