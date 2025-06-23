package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/asgardtech/spectrum-go/prism"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type breadcrumbsPage struct {
	app.Compo
}

func newBreadcrumbsPage() *breadcrumbsPage {
	return &breadcrumbsPage{}
}

func (p *breadcrumbsPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *breadcrumbsPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle(defaultTitle)
	ctx.Page().SetDescription(defaultDescription)
}

func (p *breadcrumbsPage) Render() app.UI {
	return prism.NewPage().Content(
		app.H1().Text("Breadcrumbs Component"),
		app.P().Text("Breadcrumbs display a hierarchy of links to the current page or resource in an application."),

		// Basic Breadcrumbs
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Basic Breadcrumbs"),
				app.P().Text("A simple breadcrumb navigation with multiple levels."),

				app.Div().
					Style("display", "flex").
					Style("flex-direction", "column").
					Style("gap", "20px").
					Style("margin", "20px 0").
					Body(
						sp.Breadcrumbs().
							Body(
								sp.BreadcrumbItem().
									Value("home").
									Text("Home"),
								sp.BreadcrumbItem().
									Value("products").
									Text("Products"),
								sp.BreadcrumbItem().
									Value("electronics").
									Text("Electronics"),
								sp.BreadcrumbItem().
									Value("cameras").
									Text("Cameras"),
							),
					),
			),

		// Compact Breadcrumbs
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Compact Breadcrumbs"),
				app.P().Text("Compact breadcrumbs take up less vertical space, useful in space-constrained UIs."),

				app.Div().
					Style("display", "flex").
					Style("flex-direction", "column").
					Style("gap", "20px").
					Style("margin", "20px 0").
					Body(
						sp.Breadcrumbs().
							Compact(true).
							Body(
								sp.BreadcrumbItem().
									Value("home").
									Text("Home"),
								sp.BreadcrumbItem().
									Value("products").
									Text("Products"),
								sp.BreadcrumbItem().
									Value("electronics").
									Text("Electronics"),
								sp.BreadcrumbItem().
									Value("cameras").
									Text("Cameras"),
							),
					),
			),

		// Disabled Breadcrumb Items
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Disabled Breadcrumb Items"),
				app.P().Text("Breadcrumb items can be disabled to indicate they are not clickable."),

				app.Div().
					Style("display", "flex").
					Style("flex-direction", "column").
					Style("gap", "20px").
					Style("margin", "20px 0").
					Body(
						sp.Breadcrumbs().
							Body(
								sp.BreadcrumbItem().
									Value("home").
									Text("Home"),
								sp.BreadcrumbItem().
									Value("products").
									Disabled(true).
									Text("Products"),
								sp.BreadcrumbItem().
									Value("electronics").
									Text("Electronics"),
								sp.BreadcrumbItem().
									Value("cameras").
									Text("Cameras"),
							),
					),
			),

		// Breadcrumbs with Links
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Breadcrumbs with Links"),
				app.P().Text("Breadcrumb items can function as links when using the href attribute."),

				app.Div().
					Style("display", "flex").
					Style("flex-direction", "column").
					Style("gap", "20px").
					Style("margin", "20px 0").
					Body(
						sp.Breadcrumbs().
							Body(
								sp.BreadcrumbItem().
									Href("/").
									Text("Home"),
								sp.BreadcrumbItem().
									Href("/products").
									Text("Products"),
								sp.BreadcrumbItem().
									Href("/products/electronics").
									Text("Electronics"),
								sp.BreadcrumbItem().
									Href("/products/electronics/cameras").
									Text("Cameras"),
							),
					),
			),

		// Overflowing Breadcrumbs
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Overflowing Breadcrumbs"),
				app.P().Text("When there are many breadcrumb items, they automatically collapse into a menu."),

				app.Div().
					Style("display", "flex").
					Style("flex-direction", "column").
					Style("gap", "20px").
					Style("margin", "20px 0").
					Body(
						sp.Breadcrumbs().
							Body(
								sp.BreadcrumbItem().
									Value("home").
									Text("Home"),
								sp.BreadcrumbItem().
									Value("products").
									Text("Products"),
								sp.BreadcrumbItem().
									Value("electronics").
									Text("Electronics"),
								sp.BreadcrumbItem().
									Value("cameras").
									Text("Cameras"),
								sp.BreadcrumbItem().
									Value("digital").
									Text("Digital"),
								sp.BreadcrumbItem().
									Value("dslr").
									Text("DSLR"),
								sp.BreadcrumbItem().
									Value("canon").
									Text("Canon"),
								sp.BreadcrumbItem().
									Value("eos").
									Text("EOS Series"),
							),
					),
			),

		// Breadcrumbs with Root Item
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Breadcrumbs with Root Item"),
				app.P().Text("You can designate a breadcrumb item to always display as the root, even when items are collapsed."),

				app.Div().
					Style("display", "flex").
					Style("flex-direction", "column").
					Style("gap", "20px").
					Style("margin", "20px 0").
					Body(
						sp.Breadcrumbs().
							Root(
								sp.BreadcrumbItem().
									Value("home").
									Text("Home"),
							).
							Body(
								sp.BreadcrumbItem().
									Value("products").
									Text("Products"),
								sp.BreadcrumbItem().
									Value("electronics").
									Text("Electronics"),
								sp.BreadcrumbItem().
									Value("cameras").
									Text("Cameras"),
								sp.BreadcrumbItem().
									Value("digital").
									Text("Digital"),
								sp.BreadcrumbItem().
									Value("dslr").
									Text("DSLR"),
								sp.BreadcrumbItem().
									Value("canon").
									Text("Canon"),
								sp.BreadcrumbItem().
									Value("eos").
									Text("EOS Series"),
							),
					),
			),

		// Custom Maximum Visible Items
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Custom Maximum Visible Items"),
				app.P().Text("You can control how many breadcrumb items are visible before collapsing."),

				app.Div().
					Style("display", "flex").
					Style("flex-direction", "column").
					Style("gap", "20px").
					Style("margin", "20px 0").
					Body(
						sp.Breadcrumbs().
							MaxVisibleItems(3).
							Body(
								sp.BreadcrumbItem().
									Value("home").
									Text("Home"),
								sp.BreadcrumbItem().
									Value("products").
									Text("Products"),
								sp.BreadcrumbItem().
									Value("electronics").
									Text("Electronics"),
								sp.BreadcrumbItem().
									Value("cameras").
									Text("Cameras"),
								sp.BreadcrumbItem().
									Value("digital").
									Text("Digital"),
								sp.BreadcrumbItem().
									Value("dslr").
									Text("DSLR"),
							),
					),
			),

		// Accessibility Features
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Accessibility Features"),
				app.P().Text("Breadcrumbs include accessibility features like custom labels."),

				app.Div().
					Style("display", "flex").
					Style("flex-direction", "column").
					Style("gap", "20px").
					Style("margin", "20px 0").
					Body(
						sp.Breadcrumbs().
							Label("Site navigation").
							MenuLabel("More navigation options").
							Body(
								sp.BreadcrumbItem().
									Value("home").
									Text("Home"),
								sp.BreadcrumbItem().
									Value("products").
									Text("Products"),
								sp.BreadcrumbItem().
									Value("electronics").
									Text("Electronics"),
								sp.BreadcrumbItem().
									Value("cameras").
									Text("Cameras"),
								sp.BreadcrumbItem().
									Value("digital").
									Text("Digital"),
								sp.BreadcrumbItem().
									Value("dslr").
									Text("DSLR"),
							),
					),
			),

		// Usage Guidelines
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Usage Guidelines"),
				app.P().Text("When to use breadcrumbs:"),
				app.Ul().
					Body(
						app.Li().Text("Use breadcrumbs to help users understand their current location within a hierarchical structure."),
						app.Li().Text("Use breadcrumbs to provide a quick way to navigate back to previous levels."),
						app.Li().Text("Use breadcrumbs in complex applications with multiple levels of navigation."),
					),
				app.P().Text("Best practices:"),
				app.Ul().
					Body(
						app.Li().Text("Keep breadcrumb labels short and descriptive."),
						app.Li().Text("Order breadcrumbs from highest level (root) to lowest level (current page)."),
						app.Li().Text("Use the compact variant when vertical space is limited."),
						app.Li().Text("Consider using URLs in breadcrumbs for direct navigation to specific pages."),
					),
			),
	)
}
