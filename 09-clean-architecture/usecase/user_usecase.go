package usecase

import (
	"errors"
	"github.com/tutorial/clean-architecture/domain"
	"strings"
)

// UserUseCase handles user business logic
type UserUseCase struct {
	repo domain.UserRepository
}

func NewUserUseCase(repo domain.UserRepository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (uc *UserUseCase) GetUser(id int) (*domain.User, error) {
	return uc.repo.FindByID(id)
}

func (uc *UserUseCase) CreateUser(name, email string) (*domain.User, error) {
	// Business validation
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}
	if !strings.Contains(email, "@") {
		return nil, errors.New("invalid email format")
	}

	user := &domain.User{
		Name:  name,
		Email: email,
	}

	if err := uc.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *UserUseCase) UpdateUser(id int, name, email string) error {
	user, err := uc.repo.FindByID(id)
	if err != nil {
		return err
	}

	if name != "" {
		user.Name = name
	}
	if email != "" {
		if !strings.Contains(email, "@") {
			return errors.New("invalid email format")
		}
		user.Email = email
	}

	return uc.repo.Update(user)
}

func (uc *UserUseCase) DeleteUser(id int) error {
	return uc.repo.Delete(id)
}
