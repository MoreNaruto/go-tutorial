package domain

// User represents a user entity
type User struct {
	ID    int
	Name  string
	Email string
}

// UserRepository defines the interface for user data access
type UserRepository interface {
	FindByID(id int) (*User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id int) error
}
