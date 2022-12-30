package queue

type Coder[T any] interface {
	ContentType() string
	Encode(data T) ([]byte, error)
	Decode(data []byte) (T, error)
}
