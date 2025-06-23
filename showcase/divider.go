package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/asgardtech/spectrum-go/prism"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type dividerPage struct {
	app.Compo
}

func newDividerPage() *dividerPage {
	return &dividerPage{}
}

func (p *dividerPage) Render() app.UI {
	return prism.NewPage().Content(
		app.H1().Text("Divider Component"),
		app.P().Text("Dividers bring clarity to a layout by grouping and dividing content that exists in close proximity. They can also be used to establish rhythm and hierarchy."),

		// Horizontal dividers - different sizes
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Horizontal Dividers"),
				app.Div().Class("component-row").Body(
					app.Div().Class("divider-container").Body(
						app.Text("Content above divider"),
						sp.Divider().SizeS(),
						app.Text("Small (S) divider"),
					),
					app.Div().Class("divider-container").Body(
						app.Text("Content above divider"),
						sp.Divider().SizeM(),
						app.Text("Medium (M) divider - Default"),
					),
					app.Div().Class("divider-container").Body(
						app.Text("Content above divider"),
						sp.Divider().SizeL(),
						app.Text("Large (L) divider"),
					),
				),
			),

		// Vertical dividers
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Vertical Dividers"),
				app.Div().Class("component-row vertical-demo").Style("height", "100px").Body(
					app.Text("Content before"),
					sp.Divider().SetVertical().SizeS(),
					app.Text("Small (S) divider"),
					sp.Divider().SetVertical().SizeM(),
					app.Text("Medium (M) divider"),
					sp.Divider().SetVertical().SizeL(),
					app.Text("Large (L) divider"),
				),
			),

		// Use case example - dividing sections
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Dividing Content Sections"),
				app.Div().Class("section-demo").Body(
					app.H3().Text("Section 1"),
					app.P().Text("This is the content of section 1. Dividers help create clear separation between different sections of content."),
					sp.Divider(),
					app.H3().Text("Section 2"),
					app.P().Text("This is the content of section 2. Notice how the divider creates visual separation."),
					sp.Divider(),
					app.H3().Text("Section 3"),
					app.P().Text("This is the content of section 3. The divider helps users understand where one section ends and another begins."),
				),
			),

		// Use case example - in cards
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Dividers in Cards"),
				app.Div().Class("component-row").Body(
					sp.Card().
						Heading("Card Title"),
					app.Div().
						Attr("slot", "description").
						Body(
							app.P().Text("Card content above the divider."),
							sp.Divider(),
							app.P().Text("Card content below the divider."),
							app.Div().Class("button-container").Body(
								sp.Button().Text("Primary Action").Variant(sp.ButtonVariantPrimary),
								sp.Button().Text("Secondary Action").Variant(sp.ButtonVariantSecondary),
							),
						),
				),
			),
	)
}
