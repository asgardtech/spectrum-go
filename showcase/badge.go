package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type badgePage struct {
	app.Compo
}

func newBadgePage() *badgePage {
	return &badgePage{}
}

func (p *badgePage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *badgePage) initPage(ctx app.Context) {
	ctx.Page().SetTitle(defaultTitle)
	ctx.Page().SetDescription(defaultDescription)
}

func (p *badgePage) Render() app.UI {
	return newPage().Content(
		app.H1().Text("Badge Component"),
		app.P().Text("Badges display a small amount of color-categorized metadata. They're ideal for getting a user's attention."),

		// Basic badge
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Basic Badge"),
				app.Div().Class("component-row").Body(
					sp.Badge().Text("Default Badge"),
				),
			),

		// Semantic variants
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Semantic Variants"),
				app.Div().Class("component-row").Body(
					sp.Badge().Text("Accent").Variant("accent"),
					sp.Badge().Text("Informative").Variant("informative"),
					sp.Badge().Text("Neutral").Variant("neutral"),
					sp.Badge().Text("Positive").Variant("positive"),
					sp.Badge().Text("Negative").Variant("negative"),
					sp.Badge().Text("Notice").Variant("notice"),
				),
			),

		// Color variants
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Color Variants"),
				app.Div().Class("component-row").Body(
					sp.Badge().Text("Seafoam").Variant("seafoam"),
					sp.Badge().Text("Indigo").Variant("indigo"),
					sp.Badge().Text("Purple").Variant("purple"),
					sp.Badge().Text("Fuchsia").Variant("fuchsia"),
					sp.Badge().Text("Magenta").Variant("magenta"),
					sp.Badge().Text("Yellow").Variant("yellow"),
				),
				app.Div().Class("component-row").Body(
					sp.Badge().Text("Gray").Variant("gray"),
					sp.Badge().Text("Red").Variant("red"),
					sp.Badge().Text("Orange").Variant("orange"),
					sp.Badge().Text("Chartreuse").Variant("chartreuse"),
					sp.Badge().Text("Celery").Variant("celery"),
					sp.Badge().Text("Green").Variant("green"),
					sp.Badge().Text("Cyan").Variant("cyan"),
					sp.Badge().Text("Blue").Variant("blue"),
				),
			),

		// Badge with icon
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Badge with Icon"),
				app.Div().Class("component-row").Body(
					sp.Badge().
						Text("Badge with Icon").
						Variant("positive").
						Icon(sp.Icon().Name("ui:CheckmarkCircle100")),
					sp.Badge().
						Text("Error Badge").
						Variant("negative").
						Icon(sp.Icon().Name("ui:AlertCircle100")),
				),
			),

		// Size variants
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Size Variants"),
				app.Div().Class("component-row").Body(
					sp.Badge().Text("Small").Size("s"),
					sp.Badge().Text("Medium (Default)").Size("m"),
					sp.Badge().Text("Large").Size("l"),
					sp.Badge().Text("Extra Large").Size("xl"),
				),
			),

		// Fixed position badges
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Fixed Position"),
				app.P().Text("Badges can be positioned with different edge styles to appear 'attached' to other elements."),
				app.Div().Class("badge-position-demo").Body(
					app.Div().Class("badge-container").Body(
						app.Div().Class("badge-demo-box").Text("Content"),
						sp.Badge().
							Text("Top").
							Fixed("block-start").
							Style("position", "absolute").
							Style("top", "0").
							Style("left", "50%").
							Style("transform", "translateX(-50%)"),
					),
					app.Div().Class("badge-container").Body(
						app.Div().Class("badge-demo-box").Text("Content"),
						sp.Badge().
							Text("Right").
							Fixed("inline-end").
							Style("position", "absolute").
							Style("right", "0").
							Style("top", "50%").
							Style("transform", "translateY(-50%)"),
					),
					app.Div().Class("badge-container").Body(
						app.Div().Class("badge-demo-box").Text("Content"),
						sp.Badge().
							Text("Bottom").
							Fixed("block-end").
							Style("position", "absolute").
							Style("bottom", "0").
							Style("left", "50%").
							Style("transform", "translateX(-50%)"),
					),
					app.Div().Class("badge-container").Body(
						app.Div().Class("badge-demo-box").Text("Content"),
						sp.Badge().
							Text("Left").
							Fixed("inline-start").
							Style("position", "absolute").
							Style("left", "0").
							Style("top", "50%").
							Style("transform", "translateY(-50%)"),
					),
				),
			),
	)
}
