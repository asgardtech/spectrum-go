package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/asgardtech/spectrum-go/prism"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type swatchGroupPage struct {
	app.Compo
}

func newSwatchGroupPage() *swatchGroupPage {
	return &swatchGroupPage{}
}

func (p *swatchGroupPage) Render() app.UI {
	return prism.NewLayout().Content(
		app.H1().Text("Swatch Group Component"),
		app.P().Text("Swatch Groups are used to organize and manage multiple color swatches."),

		// Basic Swatch Group
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Basic Swatch Group"),
				app.Div().Class("component-row").Body(
					sp.SwatchGroup().
						//Label("Basic Color Palette").
						Body(
							sp.Swatch().
								Color("#ff0000").
								Label("Red"),
							sp.Swatch().
								Color("#00ff00").
								Label("Green"),
							sp.Swatch().
								Color("#0000ff").
								Label("Blue"),
							sp.Swatch().
								Color("#ffff00").
								Label("Yellow"),
							sp.Swatch().
								Color("#ff00ff").
								Label("Magenta"),
							sp.Swatch().
								Color("#00ffff").
								Label("Cyan"),
						),
				),
			),

		// Selection Mode
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Selection Modes"),
				app.Div().
					Style("display", "grid").
					Style("grid-template-columns", "repeat(2, 1fr)").
					Style("gap", "20px").
					Body(
						app.Div().Body(
							app.H3().Text("Single Selection"),
							sp.SwatchGroup().
								// Label("Single Select")
								SelectsSingle().
								Body(
									sp.Swatch().
										Color("#ff5252").
										Label("Red"),
									sp.Swatch().
										Color("#4caf50").
										Label("Green").
										SetSelected(),
									sp.Swatch().
										Color("#2196f3").
										Label("Blue"),
									sp.Swatch().
										Color("#ffc107").
										Label("Yellow"),
								),
						),
						app.Div().Body(
							app.H3().Text("Multiple Selection"),
							sp.SwatchGroup().
								// Label("Multiple Select")
								SelectsMultiple().
								Body(
									sp.Swatch().
										Color("#ff5252").
										Label("Red").
										SetSelected(),
									sp.Swatch().
										Color("#4caf50").
										Label("Green"),
									sp.Swatch().
										Color("#2196f3").
										Label("Blue").
										SetSelected(),
									sp.Swatch().
										Color("#ffc107").
										Label("Yellow"),
								),
						),
					),
			),

		// Swatch Group Sizes
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Swatch Group Sizes"),
				app.Div().
					Style("display", "grid").
					Style("grid-template-columns", "repeat(2, 1fr)").
					Style("gap", "20px").
					Body(
						app.Div().Body(
							app.H3().Text("Extra Small"),
							sp.SwatchGroup().
								// Label("Extra Small").
								SizeXs().
								Body(
									sp.Swatch().
										Color("#ff0000").
										Label("Red"),
									sp.Swatch().
										Color("#00ff00").
										Label("Green"),
									sp.Swatch().
										Color("#0000ff").
										Label("Blue"),
								),
						),
						app.Div().Body(
							app.H3().Text("Small"),
							sp.SwatchGroup().
								// Label("Small").
								SizeS().
								Body(
									sp.Swatch().
										Color("#ff0000").
										Label("Red"),
									sp.Swatch().
										Color("#00ff00").
										Label("Green"),
									sp.Swatch().
										Color("#0000ff").
										Label("Blue"),
								),
						),
						app.Div().Body(
							app.H3().Text("Medium"),
							sp.SwatchGroup().
								// Label("Medium").
								SizeM().
								Body(
									sp.Swatch().
										Color("#ff0000").
										Label("Red"),
									sp.Swatch().
										Color("#00ff00").
										Label("Green"),
									sp.Swatch().
										Color("#0000ff").
										Label("Blue"),
								),
						),
						app.Div().Body(
							app.H3().Text("Large"),
							sp.SwatchGroup().
								// Label("Large").
								SizeL().
								Body(
									sp.Swatch().
										Color("#ff0000").
										Label("Red"),
									sp.Swatch().
										Color("#00ff00").
										Label("Green"),
									sp.Swatch().
										Color("#0000ff").
										Label("Blue"),
								),
						),
					),
			),

		// Border Styles
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Border Styles"),
				app.Div().
					Style("display", "grid").
					Style("grid-template-columns", "repeat(3, 1fr)").
					Style("gap", "20px").
					Body(
						app.Div().Body(
							app.H3().Text("Default"),
							sp.SwatchGroup().
								// Label("Default Border").
								Body(
									sp.Swatch().
										Color("#ffffff").
										Label("White"),
									sp.Swatch().
										Color("#f5f5f5").
										Label("Light Gray"),
									sp.Swatch().
										Color("#e0e0e0").
										Label("Gray"),
								),
						),
						app.Div().Body(
							app.H3().Text("None"),
							sp.SwatchGroup().
								// Label("No Border").
								BorderNone().
								Body(
									sp.Swatch().
										Color("#ffffff").
										Label("White"),
									sp.Swatch().
										Color("#f5f5f5").
										Label("Light Gray"),
									sp.Swatch().
										Color("#e0e0e0").
										Label("Gray"),
								),
						),
						app.Div().Body(
							app.H3().Text("Light"),
							sp.SwatchGroup().
								// Label("Light Border").
								BorderLight().
								Body(
									sp.Swatch().
										Color("#ffffff").
										Label("White"),
									sp.Swatch().
										Color("#f5f5f5").
										Label("Light Gray"),
									sp.Swatch().
										Color("#e0e0e0").
										Label("Gray"),
								),
						),
					),
			),

		// Density Options
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Density Options"),
				app.Div().
					Style("display", "grid").
					Style("grid-template-columns", "repeat(2, 1fr)").
					Style("gap", "20px").
					Body(
						app.Div().Body(
							app.H3().Text("Compact"),
							sp.SwatchGroup().
								// Label("Compact Density").
								DensityCompact().
								Body(
									sp.Swatch().
										Color("#ff0000").
										Label("Red"),
									sp.Swatch().
										Color("#00ff00").
										Label("Green"),
									sp.Swatch().
										Color("#0000ff").
										Label("Blue"),
									sp.Swatch().
										Color("#ffff00").
										Label("Yellow"),
								),
						),
						app.Div().Body(
							app.H3().Text("Spacious"),
							sp.SwatchGroup().
								// Label("Spacious Density").
								DensitySpacious().
								Body(
									sp.Swatch().
										Color("#ff0000").
										Label("Red"),
									sp.Swatch().
										Color("#00ff00").
										Label("Green"),
									sp.Swatch().
										Color("#0000ff").
										Label("Blue"),
									sp.Swatch().
										Color("#ffff00").
										Label("Yellow"),
								),
						),
					),
			),

		// Rounding Options
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Rounding Options"),
				app.Div().
					Style("display", "grid").
					Style("grid-template-columns", "repeat(3, 1fr)").
					Style("gap", "20px").
					Body(
						app.Div().Body(
							app.H3().Text("Default"),
							sp.SwatchGroup().
								// Label("Default Rounding").
								Body(
									sp.Swatch().
										Color("#ff0000").
										Label("Red"),
									sp.Swatch().
										Color("#00ff00").
										Label("Green"),
									sp.Swatch().
										Color("#0000ff").
										Label("Blue"),
								),
						),
						app.Div().Body(
							app.H3().Text("None"),
							sp.SwatchGroup().
								// Label("No Rounding").
								RoundingNone().
								Body(
									sp.Swatch().
										Color("#ff0000").
										Label("Red"),
									sp.Swatch().
										Color("#00ff00").
										Label("Green"),
									sp.Swatch().
										Color("#0000ff").
										Label("Blue"),
								),
						),
						app.Div().Body(
							app.H3().Text("Full"),
							sp.SwatchGroup().
								// Label("Full Rounding").
								RoundingFull().
								Body(
									sp.Swatch().
										Color("#ff0000").
										Label("Red"),
									sp.Swatch().
										Color("#00ff00").
										Label("Green"),
									sp.Swatch().
										Color("#0000ff").
										Label("Blue"),
								),
						),
					),
			),

		// Shape Options
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Shape Options"),
				app.Div().
					Style("display", "grid").
					Style("grid-template-columns", "repeat(2, 1fr)").
					Style("gap", "20px").
					Body(
						app.Div().Body(
							app.H3().Text("Default"),
							sp.SwatchGroup().
								// Label("Default Shape").
								Body(
									sp.Swatch().
										Color("#ff0000").
										Label("Red"),
									sp.Swatch().
										Color("#00ff00").
										Label("Green"),
									sp.Swatch().
										Color("#0000ff").
										Label("Blue"),
								),
						),
						app.Div().Body(
							app.H3().Text("Rectangle"),
							sp.SwatchGroup().
								// Label("Rectangle Shape").
								ShapeRectangle().
								Body(
									sp.Swatch().
										Color("#ff0000").
										Label("Red"),
									sp.Swatch().
										Color("#00ff00").
										Label("Green"),
									sp.Swatch().
										Color("#0000ff").
										Label("Blue"),
								),
						),
					),
			),
	)
}
