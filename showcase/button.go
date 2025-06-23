package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/asgardtech/spectrum-go/prism"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type buttonPage struct {
	app.Compo
}

func newButtonPage() *buttonPage {
	return &buttonPage{}
}

func (p *buttonPage) Render() app.UI {
	return prism.NewPage().Content(
		app.H1().Text("Button Component"),
		app.P().Text("Buttons represent actions a user can take."),

		// Basic buttons - different variants
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Button Variants"),
				app.Div().Class("component-row").Body(
					sp.Button().Text("Accent (Default)").Variant(sp.ButtonVariantAccent),
					sp.Button().Text("Primary").Variant(sp.ButtonVariantPrimary),
					sp.Button().Text("Secondary").Variant(sp.ButtonVariantSecondary),
					sp.Button().Text("Negative").Variant(sp.ButtonVariantNegative),
				),
			),

		// Treatment options
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Button Treatments"),
				app.Div().Class("component-row").Body(
					sp.Button().Text("Fill (Default)").Treatment(sp.ButtonTreatmentFill),
					sp.Button().Text("Outline").Treatment(sp.ButtonTreatmentOutline),
				),
			),

		// Quiet mode
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Quiet Buttons"),
				app.Div().Class("component-row").Body(
					sp.Button().Text("Standard").Variant(sp.ButtonVariantPrimary),
					sp.Button().Text("Quiet").Variant(sp.ButtonVariantPrimary).Quiet(true),
				),
			),

		// Disabled state
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Disabled Buttons"),
				app.Div().Class("component-row").Body(
					sp.Button().Text("Enabled").Variant(sp.ButtonVariantPrimary),
					sp.Button().Text("Disabled").Variant(sp.ButtonVariantPrimary).Disabled(true),
				),
			),

		// Buttons with icons
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Buttons with Icons"),
				app.Div().Class("component-row").Body(
					sp.Button().Text("With Icon").Variant(sp.ButtonVariantPrimary).
						Icon(sp.Icon().Name("ui:ChevronDown100")),
					sp.Button().Label("Icon Only").Variant(sp.ButtonVariantPrimary).
						Icon(sp.Icon().Name("ui:Help100")),
				),
			),

		// Active state
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Active State"),
				app.Div().Class("component-row").Body(
					sp.Button().Text("Normal").Variant(sp.ButtonVariantPrimary),
					sp.Button().Text("Active").Variant(sp.ButtonVariantPrimary).Active(true),
				),
			),

		// Link buttons
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Link Buttons"),
				app.Div().Class("component-row").Body(
					sp.Button().Text("Internal Link").Variant(sp.ButtonVariantPrimary).Href("#"),
					sp.Button().Text("External Link").Variant(sp.ButtonVariantPrimary).
						Href("https://github.com/asgardtech/spectrum-go").
						Target("_blank"),
				),
			),

		// Pending state example (static)
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Pending State"),
				app.Div().Class("component-row").Body(
					sp.Button().Text("Normal").Variant(sp.ButtonVariantPrimary),
					sp.Button().Text("Loading").Variant(sp.ButtonVariantPrimary).
						Pending(true).PendingLabel("Loading..."),
				),
			),

		// Static color
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Static Colors"),
				app.Div().Class("color-background-demo dark").Body(
					sp.Button().Text("White on Dark").StaticColor(sp.ButtonStaticColorWhite),
				),
				app.Div().Class("color-background-demo light").Body(
					sp.Button().Text("Black on Light").StaticColor(sp.ButtonStaticColorBlack),
				),
			),
	)
}
