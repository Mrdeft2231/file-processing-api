package grpcauth

import (
	"github.com/stretchr/testify/mock"
)

type MockAuth struct {
	mock.Mock
}

func (m *MockAuth) Register(username string, password string) (string, error) {
	args := m.Called(username, password)
	return args.String(0), args.Error(1)
}

func (m *MockAuth) Login(username string, password string) (string, string, error) {
	args := m.Called(username, password)
	return args.String(0), args.String(1), args.Error(2)
}

func (m *MockAuth) RefreshToken(token string) (string, error) {
	args := m.Called(token)
	return args.String(0), args.Error(1)
}

func (m *MockAuth) ValidateToken(token string) (bool, error) {
	args := m.Called(token)
	return args.Bool(0), args.Error(1)
}

func (m *MockAuth) Logout(token string) error {
	args := m.Called(token)
	return args.Error(0)
}
