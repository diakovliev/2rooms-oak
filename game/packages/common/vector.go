package common

import "github.com/oakmound/oak/v4/alg/floatgeom"

// Vector is a struct that contains the entity, the delta, the old and new positions.
type Vector struct {
	Entity Entity
	Delta  floatgeom.Point2
	Old    floatgeom.Point2
	New    floatgeom.Point2
	Speed  floatgeom.Point2
}

func (v Vector) Reverse() Vector {
	oldOld := v.Old
	oldNew := v.New
	v.New = oldOld
	v.Old = oldNew
	return v
}
