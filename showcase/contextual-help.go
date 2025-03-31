package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type contextualHelpPage struct {
	app.Compo
}

func newContextualHelpPage() *contextualHelpPage {
	return &contextualHelpPage{}
}

func (p *contextualHelpPage) Render() app.UI {
	return newPage().Content(
		app.H1().Text("Contextual Help Component"),
		app.P().Text("Contextual help shows a user extra information about the state of either an adjacent component or an entire view."),

		// Basic Contextual Help
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Basic Contextual Help"),
				app.Div().Class("component-row").Body(
					sp.ContextualHelp().
						Label("More info").
						Heading(app.H3().Text("Need help?")).
						Body(
							app.P().Text("This is a basic contextual help component that displays helpful information in a popover."),
						),
				),
			),

		// Variants
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Contextual Help Variants"),
				app.Div().Class("component-row").Body(
					sp.ContextualHelp().
						Label("Info").
						VariantInfo().
						Heading(app.H3().Text("Info Variant")).
						Body(
							app.P().Text("This uses the 'info' variant, which is the default."),
						),
					sp.ContextualHelp().
						Label("Help").
						VariantHelp().
						Heading(app.H3().Text("Help Variant")).
						Body(
							app.P().Text("This uses the 'help' variant, which changes the icon."),
						),
				),
			),

		// Placements
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Contextual Help Placements"),
				app.Div().Class("component-row").Body(
					app.Div().
						Style("display", "grid").
						Style("grid-template-columns", "repeat(4, 1fr)").
						Style("gap", "20px").
						Body(
							app.Div().Body(
								sp.ContextualHelp().
									Label("Top").
									PlacementTop().
									Heading(app.H3().Text("Top")).
									Body(
										app.P().Text("Popover placed at the top."),
									),
							),
							app.Div().Body(
								sp.ContextualHelp().
									Label("Right").
									PlacementRight().
									Heading(app.H3().Text("Right")).
									Body(
										app.P().Text("Popover placed at the right."),
									),
							),
							app.Div().Body(
								sp.ContextualHelp().
									Label("Bottom").
									PlacementBottom().
									Heading(app.H3().Text("Bottom")).
									Body(
										app.P().Text("Popover placed at the bottom."),
									),
							),
							app.Div().Body(
								sp.ContextualHelp().
									Label("Left").
									PlacementLeft().
									Heading(app.H3().Text("Left")).
									Body(
										app.P().Text("Popover placed at the left."),
									),
							),
						),
				),
			),

		// With Link
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Contextual Help with Link"),
				app.Div().Class("component-row").Body(
					sp.ContextualHelp().
						Label("Help with link").
						VariantHelp().
						Heading(app.H3().Text("Need more help?")).
						Body(
							app.P().Text("This contextual help includes a link to more information."),
						).
						Link(
							sp.Link().
								Href("https://github.com/asgardtech/spectrum-go").
								Target("_blank").
								Text("Learn more"),
						),
				),
			),

		// In Context Example
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("In Context Example"),
				app.P().Text("Contextual help is often used alongside form components to provide additional guidance."),
				app.Div().
					Class("component-row").
					Style("position", "relative").
					Body(
						app.Div().
							Style("display", "flex").
							Style("align-items", "center").
							Style("gap", "8px").
							Body(
								sp.FieldLabel().
									For("example-field").
									Text("Username"),
								sp.ContextualHelp().
									Label("Username help").
									PlacementRight().
									Heading(app.H3().Text("Choosing a Username")).
									Body(
										app.P().Text("Your username must be between 3-20 characters and can only contain letters, numbers, underscores, and hyphens."),
									),
							),
						sp.Textfield().
							Id("example-field").
							Placeholder("Enter username"),
					),
			),

		// Helper Text Options
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Helper Text with Contextual Help"),
				app.P().Text("Contextual help is often used alongside helper text."),
				app.Div().
					Class("component-row").
					Body(
						app.Div().
							Style("max-width", "300px").
							Body(
								sp.FieldLabel().
									For("password-field").
									Text("Password").
									AddToBody(
										sp.ContextualHelp().
											Label("Password requirements").
											Heading(app.H3().Text("Password Requirements")).
											Body(
												app.P().Text("Your password must contain at least:"),
												app.Ul().Body(
													app.Li().Text("8 characters"),
													app.Li().Text("One uppercase letter"),
													app.Li().Text("One lowercase letter"),
													app.Li().Text("One number"),
													app.Li().Text("One special character"),
												),
											),
									),
								sp.Textfield().
									Id("password-field").
									Type("password").
									Placeholder("Enter password"),
								sp.HelpText().
									Text("Password must be at least 8 characters"),
							),
					),
			),

		// Offset Example
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Contextual Help with Offset"),
				app.P().Text("You can adjust the position of the popover using the offset property."),
				app.Div().Class("component-row").Body(
					sp.ContextualHelp().
						Label("With offset").
						PlacementBottom().
						Offset("20").
						Heading(app.H3().Text("Offset Example")).
						Body(
							app.P().Text("This popover has an offset of 20 pixels from the trigger."),
						),
					sp.ContextualHelp().
						Label("With double offset").
						PlacementBottom().
						Offset("20, 10").
						Heading(app.H3().Text("Double Offset Example")).
						Body(
							app.P().Text("This popover has an offset of 20 pixels on the main axis and 10 pixels on the cross axis."),
						),
				),
			),
	)
}