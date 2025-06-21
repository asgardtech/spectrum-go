package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/asgardtech/spectrum-go/prism"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type tooltipPage struct {
	app.Compo
}

func newTooltipPage() *tooltipPage {
	return &tooltipPage{}
}

func (p *tooltipPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *tooltipPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle(defaultTitle)
	ctx.Page().SetDescription(defaultDescription)
}

func (p *tooltipPage) Render() app.UI {
	return prism.NewLayout().Content(
		app.H1().Text("Tooltip Component"),
		app.P().Text("Tooltips allow users to get contextual help or information about specific components when hovering or focusing on them."),

		// Basic Tooltip Examples
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Basic Tooltip Examples"),
				app.P().Text("Tooltips can be positioned at different placements relative to their target."),

				app.Div().
					Style("display", "flex").
					Style("flex-direction", "column").
					Style("gap", "30px").
					Style("margin", "20px 0").
					Body(
						// Top Placement
						app.Div().
							Style("display", "flex").
							Style("align-items", "center").
							Style("gap", "20px").
							Body(
								app.Div().
									Style("width", "150px").
									Text("Top Placement:"),
								app.Div().
									Style("position", "relative").
									Body(
										sp.Button().
											Text("Hover Me").
											Variant("primary"),
										sp.Tooltip().
											Open(true).
											Placement("top").
											Text("Tooltip on top"),
									),
							),

						// Right Placement
						app.Div().
							Style("display", "flex").
							Style("align-items", "center").
							Style("gap", "20px").
							Body(
								app.Div().
									Style("width", "150px").
									Text("Right Placement:"),
								app.Div().
									Style("position", "relative").
									Body(
										sp.Button().
											Text("Hover Me").
											Variant("primary"),
										sp.Tooltip().
											Open(true).
											Placement("right").
											Text("Tooltip on right"),
									),
							),

						// Bottom Placement
						app.Div().
							Style("display", "flex").
							Style("align-items", "center").
							Style("gap", "20px").
							Body(
								app.Div().
									Style("width", "150px").
									Text("Bottom Placement:"),
								app.Div().
									Style("position", "relative").
									Body(
										sp.Button().
											Text("Hover Me").
											Variant("primary"),
										sp.Tooltip().
											Open(true).
											Placement("bottom").
											Text("Tooltip on bottom"),
									),
							),

						// Left Placement
						app.Div().
							Style("display", "flex").
							Style("align-items", "center").
							Style("gap", "20px").
							Body(
								app.Div().
									Style("width", "150px").
									Text("Left Placement:"),
								app.Div().
									Style("position", "relative").
									Body(
										sp.Button().
											Text("Hover Me").
											Variant("primary"),
										sp.Tooltip().
											Open(true).
											Placement("left").
											Text("Tooltip on left"),
									),
							),
					),
			),

		// Tooltip Variants
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Tooltip Variants"),
				app.P().Text("Tooltips come in different variants to indicate different types of information."),

				app.Div().
					Style("display", "flex").
					Style("flex-direction", "column").
					Style("gap", "30px").
					Style("margin", "20px 0").
					Body(
						// Standard Tooltip
						app.Div().
							Style("display", "flex").
							Style("align-items", "center").
							Style("gap", "20px").
							Body(
								app.Div().
									Style("width", "150px").
									Text("Standard:"),
								app.Div().
									Style("position", "relative").
									Body(
										sp.Button().
											Text("Hover Me").
											Variant("primary"),
										sp.Tooltip().
											Open(true).
											Placement("top").
											Text("Standard tooltip"),
									),
							),

						// Info Tooltip
						app.Div().
							Style("display", "flex").
							Style("align-items", "center").
							Style("gap", "20px").
							Body(
								app.Div().
									Style("width", "150px").
									Text("Info:"),
								app.Div().
									Style("position", "relative").
									Body(
										sp.Button().
											Text("Hover Me").
											Variant("primary"),
										sp.Tooltip().
											Open(true).
											Placement("top").
											Variant("info").
											Text("Informational tooltip"),
									),
							),

						// Positive Tooltip
						app.Div().
							Style("display", "flex").
							Style("align-items", "center").
							Style("gap", "20px").
							Body(
								app.Div().
									Style("width", "150px").
									Text("Positive:"),
								app.Div().
									Style("position", "relative").
									Body(
										sp.Button().
											Text("Hover Me").
											Variant("primary"),
										sp.Tooltip().
											Open(true).
											Placement("top").
											Variant("positive").
											Text("Success tooltip"),
									),
							),

						// Negative Tooltip
						app.Div().
							Style("display", "flex").
							Style("align-items", "center").
							Style("gap", "20px").
							Body(
								app.Div().
									Style("width", "150px").
									Text("Negative:"),
								app.Div().
									Style("position", "relative").
									Body(
										sp.Button().
											Text("Hover Me").
											Variant("primary"),
										sp.Tooltip().
											Open(true).
											Placement("top").
											Variant("negative").
											Text("Error tooltip"),
									),
							),
					),
			),

		// Tooltips with Custom Content
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Tooltips with Custom Content"),
				app.P().Text("Tooltips can include custom content to enhance the displayed information."),

				app.Div().
					Style("display", "flex").
					Style("flex-direction", "column").
					Style("gap", "30px").
					Style("margin", "20px 0").
					Body(
						// Info Tooltip with Custom Content
						app.Div().
							Style("display", "flex").
							Style("align-items", "center").
							Style("gap", "20px").
							Body(
								app.Div().
									Style("width", "150px").
									Text("Info with Icon:"),
								app.Div().
									Style("position", "relative").
									Body(
										sp.Button().
											Text("Hover Me").
											Variant("primary"),
										sp.Tooltip().
											Open(true).
											Placement("top").
											Variant("info").
											Text("Info tooltip with custom icon"),
									),
							),

						// Tooltip with Rich Content
						app.Div().
							Style("display", "flex").
							Style("align-items", "center").
							Style("gap", "20px").
							Body(
								app.Div().
									Style("width", "150px").
									Text("Rich Content:"),
								app.Div().
									Style("position", "relative").
									Body(
										sp.Button().
											Text("Hover Me").
											Variant("primary"),
										sp.Tooltip().
											Open(true).
											Placement("right").
											Text("This tooltip contains formatted text with important information."),
									),
							),
					),
			),

		// Self-managed Tooltips
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Self-managed Tooltips"),
				app.P().Text("Tooltips can be self-managed, meaning they will automatically show on hover or focus of their parent element."),

				app.Div().
					Style("display", "flex").
					Style("align-items", "center").
					Style("gap", "20px").
					Style("margin", "20px 0").
					Body(
						app.Span().
							Text("Hover over the button:"),
						sp.Button().
							Text("Interactive Button").
							Variant("primary").
							Body(
								sp.Tooltip().
									SelfManaged(true).
									Placement("bottom").
									Text("This tooltip appears on hover"),
							),
					),
			),

		// Additional Features
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Additional Features"),
				app.P().Text("Tooltips support additional features like custom offset and tip padding."),

				app.Div().
					Style("display", "flex").
					Style("flex-direction", "column").
					Style("gap", "30px").
					Style("margin", "20px 0").
					Body(
						// Tooltip with Offset
						app.Div().
							Style("display", "flex").
							Style("align-items", "center").
							Style("gap", "20px").
							Body(
								app.Div().
									Style("width", "150px").
									Text("With Offset:"),
								app.Div().
									Style("position", "relative").
									Body(
										sp.Button().
											Text("Hover Me").
											Variant("primary"),
										sp.Tooltip().
											Open(true).
											Placement("top").
											Offset(10).
											Text("Tooltip with 10px offset"),
									),
							),

						// Tooltip with Long Text
						app.Div().
							Style("display", "flex").
							Style("align-items", "center").
							Style("gap", "20px").
							Body(
								app.Div().
									Style("width", "150px").
									Text("Long Text:"),
								app.Div().
									Style("position", "relative").
									Body(
										sp.Button().
											Text("Hover Me").
											Variant("primary"),
										sp.Tooltip().
											Open(true).
											Placement("right").
											Text("This is a tooltip with a longer text that demonstrates how tooltips can handle multiple lines of content."),
									),
							),
					),
			),

		// Usage Guidelines
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Usage Guidelines"),
				app.P().Text("When to use tooltips:"),
				app.Ul().
					Body(
						app.Li().Text("Use tooltips to provide additional context or explanation for UI elements."),
						app.Li().Text("Use tooltips for elements that might be unfamiliar to users or have hidden functionality."),
						app.Li().Text("Use tooltips to display keyboard shortcuts or additional information that's not essential for basic usage."),
					),
				app.P().Text("Best practices:"),
				app.Ul().
					Body(
						app.Li().Text("Keep tooltip content brief and focused."),
						app.Li().Text("Use appropriate placement to avoid covering important information."),
						app.Li().Text("Use variants (info, positive, negative) to match the context of the information."),
						app.Li().Text("Use appropriate icons to supplement variants when needed."),
					),
			),
	)
}
