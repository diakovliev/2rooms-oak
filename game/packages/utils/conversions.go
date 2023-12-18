package utils

import "fmt"

func AsPtr[T any](t T) *T {
	return &t
}

func As[T any](t any) (ret T) {
	ret, ok := t.(T)
	if !ok {
		panic(fmt.Errorf("can't cast %T to %T", t, ret))
	}
	return
}
