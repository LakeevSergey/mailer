package hasher

type EmptyHasher struct{}

func NewEmptyHasher() *EmptyHasher {
	return &EmptyHasher{}
}

func (h *EmptyHasher) Hash(data string) string {
	return ""
}

func (h *EmptyHasher) Equal(data string, hash string) bool {
	return true
}
