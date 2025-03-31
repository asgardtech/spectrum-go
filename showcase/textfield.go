package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type textfieldPage struct {
	app.Compo
}

func newTextfieldPage() *textfieldPage {
	return &textfieldPage{}
}

func (p *textfieldPage) Render() app.UI {
	return newPage().Content(
		app.H1().Text("Textfield Component"),
		app.P().Text("Textfields allow users to input, edit, and select text."),

		// Basic textfield
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Basic Textfield"),
				app.Div().Class("component-row").Body(
					sp.Textfield().
						Label("Default textfield").
						Placeholder("Enter text here"),
				),
			),

		// Textfield with help text
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Textfield with Help Text"),
				app.Div().Class("component-row").Body(
					sp.Textfield().
						Label("Email").
						Placeholder("Enter your email").
						HelpText(app.Text("We'll never share your email.")),
				),
			),

		// Required textfield
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Required Textfield"),
				app.Div().Class("component-row").Body(
					sp.Textfield().
						Label("Username").
						Placeholder("Enter username").
						Required(true),
				),
			),

		// Disabled state
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Disabled State"),
				app.Div().Class("component-row").Body(
					sp.Textfield().
						Label("Disabled field").
						Value("This field is disabled").
						Disabled(true),
				),
			),

		// Invalid state
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Invalid State"),
				app.Div().Class("component-row").Body(
					sp.Textfield().
						Label("Password").
						Placeholder("Enter password").
						Type("password").
						Invalid(true).
						NegativeHelpText(app.Text("Password must be at least 8 characters.")),
				),
			),

		// Quiet variant
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Quiet Variant"),
				app.Div().Class("component-row").Body(
					sp.Textfield().
						Label("Standard field").
						Placeholder("Standard field"),
					sp.Textfield().
						Label("Quiet field").
						Placeholder("Quiet field").
						Quiet(true),
				),
			),

		// Different sizes
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Textfield Sizes"),
				app.Div().Class("component-row").Body(
					sp.Textfield().
						Label("Small (S)").
						Placeholder("Small field").
						Size("s"),
					sp.Textfield().
						Label("Medium (M) - Default").
						Placeholder("Medium field").
						Size("m"),
					sp.Textfield().
						Label("Large (L)").
						Placeholder("Large field").
						Size("l"),
					sp.Textfield().
						Label("Extra Large (XL)").
						Placeholder("Extra large field").
						Size("xl"),
				),
			),

		// With icons - using slots
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Textfield with Icons"),
				app.Div().Class("component-row").Body(
					app.Div().Body(
						sp.Textfield().
							Label("Search").
							Placeholder("Search..."),
						app.Elem("sp-icon").
							Attr("slot", "icon").
							Attr("name", "ui:Magnifier"),
					),
				),
			),
	)
}