package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/asgardtech/spectrum-go/prism"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// topNavPage represents the Top Nav component showcase
type topNavPage struct {
	app.Compo
}

// newTopNavPage creates a new Top Nav component showcase
func newTopNavPage() *topNavPage {
	return &topNavPage{}
}

// OnNav initializes the page when navigated to
func (p *topNavPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

// initPage initializes the page
func (p *topNavPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle(defaultTitle)
	ctx.Page().SetDescription(defaultDescription)
}

// Render renders the Top Nav component showcase
func (p *topNavPage) Render() app.UI {
	return prism.NewPage().Content(
		app.H1().Text("Top Nav Component"),
		app.P().Text("The Top Nav component delivers site navigation, particularly for when that navigation will change the majority of the page's content and/or the page's URL when selected. All primary elements are directly accessible in the tab order."),

		// Basic Top Nav
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Basic Top Nav"),
				app.P().Text("A standard top navigation bar with items:"),
				app.Div().
					Class("component-example").
					Body(
						sp.TopNav().
							Label("Basic Navigation").
							Body(
								sp.TopNavItem().
									Label("Site Name").
									Href("#").
									Text("Site Name"),
								sp.TopNavItem().
									Label("Page 1").
									Href("#page-1").
									Style("margin-inline-start", "auto").
									Text("Page 1"),
								sp.TopNavItem().
									Label("Page 2").
									Href("#page-2").
									Text("Page 2"),
								sp.TopNavItem().
									Label("Page 3").
									Href("#page-3").
									Text("Page 3"),
							),
					),

				app.Pre().
					Class("language-go").
					Text(`sp.TopNav().
    Label("Basic Navigation").
    Body(
        sp.TopNavItem().
            Label("Site Name").
            Href("#").
            Text("Site Name"),
        sp.TopNavItem().
            Label("Page 1").
            Href("#page-1").
            Style("margin-inline-start", "auto").
            Text("Page 1"),
        sp.TopNavItem().
            Label("Page 2").
            Href("#page-2").
            Text("Page 2"),
        sp.TopNavItem().
            Label("Page 3").
            Href("#page-3").
            Text("Page 3"),
    )`),
			),

		// Top Nav with Action Menu
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Top Nav with Action Menu"),
				app.P().Text("A top navigation bar with navigation items and an action menu:"),
				app.Div().
					Class("component-example").
					Body(
						sp.TopNav().
							Label("Navigation with Action Menu").
							Body(
								sp.TopNavItem().
									Label("Site Name").
									Href("#").
									Text("Site Name"),
								sp.TopNavItem().
									Label("Page 1").
									Href("#page-1").
									Text("Page 1"),
								sp.TopNavItem().
									Label("Page 2").
									Href("#page-2").
									Text("Page 2"),
								sp.ActionMenu().
									Label("Account").
									Placement("bottom-end").
									Style("margin-inline-start", "auto").
									Quiet(true).
									Body(
										sp.MenuItem().
											Label("Account Settings").
											Text("Account Settings"),
										sp.MenuItem().
											Label("My Profile").
											Text("My Profile"),
										sp.MenuDivider(),
										sp.MenuItem().
											Label("Sign Out").
											Text("Sign Out"),
									),
							),
					),

				app.Pre().
					Class("language-go").
					Text(`sp.TopNav().
    Label("Navigation with Action Menu").
    Body(
        sp.TopNavItem().
            Label("Site Name").
            Href("#").
            Text("Site Name"),
        sp.TopNavItem().
            Label("Page 1").
            Href("#page-1").
            Text("Page 1"),
        sp.TopNavItem().
            Label("Page 2").
            Href("#page-2").
            Text("Page 2"),
        sp.ActionMenu().
            Label("Account").
            Placement("bottom-end").
            Style("margin-inline-start", "auto").
            Quiet(true).
            Body(
                sp.MenuItem().
                    Label("Account Settings").
                    Text("Account Settings"),
                sp.MenuItem().
                    Label("My Profile").
                    Text("My Profile"),
                sp.MenuDivider(),
                sp.MenuItem().
                    Label("Sign Out").
                    Text("Sign Out"),
            ),
    )`),
			),

		// Quiet Top Nav
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Quiet Top Nav"),
				app.P().Text("A top navigation bar without a border:"),
				app.Div().
					Class("component-example").
					Body(
						sp.TopNav().
							Label("Quiet Navigation").
							Quiet(true).
							Body(
								sp.TopNavItem().
									Label("Site Name").
									Href("#").
									Text("Site Name"),
								sp.TopNavItem().
									Label("Page 1").
									Href("#page-1").
									Style("margin-inline-start", "auto").
									Text("Page 1"),
								sp.TopNavItem().
									Label("Page 2").
									Href("#page-2").
									Text("Page 2"),
							),
					),

				app.Pre().
					Class("language-go").
					Text(`sp.TopNav().
    Label("Quiet Navigation").
    Quiet(true).
    Body(
        sp.TopNavItem().
            Label("Site Name").
            Href("#").
            Text("Site Name"),
        sp.TopNavItem().
            Label("Page 1").
            Href("#page-1").
            Style("margin-inline-start", "auto").
            Text("Page 1"),
        sp.TopNavItem().
            Label("Page 2").
            Href("#page-2").
            Text("Page 2"),
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
						app.Li().Text("For primary site navigation that affects the majority of the page content"),
						app.Li().Text("When navigation items need to be directly accessible in the tab order"),
						app.Li().Text("To provide consistent navigation across different sections of your site"),
					),
				app.H3().Text("Best practices"),
				app.Ul().
					Body(
						app.Li().Text("Keep navigation items clear and concise"),
						app.Li().Text("Use meaningful labels that describe the destination"),
						app.Li().Text("Consider using the quiet variant when you want a more subtle navigation bar"),
						app.Li().Text("Place action menus on the right side of the navigation bar"),
						app.Li().Text("Use appropriate spacing between navigation items"),
					),
			),
	)
}
