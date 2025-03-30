package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type homePage struct {
	app.Compo

	// Theme state
	CurrentTheme sp.ThemeColor
}

func newHomePage() *homePage {
	return &homePage{
		CurrentTheme: sp.ThemeColorLight, // Default to light theme
	}
}

func (p *homePage) Render() app.UI {
	// Create a main container with theme-aware styling
	return sp.Theme().
		Color(p.CurrentTheme).
		Scale(sp.ThemeScaleMedium).
		Children(

			sp.TopNav().
				Children(
					sp.TopNavItem().
						Label("Home").
						Value("home").
						Text("Home"),
					sp.ActionMenu().
						Placement("bottom-end").
						Style("margin-inline-start", "auto").
						Quiet(true).
						AddItem(sp.MenuItem().
							Label("Toggle Theme").
							Text("Toggle Theme").
							//	Icon(p.getThemeIcon()).
							OnClick(func(ctx app.Context, e app.Event) {
								app.Log("Toggle Theme", p.CurrentTheme)
								p.toggleTheme()
							}),
						),
				),

			// Main content with split view
			sp.SplitView().
				Collapsible(true).
				PrimarySize("250px").
				PrimaryMin(200).
				Resizable(false).
				PrimaryPanel(p.renderSidenav()).
				SecondaryPanel(p.renderContent()),
		)
}

// Toggle theme between light and dark
func (p *homePage) toggleTheme() {
	if p.CurrentTheme == sp.ThemeColorLight {
		p.CurrentTheme = sp.ThemeColorDark
	} else {
		p.CurrentTheme = sp.ThemeColorLight
	}
}

// Get appropriate theme icon based on current theme
func (p *homePage) getThemeIcon() app.UI {
	if p.CurrentTheme == sp.ThemeColorLight {
		return sp.Icon().Name("ui:Moon100")
	}
	return sp.Icon().Name("ui:Sun100")
}

func (p *homePage) renderSidenav() app.UI {
	// Create Sidenav with components
	return sp.Sidenav().
		Label("Component Navigation").
		AddItem(sp.SidenavItem().Label("Button").Value("button").Href("/button")).
		AddItem(sp.SidenavItem().Label("Accordion").Value("accordion").Href("/accordion")).
		AddItem(sp.SidenavItem().Label("Button Group").Value("button-group").Href("/button-group")).
		AddItem(sp.SidenavItem().Label("Checkbox").Value("checkbox").Href("/checkbox")).
		AddItem(sp.SidenavItem().Label("Dropdown").Value("dropdown").Href("/dropdown")).
		AddItem(sp.SidenavItem().Label("Label").Value("label").Href("/label"))
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
