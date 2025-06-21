package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/asgardtech/spectrum-go/prism"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// actionButtonPage represents the Action Button component showcase
type actionButtonPage struct {
	app.Compo
}

// newActionButtonPage creates a new Action Button component showcase
func newActionButtonPage() *actionButtonPage {
	return &actionButtonPage{}
}

// OnNav initializes the page when navigated to
func (p *actionButtonPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

// initPage initializes the page
func (p *actionButtonPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle(defaultTitle)
	ctx.Page().SetDescription(defaultDescription)
}

// Render renders the Action Button component showcase
func (p *actionButtonPage) Render() app.UI {
	return prism.NewLayout().Content(
		app.H1().Text("Action Button Component"),
		app.P().Text("An Action Button represents an action a user can take. It can be combined with icons, text, and various visual styles."),

		// Basic Action Button
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Basic Action Button"),
				app.P().Text("A simple action button with text:"),
				app.Div().
					Class("component-example").
					Body(
						sp.ActionButton().Text("Edit"),
					),

				app.Pre().
					Class("language-go").
					Text(`sp.ActionButton().Text("Edit")`),
			),

		// Action Button with Icon
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Action Button with Icon"),
				app.P().Text("A button with both icon and text:"),
				app.Div().
					Class("component-example").
					Body(
						sp.ActionButton().
							Icon(
								app.Elem("div").
									Attr("slot", "icon").
									Body(
										sp.Icon().Name("edit"),
									),
							).
							Text("Edit"),
					),

				app.Pre().
					Class("language-go").
					Text(`sp.ActionButton().
    Icon(
        app.Elem("div").
            Attr("slot", "icon").
            Body(
                sp.Icon().Name("edit"),
            ),
    ).
    Text("Edit")`),
			),

		// Icon-only Action Button
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Icon-only Action Button"),
				app.P().Text("A button with only an icon (requires label for accessibility):"),
				app.Div().
					Class("component-example").
					Body(
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

				app.Pre().
					Class("language-go").
					Text(`sp.ActionButton().
    Label("Edit").
    Icon(
        app.Elem("div").
            Attr("slot", "icon").
            Body(
                sp.Icon().Name("edit"),
            ),
    )`),
			),

		// Toggle Action Button
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Toggle Action Button"),
				app.P().Text("A button that toggles its selected state on click:"),
				app.Div().
					Class("component-example").
					Body(
						sp.ActionButton().
							Toggles(true).
							Text("Toggle button"),
					),

				app.Pre().
					Class("language-go").
					Text(`sp.ActionButton().
    Toggles(true).
    Text("Toggle button")`),
			),

		// Action Button with Hold Affordance
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Action Button with Hold Affordance"),
				app.P().Text("A button that shows a visual cue for longpress functionality:"),
				app.Div().
					Class("component-example").
					Body(
						sp.ActionButton().
							HoldAffordance(true).
							Icon(
								app.Elem("div").
									Attr("slot", "icon").
									Body(
										sp.Icon().Name("edit"),
									),
							),
					),

				app.Pre().
					Class("language-go").
					Text(`sp.ActionButton().
    HoldAffordance(true).
    Icon(
        app.Elem("div").
            Attr("slot", "icon").
            Body(
                sp.Icon().Name("edit"),
            ),
    )`),
			),

		// Action Button Variants
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Action Button Variants"),
				app.P().Text("Action buttons can be customized with various visual styles:"),

				// Emphasized variant
				app.H3().Text("Emphasized"),
				app.Div().
					Class("component-example").
					Body(
						sp.ActionButton().
							Emphasized(true).
							Text("Emphasized button"),
					),

				app.Pre().
					Class("language-go").
					Text(`sp.ActionButton().
    Emphasized(true).
    Text("Emphasized button")`),

				// Quiet variant
				app.H3().Text("Quiet"),
				app.Div().
					Class("component-example").
					Body(
						sp.ActionButton().
							Quiet(true).
							Text("Quiet button"),
					),

				app.Pre().
					Class("language-go").
					Text(`sp.ActionButton().
    Quiet(true).
    Text("Quiet button")`),

				// Combined variants
				app.H3().Text("Combined Variants"),
				app.Div().
					Class("component-example").
					Body(
						sp.ActionButton().
							Emphasized(true).
							Quiet(true).
							Text("Emphasized quiet button"),
					),

				app.Pre().
					Class("language-go").
					Text(`sp.ActionButton().
    Emphasized(true).
    Quiet(true).
    Text("Emphasized quiet button")`),
			),

		// Action Button States
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Action Button States"),
				app.P().Text("Action buttons can be displayed in various states:"),

				// Active state
				app.H3().Text("Active"),
				app.Div().
					Class("component-example").
					Body(
						sp.ActionButton().
							Active(true).
							Text("Active button"),
					),

				app.Pre().
					Class("language-go").
					Text(`sp.ActionButton().
    Active(true).
    Text("Active button")`),

				// Selected state
				app.H3().Text("Selected"),
				app.Div().
					Class("component-example").
					Body(
						sp.ActionButton().
							Toggles(true).
							Selected(true).
							Text("Selected button"),
					),

				app.Pre().
					Class("language-go").
					Text(`sp.ActionButton().
    Toggles(true).
    Selected(true).
    Text("Selected button")`),

				// Disabled state
				app.H3().Text("Disabled"),
				app.Div().
					Class("component-example").
					Body(
						sp.ActionButton().
							Disabled(true).
							Text("Disabled button"),
					),

				app.Pre().
					Class("language-go").
					Text(`sp.ActionButton().
    Disabled(true).
    Text("Disabled button")`),
			),

		// Action Button as Links
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Action Button as Links"),
				app.P().Text("Action buttons can function as links with the href attribute:"),
				app.Div().
					Class("component-example").
					Body(
						sp.ActionButton().
							Href("https://adobe.com").
							Target(sp.ActionButtonTarget_blank).
							Text("Visit Adobe"),
					),

				app.Pre().
					Class("language-go").
					Text(`sp.ActionButton().
    Href("https://adobe.com").
    Target(sp.ActionButtonTarget_blank).
    Text("Visit Adobe")`),
			),

		// Action Button Types
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Action Button Types"),
				app.P().Text("Action buttons can have different types in forms:"),

				// Submit button
				app.H3().Text("Submit Button"),
				app.Div().
					Class("component-example").
					Body(
						sp.ActionButton().
							TypeSubmit().
							Text("Submit Form"),
					),

				app.Pre().
					Class("language-go").
					Text(`sp.ActionButton().
    TypeSubmit().
    Text("Submit Form")`),

				// Reset button
				app.H3().Text("Reset Button"),
				app.Div().
					Class("component-example").
					Body(
						sp.ActionButton().
							TypeReset().
							Text("Reset Form"),
					),

				app.Pre().
					Class("language-go").
					Text(`sp.ActionButton().
    TypeReset().
    Text("Reset Form")`),
			),

		// Usage Guidelines
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Usage Guidelines"),
				app.H3().Text("When to use"),
				app.Ul().
					Body(
						app.Li().Text("For representing actions a user can take"),
						app.Li().Text("When you need a button with an icon"),
						app.Li().Text("For toggle functionality in a button context"),
						app.Li().Text("When buttons need to incorporate special behaviors like longpress"),
					),
				app.H3().Text("Best practices"),
				app.Ul().
					Body(
						app.Li().Text("Keep button text concise and action-oriented"),
						app.Li().Text("For icon-only buttons, always include a label for accessibility"),
						app.Li().Text("Use consistent styling for related actions"),
						app.Li().Text("Consider using ActionGroup to organize multiple ActionButtons"),
					),
			),
	)
}
