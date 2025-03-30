package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type buttonPage struct {
	app.Compo
}

func newButtonPage() *buttonPage {
	return &buttonPage{}
}

func (p *buttonPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}
func (p *buttonPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle(defaultTitle)
	ctx.Page().SetDescription(defaultDescription)
}

func (p *buttonPage) Render() app.UI {
	// Create a main container with theme-aware styling
	return newPage().Content(
		app.Div().
			Body(
				sp.Button().
					Label("Button").
					Text("Button"),
			),
	)
}
