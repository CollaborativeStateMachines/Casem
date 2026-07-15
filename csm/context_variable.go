package csm

import (
	"errors"
	"fmt"
	"strings"
)

type ContextVariable[T any] struct {
	name  string
	Value T
}

func New[T any](name string, value T) (*ContextVariable[T], error) {
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("name cannot be blank")
	}
	return &ContextVariable[T]{
		name,
		value,
	}, nil
}

func (cv *ContextVariable[T]) String() string {
	if cv == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ContextVariable(name='%s', value='%+v')", cv.name, cv.Value)
}

func (cv *ContextVariable[T]) Name() string {
	if cv == nil {
		return "<nil>"
	}
	return cv.name
}
