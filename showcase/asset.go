package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/asgardtech/spectrum-go/prism"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type assetPage struct {
	app.Compo
}

func newAssetPage() *assetPage {
	return &assetPage{}
}

func (p *assetPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *assetPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle(defaultTitle)
	ctx.Page().SetDescription(defaultDescription)
}

func (p *assetPage) Render() app.UI {
	return prism.NewLayout().Content(
		app.H1().Text("Asset Component"),
		app.P().Text("The Asset component is used to visually represent files, folders, or images in your application. It provides consistent styling and layout for different types of media content."),

		// Image Asset
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Image Asset"),
				app.P().Text("Use the Asset component to display images with proper containment and centering."),
				app.Div().
					Style("height", "128px").
					Body(
						sp.Asset().
							Body(
								app.Img().
									Src("https://picsum.photos/500/500").
									Alt("Demo Image"),
							),
					),
			),

		// File Asset
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("File Asset"),
				app.P().Text("Display file assets with or without labels."),
				app.Div().
					Style("display", "flex").
					Style("gap", "16px").
					Body(
						sp.Asset().
							Variant("file"),
						sp.Asset().
							Variant("file").
							Label("document.pdf"),
					),
			),

		// Folder Asset
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Folder Asset"),
				app.P().Text("Display folder assets with or without labels."),
				app.Div().
					Style("display", "flex").
					Style("gap", "16px").
					Body(
						sp.Asset().
							Variant("folder"),
						sp.Asset().
							Variant("folder").
							Label("Documents"),
					),
			),

		// Custom Content
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Custom Content"),
				app.P().Text("You can add custom content to the Asset component when needed."),
				app.Div().
					Style("height", "128px").
					Body(
						sp.Asset().
							Body(
								app.Div().
									Style("display", "flex").
									Style("flex-direction", "column").
									Style("align-items", "center").
									Style("justify-content", "center").
									Style("height", "100%").
									Body(
										sp.Icon().
											Name("File").
											Size("xl"),
										app.Text("Custom Asset"),
									),
							),
					),
			),
	)
}
