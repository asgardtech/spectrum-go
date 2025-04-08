package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// trayPage represents the Tray component showcase
type trayPage struct {
	app.Compo
	isDialogTrayOpen  bool
	isMenuTrayOpen    bool
	isContentTrayOpen bool
}

// newTrayPage creates a new Tray component showcase
func newTrayPage() *trayPage {
	return &trayPage{}
}

// OnNav initializes the page when navigated to
func (p *trayPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

// initPage initializes the page
func (p *trayPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle(defaultTitle)
	ctx.Page().SetDescription(defaultDescription)
}

// toggleDialogTray toggles the dialog tray visibility
func (p *trayPage) toggleDialogTray(ctx app.Context, e app.Event) {
	p.isDialogTrayOpen = !p.isDialogTrayOpen
	ctx.Dispatch(func(ctx app.Context) {})
}

// toggleMenuTray toggles the menu tray visibility
func (p *trayPage) toggleMenuTray(ctx app.Context, e app.Event) {
	p.isMenuTrayOpen = !p.isMenuTrayOpen
	ctx.Dispatch(func(ctx app.Context) {})
}

// toggleContentTray toggles the content tray visibility
func (p *trayPage) toggleContentTray(ctx app.Context, e app.Event) {
	p.isContentTrayOpen = !p.isContentTrayOpen
	ctx.Dispatch(func(ctx app.Context) {})
}

// handleTrayClose handles the close event for trays
func (p *trayPage) handleTrayClose(ctx app.Context, e app.Event) {
	p.isDialogTrayOpen = false
	p.isMenuTrayOpen = false
	p.isContentTrayOpen = false
	ctx.Dispatch(func(ctx app.Context) {})
}

// Render renders the Tray component showcase
func (p *trayPage) Render() app.UI {
	return newPage().Content(
		app.H1().Text("Tray Component"),
		app.P().Text("The Tray component is typically used to portray information on mobile devices or smaller screens. It slides in from the bottom of the viewport."),

		// Tray with Dialog
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Tray with Dialog"),
				app.P().Text("A tray containing a dialog:"),
				app.Div().
					Class("component-example").
					Body(
						sp.Button().
							Text("Toggle Dialog Tray").
							OnClick(p.toggleDialogTray),
						sp.Tray().
							Open(p.isDialogTrayOpen).
							OnClose(p.handleTrayClose).
							Body(
								sp.Dialog().
									Size("s").
									Dismissable(true).
									Heading(
										app.H2().Text("New Messages"),
									).
									Body(
										app.P().Text("You have 5 new messages."),
									),
							),
					),

				app.Pre().
					Class("language-go").
					Text(`sp.Button().
    Text("Toggle Dialog Tray").
    OnClick(toggleDialogTray),
sp.Tray().
    Open(isDialogTrayOpen).
    OnClose(handleTrayClose).
    Body(
        sp.Dialog().
            Size("s").
            Dismissable(true).
            Heading(
                app.H2().Text("New Messages"),
            ).
            Body(
                app.P().Text("You have 5 new messages."),
            ),
    )`),
			),

		// Tray with Menu
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Tray with Menu"),
				app.P().Text("A tray containing a menu:"),
				app.Div().
					Class("component-example").
					Body(
						sp.Button().
							Text("Toggle Menu Tray").
							OnClick(p.toggleMenuTray),
						sp.Tray().
							Open(p.isMenuTrayOpen).
							OnClose(p.handleTrayClose).
							Body(
								sp.Menu().
									Style("width", "100%").
									Body(
										sp.MenuItem().
											Selected(true).
											Text("Deselect"),
										sp.MenuItem().
											Text("Select Inverse"),
										sp.MenuItem().
											Focused(true).
											Text("Feather..."),
										sp.MenuItem().
											Text("Select and Mask..."),
										sp.MenuDivider(),
										sp.MenuItem().
											Text("Save Selection"),
										sp.MenuItem().
											Disabled(true).
											Text("Make Work Path"),
									),
							),
					),

				app.Pre().
					Class("language-go").
					Text(`sp.Button().
    Text("Toggle Menu Tray").
    OnClick(toggleMenuTray),
sp.Tray().
    Open(isMenuTrayOpen).
    OnClose(handleTrayClose).
    Body(
        sp.Menu().
            Style("width", "100%").
            Body(
                sp.MenuItem().
                    Selected(true).
                    Text("Deselect"),
                sp.MenuItem().
                    Text("Select Inverse"),
                sp.MenuItem().
                    Focused(true).
                    Text("Feather..."),
                sp.MenuItem().
                    Text("Select and Mask..."),
                sp.MenuDivider(),
                sp.MenuItem().
                    Text("Save Selection"),
                sp.MenuItem().
                    Disabled(true).
                    Text("Make Work Path"),
            ),
    )`),
			),

		// Tray with Custom Content
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Tray with Custom Content"),
				app.P().Text("A tray with custom content:"),
				app.Div().
					Class("component-example").
					Body(
						sp.Button().
							Text("Toggle Content Tray").
							OnClick(p.toggleContentTray),
						sp.Tray().
							Open(p.isContentTrayOpen).
							OnClose(p.handleTrayClose).
							Body(
								app.Div().
									Style("padding", "20px").
									Style("background", "var(--spectrum-global-color-gray-100)").
									Style("display", "flex").
									Style("flex-direction", "column").
									Style("gap", "12px").
									Body(
										app.H3().
											Style("margin", "0").
											Text("Custom Tray Content"),
										app.P().Text("This is a tray with custom content. You can put any component or HTML element inside a tray."),
										sp.Divider(),
										sp.ActionGroup().
											Body(
												sp.Button().
													Variant("primary").
													Text("Confirm"),
												sp.Button().
													Variant("secondary").
													Text("Cancel").
													OnClick(p.handleTrayClose),
											),
									),
							),
					),

				app.Pre().
					Class("language-go").
					Text(`sp.Button().
    Text("Toggle Content Tray").
    OnClick(toggleContentTray),
sp.Tray().
    Open(isContentTrayOpen).
    OnClose(handleTrayClose).
    Body(
        app.Div().
            Style("padding", "20px").
            Style("background", "var(--spectrum-global-color-gray-100)").
            Style("display", "flex").
            Style("flex-direction", "column").
            Style("gap", "12px").
            Body(
                app.H3().
                    Style("margin", "0").
                    Text("Custom Tray Content"),
                app.P().Text("This is a tray with custom content. You can put any component or HTML element inside a tray."),
                sp.Divider(),
                sp.ActionGroup().
                    Body(
                        sp.Button().
                            Variant("primary").
                            Text("Confirm"),
                        sp.Button().
                            Variant("secondary").
                            Text("Cancel").
                            OnClick(handleTrayClose),
                    ),
            ),
    )`),
			),

		// Usage with Mobile Devices
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Usage with Mobile Devices"),
				app.P().Text("Trays are particularly useful on mobile devices where screen space is limited. They provide a way to display additional content without navigating away from the current view."),
				app.P().Text("Common use cases include:"),
				app.Ul().
					Body(
						app.Li().Text("Displaying menus"),
						app.Li().Text("Showing details or additional information"),
						app.Li().Text("Presenting forms or configuration options"),
						app.Li().Text("Displaying notifications or alerts"),
					),
			),

		// Usage Guidelines
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Usage Guidelines"),
				app.H3().Text("When to use"),
				app.Ul().
					Body(
						app.Li().Text("For mobile or small screen experiences"),
						app.Li().Text("When proximity to the trigger is less important"),
						app.Li().Text("For displaying larger volumes of content that might overwhelm a popover"),
						app.Li().Text("When you need a temporary overlay that slides in from the bottom"),
					),
				app.H3().Text("Best practices"),
				app.Ul().
					Body(
						app.Li().Text("Provide a clear way to dismiss the tray (close button or escape key)"),
						app.Li().Text("Use appropriate sizing for mobile screens"),
						app.Li().Text("Consider the tray's content hierarchy and organization"),
						app.Li().Text("Use with dialogs, menus, or custom content as needed"),
						app.Li().Text("Consider accessibility in your implementation"),
					),
			),
	)
}
