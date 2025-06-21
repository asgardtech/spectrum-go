package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/asgardtech/spectrum-go/prism"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type actionMenuPage struct {
	app.Compo
	ctx         app.Context
	lastClicked string
}

func newActionMenuPage() *actionMenuPage {
	return &actionMenuPage{}
}

func (p *actionMenuPage) OnMount(ctx app.Context) {
	p.ctx = ctx
}

func (p *actionMenuPage) onClick(ctx app.Context, e app.Event) {
	menuItem := ctx.JSSrc()
	p.lastClicked = menuItem.Get("label").String()
	p.ctx.Update()
}

func (p *actionMenuPage) Render() app.UI {
	return prism.NewLayout().
		Content(
			app.H1().Text("Action Menu"),
			app.P().Text("An action menu offers a list of actions in a popover."),
			app.If(p.lastClicked != "",
				func() app.UI {
					return app.Div().Body(
						app.Text("Last clicked: "),
						app.Strong().Text(p.lastClicked),
					)
				},
			),
			app.H2().Text("Standard"),
			sp.ActionMenu().Body(
				sp.MenuItem().Label("Item 1").OnClick(p.onClick),
				sp.MenuItem().Label("Item 2").OnClick(p.onClick),
				sp.MenuItem().Label("Item 3").OnClick(p.onClick),
			),
			app.H2().Text("With Icons"),
			sp.ActionMenu().Body(
				sp.MenuItem().Label("Edit").Icon(sp.Icon().Name("ui:Edit")).OnClick(p.onClick),
				sp.MenuItem().Label("Copy").Icon(sp.Icon().Name("ui:Copy")).OnClick(p.onClick),
				sp.MenuItem().Label("Delete").Icon(sp.Icon().Name("ui:Delete")).OnClick(p.onClick),
			),
			app.H2().Text("With Selection"),
			sp.ActionMenu().Selects(sp.ActionMenuSelectsSingle).Body(
				sp.MenuItem().Label("Single 1").OnClick(p.onClick),
				sp.MenuItem().Label("Single 2").Selected(true).OnClick(p.onClick),
				sp.MenuItem().Label("Single 3").OnClick(p.onClick),
			),
		)
}
