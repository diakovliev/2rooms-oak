package halign

import (
	"image/color"
	"runtime"
	"time"

	"github.com/diakovliev/2rooms-oak/packages/common"
	"github.com/diakovliev/2rooms-oak/packages/layout2d"
	"github.com/diakovliev/2rooms-oak/packages/layout2d/direction"
	"github.com/diakovliev/2rooms-oak/packages/move2d"
	"github.com/diakovliev/2rooms-oak/packages/scene/scenes"
	"github.com/diakovliev/2rooms-oak/packages/stimer"
	"github.com/diakovliev/2rooms-oak/packages/ui/button"
	"github.com/diakovliev/2rooms-oak/packages/ui/debuginfo"
	"github.com/diakovliev/2rooms-oak/packages/utils"
	"github.com/diakovliev/oak/v4"
	"github.com/diakovliev/oak/v4/alg/floatgeom"
	"github.com/diakovliev/oak/v4/entities"
	oakscene "github.com/diakovliev/oak/v4/scene"
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
		Name:   scenes.HAlign,
		End:    scenes.MainMenu,
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
			"Back",
			button.Dimensions(btnWidth, btnHeight),
			button.Callback(func() {
				ctx.Window.GoToScene(s.End)
			}),
		),
	)
	menu.Apply(layout2d.VCenter)
	// Place the menu in the center of a screen
	menuRect := layout2d.AlignRect(
		layout2d.Left|layout2d.Bottom,
		viewport,
		floatgeom.Rect2{
			Min: viewport.Min,
			Max: menu.Dims(),
		}, 20,
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

func (s Scene) start(ctx *oakscene.Context) {
	debuginfo.DebugInfo(ctx)
	s.makeMenu(ctx)
	s.testHorizontalEntities(ctx, move2d.New(ctx))
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
