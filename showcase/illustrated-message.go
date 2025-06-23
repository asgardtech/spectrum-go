package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/asgardtech/spectrum-go/prism"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type illustratedMessagePage struct {
	app.Compo
}

func newIllustratedMessagePage() *illustratedMessagePage {
	return &illustratedMessagePage{}
}

func (p *illustratedMessagePage) Render() app.UI {
	return prism.NewPage().Content(
		app.H1().Text("Illustrated Message Component"),
		app.P().Text("An illustrated message displays an illustration icon and a message, usually in an empty state or on an error page."),

		// Basic Illustrated Message
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Basic Illustrated Message"),
				app.Div().Class("component-row").Body(
					sp.IllustratedMessage().
						Heading("No results found").
						Description("Try adjusting your search or filter to find what you're looking for.").
						Body(
							sp.Icon().Name("ui:MagnifyMedium"),
						),
				),
			),

		// Illustrated Message with Custom Content
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Using Content Slots"),
				app.Div().Class("component-row").Body(
					sp.IllustratedMessage().
						HeadingContent(
							app.H3().
								Style("color", "var(--spectrum-blue-900)").
								Text("Custom Heading"),
						).
						DescriptionContent(
							app.P().
								Style("font-style", "italic").
								Text("This is a custom description using the description slot."),
						).
						Body(
							sp.Icon().Name("ui:InfoMedium"),
						),
				),
			),

		// Different Icons
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Different Icons"),
				app.Div().
					Style("display", "grid").
					Style("grid-template-columns", "repeat(3, 1fr)").
					Style("gap", "20px").
					Body(
						app.Div().Body(
							sp.IllustratedMessage().
								Heading("No items").
								Description("Your list is empty.").
								Body(
									sp.Icon().Name("ui:FolderOpenMedium"),
								),
						),
						app.Div().Body(
							sp.IllustratedMessage().
								Heading("No messages").
								Description("Your inbox is empty.").
								Body(
									sp.Icon().Name("ui:MailMedium"),
								),
						),
						app.Div().Body(
							sp.IllustratedMessage().
								Heading("Error occurred").
								Description("Something went wrong.").
								Body(
									sp.Icon().Name("ui:AlertMedium"),
								),
						),
					),
			),

		// Using Custom Illustration
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Custom Illustration"),
				app.Div().Class("component-row").Body(
					sp.IllustratedMessage().
						Heading("Success!").
						Description("Your file has been uploaded.").
						Body(
							app.Div().
								Style("width", "48px").
								Style("height", "48px").
								Style("border-radius", "50%").
								Style("background-color", "var(--spectrum-green-400)").
								Style("display", "flex").
								Style("align-items", "center").
								Style("justify-content", "center").
								Body(
									sp.Icon().
										Name("ui:CheckmarkMedium").
										Style("color", "white"),
								),
						),
				),
			),

		// In a Card
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Illustrated Message in a Card"),
				app.Div().Class("component-row").Body(
					sp.Card().
						Style("width", "350px").
						Style("padding", "20px").
						HeadingContent(
							sp.IllustratedMessage().
								Heading("No notifications").
								Description("You're all caught up!").
								Body(
									sp.Icon().Name("ui:NotificationMedium"),
								),
						),
				),
			),

		// With Action Button
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("With Action Button"),
				app.Div().Class("component-row").Body(
					sp.IllustratedMessage().
						Heading("No projects").
						Description("You haven't created any projects yet.").
						Body(
							sp.Icon().Name("ui:LayersMedium"),
							app.Div().
								Style("margin-top", "20px").
								Body(
									sp.Button().
										Text("Create Project").
										Variant(sp.ButtonVariantPrimary),
								),
						),
				),
			),

		// Error State
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Error State"),
				app.Div().Class("component-row").Body(
					sp.IllustratedMessage().
						HeadingContent(
							app.H3().
								Style("color", "var(--spectrum-negative-content-color)").
								Text("Connection Error"),
						).
						DescriptionContent(
							app.P().
								Style("color", "var(--spectrum-negative-content-color)").
								Text("Unable to connect to the server. Please check your internet connection and try again."),
						).
						Body(
							sp.Icon().
								Name("ui:AlertMedium").
								Style("color", "var(--spectrum-negative-content-color)"),
							app.Div().
								Style("margin-top", "20px").
								Body(
									sp.Button().
										Text("Retry").
										Variant(sp.ButtonVariantPrimary),
								),
						),
				),
			),

		// Empty Search Results
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Empty Search Results"),
				app.Div().Class("component-row").Body(
					app.Div().
						Style("width", "100%").
						Style("max-width", "500px").
						Style("border", "1px solid var(--spectrum-gray-300)").
						Style("border-radius", "4px").
						Style("padding", "20px").
						Body(
							sp.Search().
								Label("Search").
								Placeholder("Search...").
								Style("width", "100%").
								Style("margin-bottom", "20px"),
							sp.IllustratedMessage().
								Heading("No results found").
								Description("Try different keywords or filters.").
								Body(
									sp.Icon().Name("ui:MagnifyMedium"),
								),
						),
				),
			),
	)
}
