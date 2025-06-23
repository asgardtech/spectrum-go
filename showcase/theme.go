package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/asgardtech/spectrum-go/prism"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type themePage struct {
	app.Compo
}

func newThemePage() *themePage {
	return &themePage{}
}

func (p *themePage) Render() app.UI {
	return prism.NewPage().
		Content(
			app.H1().Text("Theme"),
			app.P().Text("A component to apply a theme to its children."),
			app.H2().Text("Light Theme"),
			sp.Theme().Color(sp.ThemeColorLight).Scale(sp.ThemeScaleMedium).Body(
				sp.Button().Variant(sp.ButtonVariantPrimary).Label("Themed Button"),
			),
			app.H2().Text("Dark Theme"),
			sp.Theme().Color(sp.ThemeColorDark).Scale(sp.ThemeScaleMedium).Body(
				sp.Button().Variant(sp.ButtonVariantPrimary).Label("Themed Button"),
			),
		)
}
