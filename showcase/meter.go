package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type meterPage struct {
	app.Compo
}

func newMeterPage() *meterPage {
	return &meterPage{}
}

func (p *meterPage) Render() app.UI {
	return newPage().Content(
		app.H1().Text("Meter Component"),
		app.P().Text("Meters display progress as a percentage of a known quantity, across multiple classifications."),

		// Basic meter
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Basic Meter"),
				app.Div().Class("component-row").Body(
					sp.Meter().
						Label("Basic meter").
						Value(75).
						Max(100),
				),
			),

		// Meter variants
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Meter Variants"),
				app.Div().Class("component-row").Body(
					app.Div().Style("display", "flex").Style("flex-direction", "column").Style("gap", "24px").Body(
						sp.Meter().
							Label("Positive (Default)").
							Value(75).
							Max(100).
							Variant(sp.MeterVariantPositive),
						sp.Meter().
							Label("Negative").
							Value(75).
							Max(100).
							Variant(sp.MeterVariantNegative),
						sp.Meter().
							Label("Notice").
							Value(75).
							Max(100).
							Variant(sp.MeterVariantNotice),
						sp.Meter().
							Label("Neutral").
							Value(75).
							Max(100).
							Variant(sp.MeterVariantNeutral),
					),
				),
			),

		// Meter sizes
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Meter Sizes"),
				app.Div().Class("component-row").Body(
					app.Div().Style("display", "flex").Style("flex-direction", "column").Style("gap", "24px").Body(
						sp.Meter().
							Label("Small (S)").
							Value(75).
							Max(100).
							Size(sp.MeterSizeS),
						sp.Meter().
							Label("Medium (M) - Default").
							Value(75).
							Max(100).
							Size(sp.MeterSizeM),
						sp.Meter().
							Label("Large (L)").
							Value(75).
							Max(100).
							Size(sp.MeterSizeL),
					),
				),
			),

		// Different values
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Different Values"),
				app.Div().Class("component-row").Body(
					app.Div().Style("display", "flex").Style("flex-direction", "column").Style("gap", "24px").Body(
						sp.Meter().
							Label("Low value (25%)").
							Value(25).
							Max(100),
						sp.Meter().
							Label("Medium value (50%)").
							Value(50).
							Max(100),
						sp.Meter().
							Label("High value (90%)").
							Value(90).
							Max(100),
						sp.Meter().
							Label("Custom range (5 out of 10)").
							Value(5).
							Min(0).
							Max(10),
					),
				),
			),

		// With side label
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("With Side Label"),
				app.Div().Class("component-row").Body(
					app.Div().Style("display", "flex").Style("flex-direction", "column").Style("gap", "24px").Body(
						sp.Meter().
							Label("Available disk space").
							Value(75).
							Max(100).
							SideLabel("75 GB"),
						sp.Meter().
							Label("Battery").
							Value(25).
							Max(100).
							SideLabel("25%").
							Variant(sp.MeterVariantNegative),
						sp.Meter().
							Label("CPU usage").
							Value(90).
							Max(100).
							SideLabel("90%").
							Variant(sp.MeterVariantNotice),
					),
				),
			),

		// Over background
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Over Background"),
				app.Div().Class("component-row").Body(
					app.Div().Style("background-color", "#323232").Style("padding", "24px").Style("border-radius", "8px").Body(
						sp.Meter().
							Label("Over dark background").
							Value(75).
							Max(100).
							Overbackground(true),
					),
				),
			),

		// Custom meter label and value
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Custom Label and Value"),
				app.Div().Class("component-row").Body(
					app.Div().Style("display", "flex").Style("flex-direction", "column").Style("gap", "24px").Body(
						sp.Meter().
							Value(75).
							Max(100),
						app.Strong().
							Attr("slot", "label").
							Text("Custom formatted label"),
						app.Strong().
							Attr("slot", "value").
							Style("color", "blue").
							Text("75%"),
					),
				),
			),
	)
}