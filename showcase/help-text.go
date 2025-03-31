package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type helpTextPage struct {
	app.Compo
}

func newHelpTextPage() *helpTextPage {
	return &helpTextPage{}
}

func (p *helpTextPage) Render() app.UI {
	return newPage().Content(
		app.H1().Text("Help Text Component"),
		app.P().Text("Help Text provides context and assistance to the user about a particular field or component."),

		// Basic help text
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Basic Help Text"),
				app.Div().Class("component-row").Body(
					app.Div().Style("display", "flex").Style("flex-direction", "column").Style("gap", "8px").Body(
						sp.FieldLabel().
							For("textfield1").
							Text("Username"),
						sp.Textfield().
							Id("textfield1").
							Placeholder("Enter your username"),
						sp.HelpText().
							Text("Username must be between 3-16 characters"),
					),
				),
			),

		// Help text variants
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Help Text Variants"),
				app.Div().Class("component-row").Body(
					app.Div().Style("display", "flex").Style("flex-direction", "column").Style("gap", "24px").Body(
						app.Div().Style("display", "flex").Style("flex-direction", "column").Style("gap", "8px").Body(
							sp.FieldLabel().
								For("textfield2").
								Text("Email"),
							sp.Textfield().
								Id("textfield2").
								Placeholder("Enter your email"),
							sp.HelpText().
								Variant(sp.HelpTextVariantNeutral).
								Text("Neutral help text (default)"),
						),
						app.Div().Style("display", "flex").Style("flex-direction", "column").Style("gap", "8px").Body(
							sp.FieldLabel().
								For("textfield3").
								Text("Password"),
							sp.Textfield().
								Id("textfield3").
								Type("password").
								Invalid(true),
							sp.HelpText().
								Variant(sp.HelpTextVariantNegative).
								Text("Negative help text for validation errors"),
						),
					),
				),
			),

		// Help text with icon
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Help Text with Icon"),
				app.Div().Class("component-row").Body(
					app.Div().Style("display", "flex").Style("flex-direction", "column").Style("gap", "24px").Body(
						app.Div().Style("display", "flex").Style("flex-direction", "column").Style("gap", "8px").Body(
							sp.FieldLabel().
								For("textfield4").
								Text("Website"),
							sp.Textfield().
								Id("textfield4").
								Placeholder("Enter website URL"),
							sp.HelpText().
								SetIcon().
								Text("Include https:// for secure websites"),
						),
						app.Div().Style("display", "flex").Style("flex-direction", "column").Style("gap", "8px").Body(
							sp.FieldLabel().
								For("textfield5").
								Text("Password"),
							sp.Textfield().
								Id("textfield5").
								Type("password").
								Invalid(true),
							sp.HelpText().
								Variant(sp.HelpTextVariantNegative).
								SetIcon().
								Text("Password must include at least one number"),
						),
					),
				),
			),

		// Disabled help text
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Disabled Help Text"),
				app.Div().Class("component-row").Body(
					app.Div().Style("display", "flex").Style("flex-direction", "column").Style("gap", "8px").Body(
						sp.FieldLabel().
							For("textfield6").
							Text("Phone Number").
							Disabled(true),
						sp.Textfield().
							Id("textfield6").
							Disabled(true),
						sp.HelpText().
							SetDisabled().
							Text("Format: (000) 000-0000"),
					),
				),
			),

		// Help text sizes
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Help Text Sizes"),
				app.Div().Class("component-row").Body(
					app.Div().Style("display", "flex").Style("flex-direction", "column").Style("gap", "24px").Body(
						app.Div().Body(
							sp.HelpText().
								Size(sp.HelpTextSizeS).
								Text("Small (S) help text"),
						),
						app.Div().Body(
							sp.HelpText().
								Size(sp.HelpTextSizeM).
								Text("Medium (M) help text - Default"),
						),
						app.Div().Body(
							sp.HelpText().
								Size(sp.HelpTextSizeL).
								Text("Large (L) help text"),
						),
						app.Div().Body(
							sp.HelpText().
								Size(sp.HelpTextSizeXl).
								Text("Extra Large (XL) help text"),
						),
					),
				),
			),

		// Help text with other components
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Help Text with Different Components"),
				app.Div().Class("component-row").Body(
					app.Div().Style("display", "flex").Style("flex-direction", "column").Style("gap", "24px").Body(
						// With textarea
						app.Div().Style("display", "flex").Style("flex-direction", "column").Style("gap", "8px").Body(
							sp.FieldLabel().
								For("textarea1").
								Text("Comments"),
							sp.Textarea().
								Id("textarea1"),
							sp.HelpText().
								Text("Maximum 500 characters"),
						),
						// With checkbox
						app.Div().Style("display", "flex").Style("flex-direction", "column").Style("gap", "8px").Body(
							app.Div().Style("display", "flex").Style("gap", "8px").Body(
								sp.Checkbox().
									Id("checkbox1"),
								sp.FieldLabel().
									For("checkbox1").
									Text("Subscribe to newsletter"),
							),
							sp.HelpText().
								Text("We send updates once a month"),
						),
					),
				),
			),

		// Help text with complex content
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Help Text with Complex Content"),
				app.Div().Class("component-row").Body(
					app.Div().Style("display", "flex").Style("flex-direction", "column").Style("gap", "8px").Body(
						sp.FieldLabel().
							For("password1").
							Text("Password"),
						sp.Textfield().
							Id("password1").
							Type("password"),
						sp.HelpText().
							Body(
								app.Strong().Text("Password requirements:"),
								app.Ul().Body(
									app.Li().Text("At least 8 characters"),
									app.Li().Text("At least one uppercase letter"),
									app.Li().Text("At least one number"),
									app.Li().Text("At least one special character"),
								),
							),
					),
				),
			),
	)
}