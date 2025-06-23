package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/asgardtech/spectrum-go/prism"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type pickerButtonPage struct {
	app.Compo
}

func newPickerButtonPage() *pickerButtonPage {
	return &pickerButtonPage{}
}

func (p *pickerButtonPage) Render() app.UI {
	return prism.NewPage().Content(
		app.H1().Text("Picker Button Component"),
		app.P().Text("The Picker Button component resembles a dropdown button that can be used to trigger popover interfaces."),

		// Basic Picker Button
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Basic Picker Button"),
				app.Div().Class("component-row").Body(
					sp.PickerButton().
						Label("Select color").
						Text("Select Color"),
				),
			),

		// Picker Button with Icon
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Picker Button with Icon"),
				app.Div().Class("component-row").Body(
					sp.PickerButton().
						Label("Select color").
						Text("With Icon").
						Icon(
							sp.Icon().Name("ui:ColorSelect"),
						),
				),
			),

		// Position Options
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Position Options"),
				app.Div().Class("component-row").Body(
					sp.PickerButton().
						Label("Left position").
						Text("Left").
						PositionLeft(),
					sp.PickerButton().
						Label("Right position").
						Text("Right").
						PositionRight(),
				),
			),

		// Button Types
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Button Types"),
				app.Div().Class("component-row").Body(
					sp.PickerButton().
						Label("Button type").
						Text("Button").
						TypeButton(),
					sp.PickerButton().
						Label("Submit type").
						Text("Submit").
						TypeSubmit(),
					sp.PickerButton().
						Label("Reset type").
						Text("Reset").
						TypeReset(),
				),
			),

		// States
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Picker Button States"),
				app.Div().
					Style("display", "grid").
					Style("grid-template-columns", "repeat(2, 1fr)").
					Style("gap", "16px").
					Body(
						app.Div().Body(
							app.H3().Text("Normal"),
							sp.PickerButton().
								Label("Normal state").
								Text("Normal"),
						),
						app.Div().Body(
							app.H3().Text("Active"),
							sp.PickerButton().
								Label("Active state").
								Text("Active").
								SetActive(),
						),
						app.Div().Body(
							app.H3().Text("Disabled"),
							sp.PickerButton().
								Label("Disabled state").
								Text("Disabled").
								SetDisabled(),
						),
						app.Div().Body(
							app.H3().Text("Quiet"),
							sp.PickerButton().
								Label("Quiet style").
								Text("Quiet").
								SetQuiet(),
						),
					),
			),

		// Rounded Style
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Rounded Style"),
				app.Div().Class("component-row").Body(
					sp.PickerButton().
						Label("Regular").
						Text("Regular"),
					sp.PickerButton().
						Label("Rounded").
						Text("Rounded").
						SetRounded(),
				),
			),

		// As Link
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("As Link"),
				app.Div().Class("component-row").Body(
					sp.PickerButton().
						Label("External link").
						Text("GitHub").
						Href("https://github.com/asgardtech/spectrum-go").
						Icon(sp.Icon().Name("ui:ExternalLink")),
				),
			),

		// With Click Event
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("With Click Event"),
				app.Div().Class("component-row").Body(
					sp.PickerButton().
						Label("Click me").
						Text("Click me").
						OnClick(func(ctx app.Context, e app.Event) {
							app.Window().Get("alert").Invoke("Button was clicked!")
						}),
				),
			),

		// In a Form
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("In a Form"),
				app.Div().Class("component-row").Body(
					app.Form().
						Style("display", "flex").
						Style("gap", "10px").
						Style("align-items", "center").
						Body(
							sp.Textfield().
								Label("Search term").
								Placeholder("Enter search term"),
							sp.PickerButton().
								Label("Submit search").
								Text("Search").
								TypeSubmit().
								Icon(sp.Icon().Name("ui:Magnify")),
							sp.PickerButton().
								Label("Reset form").
								Text("Reset").
								TypeReset().
								Icon(sp.Icon().Name("ui:Close")),
						),
				),
			),
	)
}
