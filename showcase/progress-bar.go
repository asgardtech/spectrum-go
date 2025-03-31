package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type progressBarPage struct {
	app.Compo
}

func newProgressBarPage() *progressBarPage {
	return &progressBarPage{}
}

func (p *progressBarPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *progressBarPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle(defaultTitle)
	ctx.Page().SetDescription(defaultDescription)
}

func (p *progressBarPage) Render() app.UI {
	return newPage().Content(
		app.H1().Text("Progress Bar Component"),
		app.P().Text("A progress bar shows the progression of a system operation such as downloading, uploading, processing, etc. in a visual way."),

		// Basic Progress Bar
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Basic Progress Bar"),
				app.P().Text("Showing determinate progress with different values."),
				app.Div().Class("component-row").Body(
					sp.ProgressBar().
						Label("0% Complete").
						Progress(0),
					sp.ProgressBar().
						Label("25% Complete").
						Progress(25),
					sp.ProgressBar().
						Label("50% Complete").
						Progress(50),
					sp.ProgressBar().
						Label("75% Complete").
						Progress(75),
					sp.ProgressBar().
						Label("100% Complete").
						Progress(100),
				),
			),

		// Indeterminate Progress Bar
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Indeterminate Progress Bar"),
				app.P().Text("Used when the duration of the operation is unknown."),
				app.Div().Class("component-row").Body(
					sp.ProgressBar().
						Label("Loading...").
						Indeterminate(true),
				),
			),

		// Side Label
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Side Label"),
				app.P().Text("Display the label beside the progress bar instead of above it."),
				app.Div().Class("component-row").Body(
					sp.ProgressBar().
						Label("Default Label Position").
						Progress(50),
					sp.ProgressBar().
						Label("Side Label").
						Progress(50).
						SideLabel(true),
				),
			),

		// Size Variants
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Size Variants"),
				app.P().Text("Progress bars in different sizes."),
				app.Div().Class("component-row").Body(
					sp.ProgressBar().
						Label("Small (S)").
						Progress(50).
						Size("s"),
					sp.ProgressBar().
						Label("Medium (M) - Default").
						Progress(50).
						Size("m"),
					sp.ProgressBar().
						Label("Large (L)").
						Progress(50).
						Size("l"),
					sp.ProgressBar().
						Label("Extra Large (XL)").
						Progress(50).
						Size("xl"),
				),
			),

		// Over Background
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Over Background"),
				app.P().Text("Progress bar designed to be displayed on a colored background."),
				app.Div().Class("color-background-demo dark").Body(
					sp.ProgressBar().
						Label("Standard on Dark").
						Progress(50),
					sp.ProgressBar().
						Label("Over Background").
						Progress(50).
						OverBackground(true),
				),
			),

		// Static Color
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Static Color"),
				app.P().Text("Progress bar with static white color for dark backgrounds."),
				app.Div().Class("color-background-demo dark").Body(
					sp.ProgressBar().
						Label("White Progress Bar").
						Progress(50).
						StaticColor("white"),
				),
			),

		// Multiple Progress Bars
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Usage Example"),
				app.P().Text("Multiple progress bars in a real-world scenario."),
				app.Div().Class("file-upload-demo").Body(
					app.Div().Class("upload-item").Body(
						app.Div().Class("upload-info").Body(
							app.Div().Class("upload-name").Text("document.pdf"),
							app.Div().Class("upload-size").Text("2.5MB / 10MB"),
						),
						sp.ProgressBar().
							Label("Uploading document.pdf").
							Progress(25).
							SideLabel(true),
					),
					app.Div().Class("upload-item").Body(
						app.Div().Class("upload-info").Body(
							app.Div().Class("upload-name").Text("image.jpg"),
							app.Div().Class("upload-size").Text("1.8MB / 1.8MB"),
						),
						sp.ProgressBar().
							Label("Uploading image.jpg").
							Progress(100).
							SideLabel(true),
					),
					app.Div().Class("upload-item").Body(
						app.Div().Class("upload-info").Body(
							app.Div().Class("upload-name").Text("video.mp4"),
							app.Div().Class("upload-size").Text("Processing..."),
						),
						sp.ProgressBar().
							Label("Processing video.mp4").
							Indeterminate(true).
							SideLabel(true),
					),
				),
			),
	)
}
