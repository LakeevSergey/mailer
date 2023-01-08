package hasher

import (
	"github.com/stretchr/testify/mock"
)

type MockHasher struct {
	mock.Mock
}

func NewMockHasher() *MockHasher {
	return &MockHasher{}
}

func (h *MockHasher) Hash(data string) string {
	return data
}

func (h *MockHasher) Equal(data string, hash string) bool {
	return data == hash
}
