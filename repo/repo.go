package repo

import "sync"

type UserRepository struct {
	mu     sync.Mutex
	users  []User
	nextID int
}

func NewUserRepository() *UserRepository {
	return &UserRepository{nextID: 1}
}

func (r *UserRepository) AddUser(name string) User {
	r.mu.Lock()
	defer r.mu.Unlock()

	user := User{ID: r.nextID, Name: name}
	r.nextID++
	r.users = append(r.users, user)
	return user
}

func (r *UserRepository) GetAllUsers() []User {
	r.mu.Lock()
	defer r.mu.Unlock()

	return append([]User(nil), r.users...) // копия слайса
}
