package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/asgardtech/spectrum-go/prism"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// trayPage represents the Tray component showcase
type trayPage struct {
	app.Compo
	isDialogTrayOpen  bool
	isMenuTrayOpen    bool
	isContentTrayOpen bool
}

// newTrayPage creates a new Tray component showcase
func newTrayPage() *trayPage {
	return &trayPage{}
}

// OnNav initializes the page when navigated to
func (p *trayPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

// initPage initializes the page
func (p *trayPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle(defaultTitle)
	ctx.Page().SetDescription(defaultDescription)
}

// toggleDialogTray toggles the dialog tray visibility
func (p *trayPage) toggleDialogTray(ctx app.Context, e app.Event) {
	p.isDialogTrayOpen = !p.isDialogTrayOpen
	ctx.Dispatch(func(ctx app.Context) {})
}

// toggleMenuTray toggles the menu tray visibility
func (p *trayPage) toggleMenuTray(ctx app.Context, e app.Event) {
	p.isMenuTrayOpen = !p.isMenuTrayOpen
	ctx.Dispatch(func(ctx app.Context) {})
}

// toggleContentTray toggles the content tray visibility
func (p *trayPage) toggleContentTray(ctx app.Context, e app.Event) {
	p.isContentTrayOpen = !p.isContentTrayOpen
	ctx.Dispatch(func(ctx app.Context) {})
}

// handleTrayClose handles the close event for trays
func (p *trayPage) handleTrayClose(ctx app.Context, e app.Event) {
	p.isDialogTrayOpen = false
	p.isMenuTrayOpen = false
	p.isContentTrayOpen = false
	ctx.Dispatch(func(ctx app.Context) {})
}

// Render renders the Tray component showcase
func (p *trayPage) Render() app.UI {
	return prism.NewPage().
		Content(
			app.H1().Text("Tray"),
			app.P().Text("A container that slides up from the bottom of the screen."),
			app.H2().Text("Standard"),
			sp.OverlayTrigger().
				Type("modal").
				Trigger(
					sp.ActionButton().Label("Show Tray"),
				).
				ClickContent(
					sp.Tray().Body(
						sp.Dialog().Body(
							app.H1().Text("Tray Content"),
						),
					),
				),
		)
}
