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

func Round(x, y float64) Option {
	return func(b *Button) {
		b.roundX = x
		b.roundY = y
	}
}

func Callback(cb func()) Option {
	return func(b *Button) {
		b.callback = cb
	}
}
