package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/asgardtech/spectrum-go/prism"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// actionBarPage represents the Action Bar component showcase
type actionBarPage struct {
	app.Compo
	isOpen bool
}

// newActionBarPage creates a new Action Bar component showcase
func newActionBarPage() *actionBarPage {
	return &actionBarPage{}
}

// OnNav initializes the page when navigated to
func (p *actionBarPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

// initPage initializes the page
func (p *actionBarPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle(defaultTitle)
	ctx.Page().SetDescription(defaultDescription)
}

// toggleActionBar toggles the visibility of the action bar
func (p *actionBarPage) toggleActionBar(ctx app.Context, e app.Event) {
	p.isOpen = !p.isOpen
	ctx.Dispatch(func(ctx app.Context) {})
}

// Render renders the Action Bar component showcase
func (p *actionBarPage) Render() app.UI {
	return prism.NewLayout().Content(
		app.H1().Text("Action Bar Component"),
		app.P().Text("A floating action bar that is a convenient way to deliver stateful actions in cases like selection mode. It can be deployed with fixed or sticky variants."),

		// Basic Action Bar
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Basic Action Bar"),
				app.P().Text("A simple action bar with text and buttons:"),
				app.Div().
					Class("component-example").
					Body(
						sp.Button().
							Text("Toggle Action Bar").
							OnClick(p.toggleActionBar),
						sp.ActionBar().
							Open(p.isOpen).
							Text("2 selected").
							Buttons(
								app.Div().Body(
									sp.ActionButton().
										Label("Edit").
										Icon(
											app.Elem("div").
												Attr("slot", "icon").
												Body(
													sp.Icon().Name("edit"),
												),
										),
									sp.ActionButton().
										Label("More").
										Icon(
											app.Elem("div").
												Attr("slot", "icon").
												Body(
													sp.Icon().Name("more"),
												),
										),
								),
							),
					),

				app.Pre().
					Class("language-go").
					Text(`sp.ActionBar().
    Open(true).
    Text("2 selected").
    Buttons(
        app.Div().Body(
            sp.ActionButton().
                Label("Edit").
                Icon(
                    app.Elem("div").
                        Attr("slot", "icon").
                        Body(
                            sp.Icon().Name("edit"),
                        ),
                ),
            sp.ActionButton().
                Label("More").
                Icon(
                    app.Elem("div").
                        Attr("slot", "icon").
                        Body(
                            sp.Icon().Name("more"),
                        ),
                ),
        ),
    )`),
			),

		// Emphasized Action Bar
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Emphasized Action Bar"),
				app.P().Text("An action bar with additional visual emphasis:"),
				app.Div().
					Class("component-example").
					Body(
						sp.Button().
							Text("Toggle Emphasized Action Bar").
							OnClick(p.toggleActionBar),
						sp.ActionBar().
							Open(p.isOpen).
							Emphasized(true).
							Text("2 selected").
							Buttons(
								app.Div().Body(
									sp.ActionButton().
										Label("Edit").
										Icon(
											app.Elem("div").
												Attr("slot", "icon").
												Body(
													sp.Icon().Name("edit"),
												),
										),
									sp.ActionButton().
										Label("More").
										Icon(
											app.Elem("div").
												Attr("slot", "icon").
												Body(
													sp.Icon().Name("more"),
												),
										),
								),
							),
					),

				app.Pre().
					Class("language-go").
					Text(`sp.ActionBar().
    Open(true).
    Emphasized(true).
    Text("2 selected").
    Buttons(
        app.Div().Body(
            sp.ActionButton().
                Label("Edit").
                Icon(
                    app.Elem("div").
                        Attr("slot", "icon").
                        Body(
                            sp.Icon().Name("edit"),
                        ),
                ),
            sp.ActionButton().
                Label("More").
                Icon(
                    app.Elem("div").
                        Attr("slot", "icon").
                        Body(
                            sp.Icon().Name("more"),
                        ),
                ),
        ),
    )`),
			),

		// Fixed Action Bar
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Fixed Action Bar"),
				app.P().Text("An action bar positioned fixed relative to the page:"),
				app.Div().
					Class("component-example").
					Body(
						sp.Button().
							Text("Toggle Fixed Action Bar").
							OnClick(p.toggleActionBar),
						sp.ActionBar().
							Open(p.isOpen).
							VariantFixed().
							Text("2 selected").
							Buttons(
								app.Div().Body(
									sp.ActionButton().
										Label("Edit").
										Icon(
											app.Elem("div").
												Attr("slot", "icon").
												Body(
													sp.Icon().Name("edit"),
												),
										),
								),
							),
					),

				app.Pre().
					Class("language-go").
					Text(`sp.ActionBar().
    Open(true).
    VariantFixed().
    Text("2 selected").
    Buttons(
        app.Div().Body(
            sp.ActionButton().
                Label("Edit").
                Icon(
                    app.Elem("div").
                        Attr("slot", "icon").
                        Body(
                            sp.Icon().Name("edit"),
                        ),
                ),
        ),
    )`),
			),

		// Sticky Action Bar
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Sticky Action Bar"),
				app.P().Text("An action bar positioned sticky relative to scrolling content:"),
				app.Div().
					Class("component-example").
					Style("position", "relative").
					Style("max-height", "200px").
					Style("overflow", "auto").
					Body(
						app.H4().Text("Scroll down for toggle button"),
						app.P().Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."),
						sp.Button().
							Text("Toggle Sticky Action Bar").
							OnClick(p.toggleActionBar),
						sp.ActionBar().
							Open(p.isOpen).
							VariantSticky().
							Style("inset-block", "0px").
							Text("2 selected").
							Buttons(
								app.Div().Body(
									sp.ActionButton().
										Label("Edit").
										Icon(
											app.Elem("div").
												Attr("slot", "icon").
												Body(
													sp.Icon().Name("edit"),
												),
										),
									sp.ActionButton().
										Label("More").
										Icon(
											app.Elem("div").
												Attr("slot", "icon").
												Body(
													sp.Icon().Name("more"),
												),
										),
								),
							),
					),

				app.Pre().
					Class("language-go").
					Text(`sp.ActionBar().
    Open(true).
    VariantSticky().
    Style("inset-block", "0px").
    Text("2 selected").
    Buttons(
        app.Div().Body(
            sp.ActionButton().
                Label("Edit").
                Icon(
                    app.Elem("div").
                        Attr("slot", "icon").
                        Body(
                            sp.Icon().Name("edit"),
                        ),
                ),
            sp.ActionButton().
                Label("More").
                Icon(
                    app.Elem("div").
                        Attr("slot", "icon").
                        Body(
                            sp.Icon().Name("more"),
                        ),
                ),
        ),
    )`),
			),

		// Flexible Action Bar
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Flexible Action Bar"),
				app.P().Text("An action bar that sizes itself to its content:"),
				app.Div().
					Class("component-example").
					Body(
						sp.Button().
							Text("Toggle Flexible Action Bar").
							OnClick(p.toggleActionBar),
						sp.ActionBar().
							Open(p.isOpen).
							Flexible(true).
							Text("2 selected").
							Buttons(
								app.Div().Body(
									sp.ActionButton().
										Label("Edit").
										Icon(
											app.Elem("div").
												Attr("slot", "icon").
												Body(
													sp.Icon().Name("edit"),
												),
										),
									sp.ActionButton().
										Label("More").
										Icon(
											app.Elem("div").
												Attr("slot", "icon").
												Body(
													sp.Icon().Name("more"),
												),
										),
								),
							),
					),

				app.Pre().
					Class("language-go").
					Text(`sp.ActionBar().
    Open(true).
    Flexible(true).
    Text("2 selected").
    Buttons(
        app.Div().Body(
            sp.ActionButton().
                Label("Edit").
                Icon(
                    app.Elem("div").
                        Attr("slot", "icon").
                        Body(
                            sp.Icon().Name("edit"),
                        ),
                ),
            sp.ActionButton().
                Label("More").
                Icon(
                    app.Elem("div").
                        Attr("slot", "icon").
                        Body(
                            sp.Icon().Name("more"),
                        ),
                ),
        ),
    )`),
			),

		// Usage Guidelines
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Usage Guidelines"),
				app.H3().Text("When to use"),
				app.Ul().
					Body(
						app.Li().Text("For displaying actions related to selected items"),
						app.Li().Text("When you need a floating action bar that appears in response to selection"),
						app.Li().Text("For providing contextual actions in a selection mode"),
					),
				app.H3().Text("Best practices"),
				app.Ul().
					Body(
						app.Li().Text("Keep the action bar visible while items are selected"),
						app.Li().Text("Use clear and concise labels for actions"),
						app.Li().Text("Group related actions together"),
						app.Li().Text("Consider using the emphasized variant for important actions"),
						app.Li().Text("Choose the appropriate variant (fixed/sticky) based on your layout needs"),
					),
			),
	)
}
