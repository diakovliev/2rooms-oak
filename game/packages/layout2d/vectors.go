package layout2d

import "github.com/oakmound/oak/v4/alg/floatgeom"

type Vectors struct {
	Entity Entity
	Delta  floatgeom.Point2
	Old    floatgeom.Point2
	New    floatgeom.Point2
}
