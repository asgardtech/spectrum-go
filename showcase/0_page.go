package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"github.com/maxence-charriere/go-app/v10/pkg/ui"
)

type page struct {
	app.Compo

	// Theme state
	CurrentTheme sp.ThemeColor

	// Content
	PageContent []app.UI

	// Context
	ctx app.Context
}

func newPage() *page {
	return &page{
		CurrentTheme: sp.ThemeColorLight, // Default to light theme
	}
}

func (p *page) Content(v ...app.UI) *page {
	p.PageContent = app.FilterUIElems(v...)
	return p
}

func (p *page) OnMount(ctx app.Context) {
	p.ctx = ctx
}

func (p *page) Render() app.UI {
	// Create a main container with theme-aware styling
	return ui.Shell().Content(
		sp.Theme().
			Color(p.CurrentTheme).
			Scale(sp.ThemeScaleMedium).
			Body(

				sp.TopNav().
					Body(
						sp.TopNavItem().
							Label("Home").
							Text("Home"),
						sp.ActionMenu().
							Placement("bottom-end").
							Style("margin-inline-start", "auto").
							Quiet(true).
							AddToBody(
								sp.MenuItem().
									Label("Toggle Theme").
									Text("Toggle Theme").
									//	Icon(p.getThemeIcon()).
									OnClick(func(ctx app.Context, e app.Event) {
										app.Log("Current Theme", p.CurrentTheme)
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
					AddToBody(p.renderSidenav()).
					AddToBody(p.renderContent()),
			),
	)
}

// Toggle theme between light and dark
func (p *page) toggleTheme() {
	if p.CurrentTheme == sp.ThemeColorLight {
		p.CurrentTheme = sp.ThemeColorDark
	} else {
		p.CurrentTheme = sp.ThemeColorLight
	}
}

// Get appropriate theme icon based on current theme
func (p *page) getThemeIcon() app.UI {
	if p.CurrentTheme == sp.ThemeColorLight {
		return sp.Icon().Name("ui:Moon100")
	}
	return sp.Icon().Name("ui:Sun100")
}

func (p *page) renderSidenav() app.UI {
	// Create Sidenav with components
	return app.Div().
		Class("sidenav").
		Body(
			sp.Sidenav().
				Label("Component Navigation").
				AddToBody(sp.SidenavItem().Label("Button").Value("button").OnClick(
					func(ctx app.Context, e app.Event) {
						p.ctx.Navigate("/button")
					},
				)).
				AddToBody(sp.SidenavItem().Label("Accordion").Value("accordion").Href("/button")).
				AddToBody(sp.SidenavItem().Label("Button Group").Value("button-group").Href("/button-group")).
				AddToBody(sp.SidenavItem().Label("Checkbox").Value("checkbox").Href("/checkbox")).
				AddToBody(sp.SidenavItem().Label("Dropdown").Value("dropdown").Href("/dropdown")).
				AddToBody(sp.SidenavItem().Label("Label").Value("label").Href("/label")),
		)
}

func (p *page) renderContent() app.UI {
	// Content area - will be replaced with component showcase
	return app.Div().
		Body(
			p.PageContent...,
		)
}
