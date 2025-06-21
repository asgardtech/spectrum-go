package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/asgardtech/spectrum-go/prism"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type cardPage struct {
	app.Compo
}

func newCardPage() *cardPage {
	return &cardPage{}
}

func (p *cardPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *cardPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle(defaultTitle)
	ctx.Page().SetDescription(defaultDescription)
}

func (p *cardPage) Render() app.UI {
	return prism.NewLayout().Content(
		app.H1().Text("Card Component"),
		app.P().Text("The Card component is a versatile container that organizes content into distinct sections. Cards are typically used to encapsulate units of a data set, such as a gallery of image/caption pairs."),

		// Basic Card
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Basic Card"),
				app.P().Text("A standard card can contain a heading, subheading, cover photo, and footer."),
				app.Div().
					Style("max-width", "300px").
					Body(
						app.Div().Body(
							sp.Card().
								Heading("Card Heading"),
							app.Img().
								Attr("slot", "cover-photo").
								Src("https://picsum.photos/250/300").
								Alt("Sample image"),
							app.Span().
								Attr("slot", "subheading").
								Text("JPG photo"),
							app.Div().
								Attr("slot", "footer").
								Text("Footer content"),
						),
					),
			),

		// Card with Custom Heading
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Card with Custom Heading"),
				app.P().Text("For more complex headings, use the heading slot instead of the heading attribute."),
				app.Div().
					Style("max-width", "300px").
					Body(
						app.Div().Body(
							sp.Card().
								Subheading("JPG Photo"),
							app.H1().
								Attr("slot", "heading").
								Text("Custom Card Heading"),
							app.Img().
								Attr("slot", "cover-photo").
								Src("https://picsum.photos/250/300").
								Alt("Sample image"),
							app.Div().
								Attr("slot", "footer").
								Text("Footer content"),
						),
					),
			),

		// Linked Card
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Linked Card"),
				app.P().Text("Cards can act as a large anchor element when provided with an href attribute."),
				app.Div().
					Style("max-width", "300px").
					Body(
						app.Div().Body(
							sp.Card().
								Heading("Linked Card").
								Subheading("Click to open").
								Href("https://picsum.photos").
								Target("_blank"),
							app.Img().
								Attr("slot", "cover-photo").
								Src("https://picsum.photos/250/300").
								Alt("Sample image"),
						),
					),
			),

		// Horizontal Card
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Horizontal Card"),
				app.P().Text("Cards can be arranged horizontally to display content side by side."),
				app.Div().
					Style("max-width", "500px").
					Body(
						app.Div().Body(
							sp.Card().
								Horizontal(true).
								Heading("Horizontal Card").
								Subheading("Side-by-side layout"),
							app.Img().
								Attr("slot", "cover-photo").
								Src("https://picsum.photos/200/250").
								Alt("Sample image"),
							app.Div().
								Attr("slot", "description").
								Text("This card uses a horizontal layout to display content side by side, which is useful for displaying items in a list."),
						),
					),
			),

		// Quiet Card
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Quiet Card"),
				app.P().Text("Quiet cards have a subtle appearance with minimal visual noise."),
				app.Div().
					Style("max-width", "300px").
					Body(
						app.Div().Body(
							sp.Card().
								Variant("quiet").
								Heading("Quiet Card").
								Subheading("Subtle appearance"),
							app.Img().
								Attr("slot", "preview").
								Src("https://picsum.photos/200/250").
								Alt("Sample image"),
							app.Div().
								Attr("slot", "description").
								Text("Created on 10/15/2023"),
							app.Div().
								Attr("slot", "footer").
								Text("Footer content"),
						),
					),
			),

		// Gallery Card
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Gallery Card"),
				app.P().Text("Gallery cards are optimized for displaying images in a gallery format."),
				app.Div().
					Style("max-width", "300px").
					Body(
						app.Div().Body(
							sp.Card().
								Variant("gallery").
								Heading("Gallery Card").
								Subheading("Optimized for images"),
							app.Img().
								Attr("slot", "preview").
								Src("https://picsum.photos/300/200").
								Alt("Sample image"),
							app.Div().
								Attr("slot", "description").
								Text("Created on 10/15/2023"),
							app.Div().
								Attr("slot", "footer").
								Text("Footer content"),
						),
					),
			),

		// Asset Cards
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Asset Cards"),
				app.P().Text("Cards can represent file or folder assets."),
				app.Div().
					Style("display", "flex").
					Style("gap", "16px").
					Body(
						// File Asset Card
						app.Div().
							Style("max-width", "300px").
							Body(
								app.Div().Body(
									sp.Card().
										Asset("file"),
									app.Div().
										Attr("slot", "heading").
										Text("File Name.pdf"),
									app.Div().
										Attr("slot", "subheading").
										Text("PDF Document"),
									app.Div().
										Attr("slot", "description").
										Text("Modified 10/15/2023"),
									app.Div().
										Attr("slot", "footer").
										Text("2.4 MB"),
								),
							),

						// Folder Asset Card
						app.Div().
							Style("max-width", "300px").
							Body(
								app.Div().Body(
									sp.Card().
										Asset("folder"),
									app.Div().
										Attr("slot", "heading").
										Text("Project Files"),
									app.Div().
										Attr("slot", "subheading").
										Text("Folder"),
									app.Div().
										Attr("slot", "description").
										Text("Modified 10/15/2023"),
									app.Div().
										Attr("slot", "footer").
										Text("15 items"),
								),
							),
					),
			),

		// Card without Preview Image
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Card without Preview Image"),
				app.P().Text("Cards can display content without an image."),
				app.Div().
					Style("max-width", "300px").
					Body(
						app.Div().Body(
							sp.Card().
								Heading("Text-Only Card").
								Subheading("No preview image"),
							app.Div().
								Attr("slot", "description").
								Text("This card displays text content without an image, which is useful for displaying information."),
							app.Div().
								Attr("slot", "footer").
								Text("Footer content"),
						),
					),
			),
	)
}
