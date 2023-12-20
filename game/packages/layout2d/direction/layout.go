package direction

import (
	"sync"

	"github.com/diakovliev/2rooms-oak/packages/layout2d"
	"github.com/oakmound/oak/v4/alg/floatgeom"
)

// Layout is a layout
type Layout struct {
	*sync.Mutex
	entities  []layout2d.Entity
	pos       floatgeom.Point2
	alignment layout2d.Alignment
	margin    float64
	w, h      float64
	vectors   func(layout *Layout, alignment layout2d.Alignment) (ret []layout2d.Vectors)
	add       func(layout *Layout, e []layout2d.Entity)
}

// newLayout creates a new Layout with the given parameters.
//
// Parameters:
// - pos: the position of the layout.
// - margin: the margin of the layout.
// - vectors: a function that returns a slice of Delta.
// - add: a function that adds entities to the layout.
//
// Returns:
// - a pointer to the newly created Layout.
func newLayout(
	pos floatgeom.Point2,
	margin float64,
	vectors func(layout *Layout, alignment layout2d.Alignment) (ret []layout2d.Vectors),
	add func(layout *Layout, e []layout2d.Entity),
) *Layout {
	return &Layout{
		Mutex:     &sync.Mutex{},
		pos:       pos,
		margin:    margin,
		alignment: layout2d.Left,
		w:         margin,
		h:         margin,
		vectors:   vectors,
		add:       add,
	}
}

// Add adds the given entities to the layout and returns a pointer to the modified layout.
//
// It takes a variadic parameter of type Entity.
// It returns a pointer to the Layout.
func (l Layout) Add(e ...layout2d.Entity) *Layout {
	l.Mutex = &sync.Mutex{}
	l.add(&l, e)
	return &l
}

// Dims returns the dimensions of the Layout as a floatgeom.Point2.
//
// There are no parameters.
// It returns a floatgeom.Point2.
func (l Layout) Dims() floatgeom.Point2 {
	return floatgeom.Point2{l.w, l.h}
}

// X returns the X-coordinate of the Layout object.
//
// No parameters.
// Returns a float64 value.
func (l Layout) X() float64 {
	return l.pos.X()
}

// Y returns the Y coordinate of the layout.
//
// No parameters.
// Returns a float64.
func (l Layout) Y() float64 {
	return l.pos.Y()
}

// W returns the value of the W field in the Layout struct.
//
// No parameters.
// Returns a float64.
func (l Layout) W() float64 {
	return l.w
}

// H returns the value of the 'h' field in the Layout struct.
//
// No parameters.
// Returns a float64.
func (l Layout) H() float64 {
	return l.h
}

// SetPos sets the position of the VLayout to the specified point.
//
// p: The new position for the VLayout.
func (l *Layout) SetPos(p floatgeom.Point2) {
	l.Lock()
	l.pos = p
	l.Unlock()
	l.Apply(l.alignment)
}

// Vectors returns a slice of Delta objects representing the vectors in the Layout,
// based on the specified alignment.
//
// The alignment parameter specifies the desired alignment for the vectors.
//
// The return type of this function is []Delta.
func (l *Layout) Vectors(alignment layout2d.Alignment) []layout2d.Vectors {
	l.Lock()
	defer l.Unlock()
	return l.vectors(l, alignment)
}

// Apply rearranges the entities in a Layout according to the specified alignment.
//
// alignment is the alignment to use.
func (l *Layout) Apply(alignment layout2d.Alignment) {
	l.Lock()
	defer l.Unlock()
	l.alignment = alignment
	for _, vector := range l.vectors(l, alignment) {
		vector.Entity.SetPos(vector.New)
	}
}
