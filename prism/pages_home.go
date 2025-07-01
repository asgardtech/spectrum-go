package prism

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

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
			app.H1().
				Class("spectrum-Heading", "spectrum-Heading--sizeS").
				Text("Home"),
			app.P().
				Class("spectrum-Body spectrum-Body--sizeM").
				Text("This is the home page. It is the first page that the user sees when they log in."),
			app.Input().
				Type("checkbox").
				Checked(p.GetApp().GetUser().LoggedIn).
				OnChange(func(ctx app.Context, e app.Event) {
					if ctx.JSSrc().Get("checked").Bool() {
						p.GetApp().SetUser(
							User{
								Name:     "John Doe",
								LoggedIn: true,
								Email:    "john.doe@example.com",
								Avatar:   "https://github.com/shadcn.png",
							},
						)
					} else {
						p.GetApp().SetUser(User{})
					}
				}),
		)
}
