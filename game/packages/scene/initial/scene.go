package initial

import (
	"image/color"

	"github.com/diakovliev/2rooms-oak/packages/scene"
	"github.com/diakovliev/2rooms-oak/packages/utils"
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

	entLayer := 10

	testLayout := utils.NewVLayout(floatgeom.Point2{200, 200}, 3.).Add(
		entities.New(ctx,
			entities.WithDimensions(floatgeom.Point2{100, 10}),
			entities.WithColor(color.RGBA{0, 255, 0, 255}),
			entities.WithDrawLayers([]int{entLayer}),
		),
		entities.New(ctx,
			entities.WithDimensions(floatgeom.Point2{200, 10}),
			entities.WithColor(color.RGBA{0, 255, 0, 255}),
			entities.WithDrawLayers([]int{entLayer}),
		),
		entities.New(ctx,
			entities.WithDimensions(floatgeom.Point2{100, 10}),
			entities.WithColor(color.RGBA{0, 255, 0, 255}),
			entities.WithDrawLayers([]int{entLayer}),
		),
		entities.New(ctx,
			entities.WithDimensions(floatgeom.Point2{200, 10}),
			entities.WithColor(color.RGBA{0, 255, 0, 255}),
			entities.WithDrawLayers([]int{entLayer}),
		),
	)

	testLayout.Rearrange(utils.VCenter)
	//testLayout.Rearrange(utils.VLeft)
	//testLayout.Rearrange(utils.VRight)
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
