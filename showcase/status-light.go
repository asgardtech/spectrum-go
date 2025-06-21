package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/asgardtech/spectrum-go/prism"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type statusLightPage struct {
	app.Compo
}

func newStatusLightPage() *statusLightPage {
	return &statusLightPage{}
}

func (p *statusLightPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *statusLightPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle(defaultTitle)
	ctx.Page().SetDescription(defaultDescription)
}

func (p *statusLightPage) Render() app.UI {
	return prism.NewLayout().Content(
		app.H1().Text("Status Light Component"),
		app.P().Text("Status Light is a great way to convey semantic meaning, such as statuses and categories."),

		// Different Variants
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Status Light Variants"),
				app.P().Text("There are many variants to choose from in Spectrum:"),
				app.Div().Class("component-row").Body(
					sp.StatusLight().
						Variant("positive").
						Body(
							app.Text("Positive"),
						),
					sp.StatusLight().
						Variant("negative").
						Body(
							app.Text("Negative"),
						),
					sp.StatusLight().
						Variant("notice").
						Body(
							app.Text("Notice"),
						),
					sp.StatusLight().
						Variant("info").
						Body(
							app.Text("Info"),
						),
				),
				app.Div().Class("component-row").Body(
					sp.StatusLight().
						Variant("neutral").
						Body(
							app.Text("Neutral"),
						),
					sp.StatusLight().
						Variant("yellow").
						Body(
							app.Text("Yellow"),
						),
					sp.StatusLight().
						Variant("fuchsia").
						Body(
							app.Text("Fuchsia"),
						),
					sp.StatusLight().
						Variant("indigo").
						Body(
							app.Text("Indigo"),
						),
				),
				app.Div().Class("component-row").Body(
					sp.StatusLight().
						Variant("seafoam").
						Body(
							app.Text("Seafoam"),
						),
					sp.StatusLight().
						Variant("chartreuse").
						Body(
							app.Text("Chartreuse"),
						),
					sp.StatusLight().
						Variant("magenta").
						Body(
							app.Text("Magenta"),
						),
					sp.StatusLight().
						Variant("celery").
						Body(
							app.Text("Celery"),
						),
				),
				app.Div().Class("component-row").Body(
					sp.StatusLight().
						Variant("purple").
						Body(
							app.Text("Purple"),
						),
				),
			),

		// Different Sizes
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Status Light Sizes"),
				app.Div().Class("component-row").Body(
					sp.StatusLight().
						Variant("positive").
						Size("s").
						Body(
							app.Text("Small (S)"),
						),
					sp.StatusLight().
						Variant("positive").
						Size("m").
						Body(
							app.Text("Medium (M) - Default"),
						),
					sp.StatusLight().
						Variant("positive").
						Size("l").
						Body(
							app.Text("Large (L)"),
						),
					sp.StatusLight().
						Variant("positive").
						Size("xl").
						Body(
							app.Text("Extra Large (XL)"),
						),
				),
			),

		// Disabled state
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Disabled Status Light"),
				app.P().Text("A status light in a disabled state shows that a status exists, but is not available in that circumstance."),
				app.Div().Class("component-row").Body(
					sp.StatusLight().
						Variant("positive").
						Disabled(true).
						Body(
							app.Text("Disabled Positive"),
						),
					sp.StatusLight().
						Variant("negative").
						Disabled(true).
						Body(
							app.Text("Disabled Negative"),
						),
					sp.StatusLight().
						Variant("info").
						Disabled(true).
						Body(
							app.Text("Disabled Info"),
						),
				),
			),
	)
}
