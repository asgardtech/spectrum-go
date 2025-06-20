package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type actionGroupPage struct {
	app.Compo
	ctx         app.Context
	lastClicked string
}

func newActionGroupPage() *actionGroupPage {
	return &actionGroupPage{}
}

func (p *actionGroupPage) OnMount(ctx app.Context) {
	p.ctx = ctx
}

func (p *actionGroupPage) onClick(ctx app.Context, e app.Event) {
	button := ctx.JSSrc()
	p.lastClicked = button.Get("label").String()
	p.ctx.Update()
}

func (p *actionGroupPage) Render() app.UI {
	return newPage().
		Content(
			app.H1().Text("Action Group"),
			app.P().Text("An action group is a container for organizing action buttons."),
			app.If(p.lastClicked != "",
				func() app.UI {
					return app.Div().Body(
						app.Text("Last clicked: "),
						app.Strong().Text(p.lastClicked),
					)
				},
			),
			app.H2().Text("Standard"),
			sp.ActionGroup().Body(
				sp.ActionButton().Label("Button 1").OnClick(p.onClick),
				sp.ActionButton().Label("Button 2").OnClick(p.onClick),
				sp.ActionButton().Label("Button 3").OnClick(p.onClick),
			),
			app.H2().Text("Vertical"),
			sp.ActionGroup().Vertical(true).Body(
				sp.ActionButton().Label("Vertical 1").OnClick(p.onClick),
				sp.ActionButton().Label("Vertical 2").OnClick(p.onClick),
				sp.ActionButton().Label("Vertical 3").OnClick(p.onClick),
			),
			app.H2().Text("With Selection (Single)"),
			sp.ActionGroup().Selects(sp.ActionGroupSelectsSingle).Body(
				sp.ActionButton().Label("Single 1").OnClick(p.onClick),
				sp.ActionButton().Label("Single 2").Selected(true).OnClick(p.onClick),
				sp.ActionButton().Label("Single 3").OnClick(p.onClick),
			),
			app.H2().Text("With Selection (Multiple)"),
			sp.ActionGroup().Selects(sp.ActionGroupSelectsMultiple).Body(
				sp.ActionButton().Label("Multiple 1").OnClick(p.onClick),
				sp.ActionButton().Label("Multiple 2").Selected(true).OnClick(p.onClick),
				sp.ActionButton().Label("Multiple 3").Selected(true).OnClick(p.onClick),
			),
			app.H2().Text("Justified"),
			sp.ActionGroup().Justified(true).Body(
				sp.ActionButton().Label("Justified 1").OnClick(p.onClick),
				sp.ActionButton().Label("Justified 2").OnClick(p.onClick),
			),
		)
}
