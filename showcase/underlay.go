package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/asgardtech/spectrum-go/prism"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type underlayPage struct {
	app.Compo
}

func newUnderlayPage() *underlayPage {
	return &underlayPage{}
}

func (p *underlayPage) Render() app.UI {
	return prism.NewPage().
		Content(
			app.H1().Text("Underlay"),
			app.P().Text("A component that provides a backdrop for modal dialogs."),
			app.H2().Text("Standard"),
			sp.Underlay(),
		)
}
