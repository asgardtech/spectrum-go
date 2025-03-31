package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type iconPage struct {
	app.Compo
}

func newIconPage() *iconPage {
	return &iconPage{}
}

func (p *iconPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *iconPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle(defaultTitle)
	ctx.Page().SetDescription(defaultDescription)
}

func (p *iconPage) Render() app.UI {
	return newPage().Content(
		app.H1().Text("Icon Component"),
		app.P().Text("Icons represent objects, actions, and concepts in a visually compact form."),

		// Basic Icon
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Basic Icons"),
				app.P().Text("Simple icons with default settings."),
				app.Div().Class("component-row").Body(
					sp.Icon().Name("ui:ChevronDown100"),
					sp.Icon().Name("ui:ChevronUp100"),
					sp.Icon().Name("ui:ChevronLeft100"),
					sp.Icon().Name("ui:ChevronRight100"),
				),
			),

		// Icon with Label
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Accessible Icons"),
				app.P().Text("Icons with labels for accessibility."),
				app.Div().Class("component-row").Body(
					sp.Icon().Name("ui:Alert100").Label("Warning"),
					sp.Icon().Name("ui:CheckmarkCircle100").Label("Success"),
					sp.Icon().Name("ui:InfoCircle100").Label("Information"),
				),
			),

		// Size Variants
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Size Variants"),
				app.P().Text("Icons in different sizes."),
				app.Div().Class("component-row").Body(
					sp.Icon().Name("ui:Alert100").Size("xxs"),
					sp.Icon().Name("ui:Alert100").Size("xs"),
					sp.Icon().Name("ui:Alert100").Size("s"),
					sp.Icon().Name("ui:Alert100").Size("m"),
					sp.Icon().Name("ui:Alert100").Size("l"),
					sp.Icon().Name("ui:Alert100").Size("xl"),
					sp.Icon().Name("ui:Alert100").Size("xxl"),
				),
				app.Div().Class("size-labels").Body(
					app.Div().Class("size-label").Text("XXS"),
					app.Div().Class("size-label").Text("XS"),
					app.Div().Class("size-label").Text("S"),
					app.Div().Class("size-label").Text("M"),
					app.Div().Class("size-label").Text("L"),
					app.Div().Class("size-label").Text("XL"),
					app.Div().Class("size-label").Text("XXL"),
				),
			),

		// Custom Styled Icons
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Custom Styled Icons"),
				app.P().Text("Icons with custom colors and styling."),
				app.Div().Class("component-row").Body(
					sp.Icon().
						Name("ui:Alert100").
						Style("color", "red"),
					sp.Icon().
						Name("ui:CheckmarkCircle100").
						Style("color", "green"),
					sp.Icon().
						Name("ui:InfoCircle100").
						Style("color", "blue"),
					sp.Icon().
						Name("ui:Star100").
						Style("color", "gold"),
				),
			),

		// UI Icons
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("UI Icons"),
				app.P().Text("Common UI icons from the UI icon set."),
				app.Div().Class("component-row").Body(
					sp.Icon().Name("ui:Add100"),
					sp.Icon().Name("ui:Delete100"),
					sp.Icon().Name("ui:Edit100"),
					sp.Icon().Name("ui:Search100"),
					sp.Icon().Name("ui:Settings100"),
					sp.Icon().Name("ui:Help100"),
					sp.Icon().Name("ui:Preview100"),
					sp.Icon().Name("ui:Home100"),
					sp.Icon().Name("ui:MoreVertical100"),
				),
				app.Div().Class("icon-labels").Body(
					app.Div().Class("icon-label").Text("Add"),
					app.Div().Class("icon-label").Text("Delete"),
					app.Div().Class("icon-label").Text("Edit"),
					app.Div().Class("icon-label").Text("Search"),
					app.Div().Class("icon-label").Text("Settings"),
					app.Div().Class("icon-label").Text("Help"),
					app.Div().Class("icon-label").Text("Preview"),
					app.Div().Class("icon-label").Text("Home"),
					app.Div().Class("icon-label").Text("More"),
				),
			),

		// Status Icons
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Status Icons"),
				app.P().Text("Icons that represent different states."),
				app.Div().Class("component-row").Body(
					app.Div().Class("status-icon").Body(
						sp.Icon().Name("ui:CheckmarkCircle100").Style("color", "var(--spectrum-semantic-positive-color-default)"),
						app.Div().Class("status-label").Text("Success"),
					),
					app.Div().Class("status-icon").Body(
						sp.Icon().Name("ui:Alert100").Style("color", "var(--spectrum-semantic-negative-color-default)"),
						app.Div().Class("status-label").Text("Error"),
					),
					app.Div().Class("status-icon").Body(
						sp.Icon().Name("ui:AlertCircle100").Style("color", "var(--spectrum-semantic-notice-color-default)"),
						app.Div().Class("status-label").Text("Warning"),
					),
					app.Div().Class("status-icon").Body(
						sp.Icon().Name("ui:InfoCircle100").Style("color", "var(--spectrum-semantic-informative-color-default)"),
						app.Div().Class("status-label").Text("Info"),
					),
				),
			),

		// SVG as Icon
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Custom SVG Icons"),
				app.P().Text("Using custom SVG as icons."),
				app.Div().Class("component-row").Body(
					app.Raw(`
						<sp-icon>
							<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" height="24" width="24" fill="currentColor">
								<circle cx="12" cy="12" r="10" fill="none" stroke="currentColor" stroke-width="2"/>
								<path d="M12 6v6l4 2"/>
							</svg>
						</sp-icon>
					`),
					app.Raw(`
						<sp-icon>
							<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" height="24" width="24" fill="currentColor">
								<path d="M20 4H4c-1.1 0-2 .9-2 2v12c0 1.1.9 2 2 2h16c1.1 0 2-.9 2-2V6c0-1.1-.9-2-2-2zm0 14H4V6h16v12z"/>
								<path d="M8.5 13.5l2.5 3 3.5-4.5 4.5 6H5l3.5-4.5z"/>
							</svg>
						</sp-icon>
					`),
				),
			),

		// Icon in Various UI Components
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Icons in Components"),
				app.P().Text("Examples of icons used with other components."),

				// Button with Icon
				app.H3().Text("Button with Icon"),
				app.Div().Class("component-row").Body(
					sp.Button().
						Text("Download").
						Icon(sp.Icon().Name("ui:Download100")),
					sp.Button().
						Text("Settings").
						Icon(sp.Icon().Name("ui:Settings100")),
					sp.Button().
						Label("Add").
						Icon(sp.Icon().Name("ui:Add100")),
				),

				// Icon with Text
				app.H3().Text("Icon with Text"),
				app.Div().Class("component-row").Body(
					app.Div().Class("icon-text-row").Body(
						sp.Icon().Name("ui:InfoCircle100").Size("m"),
						app.Span().Text("This is important information"),
					),
					app.Div().Class("icon-text-row").Body(
						sp.Icon().Name("ui:Folder100").Size("m"),
						app.Span().Text("Project Files"),
					),
				),
			),

		// Icon Grid
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Icon Grid"),
				app.P().Text("A grid of commonly used icons."),
				app.Div().Class("icon-grid").Body(
					createIconGridItem("ui:Add100", "Add"),
					createIconGridItem("ui:Delete100", "Delete"),
					createIconGridItem("ui:Edit100", "Edit"),
					createIconGridItem("ui:Search100", "Search"),
					createIconGridItem("ui:Settings100", "Settings"),
					createIconGridItem("ui:Help100", "Help"),
					createIconGridItem("ui:Dashboard100", "Dashboard"),
					createIconGridItem("ui:Preview100", "Preview"),
					createIconGridItem("ui:Calendar100", "Calendar"),
					createIconGridItem("ui:User100", "User"),
					createIconGridItem("ui:Download100", "Download"),
					createIconGridItem("ui:Upload100", "Upload"),
				),
			),
	)
}

func createIconGridItem(iconName string, label string) app.UI {
	return app.Div().
		Class("icon-grid-item").
		Body(
			sp.Icon().Name(iconName).Size("m"),
			app.Div().Class("icon-grid-label").Text(label),
		)
}
