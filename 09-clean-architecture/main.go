package main

import (
	"fmt"
	"github.com/tutorial/clean-architecture/repository"
	"github.com/tutorial/clean-architecture/usecase"
)

func main() {
	// Dependency injection: wire up layers
	repo := repository.NewMemoryUserRepository()
	userUC := usecase.NewUserUseCase(repo)

	// Use the use case
	fmt.Println("=== Clean Architecture Demo ===")

	// Create user
	user, err := userUC.CreateUser("Alice", "alice@example.com")
	if err != nil {
		fmt.Printf("Error creating user: %v\n", err)
		return
	}
	fmt.Printf("Created user: ID=%d, Name=%s, Email=%s\n", user.ID, user.Name, user.Email)

	// Get user
	retrieved, err := userUC.GetUser(user.ID)
	if err != nil {
		fmt.Printf("Error getting user: %v\n", err)
		return
	}
	fmt.Printf("Retrieved user: %s\n", retrieved.Name)

	// Update user
	err = userUC.UpdateUser(user.ID, "Alice Smith", "")
	if err != nil {
		fmt.Printf("Error updating user: %v\n", err)
		return
	}
	fmt.Println("User updated successfully")

	// Get updated user
	updated, _ := userUC.GetUser(user.ID)
	fmt.Printf("Updated user: Name=%s\n", updated.Name)

	// Delete user
	err = userUC.DeleteUser(user.ID)
	if err != nil {
		fmt.Printf("Error deleting user: %v\n", err)
		return
	}
	fmt.Println("User deleted successfully")
}
