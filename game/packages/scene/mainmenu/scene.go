package mainmenu

import (
	"image/color"
	"runtime"
	"time"

	"github.com/diakovliev/2rooms-oak/packages/common"
	"github.com/diakovliev/2rooms-oak/packages/layout2d"
	"github.com/diakovliev/2rooms-oak/packages/layout2d/direction"
	"github.com/diakovliev/2rooms-oak/packages/layout2d/grid"
	"github.com/diakovliev/2rooms-oak/packages/move2d"
	"github.com/diakovliev/2rooms-oak/packages/scene/scenes"
	"github.com/diakovliev/2rooms-oak/packages/stimer"
	"github.com/diakovliev/2rooms-oak/packages/ui/button"
	"github.com/diakovliev/2rooms-oak/packages/ui/debuginfo"
	"github.com/diakovliev/2rooms-oak/packages/utils"
	"github.com/oakmound/oak/v4"
	"github.com/oakmound/oak/v4/alg/floatgeom"
	"github.com/oakmound/oak/v4/entities"
	oakscene "github.com/oakmound/oak/v4/scene"
	"github.com/rs/zerolog"
)

type Scene struct {
	logger zerolog.Logger
	Name   string
	End    string
}

func New(logger zerolog.Logger, w *oak.Window) (ret *Scene) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	ret = &Scene{
		logger: logger,
		Name:   scenes.MainMenu,
		End:    scenes.Debug,
	}
	if err := w.AddScene(ret.Name, ret.build()); err != nil {
		logger.Panic().Err(err).Msgf("failed to add '%s' scene", ret.Name)
	}
	return
}

func (s Scene) makeMenu(ctx *oakscene.Context) common.Entity {
	btnWidth := 100
	btnHeight := 20
	spaceBetween := 3.
	viewport := utils.ViewportRect(ctx)
	menu := direction.Vertical(
		ctx,
		viewport.Min,
		floatgeom.Point2{},
		spaceBetween,
	).Add(
		button.New(
			ctx,
			"Tests",
			button.Dimensions(btnWidth, btnHeight),
			button.Callback(func() {
				ctx.Window.GoToScene(s.End)
			}),
		),
		button.New(
			ctx,
			"Debug",
			button.Dimensions(btnWidth, btnHeight),
			button.Callback(func() {
				ctx.Window.GoToScene(s.End)
			}),
		),
		button.New(
			ctx,
			"Quit",
			button.Dimensions(btnWidth, btnHeight),
			button.Callback(func() {
				ctx.Window.Quit()
			}),
		),
	)
	menu.Apply(layout2d.VCenter)
	// Place the menu in the center of a screen
	menuRect := layout2d.AlignRect(
		layout2d.HCenter|layout2d.VCenter,
		viewport,
		floatgeom.Rect2{
			Min: viewport.Min,
			Max: menu.Dims(),
		}, 0,
	)
	menu.SetPos(menuRect.Min)
	return menu
}

