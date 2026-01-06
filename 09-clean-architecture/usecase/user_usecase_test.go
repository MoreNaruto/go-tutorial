package usecase

import (
	"github.com/tutorial/clean-architecture/domain"
	"github.com/tutorial/clean-architecture/repository"
	"testing"
)

func TestUserUseCase_CreateUser(t *testing.T) {
	repo := repository.NewMemoryUserRepository()
	uc := NewUserUseCase(repo)

	user, err := uc.CreateUser("Bob", "bob@example.com")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if user.Name != "Bob" {
		t.Errorf("Expected name Bob, got %s", user.Name)
	}

	if user.ID == 0 {
		t.Error("Expected user to have an ID")
	}
}

func TestUserUseCase_CreateUser_Validation(t *testing.T) {
	repo := repository.NewMemoryUserRepository()
	uc := NewUserUseCase(repo)

	tests := []struct {
		name      string
		userName  string
		email     string
		wantError bool
	}{
		{"valid user", "Alice", "alice@example.com", false},
		{"empty name", "", "test@example.com", true},
		{"invalid email", "Bob", "invalid-email", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := uc.CreateUser(tt.userName, tt.email)
			if tt.wantError && err == nil {
				t.Error("Expected error, got nil")
			}
			if !tt.wantError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

func TestUserUseCase_GetUser(t *testing.T) {
	repo := repository.NewMemoryUserRepository()
	uc := NewUserUseCase(repo)

	created, _ := uc.CreateUser("Charlie", "charlie@example.com")

	retrieved, err := uc.GetUser(created.ID)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if retrieved.Name != "Charlie" {
		t.Errorf("Expected name Charlie, got %s", retrieved.Name)
	}
}

func TestUserUseCase_UpdateUser(t *testing.T) {
	repo := repository.NewMemoryUserRepository()
	uc := NewUserUseCase(repo)

	user, _ := uc.CreateUser("David", "david@example.com")

	err := uc.UpdateUser(user.ID, "David Smith", "")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	updated, _ := uc.GetUser(user.ID)
	if updated.Name != "David Smith" {
		t.Errorf("Expected name David Smith, got %s", updated.Name)
	}
}

func TestUserUseCase_DeleteUser(t *testing.T) {
	repo := repository.NewMemoryUserRepository()
	uc := NewUserUseCase(repo)

	user, _ := uc.CreateUser("Eve", "eve@example.com")

	err := uc.DeleteUser(user.ID)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	_, err = uc.GetUser(user.ID)
	if err == nil {
		t.Error("Expected user to be deleted")
	}
}

// Test with mock repository
type MockRepository struct {
	users map[int]*domain.User
}

func (m *MockRepository) FindByID(id int) (*domain.User, error) {
	return m.users[id], nil
}

func (m *MockRepository) Create(user *domain.User) error {
	user.ID = 999
	m.users[user.ID] = user
	return nil
}

func (m *MockRepository) Update(user *domain.User) error {
	m.users[user.ID] = user
	return nil
}

func (m *MockRepository) Delete(id int) error {
	delete(m.users, id)
	return nil
}

func TestUserUseCase_WithMock(t *testing.T) {
	mock := &MockRepository{users: make(map[int]*domain.User)}
	uc := NewUserUseCase(mock)

	user, err := uc.CreateUser("Mock User", "mock@example.com")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if user.ID != 999 {
		t.Errorf("Expected mock ID 999, got %d", user.ID)
	}
}
