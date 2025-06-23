package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/asgardtech/spectrum-go/prism"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type alertBannerPage struct {
	app.Compo
}

func newAlertBannerPage() *alertBannerPage {
	return &alertBannerPage{}
}

func (p *alertBannerPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *alertBannerPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle(defaultTitle)
	ctx.Page().SetDescription(defaultDescription)
}

func (p *alertBannerPage) Render() app.UI {
	return prism.NewPage().Content(
		app.H1().Text("Alert Banner Component"),
		app.P().Text("Alert Banner shows pressing and high-signal messages, such as system alerts. It is meant to be noticed and prompt users to take action."),

		// Basic Alert Banner
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Basic Alert Banner"),
				app.P().Text("A simple alert banner with text content."),
				app.Div().Body(
					sp.AlertBanner().
						Open(true).
						Text("All documents in this folder have been archived"),
				),
			),

		// Dismissible Alert Banner
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Dismissible Alert Banner"),
				app.P().Text("Alert banners can include a close button to dismiss them."),
				app.Div().Body(
					sp.AlertBanner().
						Open(true).
						Dismissible(true).
						Text("All documents in this folder have been archived"),
				),
			),

		// Actionable Alert Banner
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Actionable Alert Banner"),
				app.P().Text("Alert banners can include action buttons for users to take."),
				app.Div().Body(
					sp.AlertBanner().
						Open(true).
						Dismissible(true).
						Body(
							app.Text("Your trial has expired"),
							app.Div().
								Attr("slot", "action").
								Body(
									sp.Button().
										Treatment("outline").
										StaticColor("white").
										Text("Buy now"),
								),
						),
				),
			),

		// Info Variant
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Info Variant"),
				app.P().Text("Alert banners can be styled with different variants to indicate different types of messages."),
				app.Div().Body(
					sp.AlertBanner().
						Open(true).
						Dismissible(true).
						Variant("info").
						Body(
							app.Text("Your trial will expire in 3 days"),
							app.Div().
								Attr("slot", "action").
								Body(
									sp.Button().
										Treatment("outline").
										StaticColor("white").
										Text("Buy now"),
								),
						),
				),
			),

		// Negative Variant
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Negative Variant"),
				app.P().Text("Negative variant is used for error or warning messages."),
				app.Div().Body(
					sp.AlertBanner().
						Open(true).
						Dismissible(true).
						Variant("negative").
						Text("Connection interrupted. Check your network to continue"),
				),
			),
	)
}
