package encoder

import "encoding/json"

type JSONEncoder[T any] struct {
}

func NewJSONEncoder[T any]() *JSONEncoder[T] {
	return &JSONEncoder[T]{}
}

func (c *JSONEncoder[T]) Encode(data T) ([]byte, error) {
	return json.Marshal(data)
}

func (c *JSONEncoder[T]) Decode(data []byte) (T, error) {
	var res T
	err := json.Unmarshal(data, &res)

	return res, err
}

func (c *JSONEncoder[T]) ContentType() string {
	return "application/json"
}
