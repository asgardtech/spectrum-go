package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type homePage struct {
	app.Compo

	// Theme state
	currentTheme sp.ThemeColor
}

func newHomePage() *homePage {
	return &homePage{
		currentTheme: sp.ThemeColorLight, // Default to light theme
	}
}

func (p *homePage) Render() app.UI {
	// Create a main container with theme-aware styling
	return sp.Theme().
		Color(p.currentTheme).
		Scale(sp.ThemeScaleMedium).
		Children(
			app.Elem("sp-icons-medium"),
			app.H2().Text("Spectrum Go Components"),
			// Theme toggle button
			sp.ActionButton().
				//Icon(p.getThemeIcon()).
				Quiet(true).
				Content("Toggle Theme").
				OnClick(func(ctx app.Context, e app.Event) {
					app.Log("Toggle Theme", p.currentTheme)
					p.toggleTheme()
				}),

			// Using the Button component with OnClick
			sp.Button().
				Text("Hello there!").
				Variant(sp.ButtonVariantSecondary).
				OnClick(func(ctx app.Context, e app.Event) {
					app.Log("Button clicked")
				}),

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

// Toggle theme between light and dark
func (p *homePage) toggleTheme() {
	if p.currentTheme == sp.ThemeColorLight {
		p.currentTheme = sp.ThemeColorDark
	} else {
		p.currentTheme = sp.ThemeColorLight
	}
}

// Get appropriate theme icon based on current theme
func (p *homePage) getThemeIcon() app.UI {
	if p.currentTheme == sp.ThemeColorLight {
		return sp.Icon().Name("ui:Moon100")
	}
	return sp.Icon().Name("ui:Sun100")
}

func (p *homePage) renderSidenav() app.UI {
	// Create Sidenav with components
	return sp.Sidenav().
		Label("Component Navigation").
		AddItem(
			sp.SidenavItem().
				Label("Button").
				Value("button").
				//	Icon(sp.Icon().Name("ui:Button").Slot("icon")).
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
