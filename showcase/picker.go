package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/asgardtech/spectrum-go/prism"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type pickerPage struct {
	app.Compo
	selectedValue string
}

func newPickerPage() *pickerPage {
	return &pickerPage{
		selectedValue: "option1",
	}
}

func (p *pickerPage) Render() app.UI {
	return prism.NewPage().Content(
		app.H1().Text("Picker Component"),
		app.P().Text("The Picker component allows users to select an option from a dropdown menu."),

		// Basic Picker
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Basic Picker"),
				app.Div().Class("component-row").Body(
					sp.Picker().
						Label("Select an option").
						Value("option1").
						Body(
							sp.Menu().
								Body(
									sp.MenuItem().Value("option1").Text("Option 1"),
									sp.MenuItem().Value("option2").Text("Option 2"),
									sp.MenuItem().Value("option3").Text("Option 3"),
								),
						),
				),
			),

		// Picker with Label
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Picker with Label"),
				app.Div().Class("component-row").Body(
					sp.Picker().
						Value("medium").
						AddToBody(
							app.Text("Size"),
						).
						Body(
							sp.Menu().
								Body(
									sp.MenuItem().Value("small").Text("Small"),
									sp.MenuItem().Value("medium").Text("Medium"),
									sp.MenuItem().Value("large").Text("Large"),
								),
						),
				),
			),

		// Picker with Icons
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Picker with Icons"),
				app.Div().Class("component-row").Body(
					sp.Picker().
						Label("Select a file type").
						Value("image").
						Body(
							sp.Menu().
								Body(
									sp.MenuItem().
										Value("image").
										Text("Image").
										Icon(sp.Icon().Name("ui:FileImage")),
									sp.MenuItem().
										Value("document").
										Text("Document").
										Icon(sp.Icon().Name("ui:FileDocument")),
									sp.MenuItem().
										Value("video").
										Text("Video").
										Icon(sp.Icon().Name("ui:FileVideo")),
								),
						),
				),
			),

		// Picker Icons Only
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Picker with Icons Only"),
				app.Div().Class("component-row").Body(
					sp.Picker().
						Label("Select color").
						Value("red").
						IconsOnly().
						Body(
							sp.Menu().
								Body(
									sp.MenuItem().
										Value("red").
										Text("Red").
										Icon(
											app.Div().
												Style("width", "18px").
												Style("height", "18px").
												Style("background-color", "red").
												Style("border-radius", "50%"),
										),
									sp.MenuItem().
										Value("green").
										Text("Green").
										Icon(
											app.Div().
												Style("width", "18px").
												Style("height", "18px").
												Style("background-color", "green").
												Style("border-radius", "50%"),
										),
									sp.MenuItem().
										Value("blue").
										Text("Blue").
										Icon(
											app.Div().
												Style("width", "18px").
												Style("height", "18px").
												Style("background-color", "blue").
												Style("border-radius", "50%"),
										),
								),
						),
				),
			),

		// Picker States
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Picker States"),
				app.Div().
					Style("display", "grid").
					Style("grid-template-columns", "repeat(2, 1fr)").
					Style("gap", "16px").
					Body(
						app.Div().Body(
							app.H3().Text("Disabled"),
							sp.Picker().
								Label("Disabled picker").
								SetDisabled().
								Value("option1").
								Body(
									sp.Menu().
										Body(
											sp.MenuItem().Value("option1").Text("Option 1"),
											sp.MenuItem().Value("option2").Text("Option 2"),
										),
								),
						),
						app.Div().Body(
							app.H3().Text("Invalid"),
							sp.Picker().
								Label("Invalid picker").
								SetInvalid().
								Value("option1").
								Body(
									sp.Menu().
										Body(
											sp.MenuItem().Value("option1").Text("Option 1"),
											sp.MenuItem().Value("option2").Text("Option 2"),
										),
								),
						),
						app.Div().Body(
							app.H3().Text("Quiet"),
							sp.Picker().
								Label("Quiet picker").
								SetQuiet().
								Value("option1").
								Body(
									sp.Menu().
										Body(
											sp.MenuItem().Value("option1").Text("Option 1"),
											sp.MenuItem().Value("option2").Text("Option 2"),
										),
								),
						),
						app.Div().Body(
							app.H3().Text("Readonly"),
							sp.Picker().
								Label("Readonly picker").
								SetReadonly().
								Value("option1").
								Body(
									sp.Menu().
										Body(
											sp.MenuItem().Value("option1").Text("Option 1"),
											sp.MenuItem().Value("option2").Text("Option 2"),
										),
								),
						),
					),
			),

		// Pending State
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Pending State"),
				app.Div().Class("component-row").Body(
					sp.Picker().
						Label("Loading options").
						SetPending().
						Pendinglabel("Loading...").
						Body(
							sp.Menu().
								Body(
									sp.MenuItem().Value("option1").Text("Option 1"),
								),
						),
				),
			),

		// Picker with Menu Groups
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Picker with Menu Groups"),
				app.Div().Class("component-row").Body(
					sp.Picker().
						Label("Select a fruit or vegetable").
						Value("apple").
						Body(
							sp.Menu().
								Body(
									sp.MenuGroup().
										Header(app.Text("Fruits")).
										Body(
											sp.MenuItem().Value("apple").Text("Apple"),
											sp.MenuItem().Value("banana").Text("Banana"),
											sp.MenuItem().Value("orange").Text("Orange"),
										),
									sp.MenuGroup().
										Header(app.Text("Vegetables")).
										Body(
											sp.MenuItem().Value("carrot").Text("Carrot"),
											sp.MenuItem().Value("broccoli").Text("Broccoli"),
											sp.MenuItem().Value("spinach").Text("Spinach"),
										),
								),
						),
				),
			),

		// Picker with Tooltip
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Picker with Tooltip"),
				app.Div().Class("component-row").Body(
					sp.Picker().
						Label("Select an option").
						Value("option1").
						Tooltip(
							sp.Tooltip().Text("Click to select an option from the dropdown"),
						).
						Body(
							sp.Menu().
								Body(
									sp.MenuItem().Value("option1").Text("Option 1"),
									sp.MenuItem().Value("option2").Text("Option 2"),
									sp.MenuItem().Value("option3").Text("Option 3"),
								),
						),
				),
			),
	)
}
