package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/asgardtech/spectrum-go/prism"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type alertDialogPage struct {
	app.Compo
}

func newAlertDialogPage() *alertDialogPage {
	return &alertDialogPage{}
}

func (p *alertDialogPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *alertDialogPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle(defaultTitle)
	ctx.Page().SetDescription(defaultDescription)
}

func (p *alertDialogPage) Render() app.UI {
	return prism.NewPage().Content(
		app.H1().Text("Alert Dialog Component"),
		app.P().Text("Alert Dialog displays important information that users need to acknowledge. It can be used for confirmations, warnings, errors, and other critical messages."),

		// Confirmation Dialog
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Confirmation Dialog"),
				app.P().Text("Use a confirmation dialog to ask users to confirm a choice."),
				app.Div().Body(
					sp.AlertDialog().
						Variant("confirmation").
						Body(
							app.H2().
								Attr("slot", "heading").
								Text("Disclaimer"),
							app.Text("Smart filters are nondestructive and will preserve your original images."),
							app.Div().
								Attr("slot", "button").
								Body(
									sp.Button().
										Id("cancelButton").
										Variant("secondary").
										Treatment("outline").
										Text("Cancel"),
								),
							app.Div().
								Attr("slot", "button").
								Body(
									sp.Button().
										Id("confirmButton").
										Variant("accent").
										Treatment("fill").
										Text("Enable"),
								),
						),
				),
			),

		// Information Dialog
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Information Dialog"),
				app.P().Text("Information dialogs communicate important information that users need to acknowledge."),
				app.Div().Body(
					sp.AlertDialog().
						Variant("information").
						Body(
							app.H2().
								Attr("slot", "heading").
								Text("Connect to wifi"),
							app.Text("Please connect to wifi to sync your projects or go to Settings to change your preferences."),
							app.Div().
								Attr("slot", "button").
								Body(
									sp.Button().
										Id("cancelButton").
										Variant("secondary").
										Treatment("outline").
										Text("Cancel"),
								),
							app.Div().
								Attr("slot", "button").
								Body(
									sp.Button().
										Id("confirmButton").
										Variant("primary").
										Treatment("outline").
										Text("Continue"),
								),
						),
				),
			),

		// Warning Dialog
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Warning Dialog"),
				app.P().Text("Warning dialogs communicate important information about issues that need acknowledgment."),
				app.Div().Body(
					sp.AlertDialog().
						Variant("warning").
						Body(
							app.H2().
								Attr("slot", "heading").
								Text("Unverified format"),
							app.Text("This format has not been verified and may not be viewable for some users. Do you want to continue publishing?"),
							app.Div().
								Attr("slot", "button").
								Body(
									sp.Button().
										Id("cancelButton").
										Variant("secondary").
										Treatment("outline").
										Text("Cancel"),
								),
							app.Div().
								Attr("slot", "button").
								Body(
									sp.Button().
										Id("confirmButton").
										Variant("primary").
										Treatment("outline").
										Text("Continue"),
								),
						),
				),
			),

		// Error Dialog
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Error Dialog"),
				app.P().Text("Error dialogs communicate critical information about issues that need acknowledgment."),
				app.Div().Body(
					sp.AlertDialog().
						Variant("error").
						Body(
							app.H2().
								Attr("slot", "heading").
								Text("Unable to share"),
							app.Text("An error occurred while sharing your project. Please verify the email address and try again."),
							app.Div().
								Attr("slot", "button").
								Body(
									sp.Button().
										Id("confirmButton").
										Variant("primary").
										Treatment("outline").
										Text("Continue"),
								),
						),
				),
			),

		// Destructive Dialog
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Destructive Dialog"),
				app.P().Text("Destructive dialogs are for confirming actions that will impact data or experience negatively."),
				app.Div().Body(
					sp.AlertDialog().
						Variant("destructive").
						Body(
							app.H2().
								Attr("slot", "heading").
								Text("Delete 3 documents?"),
							app.Text("Are you sure you want to delete the 3 selected documents?"),
							app.Div().
								Attr("slot", "button").
								Body(
									sp.Button().
										Id("cancelButton").
										Variant("secondary").
										Treatment("outline").
										Text("Cancel"),
								),
							app.Div().
								Attr("slot", "button").
								Body(
									sp.Button().
										Id("confirmButton").
										Variant("negative").
										Treatment("fill").
										Text("Delete"),
								),
						),
				),
			),

		// Secondary Button Dialog
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Secondary Button Dialog"),
				app.P().Text("Alert dialogs can have up to three buttons, including a secondary outline button."),
				app.Div().Body(
					sp.AlertDialog().
						Variant("secondary").
						Body(
							app.H2().
								Attr("slot", "heading").
								Text("Rate this app"),
							app.Text("If you enjoy our app, would you mind taking a moment to rate it?"),
							app.Div().
								Attr("slot", "button").
								Body(
									sp.Button().
										Id("secondaryButton").
										Variant("secondary").
										Treatment("outline").
										Text("No, thanks"),
								),
							app.Div().
								Attr("slot", "button").
								Body(
									sp.Button().
										Id("cancelButton").
										Variant("secondary").
										Treatment("outline").
										Text("Remind me later"),
								),
							app.Div().
								Attr("slot", "button").
								Body(
									sp.Button().
										Id("confirmButton").
										Variant("primary").
										Treatment("outline").
										Text("Rate now"),
								),
						),
				),
			),
	)
}
