package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type infieldButtonPage struct {
	app.Compo
	ctx     app.Context
	clicked bool
}

func newInfieldButtonPage() *infieldButtonPage {
	return &infieldButtonPage{}
}

func (p *infieldButtonPage) OnMount(ctx app.Context) {
	p.ctx = ctx
}

func (p *infieldButtonPage) onClick(ctx app.Context, e app.Event) {
	p.clicked = true
	p.ctx.Update()
}

func (p *infieldButtonPage) Render() app.UI {
	return newPage().
		Content(
			app.H1().Text("Infield Button"),
			app.P().Text("A button intended for use inside a text field or search component."),
			app.If(p.clicked, func() app.UI {
				return app.P().Text("Infield button clicked!")
			}),
			app.H2().Text("Standard"),
			sp.FieldGroup().Horizontal(true).Body(
				sp.Textfield().Value("Some text"),
				sp.InfieldButton().Label("Search").Inline(sp.InfieldButtonInlineEnd).OnClick(p.onClick),
			),
		)
}
