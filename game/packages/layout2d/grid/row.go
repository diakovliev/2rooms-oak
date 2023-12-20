package grid

import "github.com/diakovliev/2rooms-oak/packages/common"

type Row struct {
	Grid     *Grid
	entities []common.Entity
}

func (r Row) Len() int {
	return len(r.entities)
}

// W returns the width of a Row.
//
// It calculates the width of the Row by summing the widths of all entities within the Row, along with the margin between them.
// The width of each entity is obtained by calling its W() method.
//
// Returns the calculated width as a float64.
func (r Row) W() float64 {
	width := r.Grid.margin
	for _, ee := range r.entities {
		if ee == nil {
			width += r.Grid.margin
			continue
		}
		width += ee.W() + r.Grid.margin
	}
	return width
}

// H calculates the height of the Row.
//
// It iterates over the entities of the Row and checks their height
// to determine the maximum height. The height is calculated by adding
// the margin value to the height of each entity. If an entity is nil,
// the margin value is added to the height.
//
// Returns the calculated height as a float64 value.
func (r Row) H() float64 {
	height := r.Grid.margin
	for _, ee := range r.entities {
		if ee == nil {
			height += r.Grid.margin
			continue
		}
		if height < ee.H()+r.Grid.margin {
			height = ee.H() + r.Grid.margin
		}
	}
	return height
}
