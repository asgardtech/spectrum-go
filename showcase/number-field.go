package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type numberFieldPage struct {
	app.Compo
}

func newNumberFieldPage() *numberFieldPage {
	return &numberFieldPage{}
}

func (p *numberFieldPage) Render() app.UI {
	return newPage().Content(
		app.H1().Text("Number Field Component"),
		app.P().Text("Number Fields allow users to input, edit, and manipulate numbers with increment and decrement capabilities."),

		// Basic number field
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Basic Number Field"),
				app.Div().Class("component-row").Body(
					sp.NumberField().
						Label("Size").
						Value("1024").
						Min(0).
						Max(2048).
						Step(1),
				),
			),

		// Number field with different steps
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Number Field with Custom Steps"),
				app.Div().Class("component-row").Body(
					sp.NumberField().
						Label("Opacity").
						Value("0.5").
						Min(0).
						Max(1).
						Step(0.1),
				),
			),

		// Number field with hidden stepper
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Number Field with Hidden Stepper"),
				app.Div().Class("component-row").Body(
					sp.NumberField().
						Label("Percentage").
						Value("50").
						Min(0).
						Max(100).
						Step(1).
						Hidestepper(true),
				),
			),

		// Number field states: disabled
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Disabled Number Field"),
				app.Div().Class("component-row").Body(
					sp.NumberField().
						Label("Width").
						Value("300").
						Min(0).
						Disabled(true),
				),
			),

		// Number field states: invalid
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Invalid Number Field"),
				app.Div().Class("component-row").Body(
					sp.NumberField().
						Label("Temperature").
						Value("-10").
						Min(0).
						Max(100).
						Invalid(true),
					app.Div().
						Attr("slot", "negative-help-text").
						Text("Value must be between 0 and 100"),
				),
			),

		// Quiet number field
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Quiet Number Field"),
				app.Div().Class("component-row").Body(
					sp.NumberField().
						Label("Standard field").
						Value("100").
						Min(0).
						Max(1000),
					sp.NumberField().
						Label("Quiet field").
						Value("100").
						Min(0).
						Max(1000).
						Quiet(true),
				),
			),

		// Required number field
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Required Number Field"),
				app.Div().Class("component-row").Body(
					sp.NumberField().
						Label("Quantity").
						Placeholder("Enter quantity").
						Min(1).
						Required(true),
				),
			),

		// Number field with help text
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Number Field with Help Text"),
				app.Div().Class("component-row").Body(
					sp.NumberField().
						Label("Age").
						Value("30").
						Min(0).
						Max(120),
					app.Div().
						Attr("slot", "help-text").
						Text("Please enter your current age"),
				),
			),

		// Readonly number field
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Readonly Number Field"),
				app.Div().Class("component-row").Body(
					sp.NumberField().
						Label("Score").
						Value("85").
						Min(0).
						Max(100).
						Readonly(true),
				),
			),
	)
}