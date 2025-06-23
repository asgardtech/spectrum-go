package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/asgardtech/spectrum-go/prism"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type tabsPage struct {
	app.Compo

	// Track the selected tab
	selectedTabValue string
}

func newTabsPage() *tabsPage {
	return &tabsPage{
		selectedTabValue: "tab1",
	}
}

func (p *tabsPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *tabsPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle(defaultTitle)
	ctx.Page().SetDescription(defaultDescription)
}

func (p *tabsPage) handleTabChange(ctx app.Context, e app.Event) {
	// Get the selected tab value from the event
	value := e.Get("value").String()
	p.selectedTabValue = value
	ctx.Reload()
}

func (p *tabsPage) Render() app.UI {
	return prism.NewPage().Content(
		app.H1().Text("Tabs Component"),
		app.P().Text("Tabs organize content into multiple sections and allow users to navigate between them."),

		// Basic Tabs
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Basic Tabs"),
				app.P().Text("Basic tabs with horizontal orientation."),

				app.Div().
					Style("display", "flex").
					Style("flex-direction", "column").
					Style("gap", "20px").
					Style("margin", "20px 0").
					Body(
						sp.Tabs().
							Selected("tab1").
							OnChange(p.handleTabChange).
							Body(
								sp.Tab().
									Label("Home").
									Value("tab1"),
								sp.Tab().
									Label("Profile").
									Value("tab2"),
								sp.Tab().
									Label("Settings").
									Value("tab3"),
								sp.TabPanel().
									Value("tab1").
									Body(
										app.Div().
											Style("padding", "20px").
											Style("border", "1px solid var(--spectrum-global-color-gray-300)").
											Style("border-top", "none").
											Body(
												app.H3().Text("Home Content"),
												app.P().Text("This is the content for the Home tab. It appears when the Home tab is selected."),
											),
									),
								sp.TabPanel().
									Value("tab2").
									Body(
										app.Div().
											Style("padding", "20px").
											Style("border", "1px solid var(--spectrum-global-color-gray-300)").
											Style("border-top", "none").
											Body(
												app.H3().Text("Profile Content"),
												app.P().Text("This is the content for the Profile tab. It appears when the Profile tab is selected."),
											),
									),
								sp.TabPanel().
									Value("tab3").
									Body(
										app.Div().
											Style("padding", "20px").
											Style("border", "1px solid var(--spectrum-global-color-gray-300)").
											Style("border-top", "none").
											Body(
												app.H3().Text("Settings Content"),
												app.P().Text("This is the content for the Settings tab. It appears when the Settings tab is selected."),
											),
									),
							),
					),
			),

		// Tab Sizes
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Tab Sizes"),
				app.P().Text("Tabs come in different sizes: S, M (default), L, and XL."),

				app.Div().
					Style("display", "flex").
					Style("flex-direction", "column").
					Style("gap", "30px").
					Style("margin", "20px 0").
					Body(
						// Small Tabs
						app.Div().
							Body(
								app.H3().Text("Small (S)"),
								sp.Tabs().
									Size("s").
									Selected("small1").
									Body(
										sp.Tab().
											Label("Tab 1").
											Value("small1"),
										sp.Tab().
											Label("Tab 2").
											Value("small2"),
										sp.Tab().
											Label("Tab 3").
											Value("small3"),
									),
							),

						// Medium Tabs (Default)
						app.Div().
							Body(
								app.H3().Text("Medium (M) - Default"),
								sp.Tabs().
									Size("m").
									Selected("medium1").
									Body(
										sp.Tab().
											Label("Tab 1").
											Value("medium1"),
										sp.Tab().
											Label("Tab 2").
											Value("medium2"),
										sp.Tab().
											Label("Tab 3").
											Value("medium3"),
									),
							),

						// Large Tabs
						app.Div().
							Body(
								app.H3().Text("Large (L)"),
								sp.Tabs().
									Size("l").
									Selected("large1").
									Body(
										sp.Tab().
											Label("Tab 1").
											Value("large1"),
										sp.Tab().
											Label("Tab 2").
											Value("large2"),
										sp.Tab().
											Label("Tab 3").
											Value("large3"),
									),
							),

						// Extra Large Tabs
						app.Div().
							Body(
								app.H3().Text("Extra Large (XL)"),
								sp.Tabs().
									Size("xl").
									Selected("xl1").
									Body(
										sp.Tab().
											Label("Tab 1").
											Value("xl1"),
										sp.Tab().
											Label("Tab 2").
											Value("xl2"),
										sp.Tab().
											Label("Tab 3").
											Value("xl3"),
									),
							),
					),
			),

		// Tab Directions
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Tab Directions"),
				app.P().Text("Tabs can be displayed in horizontal (default) or vertical orientations."),

				app.Div().
					Style("display", "flex").
					Style("flex-direction", "column").
					Style("gap", "40px").
					Style("margin", "20px 0").
					Body(
						// Horizontal Tabs (Default)
						app.Div().
							Body(
								app.H3().Text("Horizontal (Default)"),
								sp.Tabs().
									Direction("horizontal").
									Selected("hor1").
									Body(
										sp.Tab().
											Label("Tab 1").
											Value("hor1"),
										sp.Tab().
											Label("Tab 2").
											Value("hor2"),
										sp.Tab().
											Label("Tab 3").
											Value("hor3"),
									),
							),

						// Vertical Tabs
						app.Div().
							Body(
								app.H3().Text("Vertical"),
								app.Div().
									Style("display", "flex").
									Body(
										sp.Tabs().
											Direction("vertical").
											Selected("ver1").
											Style("width", "150px").
											Body(
												sp.Tab().
													Label("Tab 1").
													Value("ver1"),
												sp.Tab().
													Label("Tab 2").
													Value("ver2"),
												sp.Tab().
													Label("Tab 3").
													Value("ver3"),
											),
										app.Div().
											Style("padding", "20px").
											Style("border", "1px solid var(--spectrum-global-color-gray-300)").
											Style("width", "100%").
											Text("Tab content would appear here"),
									),
							),

						// Vertical Right Tabs
						app.Div().
							Body(
								app.H3().Text("Vertical Right"),
								app.Div().
									Style("display", "flex").
									Body(
										app.Div().
											Style("padding", "20px").
											Style("border", "1px solid var(--spectrum-global-color-gray-300)").
											Style("width", "100%").
											Text("Tab content would appear here"),
										sp.Tabs().
											Direction("vertical-right").
											Selected("verr1").
											Style("width", "150px").
											Body(
												sp.Tab().
													Label("Tab 1").
													Value("verr1"),
												sp.Tab().
													Label("Tab 2").
													Value("verr2"),
												sp.Tab().
													Label("Tab 3").
													Value("verr3"),
											),
									),
							),
					),
			),

		// Tab Variations
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Tab Variations"),
				app.P().Text("Tabs can be displayed with different visual styles."),

				app.Div().
					Style("display", "flex").
					Style("flex-direction", "column").
					Style("gap", "40px").
					Style("margin", "20px 0").
					Body(
						// Quiet Tabs
						app.Div().
							Body(
								app.H3().Text("Quiet Tabs"),
								app.P().Text("Tabs without a border."),
								sp.Tabs().
									Quiet(true).
									Selected("quiet1").
									Body(
										sp.Tab().
											Label("Tab 1").
											Value("quiet1"),
										sp.Tab().
											Label("Tab 2").
											Value("quiet2"),
										sp.Tab().
											Label("Tab 3").
											Value("quiet3"),
									),
							),

						// Compact Tabs
						app.Div().
							Body(
								app.H3().Text("Compact Tabs"),
								app.P().Text("Tabs with reduced spacing."),
								sp.Tabs().
									Compact(true).
									Selected("compact1").
									Body(
										sp.Tab().
											Label("Tab 1").
											Value("compact1"),
										sp.Tab().
											Label("Tab 2").
											Value("compact2"),
										sp.Tab().
											Label("Tab 3").
											Value("compact3"),
									),
							),

						// Emphasized Tabs
						app.Div().
							Body(
								app.H3().Text("Emphasized Tabs"),
								app.P().Text("Tabs with emphasized styling."),
								sp.Tabs().
									Emphasized(true).
									Selected("emph1").
									Body(
										sp.Tab().
											Label("Tab 1").
											Value("emph1"),
										sp.Tab().
											Label("Tab 2").
											Value("emph2"),
										sp.Tab().
											Label("Tab 3").
											Value("emph3"),
									),
							),
					),
			),

		// Tabs with Icons
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Tabs with Icons"),
				app.P().Text("Tabs can include icons alongside labels."),

				app.Div().
					Style("display", "flex").
					Style("flex-direction", "column").
					Style("gap", "20px").
					Style("margin", "20px 0").
					Body(
						sp.Tabs().
							Selected("icon1").
							Body(
								sp.Tab().
									Label("Home").
									Value("icon1").
									Body(
										app.Elem("div").
											Attr("slot", "icon").
											Body(
												sp.Icon().Name("home"),
											),
									),
								sp.Tab().
									Label("Profile").
									Value("icon2").
									Body(
										app.Elem("div").
											Attr("slot", "icon").
											Body(
												sp.Icon().Name("user"),
											),
									),
								sp.Tab().
									Label("Settings").
									Value("icon3").
									Body(
										app.Elem("div").
											Attr("slot", "icon").
											Body(
												sp.Icon().Name("settings"),
											),
									),
							),
					),
			),

		// Disabled Tabs
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Disabled Tabs"),
				app.P().Text("Individual tabs can be disabled to prevent selection."),

				app.Div().
					Style("display", "flex").
					Style("flex-direction", "column").
					Style("gap", "20px").
					Style("margin", "20px 0").
					Body(
						sp.Tabs().
							Selected("enabled1").
							Body(
								sp.Tab().
									Label("Enabled Tab").
									Value("enabled1"),
								sp.Tab().
									Label("Disabled Tab").
									Value("disabled1").
									Disabled(true),
								sp.Tab().
									Label("Another Enabled Tab").
									Value("enabled2"),
							),
					),
			),

		// Usage Guidelines
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Usage Guidelines"),
				app.P().Text("When to use tabs:"),
				app.Ul().
					Body(
						app.Li().Text("Use tabs to organize content into logical, related sections."),
						app.Li().Text("Use tabs when users need to switch between different views while remaining in the same context."),
						app.Li().Text("Use tabs for content that doesn't need to be viewed simultaneously."),
					),
				app.P().Text("Best practices:"),
				app.Ul().
					Body(
						app.Li().Text("Keep tab labels short and descriptive."),
						app.Li().Text("Use horizontal tabs for most interfaces, vertical tabs for specialized interfaces with many tabs."),
						app.Li().Text("Limit the number of tabs to a reasonable amount (generally 2-7)."),
						app.Li().Text("Consider the appropriate tab size based on your design system and available space."),
					),
			),
	)
}
