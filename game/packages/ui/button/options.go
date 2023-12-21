package button

import (
	"github.com/oakmound/oak/v4/alg/intgeom"
)

type Option func(b *Button)

func Dimensions(w, h int) Option {
	return func(b *Button) {
		b.dims = intgeom.Point2{w, h}
	}
}
