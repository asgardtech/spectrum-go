package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type popoverPage struct {
	app.Compo
}

func newPopoverPage() *popoverPage {
	return &popoverPage{}
}

func (p *popoverPage) Render() app.UI {
	return newPage().Content(
		app.H1().Text("Popover Component"),
		app.P().Text("Popovers display content or interactive elements that are layered over other content. They generally appear on hover, focus, or click of a trigger element."),

		// Basic popover with overlay trigger
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Basic Popover"),
				app.Div().Class("component-row").Body(
					sp.OverlayTrigger().
						Type("click").
						Placement("right").
						Trigger(
							sp.Button().Text("Open Popover").Variant(sp.ButtonVariantPrimary),
						).
						ClickContent(
							sp.Popover().
								Tip(true).
								Body(
									app.H3().Text("Basic Popover"),
									app.P().Text("This is a basic popover with default styling."),
								),
						),
				),
			),

		// Dialog popover with overlay trigger
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Dialog Popover"),
				app.Div().Class("component-row").Body(
					sp.OverlayTrigger().
						Type("click").
						Placement("bottom").
						Trigger(
							sp.Button().Text("Open Dialog Popover").Variant(sp.ButtonVariantPrimary),
						).
						ClickContent(
							sp.Popover().
								Tip(true).
								Body(
									app.H3().Text("Dialog Popover"),
									app.P().Text("This popover has dialog styling with a header and footer section."),
									app.Div().Class("button-container").Body(
										sp.Button().Text("Close").Variant(sp.ButtonVariantSecondary),
										sp.Button().Text("Apply").Variant(sp.ButtonVariantPrimary),
									),
								),
						),
				),
			),

		// Placement options
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Popover Placement"),
				app.Div().Class("placement-demo").Body(
					app.Div().Class("placement-row").Body(
						sp.OverlayTrigger().
							Type("click").
							Placement("top-start").
							Trigger(
								sp.Button().Text("Top Start").Variant(sp.ButtonVariantSecondary),
							).
							ClickContent(
								sp.Popover().
									Tip(true).
									Body(
										app.P().Text("Top Start Placement"),
									),
							),
						sp.OverlayTrigger().
							Type("click").
							Placement("top").
							Trigger(
								sp.Button().Text("Top").Variant(sp.ButtonVariantSecondary),
							).
							ClickContent(
								sp.Popover().
									Tip(true).
									Body(
										app.P().Text("Top Placement"),
									),
							),
						sp.OverlayTrigger().
							Type("click").
							Placement("top-end").
							Trigger(
								sp.Button().Text("Top End").Variant(sp.ButtonVariantSecondary),
							).
							ClickContent(
								sp.Popover().
									Tip(true).
									Body(
										app.P().Text("Top End Placement"),
									),
							),
					),
					app.Div().Class("placement-row").Body(
						sp.OverlayTrigger().
							Type("click").
							Placement("bottom-start").
							Trigger(
								sp.Button().Text("Bottom Start").Variant(sp.ButtonVariantSecondary),
							).
							ClickContent(
								sp.Popover().
									Tip(true).
									Body(
										app.P().Text("Bottom Start Placement"),
									),
							),
						sp.OverlayTrigger().
							Type("click").
							Placement("bottom").
							Trigger(
								sp.Button().Text("Bottom").Variant(sp.ButtonVariantSecondary),
							).
							ClickContent(
								sp.Popover().
									Tip(true).
									Body(
										app.P().Text("Bottom Placement"),
									),
							),
						sp.OverlayTrigger().
							Type("click").
							Placement("bottom-end").
							Trigger(
								sp.Button().Text("Bottom End").Variant(sp.ButtonVariantSecondary),
							).
							ClickContent(
								sp.Popover().
									Tip(true).
									Body(
										app.P().Text("Bottom End Placement"),
									),
							),
					),
					app.Div().Class("placement-row").Body(
						sp.OverlayTrigger().
							Type("click").
							Placement("left").
							Trigger(
								sp.Button().Text("Left").Variant(sp.ButtonVariantSecondary),
							).
							ClickContent(
								sp.Popover().
									Tip(true).
									Body(
										app.P().Text("Left Placement"),
									),
							),
						sp.OverlayTrigger().
							Type("click").
							Placement("right").
							Trigger(
								sp.Button().Text("Right").Variant(sp.ButtonVariantSecondary),
							).
							ClickContent(
								sp.Popover().
									Tip(true).
									Body(
										app.P().Text("Right Placement"),
									),
							),
					),
				),
			),

		// Without tip
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Popover Without Tip"),
				app.Div().Class("component-row").Body(
					sp.OverlayTrigger().
						Type("click").
						Placement("right").
						Trigger(
							sp.Button().Text("No Tip").Variant(sp.ButtonVariantSecondary),
						).
						ClickContent(
							sp.Popover().
								Body(
									app.P().Text("This popover has no triangular tip."),
								),
						),
				),
			),
	)
}