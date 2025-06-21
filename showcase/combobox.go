package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/asgardtech/spectrum-go/prism"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type comboboxPage struct {
	app.Compo
}

func newComboboxPage() *comboboxPage {
	return &comboboxPage{}
}

func (p *comboboxPage) Render() app.UI {
	return prism.NewLayout().Content(
		app.H1().Text("Combobox Component"),
		app.P().Text("Comboboxes allow users to filter lists to only options matching a query, combining a textfield, picker button, and menu items."),

		// Basic combobox
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Basic Combobox"),
				app.Div().Class("component-row").Body(
					sp.Combobox().
						Label("Color").
						Body(
							sp.MenuItem().Label("Red").Value("red").Text("Red"),
							sp.MenuItem().Label("Green").Value("green").Text("Green"),
							sp.MenuItem().Label("Blue").Value("blue").Text("Blue"),
							sp.MenuItem().Label("Yellow").Value("yellow").Text("Yellow"),
							sp.MenuItem().Label("Purple").Value("purple").Text("Purple"),
							sp.MenuItem().Label("Orange").Value("orange").Text("Orange"),
						),
				),
			),

		// Combobox with placeholder
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Combobox with Placeholder"),
				app.Div().Class("component-row").Body(
					sp.Combobox().
						Label("Country").
						Placeholder("Select a country").
						Body(
							sp.MenuItem().Label("United States").Value("us").Text("United States"),
							sp.MenuItem().Label("Canada").Value("ca").Text("Canada"),
							sp.MenuItem().Label("Mexico").Value("mx").Text("Mexico"),
							sp.MenuItem().Label("United Kingdom").Value("uk").Text("United Kingdom"),
							sp.MenuItem().Label("France").Value("fr").Text("France"),
							sp.MenuItem().Label("Germany").Value("de").Text("Germany"),
							sp.MenuItem().Label("Japan").Value("jp").Text("Japan"),
						),
				),
			),

		// Combobox with auto-complete
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Combobox with Auto-complete"),
				app.Div().Class("component-row").Body(
					sp.Combobox().
						Label("Fruit").
						Placeholder("Start typing...").
						Autocomplete(sp.ComboboxAutocompleteList).
						Body(
							sp.MenuItem().Label("Apple").Value("apple").Text("Apple"),
							sp.MenuItem().Label("Banana").Value("banana").Text("Banana"),
							sp.MenuItem().Label("Cherry").Value("cherry").Text("Cherry"),
							sp.MenuItem().Label("Durian").Value("durian").Text("Durian"),
							sp.MenuItem().Label("Elderberry").Value("elderberry").Text("Elderberry"),
							sp.MenuItem().Label("Fig").Value("fig").Text("Fig"),
							sp.MenuItem().Label("Grape").Value("grape").Text("Grape"),
						),
				),
			),

		// Disabled combobox
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Disabled Combobox"),
				app.Div().Class("component-row").Body(
					sp.Combobox().
						Label("Size").
						Disabled(true).
						Body(
							sp.MenuItem().Label("Small").Value("s").Text("Small"),
							sp.MenuItem().Label("Medium").Value("m").Text("Medium"),
							sp.MenuItem().Label("Large").Value("l").Text("Large"),
							sp.MenuItem().Label("Extra Large").Value("xl").Text("Extra Large"),
						),
				),
			),

		// Invalid combobox
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Invalid Combobox"),
				app.Div().Class("component-row").Body(
					sp.Combobox().
						Label("Category").
						Invalid(true).
						Body(
							sp.MenuItem().Label("Option 1").Value("option1").Text("Option 1"),
							sp.MenuItem().Label("Option 2").Value("option2").Text("Option 2"),
							sp.MenuItem().Label("Option 3").Value("option3").Text("Option 3"),
						),
					app.Div().
						Attr("slot", "negative-help-text").
						Text("A valid category must be selected"),
				),
			),

		// Required combobox
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Required Combobox"),
				app.Div().Class("component-row").Body(
					sp.Combobox().
						Label("Department").
						Placeholder("Select a department").
						Required(true).
						Body(
							sp.MenuItem().Label("Engineering").Value("engineering").Text("Engineering"),
							sp.MenuItem().Label("Design").Value("design").Text("Design"),
							sp.MenuItem().Label("Marketing").Value("marketing").Text("Marketing"),
							sp.MenuItem().Label("Sales").Value("sales").Text("Sales"),
							sp.MenuItem().Label("Support").Value("support").Text("Support"),
						),
				),
			),

		// Quiet combobox
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Quiet Combobox"),
				app.Div().Class("component-row").Body(
					app.Div().Style("display", "flex").Style("flex-direction", "column").Style("gap", "16px").Body(
						sp.Combobox().
							Label("Standard combobox").
							Placeholder("Select an option").
							Body(
								sp.MenuItem().Label("Option 1").Value("option1").Text("Option 1"),
								sp.MenuItem().Label("Option 2").Value("option2").Text("Option 2"),
								sp.MenuItem().Label("Option 3").Value("option3").Text("Option 3"),
							),
						sp.Combobox().
							Label("Quiet combobox").
							Placeholder("Select an option").
							Quiet(true).
							Body(
								sp.MenuItem().Label("Option 1").Value("option1").Text("Option 1"),
								sp.MenuItem().Label("Option 2").Value("option2").Text("Option 2"),
								sp.MenuItem().Label("Option 3").Value("option3").Text("Option 3"),
							),
					),
				),
			),

		// Pending combobox
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Pending Combobox"),
				app.Div().Class("component-row").Body(
					sp.Combobox().
						Label("Loading Options").
						Placeholder("Loading...").
						Pending(true).
						Pendinglabel("Loading options...").
						Body(),
				),
			),
	)
}
