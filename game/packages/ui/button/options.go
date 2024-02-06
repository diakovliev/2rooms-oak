package button

import (
	"github.com/diakovliev/oak/v4/alg/intgeom"
)

type Option func(b *data)

func ISize(w, h int) Option {
	return func(b *data) {
		b.size = intgeom.Point2{w, h}
	}
}

func FSize(w, h float64) Option {
	return func(b *data) {
		b.size = intgeom.Point2{int(w), int(h)}
	}
}

func Round(x, y float64) Option {
	return func(b *data) {
		b.roundX = x
		b.roundY = y
	}
}

func Callback(cb func()) Option {
	return func(b *data) {
		b.callback = cb
	}
}
