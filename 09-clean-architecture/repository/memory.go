package repository

import (
	"errors"
	"github.com/tutorial/clean-architecture/domain"
)

// MemoryUserRepository implements UserRepository with in-memory storage
type MemoryUserRepository struct {
	users  map[int]*domain.User
	nextID int
}

func NewMemoryUserRepository() *MemoryUserRepository {
	return &MemoryUserRepository{
		users:  make(map[int]*domain.User),
		nextID: 1,
	}
}

func (r *MemoryUserRepository) FindByID(id int) (*domain.User, error) {
	user, ok := r.users[id]
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (r *MemoryUserRepository) Create(user *domain.User) error {
	user.ID = r.nextID
	r.users[user.ID] = user
	r.nextID++
	return nil
}

func (r *MemoryUserRepository) Update(user *domain.User) error {
	if _, ok := r.users[user.ID]; !ok {
		return errors.New("user not found")
	}
	r.users[user.ID] = user
	return nil
}

func (r *MemoryUserRepository) Delete(id int) error {
	if _, ok := r.users[id]; !ok {
		return errors.New("user not found")
	}
	delete(r.users, id)
	return nil
}
