package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type page struct {
	app.Compo

	currentTheme sp.ThemeColor

	// Content
	content []app.UI

	// Context
	ctx app.Context
}

func newPage() *page {
	return &page{
		currentTheme: sp.ThemeColorLight,
	}
}

func (p *page) Content(v ...app.UI) *page {
	p.content = app.FilterUIElems(v...)
	return p
}

func (p *page) OnMount(ctx app.Context) {
	p.ctx = ctx
	ctx.ObserveState("current-theme", &p.currentTheme)
}

func (p *page) Render() app.UI {
	// Create a main container with theme-aware styling
	return sp.Theme().
		Color(p.currentTheme).
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
									app.Log("Current Theme", p.currentTheme)
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
		)
}

// Toggle theme between light and dark
func (p *page) toggleTheme() {
	currentTheme := p.currentTheme
	if currentTheme == "" {
		currentTheme = sp.ThemeColorLight
	}

	if currentTheme == sp.ThemeColorLight {
		currentTheme = sp.ThemeColorDark
	} else {
		currentTheme = sp.ThemeColorLight
	}

	p.ctx.SetState("current-theme", currentTheme).
		Persist().
		Broadcast()
}

func (p *page) renderSidenav() app.UI {
	// Create Sidenav with components
	return app.Div().
		Class("spectrum-sidenav").
		Body(
			app.Div().
				Class("spectrum-sidenav-heading").
				Text("Components"),
			app.Div().
				Class("spectrum-sidenav-body").
				Body(
					sp.Sidenav().
						Label("Component Navigation").
						AddToBody(sp.SidenavItem().
							Label("Home").
							Value("/").
							Href("/"),
						).
						// Implemented components (mark as active)
						AddToBody(sp.SidenavItem().
							Label("Button").
							Value("button").
							Href("/button"),
						).
						AddToBody(sp.SidenavItem().
							Label("Checkbox").
							Value("checkbox").
							Href("/checkbox"),
						).
						AddToBody(sp.SidenavItem().
							Label("Badge").
							Value("badge").
							Href("/badge"),
						).
						AddToBody(sp.SidenavItem().
							Label("Radio").
							Value("radio").
							Href("/radio"),
						).
						AddToBody(sp.SidenavItem().
							Label("Switch").
							Value("switch").
							Href("/switch"),
						).
						// New components added
						AddToBody(sp.SidenavItem().
							Label("Progress Bar").
							Value("progress-bar").
							Href("/progress-bar"),
						).
						AddToBody(sp.SidenavItem().
							Label("Progress Circle").
							Value("progress-circle").
							Href("/progress-circle"),
						).
						AddToBody(sp.SidenavItem().
							Label("Link").
							Value("link").
							Href("/link"),
						).
						AddToBody(sp.SidenavItem().
							Label("Icon").
							Value("icon").
							Href("/icon"),
						).
						AddToBody(sp.SidenavItem().
							Label("Avatar").
							Value("avatar").
							Href("/avatar"),
						).
						AddToBody(sp.SidenavItem().
							Label("Status Light").
							Value("status-light").
							Href("/status-light"),
						).
						AddToBody(sp.SidenavItem().
							Label("Accordion").
							Value("accordion").
							Href("/accordion"),
						).
						AddToBody(sp.SidenavItem().
							Label("Coachmark").
							Value("coachmark").
							Href("/coachmark"),
						).
						AddToBody(sp.SidenavItem().
							Label("Alert Banner").
							Value("alert-banner").
							Href("/alert-banner"),
						).
						AddToBody(sp.SidenavItem().
							Label("Alert Dialog").
							Value("alert-dialog").
							Href("/alert-dialog"),
						).
						AddToBody(sp.SidenavItem().
							Label("Asset").
							Value("asset").
							Href("/asset"),
						).
						AddToBody(sp.SidenavItem().
							Label("Button Group").
							Value("button-group").
							Href("/button-group"),
						).
						AddToBody(sp.SidenavItem().
							Label("Card").
							Value("card").
							Href("/card"),
						).
						AddToBody(sp.SidenavItem().
							Label("Toast").
							Value("toast").
							Href("/toast"),
						).
						AddToBody(sp.SidenavItem().
							Label("Tooltip").
							Value("tooltip").
							Href("/tooltip"),
						).
						AddToBody(sp.SidenavItem().
							Label("Breadcrumbs").
							Value("breadcrumbs").
							Href("/breadcrumbs"),
						).
						AddToBody(sp.SidenavItem().
							Label("Tabs").
							Value("tabs").
							Href("/tabs"),
						).
						AddToBody(sp.SidenavItem().
							Label("Tags").
							Value("tags").
							Href("/tags"),
						).
						AddToBody(sp.SidenavItem().
							Label("Divider").
							Value("divider").
							Href("/divider"),
						).
						AddToBody(sp.SidenavItem().
							Label("Textfield").
							Value("textfield").
							Href("/textfield"),
						).
						AddToBody(sp.SidenavItem().
							Label("Textarea").
							Value("textarea").
							Href("/textarea"),
						).
						AddToBody(sp.SidenavItem().
							Label("Popover").
							Value("popover").
							Href("/popover"),
						).
						AddToBody(sp.SidenavItem().
							Label("Menu").
							Value("menu").
							Href("/menu"),
						).
						AddToBody(sp.SidenavItem().
							Label("Meter").
							Value("meter").
							Href("/meter"),
						).
						AddToBody(sp.SidenavItem().
							Label("Search").
							Value("search").
							Href("/search"),
						).
						AddToBody(sp.SidenavItem().
							Label("Field Label").
							Value("field-label").
							Href("/field-label"),
						).
						AddToBody(sp.SidenavItem().
							Label("Help Text").
							Value("help-text").
							Href("/help-text"),
						).
						AddToBody(sp.SidenavItem().
							Label("Slider").
							Value("slider").
							Href("/slider"),
						).
						AddToBody(sp.SidenavItem().
							Label("Sidenav").
							Value("sidenav").
							Href("/sidenav"),
						).
						AddToBody(sp.SidenavItem().
							Label("Number Field").
							Value("number-field").
							Href("/number-field"),
						).
						AddToBody(sp.SidenavItem().
							Label("Combobox").
							Value("combobox").
							Href("/combobox"),
						).
						AddToBody(sp.SidenavItem().
							Label("Split View").
							Value("split-view").
							Href("/split-view"),
						).
						AddToBody(sp.SidenavItem().
							Label("Dialog").
							Value("dialog").
							Href("/dialog"),
						).
						AddToBody(sp.SidenavItem().
							Label("Contextual Help").
							Value("contextual-help").
							Href("/contextual-help"),
						).
						AddToBody(sp.SidenavItem().
							Label("Field Group").
							Value("field-group").
							Href("/field-group"),
						).
						AddToBody(sp.SidenavItem().
							Label("Illustrated Message").
							Value("illustrated-message").
							Href("/illustrated-message"),
						).
						AddToBody(sp.SidenavItem().
							Label("Thumbnail").
							Value("thumbnail").
							Href("/thumbnail"),
						).
						AddToBody(sp.SidenavItem().
							Label("Picker").
							Value("picker").
							Href("/picker"),
						).
						AddToBody(sp.SidenavItem().
							Label("Picker Button").
							Value("picker-button").
							Href("/picker-button"),
						).
						AddToBody(sp.SidenavItem().
							Label("Swatch").
							Value("swatch").
							Href("/swatch"),
						).
						AddToBody(sp.SidenavItem().
							Label("Swatch Group").
							Value("swatch-group").
							Href("/swatch-group"),
						).
						// Not yet implemented components
						AddToBody(sp.SidenavItem().Label("Accordion Item").Value("accordion-item").Disabled(true)).
						AddToBody(sp.SidenavItem().Label("Action Bar").Value("action-bar").Disabled(true)).
						AddToBody(sp.SidenavItem().
							Label("Action Button").
							Value("action-button").
							Href("/action-button"),
						).
						AddToBody(sp.SidenavItem().Label("Action Group").Value("action-group").Disabled(true)).
						AddToBody(sp.SidenavItem().
							Label("Action Menu").
							Value("action-menu").
							Href("/action-menu"),
						),
				),
		)
}

func (p *page) renderContent() app.UI {
	// Content area - will be replaced with component showcase
	return app.Div().
		Body(
			p.content...,
		)
}
