package direction

import (
	"sync"

	"github.com/diakovliev/2rooms-oak/packages/common"
	"github.com/diakovliev/2rooms-oak/packages/layout2d"
	"github.com/diakovliev/oak/v4/alg/floatgeom"
	"github.com/diakovliev/oak/v4/event"
	oakscene "github.com/diakovliev/oak/v4/scene"
)

// Layout is a layout
type Layout struct {
	*sync.Mutex
	cid       event.CallerID
	entities  []common.Entity
	pos       floatgeom.Point2
	speed     floatgeom.Point2
	alignment layout2d.Alignment
	margin    float64
	w, h      float64
	vectors   func(layout *Layout, alignment layout2d.Alignment) (ret []common.Vector)
	add       func(layout *Layout, e []common.Entity)
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
	ctx *oakscene.Context,
	pos floatgeom.Point2,
	speed floatgeom.Point2,
	margin float64,
	vectors func(layout *Layout, alignment layout2d.Alignment) (ret []common.Vector),
	add func(layout *Layout, e []common.Entity),
) (ret *Layout) {
	ret = &Layout{
		Mutex:     &sync.Mutex{},
		pos:       pos,
		speed:     speed,
		margin:    margin,
		alignment: layout2d.Left,
		w:         margin,
		h:         margin,
		vectors:   vectors,
		add:       add,
	}
	ret.cid = ctx.CallerMap.Register(ret)
	return
}

func (l Layout) CID() event.CallerID {
	l.Lock()
	defer l.Unlock()
	return l.cid
}

// Add adds the given entities to the layout and returns a pointer to the modified layout.
//
// It takes a variadic parameter of type Entity.
// It returns a pointer to the Layout.
func (l Layout) Add(e ...common.Entity) *Layout {
	l.Mutex = &sync.Mutex{}
	l.Lock()
	defer l.Unlock()
	l.add(&l, e)
	return &l
}

// Dims returns the dimensions of the Layout as a floatgeom.Point2.
//
// There are no parameters.
// It returns a floatgeom.Point2.
func (l Layout) Dims() floatgeom.Point2 {
	l.Lock()
	defer l.Unlock()
	return floatgeom.Point2{l.w, l.h}
}

// X returns the X-coordinate of the Layout object.
//
// No parameters.
// Returns a float64 value.
func (l Layout) X() float64 {
	l.Lock()
	defer l.Unlock()
	return l.pos.X()
}

// Y returns the Y coordinate of the layout.
//
// No parameters.
// Returns a float64.
func (l Layout) Y() float64 {
	l.Lock()
	defer l.Unlock()
	return l.pos.Y()
}

// W returns the value of the W field in the Layout struct.
//
// No parameters.
// Returns a float64.
func (l Layout) W() float64 {
	l.Lock()
	defer l.Unlock()
	return l.w
}

// H returns the value of the 'h' field in the Layout struct.
//
// No parameters.
// Returns a float64.
func (l Layout) H() float64 {
	l.Lock()
	defer l.Unlock()
	return l.h
}

// SetPos sets the position of the Layout to the specified point.
//
// p: The new position for the Layout.
func (l *Layout) SetPos(pos floatgeom.Point2) {
	l.Lock()
	defer l.Unlock()
	l.pos = pos
	l.apply(l.alignment)
}

// Vectors returns a slice of Delta objects representing the vectors in the Layout,
// based on the specified alignment.
//
// The alignment parameter specifies the desired alignment for the vectors.
//
// The return type of this function is []layout2d.Vectors.
func (l *Layout) Vectors(alignment layout2d.Alignment) []common.Vector {
	l.Lock()
	defer l.Unlock()
	l.alignment = alignment
	return l.vectors(l, alignment)
}

func (l *Layout) apply(alignment layout2d.Alignment) {
	l.alignment = alignment
	for _, vector := range l.vectors(l, alignment) {
		vector.Entity.SetPos(vector.New)
	}
}

// Apply rearranges the entities in a Layout according to the specified alignment.
//
// alignment is the alignment to use.
func (l *Layout) Apply(alignment layout2d.Alignment) {
	l.Lock()
	defer l.Unlock()
	l.apply(alignment)
}
