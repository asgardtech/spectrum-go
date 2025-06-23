package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/asgardtech/spectrum-go/prism"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type swatchPage struct {
	app.Compo
}

func newSwatchPage() *swatchPage {
	return &swatchPage{}
}

func (p *swatchPage) Render() app.UI {
	return prism.NewPage().Content(
		app.H1().Text("Swatch Component"),
		app.P().Text("Swatches are used to display, select and compare colors."),

		// Basic Swatches
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Basic Swatches"),
				app.Div().Class("component-row").Body(
					sp.Swatch().
						Color("#f00").
						Label("Red"),
					sp.Swatch().
						Color("#0f0").
						Label("Green"),
					sp.Swatch().
						Color("#00f").
						Label("Blue"),
					sp.Swatch().
						Color("#ff0").
						Label("Yellow"),
					sp.Swatch().
						Color("#f0f").
						Label("Magenta"),
					sp.Swatch().
						Color("#0ff").
						Label("Cyan"),
				),
			),

		// Swatch Sizes
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Swatch Sizes"),
				app.Div().Class("component-row").Body(
					sp.Swatch().
						Color("#f00").
						Label("Extra Small").
						SizeXs(),
					sp.Swatch().
						Color("#f00").
						Label("Small").
						SizeS(),
					sp.Swatch().
						Color("#f00").
						Label("Medium").
						SizeM(),
					sp.Swatch().
						Color("#f00").
						Label("Large").
						SizeL(),
				),
			),

		// Swatch Borders
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Swatch Borders"),
				app.Div().Class("component-row").Body(
					sp.Swatch().
						Color("#ffffff").
						Label("Default Border"),
					sp.Swatch().
						Color("#ffffff").
						Label("No Border").
						BorderNone(),
					sp.Swatch().
						Color("#ffffff").
						Label("Light Border").
						BorderLight(),
				),
			),

		// Swatch Rounding
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Swatch Rounding"),
				app.Div().Class("component-row").Body(
					sp.Swatch().
						Color("#f00").
						Label("Default Rounding"),
					sp.Swatch().
						Color("#f00").
						Label("No Rounding").
						RoundingNone(),
					sp.Swatch().
						Color("#f00").
						Label("Full Rounding").
						RoundingFull(),
				),
			),

		// Swatch Shape
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Swatch Shape"),
				app.Div().Class("component-row").Body(
					sp.Swatch().
						Color("#f00").
						Label("Default Shape"),
					sp.Swatch().
						Color("#f00").
						Label("Rectangle Shape").
						ShapeRectangle(),
				),
			),

		// Swatch States
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Swatch States"),
				app.Div().
					Style("display", "grid").
					Style("grid-template-columns", "repeat(2, 1fr)").
					Style("gap", "16px").
					Body(
						app.Div().Body(
							app.H3().Text("Normal"),
							sp.Swatch().
								Color("#f00").
								Label("Normal swatch"),
						),
						app.Div().Body(
							app.H3().Text("Selected"),
							sp.Swatch().
								Color("#f00").
								Label("Selected swatch").
								SetSelected(),
						),
						app.Div().Body(
							app.H3().Text("Disabled"),
							sp.Swatch().
								Color("#f00").
								Label("Disabled swatch").
								SetDisabled(),
						),
						app.Div().Body(
							app.H3().Text("Nothing"),
							sp.Swatch().
								Label("No color").
								SetNothing(),
						),
					),
			),

		// Mixed Value Swatch
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Mixed Value Swatch"),
				app.P().Text("Used when the swatch represents more than one color."),
				app.Div().Class("component-row").Body(
					sp.Swatch().
						Label("Mixed colors").
						SetMixedValue(),
				),
			),

		// Swatch Group
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Swatch Group"),
				app.P().Text("Swatches can be grouped together to make a color selection tool."),
				app.Div().Class("component-row").Body(
					sp.SwatchGroup().
						// Label("Color Palette")
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

		// Selectable Swatch Group
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Selectable Swatch Group"),
				app.Div().Class("component-row").Body(
					sp.SwatchGroup().
						// Label("Select Color")
						SelectsSingle().
						Body(
							sp.Swatch().
								Color("#ff0000").
								Label("Red"),
							sp.Swatch().
								Color("#00ff00").
								Label("Green").
								SetSelected(),
							sp.Swatch().
								Color("#0000ff").
								Label("Blue"),
							sp.Swatch().
								Color("#ffff00").
								Label("Yellow"),
						),
				),
			),

		// Color Transparency
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Color Transparency"),
				app.P().Text("Swatch can display colors with transparency."),
				app.Div().Class("component-row").Body(
					sp.Swatch().
						Color("rgba(255, 0, 0, 1)").
						Label("100% Opacity"),
					sp.Swatch().
						Color("rgba(255, 0, 0, 0.75)").
						Label("75% Opacity"),
					sp.Swatch().
						Color("rgba(255, 0, 0, 0.5)").
						Label("50% Opacity"),
					sp.Swatch().
						Color("rgba(255, 0, 0, 0.25)").
						Label("25% Opacity"),
					sp.Swatch().
						Color("rgba(255, 0, 0, 0)").
						Label("0% Opacity"),
				),
			),
	)
}
