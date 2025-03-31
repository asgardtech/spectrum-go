package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type buttonGroupPage struct {
	app.Compo
}

func newButtonGroupPage() *buttonGroupPage {
	return &buttonGroupPage{}
}

func (p *buttonGroupPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *buttonGroupPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle(defaultTitle)
	ctx.Page().SetDescription(defaultDescription)
}

func (p *buttonGroupPage) Render() app.UI {
	return newPage().Content(
		app.H1().Text("Button Group Component"),
		app.P().Text("The Button Group component organizes related buttons together with consistent spacing. It supports both horizontal and vertical layouts."),

		// Basic Button Group (Horizontal)
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Horizontal Button Group"),
				app.P().Text("By default, Button Group organizes buttons horizontally. This is useful for toolbars, form actions, and horizontal navigation."),
				sp.ButtonGroup().
					Body(
						sp.Button().Text("Button 1"),
						sp.Button().Text("Longer Button 2"),
						sp.Button().Text("Short 3"),
					),
			),

		// Vertical Button Group
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Vertical Button Group"),
				app.P().Text("Use vertical orientation when horizontal space is limited or for stacked navigation."),
				sp.ButtonGroup().
					Vertical(true).
					Body(
						sp.Button().Text("Button 1"),
						sp.Button().Text("Longer Button 2"),
						sp.Button().Text("Short 3"),
					),
			),

		// Button Group with Different Button Variants
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Button Group with Different Button Variants"),
				app.P().Text("Button Groups can contain different button variants for different purposes."),
				sp.ButtonGroup().
					Body(
						sp.Button().
							Variant("primary").
							Treatment("fill").
							Text("Primary"),
						sp.Button().
							Variant("secondary").
							Treatment("outline").
							Text("Secondary"),
						sp.Button().
							Variant("negative").
							Treatment("fill").
							Text("Negative"),
					),
			),

		// Button Group with Icons
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Button Group with Icons"),
				app.P().Text("Buttons within a group can include icons for visual cues."),
				sp.ButtonGroup().
					Body(
						sp.Button().
							Body(
								app.Div().
									Attr("slot", "icon").
									Body(
										sp.Icon().Name("Edit"),
									),
								app.Text("Edit"),
							),
						sp.Button().
							Body(
								app.Div().
									Attr("slot", "icon").
									Body(
										sp.Icon().Name("Copy"),
									),
								app.Text("Copy"),
							),
						sp.Button().
							Body(
								app.Div().
									Attr("slot", "icon").
									Body(
										sp.Icon().Name("Delete"),
									),
								app.Text("Delete"),
							),
					),
			),

		// Button Group with Disabled Buttons
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Button Group with Disabled Buttons"),
				app.P().Text("Individual buttons within a group can be disabled."),
				sp.ButtonGroup().
					Body(
						sp.Button().Text("Enabled"),
						sp.Button().Text("Disabled").Disabled(true),
						sp.Button().Text("Enabled"),
					),
			),
	)
}
