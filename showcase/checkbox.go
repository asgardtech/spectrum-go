package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type checkboxPage struct {
	app.Compo
}

func newCheckboxPage() *checkboxPage {
	return &checkboxPage{}
}

func (p *checkboxPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *checkboxPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle(defaultTitle)
	ctx.Page().SetDescription(defaultDescription)
}

func (p *checkboxPage) Render() app.UI {
	return newPage().Content(
		app.H1().Text("Checkbox Component"),
		app.P().Text("Checkboxes allow users to select multiple items from a list of independent options, or to mark an individual option as selected."),

		// Basic checkbox
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Basic Checkbox"),
				app.Div().Class("component-row").Body(
					sp.Checkbox().
						Body(
							app.Text("Unchecked checkbox"),
						),
					sp.Checkbox().
						Body(
							app.Text("Checked checkbox"),
						).Checked(true),
				),
			),

		// Indeterminate state
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Indeterminate State"),
				app.Div().Class("component-row").Body(
					sp.Checkbox().
						Body(
							app.Text("Indeterminate checkbox"),
						).Indeterminate(true),
				),
			),

		// Emphasized style
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Emphasized Style"),
				app.Div().Class("component-row").Body(
					sp.Checkbox().
						Body(
							app.Text("Standard checkbox"),
						).Checked(true),
					sp.Checkbox().
						Body(
							app.Text("Emphasized checkbox"),
						).Emphasized(true).Checked(true),
				),
			),

		// Invalid state
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Invalid State"),
				app.Div().Class("component-row").Body(
					sp.Checkbox().
						Body(
							app.Text("Invalid checkbox"),
						).Invalid(true),
				),
			),

		// Disabled state
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Disabled State"),
				app.Div().Class("component-row").Body(
					sp.Checkbox().
						Body(
							app.Text("Disabled unchecked"),
						).Disabled(true),
					sp.Checkbox().
						Body(
							app.Text("Disabled checked"),
						).Disabled(true).Checked(true),
				),
			),

		// Readonly state
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Readonly State"),
				app.Div().Class("component-row").Body(
					sp.Checkbox().
						Body(
							app.Text("Readonly unchecked"),
						).Readonly(true),
					sp.Checkbox().
						Body(
							app.Text("Readonly checked"),
						).Readonly(true).Checked(true),
				),
			),

		// Sizes
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Checkbox Sizes"),
				app.Div().Class("component-row").Body(
					sp.Checkbox().
						Body(
							app.Text("Small (S)"),
						).Size("s"),
					sp.Checkbox().
						Body(
							app.Text("Medium (M) - Default"),
						).Size("m"),
					sp.Checkbox().
						Body(
							app.Text("Large (L)"),
						).Size("l"),
					sp.Checkbox().
						Body(
							app.Text("Extra Large (XL)"),
						).Size("xl"),
				),
			),
	)
}
