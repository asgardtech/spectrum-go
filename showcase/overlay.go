package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/asgardtech/spectrum-go/prism"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// overlayPage represents the Overlay component showcase
type overlayPage struct {
	app.Compo
	isBasicOpen     bool
	isModalOpen     bool
	isPlacementOpen bool
	isCustomOpen    bool
}

// newOverlayPage creates a new Overlay component showcase
func newOverlayPage() *overlayPage {
	return &overlayPage{}
}

// OnNav initializes the page when navigated to
func (p *overlayPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

// initPage initializes the page
func (p *overlayPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle(defaultTitle)
	ctx.Page().SetDescription(defaultDescription)
}

// toggleBasicOverlay toggles the basic overlay
func (p *overlayPage) toggleBasicOverlay(ctx app.Context, e app.Event) {
	p.isBasicOpen = !p.isBasicOpen
	ctx.Dispatch(func(ctx app.Context) {})
}

// toggleModalOverlay toggles the modal overlay
func (p *overlayPage) toggleModalOverlay(ctx app.Context, e app.Event) {
	p.isModalOpen = !p.isModalOpen
	ctx.Dispatch(func(ctx app.Context) {})
}

// togglePlacementOverlay toggles the placement overlay
func (p *overlayPage) togglePlacementOverlay(ctx app.Context, e app.Event) {
	p.isPlacementOpen = !p.isPlacementOpen
	ctx.Dispatch(func(ctx app.Context) {})
}

// toggleCustomOverlay toggles the custom overlay
func (p *overlayPage) toggleCustomOverlay(ctx app.Context, e app.Event) {
	p.isCustomOpen = !p.isCustomOpen
	ctx.Dispatch(func(ctx app.Context) {})
}

// handleOverlayClosed handles when an overlay is closed
func (p *overlayPage) handleOverlayClosed(ctx app.Context, e app.Event) {
	p.isBasicOpen = false
	p.isModalOpen = false
	p.isPlacementOpen = false
	p.isCustomOpen = false
	ctx.Dispatch(func(ctx app.Context) {})
}

// Render renders the Overlay component showcase
func (p *overlayPage) Render() app.UI {
	return prism.NewPage().Content(
		app.H1().Text("Overlay Component"),
		app.P().Text("The Overlay component is used to decorate content that you would like to present to your visitors as 'overlaid' on the rest of the application, such as dialogs, pickers, tooltips, and context menus."),

		// Basic Overlay
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Basic Overlay"),
				app.P().Text("A simple overlay that shows content on top of the page:"),
				app.Div().
					Class("component-example").
					Body(
						app.Div().
							Style("position", "relative").
							Style("min-height", "200px").
							Body(
								sp.Button().
									Id("basic-trigger").
									Text("Toggle Basic Overlay").
									OnClick(p.toggleBasicOverlay),
								sp.Overlay().
									Open(p.isBasicOpen).
									OnSpClosed(p.handleOverlayClosed).
									Body(
										app.Div().
											Style("background", "var(--spectrum-global-color-gray-50)").
											Style("padding", "20px").
											Style("border-radius", "4px").
											Style("box-shadow", "0 4px 16px rgba(0,0,0,0.2)").
											Body(
												app.H3().Text("Basic Overlay Content"),
												app.P().Text("This is a basic overlay with custom content."),
												sp.Button().
													Text("Close").
													OnClick(p.handleOverlayClosed),
											),
									),
							),
					),

				app.Pre().
					Class("language-go").
					Text(`sp.Button().
    Id("basic-trigger").
    Text("Toggle Basic Overlay").
    OnClick(toggleBasicOverlay),
sp.Overlay().
    Open(isBasicOpen).
    OnSpClosed(handleOverlayClosed).
    Body(
        app.Div().
            Style("background", "var(--spectrum-global-color-gray-50)").
            Style("padding", "20px").
            Style("border-radius", "4px").
            Style("box-shadow", "0 4px 16px rgba(0,0,0,0.2)").
            Body(
                app.H3().Text("Basic Overlay Content"),
                app.P().Text("This is a basic overlay with custom content."),
                sp.Button().
                    Text("Close").
                    OnClick(handleOverlayClosed),
            ),
    )`),
			),

		// Modal Overlay
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Modal Overlay"),
				app.P().Text("A modal overlay that requires user interaction before dismissal:"),
				app.Div().
					Class("component-example").
					Body(
						app.Div().
							Style("position", "relative").
							Style("min-height", "200px").
							Body(
								sp.Button().
									Id("modal-trigger").
									Text("Toggle Modal Overlay").
									OnClick(p.toggleModalOverlay),
								sp.Overlay().
									Open(p.isModalOpen).
									Type(sp.OverlayTypeModal).
									OnSpClosed(p.handleOverlayClosed).
									Body(
										sp.Dialog().
											Size("s").
											Dismissable(true).
											Heading(
												app.H2().Text("Sign In Form"),
											).
											Body(
												app.P().Text("Please enter your email to sign in:"),
												sp.FieldLabel().
													For("email-input").
													Text("Email"),
												sp.Textfield().
													Id("email-input").
													Placeholder("example@email.com"),
												sp.ActionButton().
													Style("margin-top", "20px").
													Label("Sign In").
													OnClick(p.handleOverlayClosed),
											),
									),
							),
					),

				app.Pre().
					Class("language-go").
					Text(`sp.Button().
    Id("modal-trigger").
    Text("Toggle Modal Overlay").
    OnClick(toggleModalOverlay),
sp.Overlay().
    Open(isModalOpen).
    Type(sp.OverlayTypeModal).
    OnSpClosed(handleOverlayClosed).
    Body(
        sp.Dialog().
            Size("s").
            Dismissable(true).
            Heading(
                app.H2().Text("Sign In Form"),
            ).
            Body(
                app.P().Text("Please enter your email to sign in:"),
                sp.FieldLabel().
                    For("email-input").
                    Text("Email"),
                sp.Textfield().
                    Id("email-input").
                    Placeholder("example@email.com"),
                sp.ActionButton().
                    Style("margin-top", "20px").
                    Label("Sign In").
                    OnClick(handleOverlayClosed),
            ),
    )`),
			),

		// Overlay with Placement
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Overlay with Placement"),
				app.P().Text("An overlay positioned relative to its trigger:"),
				app.Div().
					Class("component-example").
					Body(
						app.Div().
							Style("position", "relative").
							Style("min-height", "200px").
							Style("display", "flex").
							Style("justify-content", "center").
							Style("align-items", "center").
							Body(
								sp.Button().
									Id("placement-trigger").
									Text("Toggle Positioned Overlay").
									OnClick(p.togglePlacementOverlay),
								sp.Overlay().
									Open(p.isPlacementOpen).
									Placement("top").
									Offset(10).
									OnSpClosed(p.handleOverlayClosed).
									Body(
										app.Div().
											Style("background", "var(--spectrum-global-color-gray-50)").
											Style("padding", "15px").
											Style("border-radius", "4px").
											Style("box-shadow", "0 4px 16px rgba(0,0,0,0.2)").
											Body(
												app.P().
													Style("margin", "0").
													Text("This overlay is positioned above the trigger with an offset."),
											),
									),
							),
					),

				app.Pre().
					Class("language-go").
					Text(`sp.Button().
    Id("placement-trigger").
    Text("Toggle Positioned Overlay").
    OnClick(togglePlacementOverlay),
sp.Overlay().
    Open(isPlacementOpen).
    Placement("top").
    Offset(10).
    OnSpClosed(handleOverlayClosed).
    Body(
        app.Div().
            Style("background", "var(--spectrum-global-color-gray-50)").
            Style("padding", "15px").
            Style("border-radius", "4px").
            Style("box-shadow", "0 4px 16px rgba(0,0,0,0.2)").
            Body(
                app.P().
                    Style("margin", "0").
                    Text("This overlay is positioned above the trigger with an offset."),
            ),
    )`),
			),

		// Custom Overlay with Menu
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Overlay with Menu"),
				app.P().Text("An overlay containing a menu:"),
				app.Div().
					Class("component-example").
					Body(
						app.Div().
							Style("position", "relative").
							Style("min-height", "250px").
							Body(
								sp.Button().
									Id("custom-trigger").
									Text("Toggle Menu Overlay").
									OnClick(p.toggleCustomOverlay),
								sp.Overlay().
									Open(p.isCustomOpen).
									Placement("bottom-start").
									OnSpClosed(p.handleOverlayClosed).
									Body(
										sp.Menu().
											Style("min-width", "200px").
											Body(
												sp.MenuItem().
													Text("Edit").
													Icon(
														app.Elem("div").
															Attr("slot", "icon").
															Body(
																sp.Icon().Name("edit"),
															),
													),
												sp.MenuItem().
													Text("Copy").
													Icon(
														app.Elem("div").
															Attr("slot", "icon").
															Body(
																sp.Icon().Name("copy"),
															),
													),
												sp.MenuItem().
													Text("Delete").
													Icon(
														app.Elem("div").
															Attr("slot", "icon").
															Body(
																sp.Icon().Name("delete"),
															),
													),
												sp.MenuDivider(),
												sp.MenuItem().
													Text("Share").
													Icon(
														app.Elem("div").
															Attr("slot", "icon").
															Body(
																sp.Icon().Name("share"),
															),
													),
											),
									),
							),
					),

				app.Pre().
					Class("language-go").
					Text(`sp.Button().
    Id("custom-trigger").
    Text("Toggle Menu Overlay").
    OnClick(toggleCustomOverlay),
sp.Overlay().
    Open(isCustomOpen).
    Placement("bottom-start").
    OnSpClosed(handleOverlayClosed).
    Body(
        sp.Menu().
            Style("min-width", "200px").
            Body(
                sp.MenuItem().
                    Text("Edit").
                    Icon(
                        app.Elem("div").
                            Attr("slot", "icon").
                            Body(
                                sp.Icon().Name("edit"),
                            ),
                    ),
                sp.MenuItem().
                    Text("Copy").
                    Icon(
                        app.Elem("div").
                            Attr("slot", "icon").
                            Body(
                                sp.Icon().Name("copy"),
                            ),
                    ),
                sp.MenuItem().
                    Text("Delete").
                    Icon(
                        app.Elem("div").
                            Attr("slot", "icon").
                            Body(
                                sp.Icon().Name("delete"),
                            ),
                    ),
                sp.MenuDivider(),
                sp.MenuItem().
                    Text("Share").
                    Icon(
                        app.Elem("div").
                            Attr("slot", "icon").
                            Body(
                                sp.Icon().Name("share"),
                            ),
                    ),
            ),
    )`),
			),

		// Overlay Types
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Overlay Types"),
				app.P().Text("The Overlay component supports several types that determine its behavior:"),
				app.Ul().
					Body(
						app.Li().
							Body(
								app.Strong().Text("Auto: "),
								app.Text("The default type. The overlay will adapt its behavior based on its content."),
							),
						app.Li().
							Body(
								app.Strong().Text("Modal: "),
								app.Text("Used when you need to ensure the user interacts with the overlay before continuing. Focus is trapped inside the overlay while it's open."),
							),
						app.Li().
							Body(
								app.Strong().Text("Page: "),
								app.Text("Similar to modal, but cannot be closed via the Escape key. Useful for full-screen menus or important interactions."),
							),
						app.Li().
							Body(
								app.Strong().Text("Hint: "),
								app.Text("Like tooltips, these exist outside the tab order and should not contain interactive content."),
							),
						app.Li().
							Body(
								app.Strong().Text("Manual: "),
								app.Text("For custom overlay behavior that doesn't fit the other types."),
							),
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
						app.Li().Text("For dialogs, popovers, tooltips, and context menus"),
						app.Li().Text("When you need to display content that overlays the rest of the application"),
						app.Li().Text("For creating modal experiences that require user interaction"),
						app.Li().Text("When positioning content relative to a trigger element"),
					),
				app.H3().Text("Best practices"),
				app.Ul().
					Body(
						app.Li().Text("Choose the appropriate overlay type based on your interaction requirements"),
						app.Li().Text("Use placement to position overlays relative to their triggers"),
						app.Li().Text("Provide clear ways to dismiss the overlay when appropriate"),
						app.Li().Text("Consider accessibility in your overlay implementation"),
						app.Li().Text("Use offset to create appropriate spacing between the overlay and its trigger"),
					),
			),
	)
}
