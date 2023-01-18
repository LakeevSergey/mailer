package queue

type Encoder[T any] interface {
	ContentType() string
	Encode(data T) ([]byte, error)
	Decode(data []byte) (T, error)
}
