package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type accordionPage struct {
	app.Compo
}

func newAccordionPage() *accordionPage {
	return &accordionPage{}
}

func (p *accordionPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *accordionPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle(defaultTitle)
	ctx.Page().SetDescription(defaultDescription)
}

func (p *accordionPage) Render() app.UI {
	return newPage().Content(
		app.H1().Text("Accordion Component"),
		app.P().Text("The Accordion component contains a list of items that can be expanded or collapsed to reveal additional content."),

		// Basic accordion
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Basic Accordion"),
				app.P().Text("By default, only one item can be expanded at a time."),
				app.Div().Body(
					sp.Accordion().Body(
						sp.AccordionItem().
							Label("Section 1").
							Body(
								app.Div().Text("Content for Section 1. This is the expanded content of the accordion item."),
							),
						sp.AccordionItem().
							Label("Section 2").
							Body(
								app.Div().Text("Content for Section 2. Click on the header to expand this section."),
							),
						sp.AccordionItem().
							Label("Section 3").
							Body(
								app.Div().Text("Content for Section 3. Each section can contain any content."),
							),
						sp.AccordionItem().
							Label("Disabled Section").
							Disabled(true).
							Body(
								app.Div().Text("This section is disabled and cannot be expanded."),
							),
					),
				),
			),

		// Allow multiple
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Allow Multiple"),
				app.P().Text("Setting 'allow-multiple' allows multiple items to be expanded at the same time."),
				app.Div().Body(
					sp.Accordion().
						AllowMultiple(true).
						Body(
							sp.AccordionItem().
								Label("Section 1").
								Body(
									app.Div().Text("Content for Section 1. Multiple sections can be open at once."),
								),
							sp.AccordionItem().
								Label("Section 2").
								Body(
									app.Div().Text("Content for Section 2. Try opening this while Section 1 is open."),
								),
							sp.AccordionItem().
								Label("Section 3").
								Body(
									app.Div().Text("Content for Section 3. You can open all sections at once."),
								),
						),
				),
			),

		// Different sizes
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Accordion Sizes"),
				app.P().Text("Accordions come in different sizes: s, m, l, and xl."),
				app.H3().Text("Small (S)"),
				app.Div().Body(
					sp.Accordion().
						Size("s").
						Body(
							sp.AccordionItem().
								Label("Small Accordion Item").
								Body(
									app.Div().Text("Content for small accordion item."),
								),
							sp.AccordionItem().
								Label("Another Small Item").
								Body(
									app.Div().Text("More content for small accordion."),
								),
						),
				),
				app.H3().Text("Medium (M) - Default"),
				app.Div().Body(
					sp.Accordion().
						Size("m").
						Body(
							sp.AccordionItem().
								Label("Medium Accordion Item").
								Body(
									app.Div().Text("Content for medium accordion item."),
								),
							sp.AccordionItem().
								Label("Another Medium Item").
								Body(
									app.Div().Text("More content for medium accordion."),
								),
						),
				),
				app.H3().Text("Large (L)"),
				app.Div().Body(
					sp.Accordion().
						Size("l").
						Body(
							sp.AccordionItem().
								Label("Large Accordion Item").
								Body(
									app.Div().Text("Content for large accordion item."),
								),
							sp.AccordionItem().
								Label("Another Large Item").
								Body(
									app.Div().Text("More content for large accordion."),
								),
						),
				),
				app.H3().Text("Extra Large (XL)"),
				app.Div().Body(
					sp.Accordion().
						Size("xl").
						Body(
							sp.AccordionItem().
								Label("Extra Large Accordion Item").
								Body(
									app.Div().Text("Content for extra large accordion item."),
								),
							sp.AccordionItem().
								Label("Another Extra Large Item").
								Body(
									app.Div().Text("More content for extra large accordion."),
								),
						),
				),
			),

		// Density variations
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Accordion Density"),
				app.P().Text("Accordions can have different density settings: compact or spacious."),
				app.H3().Text("Compact Density"),
				app.Div().Body(
					sp.Accordion().
						Density("compact").
						Body(
							sp.AccordionItem().
								Label("Compact Item 1").
								Body(
									app.Div().Text("Content for compact density accordion."),
								),
							sp.AccordionItem().
								Label("Compact Item 2").
								Body(
									app.Div().Text("More content for compact density accordion."),
								),
						),
				),
				app.H3().Text("Spacious Density"),
				app.Div().Body(
					sp.Accordion().
						Density("spacious").
						Body(
							sp.AccordionItem().
								Label("Spacious Item 1").
								Body(
									app.Div().Text("Content for spacious density accordion."),
								),
							sp.AccordionItem().
								Label("Spacious Item 2").
								Body(
									app.Div().Text("More content for spacious density accordion."),
								),
						),
				),
			),
	)
}
