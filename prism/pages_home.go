package prism

import "github.com/maxence-charriere/go-app/v10/pkg/app"

type HomePage struct {
	app.Compo
	Page
}

func HomePageConstructor(options PageOptions) PageConstructor {
	return func() IPage {
		return &HomePage{
			Page: *NewPage(options),
		}
	}
}

func (p *HomePage) Render() app.UI {
	return p.
		Content(
			app.H1().Text("Home"),
		)
}
