package grid

import (
	"sync"

	"github.com/diakovliev/2rooms-oak/packages/layout2d"
	"github.com/oakmound/oak/v4/alg/floatgeom"
)

// Grid is a grid
type Grid struct {
	*sync.Mutex
	rows      []Row
	pos       floatgeom.Point2
	alignment layout2d.Alignment
	w, h      float64
	margin    float64
}

// New returns a new Grid instance.
//
// Parameters:
//   - pos: the position of the Grid.
//   - margin: the margin of the Grid.
//
// Returns:
//   - *Grid: a pointer to the created Grid instance.
func New(pos floatgeom.Point2, margin float64) *Grid {
	return &Grid{
		Mutex:     &sync.Mutex{},
		pos:       pos,
		margin:    margin,
		alignment: layout2d.Left | layout2d.Top,
		w:         margin,
		h:         margin,
		rows:      []Row{},
	}
}

// update updates the dimensions of the Grid.
//
// It iterates over the entities in the Grid and updates the width and height
// based on the maximum width of the columns and the maximum height of the rows.
// The width and height are stored in the Grid's w and h fields respectively.
func (g *Grid) update() {
	g.w = 2 * g.margin
	g.h = 2 * g.margin
	for i := 0; i < len(g.rows); i++ {
		for j := 0; j < len(g.rows[i].entities); j++ {
			if g.h < g.Column(j).H() {
				g.h = g.Column(j).H()
			}
		}
		if g.w < g.Row(i).W() {
			g.w = g.Row(i).W()
		}
	}
}

// Init creates a new Grid with the specified number of rows and columns, and initializes it with entities
// generated by the provided provider function. The Grid is returned as a pointer.
//
// Parameters:
// - rows: the number of rows in the Grid
// - cols: the number of columns in the Grid
// - provider: a function that generates entities for each cell in the Grid, based on the row and column index
//
// Return:
// - *Grid: a pointer to the newly created Grid
func (g *Grid) Init(rows, cols int, provider func(row, col int) layout2d.Entity) *Grid {
	g.Mutex = &sync.Mutex{}
	g.rows = make([]Row, rows)
	for i := 0; i < rows; i++ {
		g.rows[i] = Row{g, make([]layout2d.Entity, cols)}
		for j := 0; j < cols; j++ {
			if provider == nil {
				g.rows[i].entities[j] = nil
			} else {
				g.rows[i].entities[j] = provider(i, j)
			}
		}
	}
	g.update()
	return g
}

// Set sets the given entity at the specified row and column in the Grid.
//
// Parameters:
// - row: the row index in the Grid where the entity will be set.
// - col: the column index in the Grid where the entity will be set.
// - entity: the entity to be set in the Grid.
//
// Return:
// - *Grid: the modified Grid.
func (g *Grid) Set(row, col int, entity layout2d.Entity) *Grid {
	g.Lock()
	defer g.Unlock()
	g.rows[row].entities[col] = entity
	g.update()
	return g
}

// Get returns the layout2d.Entity at the specified row and column indices in the Grid.
//
// Parameters:
// - row: The row index of the desired entity.
// - col: The column index of the desired entity.
//
// Returns:
// - layout2d.Entity: The entity located at the specified row and column indices.
func (g Grid) Get(row, col int) layout2d.Entity {
	g.Lock()
	defer g.Unlock()
	return g.rows[row].entities[col]
}

// Row returns a Row object representing the entities in the grid at the specified index.
//
// Parameters:
//   - index: The index of the row.
//
// Returns:
//   - Row: A Row object containing the entities at the specified index.
func (g *Grid) Row(index int) Row {
	g.Lock()
	defer g.Unlock()
	if index >= len(g.rows) {
		return Row{g, []layout2d.Entity{}}
	}
	return Row{g, g.rows[index].entities}
}

// Column returns the column at the specified index in the Grid.
//
// The index parameter specifies the index of the column to retrieve.
// The function returns a Column object representing the column at the given index.
func (g *Grid) Column(index int) Column {
	g.Lock()
	defer g.Unlock()
	column := make([]layout2d.Entity, 0, len(g.rows))
	for _, row := range g.rows {
		if index >= len(row.entities) {
			column = append(column, nil)
		} else {
			column = append(column, row.entities[index])
		}
	}
	return Column{g, column}
}

func (g Grid) vectors(alignment layout2d.Alignment) (ret []layout2d.Vectors) {
	top := g.margin
	for _, row := range g.rows {
		left := g.margin
		for _, entity := range row.entities {
			if entity == nil {
				left += g.margin
				continue
			}
			w := entity.W()
			h := entity.H()
			oldPos := floatgeom.Point2{entity.X(), entity.Y()}
			newX := oldPos.X()
			newY := oldPos.Y()
			if g.alignment&layout2d.Left|layout2d.Right|layout2d.HCenter != 0 {
				switch {
				case g.alignment&layout2d.Left == layout2d.Left:
					newX = left + g.margin
				case g.alignment&layout2d.HCenter == layout2d.HCenter:
					newX = left + g.w/2 - w/2
				case g.alignment&layout2d.Right == layout2d.Right:
					newX = left + g.w - w - g.margin
				}
			}
			if g.alignment&layout2d.Top|layout2d.Bottom|layout2d.VCenter != 0 {
				switch {
				case g.alignment&layout2d.Top == layout2d.Top:
					newY = top + g.margin
				case g.alignment&layout2d.VCenter == layout2d.VCenter:
					newY = top + (g.h-h)/2
				case g.alignment&layout2d.Bottom == layout2d.Bottom:
					newY = top + g.h - h - g.margin
				}
			}
			newPos := floatgeom.Point2{newX, newY}
			ret = append(ret, layout2d.Vectors{
				Entity: entity,
				Delta:  newPos.Sub(oldPos),
				Old:    oldPos,
				New:    newPos,
			})
			left += entity.W() + g.margin
		}
		top += row.H() + g.margin
	}
	return
}

func (g *Grid) Apply(alignment layout2d.Alignment) {
	g.Lock()
	defer g.Unlock()
	g.alignment = alignment
	for _, vector := range g.vectors(alignment) {
		vector.Entity.SetPos(vector.New)
	}
}
