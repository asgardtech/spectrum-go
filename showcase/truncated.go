package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/asgardtech/spectrum-go/prism"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type truncatedPage struct {
	app.Compo
}

func newTruncatedPage() *truncatedPage {
	return &truncatedPage{}
}

func (p *truncatedPage) Render() app.UI {
	return prism.NewPage().
		Content(
			app.H1().Text("Truncated"),
			app.P().Text("A component for truncating text."),
			app.H2().Text("Standard"),
			sp.Truncated().Text("This is a very long text that will be truncated."),
		)
}
