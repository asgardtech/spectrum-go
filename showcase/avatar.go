package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/asgardtech/spectrum-go/prism"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type avatarPage struct {
	app.Compo
}

func newAvatarPage() *avatarPage {
	return &avatarPage{}
}

func (p *avatarPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *avatarPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle(defaultTitle)
	ctx.Page().SetDescription(defaultDescription)
}

func (p *avatarPage) Render() app.UI {
	return prism.NewPage().Content(
		app.H1().Text("Avatar Component"),
		app.P().Text("An Avatar is a great way to feature a visual representation of a user."),

		// Different Sizes
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Avatar Sizes"),
				app.P().Text("Avatars come in various sizes from 50 to 700."),
				app.Div().Class("component-row").Body(
					sp.Avatar().
						Size("50").
						Label("User 50").
						Src("https://picsum.photos/500/500?random=1"),
					sp.Avatar().
						Size("75").
						Label("User 75").
						Src("https://picsum.photos/500/500?random=2"),
					sp.Avatar().
						Size("100").
						Label("User 100").
						Src("https://picsum.photos/500/500?random=3"),
				),
				app.Div().Class("component-row").Body(
					sp.Avatar().
						Size("200").
						Label("User 200").
						Src("https://picsum.photos/500/500?random=4"),
					sp.Avatar().
						Size("300").
						Label("User 300").
						Src("https://picsum.photos/500/500?random=5"),
				),
			),

		// Avatar as a link
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Avatar as a Link"),
				app.P().Text("Avatars can be made clickable by adding an href attribute."),
				app.Div().Class("component-row").Body(
					sp.Avatar().
						Size("100").
						Label("Clickable Avatar").
						Src("https://picsum.photos/500/500?random=6").
						Href("https://example.com").
						Target("_blank"),
				),
			),

		// Disabled Avatar
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Disabled Avatar"),
				app.P().Text("Avatars can be disabled."),
				app.Div().Class("component-row").Body(
					sp.Avatar().
						Size("100").
						Label("Disabled Avatar").
						Src("https://picsum.photos/500/500?random=7").
						Disabled(true),
				),
			),

		// Fallback for missing image
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Fallback for Missing Image"),
				app.P().Text("When an image isn't available, the avatar will display initials derived from the label."),
				app.Div().Class("component-row").Body(
					sp.Avatar().
						Size("100").
						Label("John Doe").
						Src("https://invalid-url/no-image.jpg"),
				),
			),
	)
}
