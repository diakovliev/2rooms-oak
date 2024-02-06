package mainmenu

import (
	"runtime"

	"github.com/diakovliev/2rooms-oak/packages/common"
	"github.com/diakovliev/2rooms-oak/packages/layout2d"
	"github.com/diakovliev/2rooms-oak/packages/layout2d/direction"
	"github.com/diakovliev/2rooms-oak/packages/scene/scenes"
	"github.com/diakovliev/2rooms-oak/packages/ui/button"
	"github.com/diakovliev/2rooms-oak/packages/ui/debuginfo"
	"github.com/diakovliev/2rooms-oak/packages/utils"
	"github.com/diakovliev/oak/v4"
	"github.com/diakovliev/oak/v4/alg/floatgeom"
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
		Name:   scenes.MainMenu,
		End:    scenes.MainMenu,
	}
	if err := w.AddScene(ret.Name, ret.build()); err != nil {
		logger.Panic().Err(err).Msgf("failed to add '%s' scene", ret.Name)
	}
	return
}

func (s Scene) makeMenu(ctx *oakscene.Context) common.Entity {
	btnWidth := 140
	btnHeight := 26
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
			"Grid align test",
			button.ISize(btnWidth, btnHeight),
			button.Callback(func() {
				ctx.Window.GoToScene(scenes.GAlign)
			}),
		),
		button.New(
			ctx,
			"Horizontal align test",
			button.ISize(btnWidth, btnHeight),
			button.Callback(func() {
				ctx.Window.GoToScene(scenes.HAlign)
			}),
		),
		button.New(
			ctx,
			"Vertical align test",
			button.ISize(btnWidth, btnHeight),
			button.Callback(func() {
				ctx.Window.GoToScene(scenes.VAlign)
			}),
		),
		button.New(
			ctx,
			"Quit",
			button.ISize(btnWidth, btnHeight),
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
