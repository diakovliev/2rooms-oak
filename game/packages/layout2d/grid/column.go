package grid

import "github.com/diakovliev/2rooms-oak/packages/common"

type Column struct {
	Grid     *Grid
	entities []common.Entity
}

func (c Column) Len() int {
	return len(c.entities)
}

func (c Column) H() float64 {
	height := c.Grid.margin
	for _, ee := range c.entities {
		if ee == nil {
			height += c.Grid.margin
			continue
		}
		height += ee.H() + c.Grid.margin
	}
	return height
}
