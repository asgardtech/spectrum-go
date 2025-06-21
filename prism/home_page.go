package prism

import "github.com/maxence-charriere/go-app/v10/pkg/app"

type homePage struct {
	app.Compo
}

func NewHomePage() *homePage {
	return &homePage{}
}

func (p *homePage) Render() app.UI {
	return NewLayout().
		WithTitle("Home").
		WithDescription("Home page").
		Content(
			app.H1().Text("Home"),
		)
}
