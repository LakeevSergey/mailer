package encoder

import "encoding/json"

type JSONEncoder[T any] struct {
}

func NewJSONEncoder[T any]() *JSONEncoder[T] {
	return &JSONEncoder[T]{}
}

func (e *JSONEncoder[T]) Encode(data T) ([]byte, error) {
	return json.Marshal(data)
}

func (e *JSONEncoder[T]) Decode(data []byte) (T, error) {
	var res T
	err := json.Unmarshal(data, &res)

	return res, err
}
