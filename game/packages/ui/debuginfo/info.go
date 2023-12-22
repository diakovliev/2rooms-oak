package debuginfo

import (
	"fmt"
	"time"

	"github.com/diakovliev/2rooms-oak/packages/utils"
	"github.com/oakmound/oak/v4"
	"github.com/oakmound/oak/v4/alg/floatgeom"
	"github.com/oakmound/oak/v4/alg/intgeom"
	"github.com/oakmound/oak/v4/entities"
	"github.com/oakmound/oak/v4/event"
	"github.com/oakmound/oak/v4/mouse"
	oakscene "github.com/oakmound/oak/v4/scene"
	"github.com/oakmound/oak/v4/timing"
)

type debugCoords struct {
	Label string
	X, Y  int
}

func newMouseCoords(ctx *oakscene.Context) (ret *debugCoords) {
	ret = &debugCoords{
		Label: "M: ",
	}
	event.GlobalBind(ctx, mouse.Drag, ret.onMouseDrag)
	return
}

func newViewportPosition(ctx *oakscene.Context) (ret *debugCoords) {
	ret = &debugCoords{
		Label: "V: ",
	}
	event.GlobalBind(ctx, oak.ViewportUpdate, ret.onViewportUpdate)
	return
}

func (dc *debugCoords) onViewportUpdate(ev intgeom.Point2) event.Response {
	dc.X = ev.X()
	dc.Y = ev.Y()
	return event.ResponseNone
}

func (dc *debugCoords) onMouseDrag(ev *mouse.Event) event.Response {
	dc.X = int(ev.Point2.X())
	dc.Y = int(ev.Point2.Y())
	return event.ResponseNone
}

func (dc debugCoords) String() string {
	return fmt.Sprintf("%s%d, %d", dc.Label, dc.X, dc.Y)
}

type debugFPS struct {
	fps       int
	lastTime  time.Time
	Smoothing float64
}

func newFPS(ctx *oakscene.Context) (ret *debugFPS) {
	ret = &debugFPS{
		Smoothing: .25,
	}
	event.GlobalBind(ctx, event.Enter, ret.update)
	return
}

func (dfps *debugFPS) update(_ event.EnterPayload) event.Response {
	t := time.Now()
	dfps.fps = int((timing.FPS(dfps.lastTime, t) * dfps.Smoothing) + (float64(dfps.fps) * (1 - dfps.Smoothing)))
	dfps.lastTime = t
	return event.ResponseNone
}

func (dfps *debugFPS) String() string {
	return fmt.Sprintf("FPS: %d", dfps.fps)
}

func DebugInfo(ctx *oakscene.Context) {

	debugLayer := 1000
	debugLayers := []int{debugLayer, debugLayer + 1}
	debugMaxWidth := 200.
	debugMargin := 3.

	layout := utils.NewVTextLayout(nil, debugMaxWidth, debugMargin).Add(
		newFPS(ctx),
		newMouseCoords(ctx),
		newViewportPosition(ctx),
	)

	panel := entities.New(ctx,
		entities.WithDimensions(layout.GetFDims()),
		// entities.WithRenderable(
		// 	render.NewColorBox(
		// 		int(layout.W()),
		// 		int(layout.H()),
		// 		color.RGBA{R: 60, G: 60, B: 60, A: 255},
		// 	),
		// ),
		entities.WithChild(
			entities.WithRenderable(layout.Renderable()),
			entities.WithDimensions(layout.GetFDims()),
			entities.WithDrawLayers(debugLayers),
			entities.WithChild(),
		),
	)

	event.Bind(ctx, oak.ViewportUpdate, panel, func(e *entities.Entity, ev intgeom.Point2) event.Response {
		// Move the panel on viewport update
		panel.SetPos(
			floatgeom.Point2{
				float64(ev.X()),
				float64(ev.Y()),
			},
		)
		return event.ResponseNone
	})
}