func (s Scene) testHorizontalEntities(ctx *oakscene.Context, mover *move2d.Manager) common.Entity {
	entLayer := 10

	speed := floatgeom.Point2{5, 5}
	color := color.RGBA{0, 255, 0, 255}
	layers := []int{entLayer}
	pos := floatgeom.Point2{100, 200}

	commonOpts := entities.And(
		entities.WithSpeed(speed),
		entities.WithColor(color),
		entities.WithDrawLayers(layers),
	)

	testHorizontalLayout := direction.Horizontal(ctx, pos, speed, 3.).Add(
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

	alignments := []layout2d.Alignment{
		layout2d.HCenter,
		layout2d.Top,
		layout2d.Bottom,
	}

	index := 0
	stimer.New(
		ctx,
		time.Second,
		func() {
			vectors := testHorizontalLayout.Vectors(alignments[index%len(alignments)])
			for _, vector := range vectors {
				mover.Start(vector, vector.Speed)
			}
			index++
		},
	)

	return testHorizontalLayout
}

func (s Scene) testVerticalEntities(ctx *oakscene.Context, mover *move2d.Manager) common.Entity {
	entLayer := 10

	speed := floatgeom.Point2{5, 5}
	color := color.RGBA{0, 255, 0, 255}
	layers := []int{entLayer}
	pos := floatgeom.Point2{200, 200}

	commonOpts := entities.And(
		entities.WithSpeed(speed),
		entities.WithColor(color),
		entities.WithDrawLayers(layers),
	)

	testVerticalLayout := direction.Vertical(ctx, pos, speed, 3.).Add(
		s.testHorizontalEntities(ctx, mover),
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

	alignments := []layout2d.Alignment{
		layout2d.VCenter,
		layout2d.Left,
		layout2d.Right,
	}

	index := 0
	stimer.New(
		ctx,
		time.Second,
		func() {
			vectors := testVerticalLayout.Vectors(alignments[index%len(alignments)])
			for _, vector := range vectors {
				mover.Start(vector, vector.Speed)
			}
			index++
		},
	)

	return testVerticalLayout
}

func (s Scene) testGrid(ctx *oakscene.Context, mover *move2d.Manager) common.Entity {

	bounds := ctx.Window.Bounds()

	entLayer := 10

	pos := floatgeom.Point2{float64(bounds.X() / 4), float64(bounds.Y() / 4)}
	initialPos := floatgeom.Point2{float64(bounds.X() / 2), float64(bounds.Y() / 2)}
	speed := floatgeom.Point2{5, 5}
	layers := []int{entLayer}

	commonOpts := entities.And(
		entities.WithSpeed(speed),
		entities.WithDrawLayers(layers),
	)

	index := 0
	grid := grid.New(ctx, pos, speed, 2).Init(20, 20, func(row, col int) common.Entity {
		dims := floatgeom.Point2{10, 10}
		clr := color.RGBA{0, 255, 0, 255}
		if index%3 == 1 {
			dims = floatgeom.Point2{6, 6}
			clr = color.RGBA{255, 255, 255, 255}
		}
		index++
		return entities.New(ctx,
			entities.WithDimensions(dims),
			entities.WithPosition(initialPos),
			entities.WithColor(clr),
			commonOpts,
		)
	})

	vectors := grid.Vectors(layout2d.HCenter | layout2d.VCenter)
	//vectors := grid.Vectors(layout2d.Left | layout2d.Top)
	//vectors := grid.Vectors(layout2d.Right | layout2d.Bottom)
	reverse := make([]common.Vector, 0)
	for _, vector := range vectors {
		reverse = append(reverse, vector.Reverse())
	}

	actions := [][]common.Vector{
		vectors,
		reverse,
	}

	index = 0
	stimer.New(
		ctx,
		time.Second,
		func() {
			vectors = actions[index%len(actions)]
			for _, vector := range vectors {
				mover.Start(vector, vector.Speed)
			}
			index++
		},
	)

	return grid
}

func (s Scene) start(ctx *oakscene.Context) {
	debuginfo.DebugInfo(ctx)
	s.makeMenu(ctx)
	// // cursor
	// cursorLayer := 100

	// cursor := entities.New(ctx,
	// 	//entities.WithUseMouseTree(true),
	// 	entities.WithRect(floatgeom.NewRect2WH(10, 10, 10, 10)),
	// 	entities.WithColor(color.RGBA{255, 255, 255, 255}),
	// 	entities.WithSpeed(floatgeom.Point2{5, 5}),
	// 	// Over all
	// 	entities.WithDrawLayers([]int{cursorLayer}),
	// )

	// event.Bind(ctx, event.Enter, cursor, func(c *entities.Entity, ev event.EnterPayload) event.Response {
	// 	return event.ResponseNone
	// })

	// event.Bind(ctx, mouse.Drag, cursor, func(c *entities.Entity, ev *mouse.Event) event.Response {
	// 	c.SetPos(ev.Point2)
	// 	return event.ResponseNone
	// })

	//mover := move2d.New(ctx)

	//s.testHorizontalEntities(ctx, mover)
	//s.testVerticalEntities(ctx, mover)
	//s.testGrid(ctx, mover)
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
