package utils

import "fmt"

type Stringer struct {
	str string
}

func (s Stringer) String() string {
	return s.str
}

func S(s string) fmt.Stringer {
	return Stringer{s}
}
