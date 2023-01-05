package sign

type HashChecker interface {
	Equal(data string, hash string) bool
}
