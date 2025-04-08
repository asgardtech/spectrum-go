package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// overlayTriggerPage represents the Overlay Trigger component showcase
type overlayTriggerPage struct {
	app.Compo
}

// newOverlayTriggerPage creates a new Overlay Trigger component showcase
func newOverlayTriggerPage() *overlayTriggerPage {
	return &overlayTriggerPage{}
}

// OnNav initializes the page when navigated to
func (p *overlayTriggerPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

// initPage initializes the page
func (p *overlayTriggerPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle(defaultTitle)
	ctx.Page().SetDescription(defaultDescription)
}

// handleOverlayOpened handles when an overlay is opened
func (p *overlayTriggerPage) handleOverlayOpened(ctx app.Context, e app.Event) {
	app.Log("Overlay opened")
}

// handleOverlayClosed handles when an overlay is closed
func (p *overlayTriggerPage) handleOverlayClosed(ctx app.Context, e app.Event) {
	app.Log("Overlay closed")
}

// Render renders the Overlay Trigger component showcase
func (p *overlayTriggerPage) Render() app.UI {
	return newPage().Content(
		app.H1().Text("Overlay Trigger Component"),
		app.P().Text("The Overlay Trigger component supports the delivery of temporary overlay content based on interaction with a persistent trigger element."),

		// Basic Overlay Trigger with Click Content
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Basic Overlay Trigger with Click Content"),
				app.P().Text("An overlay trigger that displays content when clicked:"),
				app.Div().
					Class("component-example").
					Body(
						sp.OverlayTrigger().
							Placement("bottom").
							OnSpOpened(p.handleOverlayOpened).
							OnSpClosed(p.handleOverlayClosed).
							Trigger(
								sp.Button().
									Variant("primary").
									Text("Click for Popover"),
							).
							ClickContent(
								sp.Popover().
									Placement("bottom").
									Tip(true).
									Body(
										sp.Dialog().
											Nodivider(true).
											Body(
												app.H3().Text("Click Content"),
												app.P().Text("This content appears when you click the button."),
												sp.Button().
													Text("Close").
													OnClick(func(ctx app.Context, e app.Event) {
														// This will trigger the sp-closed event
														e.Get("target").Call("dispatchEvent", app.Window().Get("CustomEvent").New("close", map[string]interface{}{
															"bubbles":  true,
															"composed": true,
														}))
													}),
											),
									),
							),
					),

				app.Pre().
					Class("language-go").
					Text(`sp.OverlayTrigger().
    Placement("bottom").
    OnSpOpened(handleOverlayOpened).
    OnSpClosed(handleOverlayClosed).
    Trigger(
        sp.Button().
            Variant("primary").
            Text("Click for Popover"),
    ).
    ClickContent(
        sp.Popover().
            Placement("bottom").
            Tip(true).
            Body(
                sp.Dialog().
                    Nodivider(true).
                    Body(
                        app.H3().Text("Click Content"),
                        app.P().Text("This content appears when you click the button."),
                        sp.Button().
                            Text("Close").
                            OnClick(func(ctx app.Context, e app.Event) {
                                // This will trigger the sp-closed event
                                e.Get("target").Call("dispatchEvent", app.Window().Get("CustomEvent").New("close", map[string]interface{}{
                                    "bubbles":  true,
                                    "composed": true,
                                }))
                            }),
                    ),
            ),
    )`),
			),

		// Overlay Trigger with Hover Content
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Overlay Trigger with Hover Content"),
				app.P().Text("An overlay trigger that displays content when hovered:"),
				app.Div().
					Class("component-example").
					Body(
						sp.OverlayTrigger().
							Placement("top").
							Trigger(
								sp.Button().
									Text("Hover for Tooltip"),
							).
							HoverContent(
								sp.Tooltip().
									Placement("top").
									Body(
										app.Text("This is a tooltip that appears on hover"),
									),
							),
					),

				app.Pre().
					Class("language-go").
					Text(`sp.OverlayTrigger().
    Placement("top").
    Trigger(
        sp.Button().
            Text("Hover for Tooltip"),
    ).
    HoverContent(
        sp.Tooltip().
            Placement("top").
            Body(
                app.Text("This is a tooltip that appears on hover"),
            ),
    )`),
			),

		// Overlay Trigger with Multiple Interaction Types
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Overlay Trigger with Multiple Interaction Types"),
				app.P().Text("An overlay trigger that responds to both click and hover:"),
				app.Div().
					Class("component-example").
					Body(
						sp.OverlayTrigger().
							Placement("right").
							Triggeredby("click hover").
							Trigger(
								sp.Button().
									Text("Hover and Click"),
							).
							ClickContent(
								sp.Popover().
									Placement("right").
									Tip(true).
									Body(
										sp.Dialog().
											Nodivider(true).
											Body(
												app.H3().Text("Click Content"),
												app.P().Text("This appears when you click."),
											),
									),
							).
							HoverContent(
								sp.Tooltip().
									Placement("top").
									Body(
										app.Text("Hover tooltip - click for more options"),
									),
							),
					),

				app.Pre().
					Class("language-go").
					Text(`sp.OverlayTrigger().
    Placement("right").
    Triggeredby("click hover").
    Trigger(
        sp.Button().
            Text("Hover and Click"),
    ).
    ClickContent(
        sp.Popover().
            Placement("right").
            Tip(true).
            Body(
                sp.Dialog().
                    Nodivider(true).
                    Body(
                        app.H3().Text("Click Content"),
                        app.P().Text("This appears when you click."),
                    ),
            ),
    ).
    HoverContent(
        sp.Tooltip().
            Placement("top").
            Body(
                app.Text("Hover tooltip - click for more options"),
            ),
    )`),
			),

		// Overlay Trigger with Longpress Interaction
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Overlay Trigger with Longpress Interaction"),
				app.P().Text("An overlay trigger that displays content when long-pressed (particularly useful for mobile devices):"),
				app.Div().
					Class("component-example").
					Body(
						sp.OverlayTrigger().
							Placement("bottom").
							Triggeredby("longpress").
							Trigger(
								sp.Button().
									Text("Press and Hold"),
							).
							LongpressContent(
								sp.Popover().
									Placement("bottom").
									Tip(true).
									Body(
										sp.Dialog().
											Nodivider(true).
											Body(
												app.H3().Text("Longpress Content"),
												app.P().Text("This appears when you press and hold."),
											),
									),
							).
							LongpressDescribedbyDescriptor(
								app.Div().
									Text("Press and hold to reveal more options"),
							),
					),

				app.Pre().
					Class("language-go").
					Text(`sp.OverlayTrigger().
    Placement("bottom").
    Triggeredby("longpress").
    Trigger(
        sp.Button().
            Text("Press and Hold"),
    ).
    LongpressContent(
        sp.Popover().
            Placement("bottom").
            Tip(true).
            Body(
                sp.Dialog().
                    Nodivider(true).
                    Body(
                        app.H3().Text("Longpress Content"),
                        app.P().Text("This appears when you press and hold."),
                    ),
            ),
    ).
    LongpressDescribedbyDescriptor(
        app.Div().
            Text("Press and hold to reveal more options"),
    )`),
			),

		// Overlay Trigger with Different Placements
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Overlay Trigger with Different Placements"),
				app.P().Text("Overlay triggers can position their content in different locations relative to the trigger:"),
				app.Div().
					Class("component-example").
					Style("display", "flex").
					Style("flex-wrap", "wrap").
					Style("gap", "10px").
					Style("padding", "20px").
					Body(
						// Top placement
						sp.OverlayTrigger().
							PlacementTop().
							Trigger(
								sp.Button().
									Text("Top"),
							).
							ClickContent(
								sp.Tooltip().
									Placement("top").
									Body(
										app.Text("Top placement"),
									),
							),
						// Right placement
						sp.OverlayTrigger().
							PlacementRight().
							Trigger(
								sp.Button().
									Text("Right"),
							).
							ClickContent(
								sp.Tooltip().
									Placement("right").
									Body(
										app.Text("Right placement"),
									),
							),
						// Bottom placement
						sp.OverlayTrigger().
							PlacementBottom().
							Trigger(
								sp.Button().
									Text("Bottom"),
							).
							ClickContent(
								sp.Tooltip().
									Placement("bottom").
									Body(
										app.Text("Bottom placement"),
									),
							),
						// Left placement
						sp.OverlayTrigger().
							PlacementLeft().
							Trigger(
								sp.Button().
									Text("Left"),
							).
							ClickContent(
								sp.Tooltip().
									Placement("left").
									Body(
										app.Text("Left placement"),
									),
							),
					),

				app.Pre().
					Class("language-go").
					Text(`// Top placement
sp.OverlayTrigger().
    PlacementTop().
    Trigger(
        sp.Button().
            Text("Top"),
    ).
    ClickContent(
        sp.Tooltip().
            Placement("top").
            Body(
                app.Text("Top placement"),
            ),
    )

// More placement examples (Right, Bottom, Left) follow the same pattern`),
			),

		// Customizing Focus Behavior
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Customizing Focus Behavior"),
				app.P().Text("Control how focus is handled when the overlay opens:"),
				app.Div().
					Class("component-example").
					Body(
						sp.OverlayTrigger().
							Placement("bottom").
							ReceivesfocusTrue().
							Trigger(
								sp.Button().
									Text("Focus the Overlay"),
							).
							ClickContent(
								sp.Popover().
									Placement("bottom").
									Tip(true).
									Body(
										sp.Dialog().
											Nodivider(true).
											Body(
												app.H3().Text("Focused Content"),
												app.P().Text("This overlay receives focus when opened."),
												sp.Textfield().Placeholder("I'll receive focus"),
											),
									),
							),
					),

				app.Pre().
					Class("language-go").
					Text(`sp.OverlayTrigger().
    Placement("bottom").
    ReceivesfocusTrue().
    Trigger(
        sp.Button().
            Text("Focus the Overlay"),
    ).
    ClickContent(
        sp.Popover().
            Placement("bottom").
            Tip(true).
            Body(
                sp.Dialog().
                    Nodivider(true).
                    Body(
                        app.H3().Text("Focused Content"),
                        app.P().Text("This overlay receives focus when opened."),
                        sp.Textfield().Placeholder("I'll receive focus"),
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
						app.Li().Text("For displaying contextual information or actions when a user interacts with a trigger element"),
						app.Li().Text("When you need to display additional content without navigating to a new page"),
						app.Li().Text("For providing tooltips, popovers, menus, or dialog content based on user interaction"),
						app.Li().Text("When you need to support multiple interaction types (click, hover, longpress) for the same trigger"),
					),
				app.H3().Text("Best practices"),
				app.Ul().
					Body(
						app.Li().Text("Use the appropriate interaction type for your content (hover for tooltips, click for menus, etc.)"),
						app.Li().Text("Specify the triggeredBy attribute to optimize performance when you know which interaction types you'll use"),
						app.Li().Text("Choose appropriate placement based on the screen space available and the user's context"),
						app.Li().Text("Ensure the trigger element is interactive and can receive focus for accessibility"),
						app.Li().Text("Provide clear ways to dismiss overlay content, especially for click and longpress interactions"),
					),
			),
	)
}
