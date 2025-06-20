package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type colorPage struct {
	app.Compo
	ctx   app.Context
	color string
}

func newColorPage() *colorPage {
	return &colorPage{
		color: "#ff0000",
	}
}

func (p *colorPage) OnMount(ctx app.Context) {
	p.ctx = ctx
}

func (p *colorPage) onColorChange(ctx app.Context, e app.Event) {
	colorComponent := ctx.JSSrc()
	p.color = colorComponent.Get("color").String()
	p.ctx.Update()
}

func (p *colorPage) Render() app.UI {
	return newPage().
		Content(
			app.H1().Text("Color Components"),
			app.P().Text("Components for color selection."),
			app.H2().Text("Color Area"),
			sp.ColorArea().
				Style("width", "200px").
				Style("height", "200px").
				Color(p.color).
				OnInput(p.onColorChange),
			app.H2().Text("Color Slider"),
			sp.ColorSlider().
				Style("width", "200px").
				Color(p.color).
				OnInput(p.onColorChange),
			app.H2().Text("Color Wheel"),
			sp.ColorWheel().
				Style("width", "200px").
				Color(p.color).
				OnInput(p.onColorChange),
			app.If(p.color != "",
				func() app.UI {
					return app.Div().Body(
						app.Text("Current color: "),
						app.Strong().Text(p.color),
						app.Div().Style("width", "50px").
							Style("height", "50px").
							Style("background-color", p.color),
					)
				},
			),
		)
}
