package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/asgardtech/spectrum-go/prism"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type splitViewPage struct {
	app.Compo
}

func newSplitViewPage() *splitViewPage {
	return &splitViewPage{}
}

func (p *splitViewPage) Render() app.UI {
	return prism.NewLayout().Content(
		app.H1().Text("Split View Component"),
		app.P().Text("Split View delivers two elements in a horizontal or vertical orientation that distributes the available space. When resizable, it provides a way for users to customize the distribution."),

		// Basic horizontal split view
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Basic Horizontal Split View"),
				app.Div().Class("component-row").Style("height", "300px").Body(
					sp.SplitView().
						Body(
							app.Div().
								Style("padding", "16px").
								Style("background-color", "var(--spectrum-alias-background-color-primary)").
								Text("Left panel content"),
							app.Div().
								Style("padding", "16px").
								Style("background-color", "var(--spectrum-alias-background-color-secondary)").
								Text("Right panel content"),
						),
				),
			),

		// Resizable horizontal split view
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Resizable Horizontal Split View"),
				app.Div().Class("component-row").Style("height", "300px").Body(
					sp.SplitView().
						Resizable(true).
						PrimaryMin(100).
						SecondaryMin(100).
						PrimarySize("200px").
						Label("Resize the horizontal panels").
						Body(
							app.Div().
								Style("padding", "16px").
								Style("background-color", "var(--spectrum-alias-background-color-primary)").
								Text("Left panel content (resizable)"),
							app.Div().
								Style("padding", "16px").
								Style("background-color", "var(--spectrum-alias-background-color-secondary)").
								Text("Right panel content (resizable)"),
						),
				),
			),

		// Vertical split view
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Vertical Split View"),
				app.Div().Class("component-row").Style("height", "400px").Body(
					sp.SplitView().
						Vertical(true).
						Body(
							app.Div().
								Style("padding", "16px").
								Style("background-color", "var(--spectrum-alias-background-color-primary)").
								Text("Top panel content"),
							app.Div().
								Style("padding", "16px").
								Style("background-color", "var(--spectrum-alias-background-color-secondary)").
								Text("Bottom panel content"),
						),
				),
			),

		// Resizable vertical split view
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Resizable Vertical Split View"),
				app.Div().Class("component-row").Style("height", "400px").Body(
					sp.SplitView().
						Vertical(true).
						Resizable(true).
						PrimaryMin(50).
						SecondaryMin(50).
						PrimarySize("150px").
						Label("Resize the vertical panels").
						Body(
							app.Div().
								Style("padding", "16px").
								Style("background-color", "var(--spectrum-alias-background-color-primary)").
								Text("Top panel content (resizable)"),
							app.Div().
								Style("padding", "16px").
								Style("background-color", "var(--spectrum-alias-background-color-secondary)").
								Text("Bottom panel content (resizable)"),
						),
				),
			),

		// Collapsible split view
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Collapsible Split View"),
				app.Div().Class("component-row").Style("height", "300px").Body(
					sp.SplitView().
						Resizable(true).
						Collapsible(true).
						PrimaryMin(0).
						SecondaryMin(0).
						PrimarySize("200px").
						Label("Resize or collapse the panels").
						Body(
							app.Div().
								Style("padding", "16px").
								Style("background-color", "var(--spectrum-alias-background-color-primary)").
								Text("Left panel content (collapsible)"),
							app.Div().
								Style("padding", "16px").
								Style("background-color", "var(--spectrum-alias-background-color-secondary)").
								Text("Right panel content (collapsible)"),
						),
				),
			),

		// Percentage-based sizing
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Percentage-based Sizing"),
				app.Div().Class("component-row").Style("height", "300px").Body(
					sp.SplitView().
						Resizable(true).
						PrimarySize("30%").
						Label("Resize the panels").
						Body(
							app.Div().
								Style("padding", "16px").
								Style("background-color", "var(--spectrum-alias-background-color-primary)").
								Text("30% width panel"),
							app.Div().
								Style("padding", "16px").
								Style("background-color", "var(--spectrum-alias-background-color-secondary)").
								Text("70% width panel"),
						),
				),
			),

		// Complex content example
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Split View with Complex Content"),
				app.Div().Class("component-row").Style("height", "400px").Body(
					sp.SplitView().
						Resizable(true).
						PrimarySize("250px").
						PrimaryMin(150).
						Label("Resize navigation panel").
						Body(
							// Navigation panel
							app.Div().
								Style("height", "100%").
								Style("background-color", "var(--spectrum-alias-background-color-primary)").
								Body(
									app.H3().
										Style("margin", "8px").
										Text("Navigation"),
									sp.Sidenav().
										Label("Demo Navigation").
										Style("height", "calc(100% - 40px)").
										Body(
											sp.SidenavItem().
												Label("Dashboard").
												Text("Dashboard"),
											sp.SidenavItem().
												Label("Projects").
												Text("Projects"),
											sp.SidenavItem().
												Label("Tasks").
												Text("Tasks"),
											sp.SidenavItem().
												Label("Settings").
												Text("Settings"),
										),
								),
							// Content panel
							app.Div().
								Style("padding", "16px").
								Style("height", "100%").
								Style("overflow", "auto").
								Style("background-color", "var(--spectrum-alias-background-color-secondary)").
								Body(
									app.H3().Text("Content Area"),
									app.P().Text("This is an example of a split view with more complex content. The left panel contains a navigation menu, while the right panel displays the main content."),
									app.P().Text("You can resize the panels by dragging the divider between them."),
									app.P().Text("This layout is commonly used in applications with navigation sidebars."),
									sp.Divider(),
									app.H4().Text("Features"),
									sp.Card().
										Style("margin-bottom", "16px").
										Heading("Resizable Panels"),
									app.Div().
										Attr("slot", "description").
										Text("Users can customize the layout according to their preference."),
									sp.Card().
										Style("margin-bottom", "16px").
										Heading("Collapsible Panels"),
									app.Div().
										Attr("slot", "description").
										Text("Panels can be collapsed to maximize available space."),
									sp.Card().
										Heading("Vertical or Horizontal"),
									app.Div().
										Attr("slot", "description").
										Text("Split views can be arranged either vertically or horizontally."),
								),
						),
				),
			),
	)
}
