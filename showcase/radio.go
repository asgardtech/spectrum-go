package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/asgardtech/spectrum-go/prism"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type radioPage struct {
	app.Compo
}

func newRadioPage() *radioPage {
	return &radioPage{}
}

func (p *radioPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *radioPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle(defaultTitle)
	ctx.Page().SetDescription(defaultDescription)
}

func (p *radioPage) Render() app.UI {
	return prism.NewPage().Content(
		app.H1().Text("Radio Components"),
		app.P().Text("Radio buttons allow users to select a single option from a list of mutually exclusive options."),

		// Individual Radio Buttons
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Individual Radio Buttons"),
				app.P().Text("Radio buttons can be used individually, but are typically used in a radio group."),
				app.Div().Class("component-row").Body(
					sp.Radio().
						Value("option1").
						Text("Unchecked Radio"),
					sp.Radio().
						Value("option2").
						Text("Checked Radio").
						Checked(true),
				),
			),

		// Radio Button States
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Radio Button States"),
				app.Div().Class("component-row").Body(
					sp.Radio().
						Value("standard").
						Text("Standard").
						Checked(true),
					sp.Radio().
						Value("emphasized").
						Text("Emphasized").
						Emphasized(true).
						Checked(true),
					sp.Radio().
						Value("invalid").
						Text("Invalid").
						Invalid(true),
					sp.Radio().
						Value("disabled").
						Text("Disabled").
						Disabled(true),
					sp.Radio().
						Value("readonly").
						Text("Readonly").
						Readonly(true).
						Checked(true),
				),
			),

		// Radio Button Sizes
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Radio Button Sizes"),
				app.Div().Class("component-row").Body(
					sp.Radio().
						Value("small").
						Text("Small").
						Size("s"),
					sp.Radio().
						Value("medium").
						Text("Medium (Default)").
						Size("m"),
					sp.Radio().
						Value("large").
						Text("Large").
						Size("l"),
					sp.Radio().
						Value("xlarge").
						Text("Extra Large").
						Size("xl"),
				),
			),

		// Basic Radio Group
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Basic Radio Group"),
				app.P().Text("Radio groups manage a collection of radio buttons, ensuring only one is selected."),
				app.Div().Class("component-row").Body(
					app.Div().
						Body(
							app.Raw(`
								<sp-radio-group label="Choose a fruit" selected="banana" name="fruit">
									<sp-radio value="apple">Apple</sp-radio>
									<sp-radio value="banana">Banana</sp-radio>
									<sp-radio value="orange">Orange</sp-radio>
								</sp-radio-group>
							`),
						),
				),
			),

		// Radio Group with Help Text
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Radio Group with Help Text"),
				app.Div().Class("component-row").Body(
					app.Div().
						Body(
							app.Raw(`
								<sp-radio-group label="Choose a plan" selected="premium" name="plan">
									<sp-radio value="free">Free</sp-radio>
									<sp-radio value="basic">Basic</sp-radio>
									<sp-radio value="premium">Premium</sp-radio>
									<div slot="help-text">Select the plan that works best for you.</div>
								</sp-radio-group>
							`),
						),
				),
			),

		// Invalid Radio Group
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Invalid Radio Group"),
				app.Div().Class("component-row").Body(
					app.Div().
						Body(
							app.Raw(`
								<sp-radio-group label="Choose a country" invalid name="country">
									<sp-radio value="usa" invalid>United States</sp-radio>
									<sp-radio value="canada" invalid>Canada</sp-radio>
									<sp-radio value="mexico" invalid>Mexico</sp-radio>
									<div slot="negative-help-text">
										<span style="display: flex; align-items: center;">
											<sp-icon-alert size="s" style="margin-right: 8px;"></sp-icon-alert>
											Please select a valid country.
										</span>
									</div>
								</sp-radio-group>
							`),
						),
				),
			),

		// Horizontal Radio Group
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Horizontal Radio Group"),
				app.Div().Class("component-row").Body(
					app.Div().
						Body(
							app.Raw(`
								<sp-radio-group label="Choose a color" horizontal selected="green" name="color">
									<sp-radio value="red">Red</sp-radio>
									<sp-radio value="green">Green</sp-radio>
									<sp-radio value="blue">Blue</sp-radio>
								</sp-radio-group>
							`),
						),
				),
			),

		// Vertical Radio Group (explicit)
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Vertical Radio Group (explicit)"),
				app.Div().Class("component-row").Body(
					app.Div().
						Body(
							app.Raw(`
								<sp-radio-group label="Choose a size" vertical selected="medium" name="size">
									<sp-radio value="small">Small</sp-radio>
									<sp-radio value="medium">Medium</sp-radio>
									<sp-radio value="large">Large</sp-radio>
									<sp-radio value="xlarge">Extra Large</sp-radio>
								</sp-radio-group>
							`),
						),
				),
			),
	)
}
