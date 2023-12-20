package initial

import (
	"image/color"
	"time"

	"github.com/diakovliev/2rooms-oak/packages/layout2d"
	"github.com/diakovliev/2rooms-oak/packages/layout2d/direction"
	"github.com/diakovliev/2rooms-oak/packages/layout2d/grid"
	"github.com/diakovliev/2rooms-oak/packages/scene"
	"github.com/diakovliev/2rooms-oak/packages/window"
	"github.com/oakmound/oak/v4/alg/floatgeom"
	"github.com/oakmound/oak/v4/entities"
	"github.com/oakmound/oak/v4/event"
	"github.com/oakmound/oak/v4/mouse"
	oakscene "github.com/oakmound/oak/v4/scene"
	"github.com/rs/zerolog"
)

type Scene struct {
	logger zerolog.Logger
	Name   string
	End    string
}

func New(logger zerolog.Logger, w *window.Window) (ret *Scene) {
	ret = &Scene{
		logger: logger,
		Name:   scene.Initial,
		End:    scene.Debug,
	}
	if err := w.AddScene(ret.Name, ret.build()); err != nil {
		logger.Panic().Err(err).Msgf("failed to add '%s' scene", ret.Name)
	}
	return
}

func (s Scene) testHorizontalEntities(ctx *oakscene.Context) layout2d.Entity {
	entLayer := 10

	speed := floatgeom.Point2{5, 5}
	color := color.RGBA{0, 255, 0, 255}
	layers := []int{entLayer}

	commonOpts := entities.And(
		entities.WithSpeed(speed),
		entities.WithColor(color),
		entities.WithDrawLayers(layers),
	)

	testHorizontalLayout := direction.Horizontal(floatgeom.Point2{100, 200}, 3.).Add(
		entities.New(ctx,
			entities.WithDimensions(floatgeom.Point2{10, 100}),
			commonOpts,
		),
		entities.New(ctx,
			entities.WithDimensions(floatgeom.Point2{10, 200}),
			commonOpts,
		),
		entities.New(ctx,
			entities.WithDimensions(floatgeom.Point2{10, 100}),
			commonOpts,
		),
		entities.New(ctx,
			entities.WithDimensions(floatgeom.Point2{10, 200}),
			commonOpts,
		),
	)

	timer := time.NewTicker(time.Second)
	alignments := []layout2d.Alignment{
		layout2d.HCenter,
		layout2d.Top,
		layout2d.Bottom,
	}
	index := 0
	testHorizontalLayout.Apply(alignments[index])
	go func() {
		for {
			<-timer.C
			testHorizontalLayout.Apply(alignments[index%len(alignments)])
			index++
		}
	}()

	return testHorizontalLayout
}

func (s Scene) testVerticalEntities(ctx *oakscene.Context) layout2d.Entity {
	entLayer := 10

	speed := floatgeom.Point2{5, 5}
	color := color.RGBA{0, 255, 0, 255}
	layers := []int{entLayer}

	commonOpts := entities.And(
		entities.WithSpeed(speed),
		entities.WithColor(color),
		entities.WithDrawLayers(layers),
	)

	testVerticalLayout := direction.Vertical(floatgeom.Point2{200, 200}, 3.).Add(
		s.testHorizontalEntities(ctx),
		entities.New(ctx,
			entities.WithDimensions(floatgeom.Point2{100, 10}),
			commonOpts,
		),
		entities.New(ctx,
			entities.WithDimensions(floatgeom.Point2{200, 10}),
			commonOpts,
		),
		entities.New(ctx,
			entities.WithDimensions(floatgeom.Point2{100, 10}),
			commonOpts,
		),
		entities.New(ctx,
			entities.WithDimensions(floatgeom.Point2{200, 10}),
			commonOpts,
		),
	)

	timer := time.NewTicker(time.Second)
	alignments := []layout2d.Alignment{
		layout2d.VCenter,
		layout2d.Left,
		layout2d.Right,
	}
	index := 0
	testVerticalLayout.Apply(alignments[index])
	go func() {
		for {
			<-timer.C
			testVerticalLayout.Apply(alignments[index%len(alignments)])
			index++
		}
	}()

	return testVerticalLayout
}

func (s Scene) testGrid(ctx *oakscene.Context) {

	entLayer := 10

	pos := floatgeom.Point2{5, 5}
	speed := floatgeom.Point2{5, 5}
	color := color.RGBA{0, 255, 0, 255}
	layers := []int{entLayer}

	commonOpts := entities.And(
		entities.WithSpeed(speed),
		entities.WithColor(color),
		entities.WithDrawLayers(layers),
	)

	grid := grid.New(pos, 2).Init(20, 20, func(row, col int) layout2d.Entity {
		return entities.New(ctx,
			entities.WithDimensions(floatgeom.Point2{10, 10}),
			commonOpts,
		)
	})

	grid.Apply(layout2d.HCenter | layout2d.VCenter)
}

func (s Scene) start(ctx *oakscene.Context) {

	debugEntity(ctx)

	// cursor
	cursorLayer := 100

	cursor := entities.New(ctx,
		//entities.WithUseMouseTree(true),
		entities.WithRect(floatgeom.NewRect2WH(10, 10, 10, 10)),
		entities.WithColor(color.RGBA{255, 255, 255, 255}),
		entities.WithSpeed(floatgeom.Point2{5, 5}),
		// Over all
		entities.WithDrawLayers([]int{cursorLayer}),
	)

	event.Bind(ctx, event.Enter, cursor, func(c *entities.Entity, ev event.EnterPayload) event.Response {
		return event.ResponseNone
	})

	event.Bind(ctx, mouse.Drag, cursor, func(c *entities.Entity, ev *mouse.Event) event.Response {
		c.SetPos(ev.Point2)
		return event.ResponseNone
	})

	//s.testHorizontalEntities(ctx)
	s.testVerticalEntities(ctx)
	s.testGrid(ctx)
}

func (s Scene) end() (string, *oakscene.Result) {
	return s.End, nil
}

func (s Scene) build() oakscene.Scene {
	return oakscene.Scene{
		Start: s.start,
		End:   s.end,
	}
}
