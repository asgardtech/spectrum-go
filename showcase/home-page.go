package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type homePage struct {
	app.Compo
}

func newHomePage() *homePage {
	return &homePage{}
}

func (p *homePage) Render() app.UI {
	// Create a main container with theme-aware styling
	return app.Div().
		Class("app-container").
		Body(
			// Header with theme toggle
			app.Header().
				Body(
					app.H2().Text("Spectrum Go Components"),
					// Using the Button component with OnClick
					sp.Button().
						Text("Hello there!").
						Variant(sp.ButtonVariantSecondary).
						OnClick(func(ctx app.Context, e app.Event) {
							app.Log("Button clicked")
						}),
				),
			// Main content with split view
			app.Div().
				Body(
					sp.SplitView().
						PrimarySize("250px").
						PrimaryMin("200px").
						Resizable(true).
						PrimaryPanel(p.renderSidenav()).
						SecondaryPanel(p.renderContent()),
				),
		)
}

func (p *homePage) renderSidenav() app.UI {
	// Create Sidenav with components
	return sp.Sidenav().
		Label("Component Navigation").
		AddItem(
			sp.SidenavItem().
				Label("Button").
				Value("button").
				Icon(sp.Icon().Name("ui:Button").Slot("icon")).
				Selected(true),
		)
}

func (p *homePage) renderContent() app.UI {
	// Content area - will be replaced with component showcase
	return app.Div().
		Body(
			app.H1().
				Text("Spectrum Go Component Showcase"),
			app.P().
				Text("Select a component from the sidebar to view its showcase."),
			app.Div().
				Body(
					sp.Button().
						Text("Primary Button").
						Variant(sp.ButtonVariantPrimary),
				),
		)
}
