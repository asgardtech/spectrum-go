package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/asgardtech/spectrum-go/prism"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type linkPage struct {
	app.Compo
}

func newLinkPage() *linkPage {
	return &linkPage{}
}

func (p *linkPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *linkPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle(defaultTitle)
	ctx.Page().SetDescription(defaultDescription)
}

func (p *linkPage) Render() app.UI {
	return prism.NewPage().Content(
		app.H1().Text("Link Component"),
		app.P().Text("Links allow users to navigate to a different location. They can be presented in-line inside a paragraph or as standalone text."),

		// Basic Link
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Basic Link"),
				app.P().Text("A standard link with default styling."),
				app.Div().Class("component-row").Body(
					sp.Link().
						Href("#").
						Text("Default Link"),
				),
			),

		// Secondary Link
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Secondary Link"),
				app.P().Text("A link with secondary visual style that better matches surrounding text."),
				app.Div().Class("component-row").Body(
					app.P().Body(
						app.Text("This is some paragraph text with a "),
						sp.Link().
							Href("#").
							Text("standard link"),
						app.Text(" and a "),
						sp.Link().
							Href("#").
							Variant("secondary").
							Text("secondary link"),
						app.Text(" embedded within it."),
					),
				),
			),

		// Quiet Link
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Quiet Link"),
				app.P().Text("A link without an underline."),
				app.Div().Class("component-row").Body(
					sp.Link().
						Href("#").
						Text("Standard Link"),
					sp.Link().
						Href("#").
						Quiet(true).
						Text("Quiet Link"),
				),
			),

		// Disabled Link
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Disabled Link"),
				app.P().Text("A link that cannot be interacted with."),
				app.Div().Class("component-row").Body(
					sp.Link().
						Href("#").
						Text("Enabled Link"),
					sp.Link().
						Href("#").
						Disabled(true).
						Text("Disabled Link"),
				),
			),

		// Static Color Links
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Static Color Links"),
				app.P().Text("Links with specific colors for different backgrounds."),
				app.Div().Class("color-background-demo dark").Body(
					app.P().
						Style("color", "white").
						Body(
							app.Text("This is text with a "),
							sp.Link().
								Href("#").
								StaticColor("white").
								Text("white link"),
							app.Text(" on a dark background."),
						),
				),
				app.Div().Class("color-background-demo light").Body(
					app.P().
						Style("color", "black").
						Body(
							app.Text("This is text with a "),
							sp.Link().
								Href("#").
								StaticColor("black").
								Text("black link"),
							app.Text(" on a light background."),
						),
				),
			),

		// Link with Target
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Link with Target"),
				app.P().Text("Links that open in different contexts."),
				app.Div().Class("component-row").Body(
					sp.Link().
						Href("https://github.com/asgardtech/spectrum-go").
						Text("Open in same window"),
					sp.Link().
						Href("https://github.com/asgardtech/spectrum-go").
						Target("_blank").
						Text("Open in new window"),
				),
			),

		// Download Link
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Download Link"),
				app.P().Text("A link that prompts a file download."),
				app.Div().Class("component-row").Body(
					sp.Link().
						Href("/file-example.pdf").
						Download("example.pdf").
						Text("Download file"),
				),
			),

		// Link with Relationship
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Link with Relationship Attributes"),
				app.P().Text("Links with relationship attributes for security and SEO."),
				app.Div().Class("component-row").Body(
					sp.Link().
						Href("https://github.com/asgardtech/spectrum-go").
						Target("_blank").
						Rel("noopener noreferrer").
						Text("Secure external link"),
				),
			),

		// Complex Examples
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Usage Examples"),

				// Link in paragraph
				app.H3().Text("Links in Text"),
				app.P().Body(
					app.Text("Adobe Spectrum is a design system that provides components for consistent user experiences. You can learn more about "),
					sp.Link().
						Href("https://spectrum.adobe.com/").
						Target("_blank").
						Rel("noopener").
						Text("Adobe Spectrum"),
					app.Text(" or check out the "),
					sp.Link().
						Href("https://opensource.adobe.com/spectrum-web-components/").
						Target("_blank").
						Rel("noopener").
						Text("Spectrum Web Components"),
					app.Text(" documentation."),
				),

				// Navigation links
				app.H3().Text("Navigation Links"),
				app.Div().Class("nav-links-example").Body(
					sp.Link().
						Href("#").
						Text("Home"),
					app.Text(" | "),
					sp.Link().
						Href("#").
						Text("Products"),
					app.Text(" | "),
					sp.Link().
						Href("#").
						Text("Services"),
					app.Text(" | "),
					sp.Link().
						Href("#").
						Text("About"),
					app.Text(" | "),
					sp.Link().
						Href("#").
						Text("Contact"),
				),

				// Link with icon (using HTML directly)
				app.H3().Text("Link with Icon"),
				app.Div().Class("component-row").Body(
					app.Raw(`
						<sp-link href="#">
							<sp-icon-info slot="icon"></sp-icon-info>
							Link with icon
						</sp-link>
					`),
				),
			),
	)
}
