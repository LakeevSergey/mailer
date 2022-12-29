package coder

import "encoding/json"

type JSONCoder[T any] struct {
}

func NewJSONCoder[T any]() *JSONCoder[T] {
	return &JSONCoder[T]{}
}

func (c *JSONCoder[T]) Encode(data T) ([]byte, error) {
	return json.Marshal(data)
}

func (c *JSONCoder[T]) Decode(data []byte) (T, error) {
	var res T
	err := json.Unmarshal(data, &res)

	return res, err
}

func (c *JSONCoder[T]) ContentType() string {
	return "application/json"
}
