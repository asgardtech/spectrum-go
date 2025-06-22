package prism

import "github.com/maxence-charriere/go-app/v10/pkg/app"

type HomePage struct {
	app.Compo
}

func NewHomePage() *HomePage {
	return &HomePage{}
}

func (p *HomePage) Render() app.UI {
	return NewLayout().
		WithTitle("Home").
		WithDescription("Home page").
		Content(
			app.H1().Text("Home"),
		)
}
