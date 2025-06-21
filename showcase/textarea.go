package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/asgardtech/spectrum-go/prism"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type textareaPage struct {
	app.Compo
}

func newTextareaPage() *textareaPage {
	return &textareaPage{}
}

func (p *textareaPage) Render() app.UI {
	return prism.NewLayout().Content(
		app.H1().Text("Textarea Component"),
		app.P().Text("Textareas allow users to input, edit, and select multiple lines of text."),

		// Basic textarea
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Basic Textarea"),
				app.Div().Class("component-row").Body(
					sp.Textarea().
						Label("Default textarea").
						Placeholder("Enter text here"),
				),
			),

		// Textarea with help text
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Textarea with Help Text"),
				app.Div().Class("component-row").Body(
					sp.Textarea().
						Label("Description").
						Placeholder("Enter description").
						HelpText(app.Text("Provide a detailed description.")),
				),
			),

		// Required textarea
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Required Textarea"),
				app.Div().Class("component-row").Body(
					sp.Textarea().
						Label("Feedback").
						Placeholder("Enter your feedback").
						Required(true),
				),
			),

		// Disabled state
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Disabled State"),
				app.Div().Class("component-row").Body(
					sp.Textarea().
						Label("Disabled textarea").
						Value("This textarea is disabled").
						Disabled(true),
				),
			),

		// Invalid state
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Invalid State"),
				app.Div().Class("component-row").Body(
					sp.Textarea().
						Label("Comments").
						Placeholder("Enter comments").
						Invalid(true).
						NegativeHelpText(app.Text("Comments must be appropriate and relevant.")),
				),
			),

		// Quiet variant
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Quiet Variant"),
				app.Div().Class("component-row").Body(
					sp.Textarea().
						Label("Standard textarea").
						Placeholder("Standard textarea"),
					sp.Textarea().
						Label("Quiet textarea").
						Placeholder("Quiet textarea").
						Quiet(true),
				),
			),

		// Different sizes
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Textarea Sizes"),
				app.Div().Class("component-row").Body(
					sp.Textarea().
						Label("Small (S)").
						Placeholder("Small textarea").
						Size("s"),
					sp.Textarea().
						Label("Medium (M) - Default").
						Placeholder("Medium textarea").
						Size("m"),
					sp.Textarea().
						Label("Large (L)").
						Placeholder("Large textarea").
						Size("l"),
					sp.Textarea().
						Label("Extra Large (XL)").
						Placeholder("Extra large textarea").
						Size("xl"),
				),
			),

		// With rows
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Textarea with Specified Rows"),
				app.Div().Class("component-row").Body(
					sp.Textarea().
						Label("3 Rows").
						Placeholder("3 rows textarea").
						Rows(3),
					sp.Textarea().
						Label("5 Rows").
						Placeholder("5 rows textarea").
						Rows(5),
				),
			),

		// Grow with content
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Grow with Content"),
				app.Div().Class("component-row").Body(
					sp.Textarea().
						Label("Auto-growing textarea").
						Placeholder("Type multiple lines to see this grow...").
						SetGrows(),
				),
			),
	)
}
