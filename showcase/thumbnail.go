package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type thumbnailPage struct {
	app.Compo
}

func newThumbnailPage() *thumbnailPage {
	return &thumbnailPage{}
}

func (p *thumbnailPage) Render() app.UI {
	return newPage().Content(
		app.H1().Text("Thumbnail Component"),
		app.P().Text("Thumbnails can be used to display a preview of an image, layer, or effect."),

		// Basic Thumbnail
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Basic Thumbnail"),
				app.P().Text("Basic thumbnails with different sizes."),
				app.Div().
					Class("component-row").
					Body(
						sp.Thumbnail().
							Size100().
							Image(
								app.Img().
									Src("https://picsum.photos/id/237/200").
									Alt("Thumbnail sample"),
							),
						sp.Thumbnail().
							Size200().
							Image(
								app.Img().
									Src("https://picsum.photos/id/1025/400").
									Alt("Thumbnail sample"),
							),
						sp.Thumbnail().
							Size300().
							Image(
								app.Img().
									Src("https://picsum.photos/id/1015/600").
									Alt("Thumbnail sample"),
							),
					),
			),

		// Cover Mode
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Cover Mode"),
				app.P().Text("When cover is set, the image will fill the entire thumbnail space."),
				app.Div().
					Class("component-row").
					Body(
						sp.Thumbnail().
							Size200().
							Image(
								app.Img().
									Src("https://picsum.photos/id/1084/400").
									Alt("Thumbnail without cover"),
							),
						sp.Thumbnail().
							Size200().
							SetCover().
							Image(
								app.Img().
									Src("https://picsum.photos/id/1084/400").
									Alt("Thumbnail with cover"),
							),
					),
			),

		// States
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Thumbnail States"),
				app.P().Text("Thumbnails can have different states such as selected, focused, and disabled."),
				app.Div().
					Class("component-row").
					Body(
						app.Div().
							Style("display", "flex").
							Style("flex-direction", "column").
							Style("align-items", "center").
							Style("margin-right", "20px").
							Body(
								sp.Thumbnail().
									Size100().
									Image(
										app.Img().
											Src("https://picsum.photos/id/1055/200").
											Alt("Normal thumbnail"),
									),
								app.Span().Text("Normal"),
							),
						app.Div().
							Style("display", "flex").
							Style("flex-direction", "column").
							Style("align-items", "center").
							Style("margin-right", "20px").
							Body(
								sp.Thumbnail().
									Size100().
									SetSelected().
									Image(
										app.Img().
											Src("https://picsum.photos/id/1055/200").
											Alt("Selected thumbnail"),
									),
								app.Span().Text("Selected"),
							),
						app.Div().
							Style("display", "flex").
							Style("flex-direction", "column").
							Style("align-items", "center").
							Style("margin-right", "20px").
							Body(
								sp.Thumbnail().
									Size100().
									SetFocused().
									Image(
										app.Img().
											Src("https://picsum.photos/id/1055/200").
											Alt("Focused thumbnail"),
									),
								app.Span().Text("Focused"),
							),
						app.Div().
							Style("display", "flex").
							Style("flex-direction", "column").
							Style("align-items", "center").
							Body(
								sp.Thumbnail().
									Size100().
									SetDisabled().
									Image(
										app.Img().
											Src("https://picsum.photos/id/1055/200").
											Alt("Disabled thumbnail"),
									),
								app.Span().Text("Disabled"),
							),
					),
			),

		// Layer Style
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Layer Style"),
				app.P().Text("Thumbnails can be displayed with a layer style for use in layer management interfaces."),
				app.Div().
					Class("component-row").
					Body(
						sp.Thumbnail().
							Size100().
							SetLayer().
							Image(
								app.Img().
									Src("https://picsum.photos/id/1036/300").
									Alt("Layer style thumbnail"),
							),
						sp.Thumbnail().
							Size100().
							SetLayer().
							SetSelected().
							Image(
								app.Img().
									Src("https://picsum.photos/id/1039/300").
									Alt("Selected layer style thumbnail"),
							),
					),
			),

		// Custom Background
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Custom Background"),
				app.P().Text("You can customize the letterboxing background."),
				app.Div().
					Class("component-row").
					Body(
						sp.Thumbnail().
							Size200().
							Background("#eaeaea").
							Image(
								app.Img().
									Src("https://picsum.photos/id/1010/200/200").
									Alt("Thumbnail with gray background"),
							),
						sp.Thumbnail().
							Size200().
							Background("#e0f2ff").
							Image(
								app.Img().
									Src("https://picsum.photos/id/1010/200/200").
									Alt("Thumbnail with light blue background"),
							),
						sp.Thumbnail().
							Size200().
							Background("repeating-conic-gradient(#606dbc, #606dbc 15%, #465298 0, #465298 30%)").
							Image(
								app.Img().
									Src("https://picsum.photos/id/1010/200/200").
									Alt("Thumbnail with pattern background"),
							),
					),
			),

		// In Grid
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Thumbnails in Grid"),
				app.P().Text("Thumbnails are often arranged in a grid layout."),
				app.Div().
					Style("display", "grid").
					Style("grid-template-columns", "repeat(3, 1fr)").
					Style("gap", "16px").
					Body(
						sp.Thumbnail().
							Size100().
							Image(
								app.Img().
									Src("https://picsum.photos/id/1003/300").
									Alt("Thumbnail 1"),
							),
						sp.Thumbnail().
							Size100().
							Image(
								app.Img().
									Src("https://picsum.photos/id/1004/300").
									Alt("Thumbnail 2"),
							),
						sp.Thumbnail().
							Size100().
							Image(
								app.Img().
									Src("https://picsum.photos/id/1005/300").
									Alt("Thumbnail 3"),
							),
						sp.Thumbnail().
							Size100().
							Image(
								app.Img().
									Src("https://picsum.photos/id/1006/300").
									Alt("Thumbnail 4"),
							),
						sp.Thumbnail().
							Size100().
							SetSelected().
							Image(
								app.Img().
									Src("https://picsum.photos/id/1008/300").
									Alt("Thumbnail 5"),
							),
						sp.Thumbnail().
							Size100().
							Image(
								app.Img().
									Src("https://picsum.photos/id/1009/300").
									Alt("Thumbnail 6"),
							),
					),
			),
	)
}