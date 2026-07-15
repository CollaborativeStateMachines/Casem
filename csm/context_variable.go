package csm

import (
	"errors"
	"fmt"
	"strings"
)

type ContextVariable interface {
	String() string
}

type typedVariable[T any] struct {
	name  string
	value T
}

func NewContextVariable[T any](name string, value T) (ContextVariable, error) {
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("name cannot be blank")
	}
	return &typedVariable[T]{
		name:  name,
		value: value,
	}, nil
}

func (cv *typedVariable[T]) String() string {
	if cv == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ContextVariable(name=%q, value=%+v)", cv.name, cv.value)
}
