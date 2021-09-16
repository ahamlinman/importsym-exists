package value

type Value[T any] struct {
	value T
}

func NewValue[T any](value T) *Value[T] {
	return &Value[T]{value}
}

func (v *Value[T]) Get() T {
	return v.value
}
