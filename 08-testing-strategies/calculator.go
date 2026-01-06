package main

import "errors"

// Calculator performs mathematical operations
type Calculator struct{}

func NewCalculator() *Calculator {
	return &Calculator{}
}

func (c *Calculator) Add(a, b int) int {
	return a + b
}

func (c *Calculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// DataStore interface for testability
type DataStore interface {
	Get(key string) (string, error)
	Set(key string, value string) error
}

// UserService demonstrates dependency injection for testing
type UserService struct {
	store DataStore
}

func NewUserService(store DataStore) *UserService {
	return &UserService{store: store}
}

func (s *UserService) GetUsername(userID string) (string, error) {
	return s.store.Get("user:" + userID)
}

func (s *UserService) SaveUsername(userID, username string) error {
	return s.store.Set("user:"+userID, username)
}
