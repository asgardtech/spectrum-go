package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/asgardtech/spectrum-go/prism"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type fieldGroupPage struct {
	app.Compo
}

func newFieldGroupPage() *fieldGroupPage {
	return &fieldGroupPage{}
}

func (p *fieldGroupPage) Render() app.UI {
	return prism.NewPage().Content(
		app.H1().Text("Field Group Component"),
		app.P().Text("A field group is used to layout a group of fields, usually checkbox elements. It can be leveraged for vertical or horizontal organization of the fields that are supplied as its children."),

		// Basic Field Group with Checkboxes
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Basic Field Group with Checkboxes"),
				app.Div().Class("component-row").Body(
					sp.FieldGroup().
						Label("Notification Preferences").
						Body(
							sp.Checkbox().
								Text("Email notifications"),
							sp.Checkbox().
								Text("SMS notifications"),
							sp.Checkbox().
								Text("Push notifications"),
						),
				),
			),

		// Vertical Field Group
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Vertical Field Group"),
				app.Div().Class("component-row").Body(
					sp.FieldGroup().
						Label("Payment Options").
						SetVertical().
						Body(
							sp.Radio().
								Text("Credit Card"),
							sp.Radio().
								Text("PayPal"),
							sp.Radio().
								Text("Bank Transfer"),
						),
				),
			),

		// Horizontal Field Group
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Horizontal Field Group"),
				app.Div().Class("component-row").Body(
					sp.FieldGroup().
						Label("Frequency").
						SetHorizontal().
						Body(
							sp.Radio().
								Text("Daily"),
							sp.Radio().
								Text("Weekly"),
							sp.Radio().
								Text("Monthly"),
						),
				),
			),

		// Field Group with Help Text
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Field Group with Help Text"),
				app.Div().Class("component-row").Body(
					sp.FieldGroup().
						Label("Access Permissions").
						Body(
							sp.Checkbox().
								Text("Read"),
							sp.Checkbox().
								Text("Write"),
							sp.Checkbox().
								Text("Execute"),
						).
						HelpText(
							sp.HelpText().Text("Select the permissions you want to grant."),
						),
				),
			),

		// Invalid Field Group with Negative Help Text
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Invalid Field Group"),
				app.Div().Class("component-row").Body(
					sp.FieldGroup().
						Label("Required Fields").
						SetInvalid().
						Body(
							sp.Checkbox().
								Text("I agree to the terms and conditions"),
							sp.Checkbox().
								Text("I want to receive the newsletter"),
						).
						NegativeHelpText(
							sp.HelpText().Text("You must agree to the terms and conditions to continue."),
						),
				),
			),

		// Field Group with Textfields
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Field Group with Textfields"),
				app.Div().Class("component-row").Body(
					sp.FieldGroup().
						Label("Name").
						Body(
							app.Div().
								Style("display", "flex").
								Style("gap", "10px").
								Body(
									sp.Textfield().
										Label("First Name").
										Placeholder("First Name"),
									sp.Textfield().
										Label("Last Name").
										Placeholder("Last Name"),
								),
						),
				),
			),

		// Complex Example with Mixed Components
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Complex Example with Mixed Components"),
				app.Div().Class("component-row").Body(
					sp.FieldGroup().
						Label("Contact Information").
						SetVertical().
						Body(
							app.Div().
								Style("margin-bottom", "10px").
								Body(
									sp.FieldLabel().Text("Email"),
									sp.Textfield().
										Placeholder("Enter your email"),
								),
							app.Div().
								Style("margin-bottom", "10px").
								Body(
									sp.FieldLabel().Text("Phone"),
									sp.Textfield().
										Placeholder("Enter your phone number"),
								),
							app.Div().
								Style("margin-bottom", "10px").
								Body(
									sp.FieldLabel().Text("Contact Preferences"),
									sp.FieldGroup().
										SetHorizontal().
										Body(
											sp.Checkbox().
												Text("Email"),
											sp.Checkbox().
												Text("Phone"),
											sp.Checkbox().
												Text("SMS"),
										),
								),
						).
						HelpText(
							sp.HelpText().Text("We'll only use this information to contact you about your account."),
						),
				),
			),
	)
}
