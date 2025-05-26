package repo

type FakeUserRepository struct {
	users  []User
	nextID int
}

func NewFakeUserRepository() *FakeUserRepository {
	return &FakeUserRepository{nextID: 100}
}

func (r *FakeUserRepository) AddUser(name string) User {
	user := User{ID: r.nextID, Name: "fake_" + name}
	r.nextID++
	r.users = append(r.users, user)
	return user
}

func (r *FakeUserRepository) GetAllUsers() []User {
	return r.users
}
