package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/asgardtech/spectrum-go/prism"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type switchPage struct {
	app.Compo
}

func newSwitchPage() *switchPage {
	return &switchPage{}
}

func (p *switchPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *switchPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle(defaultTitle)
	ctx.Page().SetDescription(defaultDescription)
}

func (p *switchPage) Render() app.UI {
	return prism.NewPage().Content(
		app.H1().Text("Switch Component"),
		app.P().Text("A switch is used to turn an option on or off. Switches allow users to select the state of a single option at a time."),

		// Basic switches
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Basic Switch"),
				app.Div().Class("component-row").Body(
					sp.Switch().Text("Off (Default)"),
					sp.Switch().Text("On").Checked(true),
				),
			),

		// Emphasized style
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Emphasized Style"),
				app.Div().Class("component-row").Body(
					sp.Switch().Text("Standard").Checked(true),
					sp.Switch().Text("Emphasized").Emphasized(true).Checked(true),
				),
			),

		// Disabled state
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Disabled State"),
				app.Div().Class("component-row").Body(
					sp.Switch().Text("Disabled Off").Disabled(true),
					sp.Switch().Text("Disabled On").Disabled(true).Checked(true),
				),
			),

		// Readonly state
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Readonly State"),
				app.Div().Class("component-row").Body(
					sp.Switch().Text("Readonly Off").Readonly(true),
					sp.Switch().Text("Readonly On").Readonly(true).Checked(true),
				),
			),

		// Size variants
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Size Variants"),
				app.Div().Class("component-row").Body(
					sp.Switch().Text("Small").Size("s").Checked(true),
					sp.Switch().Text("Medium (Default)").Size("m").Checked(true),
					sp.Switch().Text("Large").Size("l").Checked(true),
					sp.Switch().Text("Extra Large").Size("xl").Checked(true),
				),
			),

		// Form integration example
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Form Integration"),
				app.P().Text("Switches can be integrated into forms with proper name attributes."),
				app.Div().Class("component-row").Body(
					app.Form().
						Target("_blank").
						Method("get").
						Action("#").
						OnSubmit(func(ctx app.Context, e app.Event) {
							e.PreventDefault()
							app.Log("Form submitted")
						}).
						Body(
							app.Div().
								Class("fieldset").
								Style("border", "none").
								Style("padding", "0").
								Body(
									app.Div().Class("legend").Text("Notification Preferences"),
									app.Div().Class("form-row").Body(
										sp.Switch().Text("Email notifications").Name("email_notify").Checked(true),
									),
									app.Div().Class("form-row").Body(
										sp.Switch().Text("Push notifications").Name("push_notify"),
									),
									app.Div().Class("form-row").Body(
										sp.Switch().Text("Marketing emails").Name("marketing_emails"),
									),
									app.Div().Class("form-row").Body(
										sp.Button().
											Type("submit").
											Text("Save Preferences").
											Variant(sp.ButtonVariantPrimary),
									),
								),
						),
				),
			),
	)
}
