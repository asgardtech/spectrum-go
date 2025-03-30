package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type checkboxPage struct {
	app.Compo
}

func newCheckboxPage() *checkboxPage {
	return &checkboxPage{}
}

func (p *checkboxPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *checkboxPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle(defaultTitle)
	ctx.Page().SetDescription(defaultDescription)
}

func (p *checkboxPage) Render() app.UI {
	// Create a main container with theme-aware styling
	return newPage().Content(
		app.Div().
			Body(
				sp.Checkbox().
					Body(
						app.Text("Checkbox"),
					).Checked(true),
			),
	)
}
