package config_test

import "github.com/stretchr/testify/mock"

type MockFileReader struct {
	mock.Mock
}

func (m *MockFileReader) ReadAll(filename string) ([]byte, error) {
	args := m.Called(filename)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockFileReader) Exists(filename string) bool {
	args := m.Called(filename)
	return args.Bool(0)
}
