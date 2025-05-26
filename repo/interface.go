package repo

type UserRepoInterface interface {
	AddUser(name string) User
	GetAllUsers() []User
}
