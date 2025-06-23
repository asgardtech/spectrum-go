package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/asgardtech/spectrum-go/prism"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type opacityCheckerboardPage struct {
	app.Compo
}

func newOpacityCheckerboardPage() *opacityCheckerboardPage {
	return &opacityCheckerboardPage{}
}

func (p *opacityCheckerboardPage) Render() app.UI {
	return prism.NewPage().
		Content(
			app.H1().Text("Opacity Checkerboard"),
			app.P().Text("A checkerboard pattern to display behind transparent images."),
			app.H2().Text("Standard"),
			sp.OpacityCheckerboard().Style("width", "200px").Style("height", "100px"),
		)
}
