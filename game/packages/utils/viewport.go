package utils

import (
	"github.com/oakmound/oak/v4/alg/floatgeom"
	oakscene "github.com/oakmound/oak/v4/scene"
)

func ViewportRect(ctx *oakscene.Context) (ret floatgeom.Rect2) {
	bounds := ctx.Window.Bounds()
	viewport := ctx.Window.Viewport()
	ret = floatgeom.Rect2{
		Min: floatgeom.Point2{float64(viewport.X()), float64(viewport.Y())},
		Max: floatgeom.Point2{float64(bounds.X()), float64(bounds.Y())},
	}
	return
}
