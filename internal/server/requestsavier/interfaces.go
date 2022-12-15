package requestsavier

type Encoder[T any] interface {
	Encode(data T) ([]byte, error)
}
