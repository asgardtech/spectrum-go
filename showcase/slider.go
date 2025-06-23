package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/asgardtech/spectrum-go/prism"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type sliderPage struct {
	app.Compo
}

func newSliderPage() *sliderPage {
	return &sliderPage{}
}

func (p *sliderPage) Render() app.UI {
	return prism.NewPage().Content(
		app.H1().Text("Slider Component"),
		app.P().Text("Sliders allow users to select a value within a range by moving a handle."),

		// Basic slider
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Basic Slider"),
				app.Div().Class("component-row").Body(
					sp.Slider().
						Label("Basic slider").
						Min(0).
						Max(100).
						Value(50),
				),
			),

		// Slider variants
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Slider Variants"),
				app.Div().Class("component-row").Body(
					app.Div().Style("display", "flex").Style("flex-direction", "column").Style("gap", "24px").Body(
						sp.Slider().
							Label("Filled (Default)").
							Min(0).
							Max(100).
							Value(50).
							Variant(sp.SliderVariantFilled),
						sp.Slider().
							Label("Tick").
							Min(0).
							Max(100).
							Value(50).
							Variant(sp.SliderVariantTick).
							TickStep(10),
						sp.Slider().
							Label("Ramp").
							Min(0).
							Max(100).
							Value(50).
							Variant(sp.SliderVariantRamp),
						sp.Slider().
							Label("Range").
							Min(0).
							Max(100).
							Value(50).
							Variant(sp.SliderVariantRange),
					),
				),
			),

		// Slider sizes
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Slider Sizes"),
				app.Div().Class("component-row").Body(
					app.Div().Style("display", "flex").Style("flex-direction", "column").Style("gap", "24px").Body(
						sp.Slider().
							Label("Small (S)").
							Min(0).
							Max(100).
							Value(50).
							Size(sp.SliderSizeS),
						sp.Slider().
							Label("Medium (M) - Default").
							Min(0).
							Max(100).
							Value(50).
							Size(sp.SliderSizeM),
						sp.Slider().
							Label("Large (L)").
							Min(0).
							Max(100).
							Value(50).
							Size(sp.SliderSizeL),
						sp.Slider().
							Label("Extra Large (XL)").
							Min(0).
							Max(100).
							Value(50).
							Size(sp.SliderSizeXl),
					),
				),
			),

		// Editable slider
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Editable Slider"),
				app.Div().Class("component-row").Body(
					sp.Slider().
						Label("Editable slider").
						Min(0).
						Max(100).
						Value(50).
						Editable(true),
				),
			),

		// Slider with tick labels
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Slider with Tick Labels"),
				app.Div().Class("component-row").Body(
					sp.Slider().
						Label("Slider with tick labels").
						Min(0).
						Max(100).
						Value(50).
						Variant(sp.SliderVariantTick).
						TickStep(20).
						TickLabels(true),
				),
			),

		// Disabled slider
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Disabled Slider"),
				app.Div().Class("component-row").Body(
					sp.Slider().
						Label("Disabled slider").
						Min(0).
						Max(100).
						Value(50).
						Disabled(true),
				),
			),

		// Quiet slider
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Quiet Slider"),
				app.Div().Class("component-row").Body(
					sp.Slider().
						Label("Standard slider").
						Min(0).
						Max(100).
						Value(50),
					sp.Slider().
						Label("Quiet slider").
						Min(0).
						Max(100).
						Value(50).
						Quiet(true),
				),
			),

		// Label visibility
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Label Visibility"),
				app.Div().Class("component-row").Body(
					app.Div().Style("display", "flex").Style("flex-direction", "column").Style("gap", "24px").Body(
						sp.Slider().
							Label("Text (Default)").
							Min(0).
							Max(100).
							Value(50).
							LabelVisibility(sp.SliderLabelVisibilityText),
						sp.Slider().
							Label("Value").
							Min(0).
							Max(100).
							Value(50).
							LabelVisibility(sp.SliderLabelVisibilityValue),
						sp.Slider().
							Label("None").
							Min(0).
							Max(100).
							Value(50).
							LabelVisibility(sp.SliderLabelVisibilityNone),
					),
				),
			),

		// Slider with fill start
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Slider with Fill Start"),
				app.Div().Class("component-row").Body(
					sp.Slider().
						Label("Fill start at 25").
						Min(0).
						Max(100).
						Value(75).
						FillStart(25).
						Variant(sp.SliderVariantFilled),
				),
			),
	)
}
