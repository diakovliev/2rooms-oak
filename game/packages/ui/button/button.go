package button

import (
	"github.com/diakovliev/oak/v4/entities"
	"github.com/diakovliev/oak/v4/event"
	"github.com/diakovliev/oak/v4/mouse"
	"github.com/diakovliev/oak/v4/render"
	"github.com/diakovliev/oak/v4/render/mod"
	oakscene "github.com/diakovliev/oak/v4/scene"
)

type Button struct {
	*entities.Entity
}

func New(ctx *oakscene.Context, text string, opts ...Option) (ret *Button) {

	layer := 1000
	drawLayers := []int{layer + 0, layer + 1, layer + 2, layer + 3}

	button := newData(text)
	for _, opt := range opts {
		opt(button)
	}

	ret = &Button{
		Entity: entities.New(
			ctx,
			entities.WithData(button),
			entities.WithUseMouseTree(true),
			entities.WithDimensions(button.dimensions()),
			entities.WithDrawLayers(drawLayers),
			entities.WithRenderable(button.getBackRenderable()),
			entities.WithMod(mod.CutRound(button.roundX, button.roundY)),
			entities.WithChild(
				entities.WithDimensions(button.dimensions()),
				entities.WithRenderable(button.getSwFocus()),
				entities.WithDrawLayers(drawLayers[1:]),
				entities.WithMod(mod.CutRound(button.roundX, button.roundY)),
			),
			entities.WithChild(
				entities.WithPosition(button.focusMarginPoint()),
				entities.WithDimensions(button.innerDimensions()),
				entities.WithRenderable(button.getSwBack()),
				entities.WithDrawLayers(drawLayers[2:]),
				entities.WithMod(mod.CutRound(button.roundX, button.roundY)),
			),
			entities.WithChild(
				entities.WithPosition(button.focusMarginPoint()),
				entities.WithDimensions(button.innerDimensions()),
				entities.WithRenderable(button.getSwLabel()),
				entities.WithDrawLayers(drawLayers[3:]),
				entities.WithMod(mod.CutRound(button.roundX, button.roundY)),
			),
		),
	}

	event.Bind(ctx, mouse.DragOn, ret.Entity, ret.mouseDragOn)
	event.Bind(ctx, mouse.PressOn, ret.Entity, ret.mousePressOn)
	event.Bind(ctx, mouse.ReleaseOn, ret.Entity, ret.mouseReleaseOn)
	event.Bind(ctx, mouse.Drag, ret.Entity, ret.mouseDrag)

	return
}

func (b *Button) mouseDragOn(e *entities.Entity, me *mouse.Event) (ret event.Response) {
	me.StopPropagation = true
	return
}

func (b *Button) mousePressOn(e *entities.Entity, me *mouse.Event) (ret event.Response) {
	me.StopPropagation = true
	d := entities.MustData[*data](b.Entity)
	if d.isState(Disabled) {
		return
	}
	b.SetState(Down)
	return
}

func (b *Button) mouseReleaseOn(e *entities.Entity, me *mouse.Event) (ret event.Response) {
	me.StopPropagation = true
	d := entities.MustData[*data](b.Entity)
	if d.isState(Disabled) {
		return
	}
	d.callCallback()
	b.SetState(Up)
	return
}

func (b *Button) mouseDrag(e *entities.Entity, me *mouse.Event) (ret event.Response) {
	me.StopPropagation = true
	d := entities.MustData[*data](b.Entity)
	if d.isState(Disabled) {
		return
	}
	if !e.Rect.Contains(me.Point2) {
		b.SetState(Up)
		d.setFocus(NonFocused)
	} else {
		d.setFocus(Focused)
	}
	return
}

func (b *Button) SetState(state State) {
	entities.MustData[*data](b.Entity).setState(state)
}

func (b *Button) Font() *render.Font {
	d := entities.MustData[*data](b.Entity)
	return d.fonts[d.state]
}
