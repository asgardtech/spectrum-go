package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/asgardtech/spectrum-go/prism"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type progressCirclePage struct {
	app.Compo
}

func newProgressCirclePage() *progressCirclePage {
	return &progressCirclePage{}
}

func (p *progressCirclePage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *progressCirclePage) initPage(ctx app.Context) {
	ctx.Page().SetTitle(defaultTitle)
	ctx.Page().SetDescription(defaultDescription)
}

func (p *progressCirclePage) Render() app.UI {
	return prism.NewLayout().Content(
		app.H1().Text("Progress Circle Component"),
		app.P().Text("A progress circle shows the progression of a system operation in a circular visual representation."),

		// Basic Progress Circle
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Basic Progress Circle"),
				app.P().Text("Showing determinate progress with different values."),
				app.Div().Class("component-row").Body(
					sp.ProgressCircle().
						Label("0% Complete").
						Progress(0),
					sp.ProgressCircle().
						Label("25% Complete").
						Progress(25),
					sp.ProgressCircle().
						Label("50% Complete").
						Progress(50),
					sp.ProgressCircle().
						Label("75% Complete").
						Progress(75),
					sp.ProgressCircle().
						Label("100% Complete").
						Progress(100),
				),
			),

		// Indeterminate Progress Circle
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Indeterminate Progress Circle"),
				app.P().Text("Used when the duration of the operation is unknown."),
				app.Div().Class("component-row").Body(
					sp.ProgressCircle().
						Label("Loading...").
						Indeterminate(true),
				),
			),

		// Size Variants
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Size Variants"),
				app.P().Text("Progress circles in different sizes."),
				app.Div().Class("component-row").Body(
					sp.ProgressCircle().
						Label("Small (S)").
						Progress(50).
						Size("s"),
					sp.ProgressCircle().
						Label("Medium (M) - Default").
						Progress(50).
						Size("m"),
					sp.ProgressCircle().
						Label("Large (L)").
						Progress(50).
						Size("l"),
				),
			),

		// Static Color on Dark Background
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Static Color"),
				app.P().Text("Progress circle with static white color for dark backgrounds."),
				app.Div().Class("color-background-demo dark").Body(
					sp.ProgressCircle().
						Label("White Progress Circle").
						Progress(50).
						StaticColor("white"),
				),
			),

		// Usage Examples
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Usage Examples"),
				app.P().Text("Common patterns for using progress circles."),

				// Example 1: Button with Loading State
				app.H3().Text("Button with Loading State"),
				app.Div().Class("component-row").Body(
					app.Div().Class("loading-button").Body(
						sp.Button().
							Disabled(true).
							Text("Loading..."),
						sp.ProgressCircle().
							Label("Loading").
							Size("s").
							Indeterminate(true),
					),
				),

				// Example 2: Card Loading State
				app.H3().Text("Card Loading State"),
				app.Div().Class("component-row").Body(
					app.Div().Class("loading-card").Body(
						app.Div().Class("card-title").Text("Loading Content"),
						app.Div().Class("card-progress").Body(
							sp.ProgressCircle().
								Label("Loading card content").
								Progress(75),
						),
						app.Div().Class("card-text").Text("75% Complete"),
					),
				),

				// Example 3: Inline with Text
				app.H3().Text("Inline with Text"),
				app.Div().Class("component-row").Body(
					app.Div().Class("inline-progress").Body(
						app.Span().Text("Fetching results"),
						sp.ProgressCircle().
							Label("Fetching").
							Size("s").
							Indeterminate(true),
					),
				),
			),

		// Multiple Progress Circles
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Multiple Operations"),
				app.P().Text("Tracking multiple operations with progress circles."),
				app.Div().Class("dashboard-demo").Body(
					app.Div().Class("dashboard-item").Body(
						sp.ProgressCircle().
							Label("Upload Progress").
							Progress(30),
						app.Div().Class("operation-label").Text("Uploading Files"),
						app.Div().Class("operation-details").Text("30% Complete"),
					),
					app.Div().Class("dashboard-item").Body(
						sp.ProgressCircle().
							Label("Download Progress").
							Progress(60),
						app.Div().Class("operation-label").Text("Downloading Updates"),
						app.Div().Class("operation-details").Text("60% Complete"),
					),
					app.Div().Class("dashboard-item").Body(
						sp.ProgressCircle().
							Label("Processing").
							Indeterminate(true),
						app.Div().Class("operation-label").Text("Processing Data"),
						app.Div().Class("operation-details").Text("Please wait..."),
					),
				),
			),
	)
}
