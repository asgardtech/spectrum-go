package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/asgardtech/spectrum-go/prism"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type fieldLabelPage struct {
	app.Compo
}

func newFieldLabelPage() *fieldLabelPage {
	return &fieldLabelPage{}
}

func (p *fieldLabelPage) Render() app.UI {
	return prism.NewLayout().Content(
		app.H1().Text("Field Label Component"),
		app.P().Text("Field labels provide accessible labeling for form elements."),

		// Basic field label
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Basic Field Label"),
				app.Div().Class("component-row").Body(
					app.Div().Style("display", "flex").Style("flex-direction", "column").Style("gap", "8px").Body(
						sp.FieldLabel().
							For("textfield1").
							Text("First Name"),
						sp.Textfield().
							Id("textfield1").
							Placeholder("Enter your first name"),
					),
				),
			),

		// Required field label
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Required Field Label"),
				app.Div().Class("component-row").Body(
					app.Div().Style("display", "flex").Style("flex-direction", "column").Style("gap", "8px").Body(
						sp.FieldLabel().
							For("textfield2").
							Text("Email Address").
							Required(true),
						sp.Textfield().
							Id("textfield2").
							Placeholder("Enter your email").
							Required(true),
					),
				),
			),

		// Disabled field label
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Disabled Field Label"),
				app.Div().Class("component-row").Body(
					app.Div().Style("display", "flex").Style("flex-direction", "column").Style("gap", "8px").Body(
						sp.FieldLabel().
							For("textfield3").
							Text("Username").
							Disabled(true),
						sp.Textfield().
							Id("textfield3").
							Value("johndoe").
							Disabled(true),
					),
				),
			),

		// Field label sizes
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Field Label Sizes"),
				app.Div().Class("component-row").Body(
					app.Div().Style("display", "flex").Style("flex-direction", "column").Style("gap", "24px").Body(
						app.Div().Body(
							sp.FieldLabel().
								For("textfield4").
								Text("Small (S)").
								Size(sp.FieldLabelSizeS),
							sp.Textfield().
								Id("textfield4").
								Size("s"),
						),
						app.Div().Body(
							sp.FieldLabel().
								For("textfield5").
								Text("Medium (M) - Default").
								Size(sp.FieldLabelSizeM),
							sp.Textfield().
								Id("textfield5").
								Size("m"),
						),
						app.Div().Body(
							sp.FieldLabel().
								For("textfield6").
								Text("Large (L)").
								Size(sp.FieldLabelSizeL),
							sp.Textfield().
								Id("textfield6").
								Size("l"),
						),
						app.Div().Body(
							sp.FieldLabel().
								For("textfield7").
								Text("Extra Large (XL)").
								Size(sp.FieldLabelSizeXl),
							sp.Textfield().
								Id("textfield7").
								Size("xl"),
						),
					),
				),
			),

		// Side-aligned field label
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Side-Aligned Field Label"),
				app.Div().Class("component-row").Body(
					app.Div().Style("display", "flex").Style("flex-direction", "column").Style("gap", "24px").Body(
						app.Div().Style("display", "flex").Style("align-items", "center").Style("gap", "16px").Body(
							sp.FieldLabel().
								For("textfield8").
								Text("Start Aligned").
								SideAligned(sp.FieldLabelSideAlignedStart),
							sp.Textfield().
								Id("textfield8").
								Style("flex", "1"),
						),
						app.Div().Style("display", "flex").Style("align-items", "center").Style("gap", "16px").Body(
							sp.FieldLabel().
								For("textfield9").
								Text("End Aligned").
								SideAligned(sp.FieldLabelSideAlignedEnd),
							sp.Textfield().
								Id("textfield9").
								Style("flex", "1"),
						),
					),
				),
			),

		// Field label with custom HTML
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Field Label with Custom HTML"),
				app.Div().Class("component-row").Body(
					app.Div().Style("display", "flex").Style("flex-direction", "column").Style("gap", "8px").Body(
						sp.FieldLabel().
							For("textfield10").
							Body(
								app.Strong().Text("Important Field"),
								app.Span().Text(" (required)").Style("color", "red"),
							),
						sp.Textfield().
							Id("textfield10").
							Required(true),
					),
				),
			),

		// Field label with other form components
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Field Label with Different Components"),
				app.Div().Class("component-row").Body(
					app.Div().Style("display", "flex").Style("flex-direction", "column").Style("gap", "24px").Body(
						// With textarea
						app.Div().Style("display", "flex").Style("flex-direction", "column").Style("gap", "8px").Body(
							sp.FieldLabel().
								For("textarea1").
								Text("Comments").
								Required(true),
							sp.Textarea().
								Id("textarea1").
								Required(true),
						),
						// With select/picker
						app.Div().Style("display", "flex").Style("flex-direction", "column").Style("gap", "8px").Body(
							sp.FieldLabel().
								For("picker1").
								Text("Country"),
							sp.Picker().
								Id("picker1").
								Label("Select a country").
								Body(
									sp.MenuItem().Text("United States").Value("us"),
									sp.MenuItem().Text("Canada").Value("ca"),
									sp.MenuItem().Text("Mexico").Value("mx"),
								),
						),
						// With checkbox
						app.Div().Style("display", "flex").Style("gap", "8px").Body(
							sp.Checkbox().
								Id("checkbox1"),
							sp.FieldLabel().
								For("checkbox1").
								Text("I agree to the terms"),
						),
					),
				),
			),
	)
}
