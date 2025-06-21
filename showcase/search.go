package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/asgardtech/spectrum-go/prism"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type searchPage struct {
	app.Compo
}

func newSearchPage() *searchPage {
	return &searchPage{}
}

func (p *searchPage) Render() app.UI {
	return prism.NewLayout().Content(
		app.H1().Text("Search Component"),
		app.P().Text("The Search component is a text field with built-in clear and search buttons to help users find content quickly."),

		// Basic search
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Basic Search"),
				app.Div().Class("component-row").Body(
					sp.Search().
						Label("Search").
						Placeholder("Search..."),
				),
			),

		// Search with pre-filled value
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Search with Value"),
				app.Div().Class("component-row").Body(
					sp.Search().
						Label("Search").
						Placeholder("Search...").
						Value("Adobe Spectrum"),
				),
			),

		// Disabled search
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Disabled Search"),
				app.Div().Class("component-row").Body(
					sp.Search().
						Label("Search").
						Placeholder("Search...").
						Disabled(true),
				),
			),

		// Search with help text
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Search with Help Text"),
				app.Div().Class("component-row").Body(
					sp.Search().
						Label("Search").
						Placeholder("Search..."),
					app.Div().
						Attr("slot", "help-text").
						Text("Enter keywords to search"),
				),
			),

		// Invalid search
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Invalid Search"),
				app.Div().Class("component-row").Body(
					sp.Search().
						Label("Search").
						Placeholder("Search...").
						Invalid(true).
						Value("@#$%"),
					app.Div().
						Attr("slot", "negative-help-text").
						Text("Special characters not allowed"),
				),
			),

		// Required search
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Required Search"),
				app.Div().Class("component-row").Body(
					sp.Search().
						Label("Search").
						Placeholder("Required search...").
						Required(true),
				),
			),

		// Quiet search
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Quiet Search"),
				app.Div().Class("component-row").Body(
					sp.Search().
						Label("Standard search").
						Placeholder("Standard search..."),
					sp.Search().
						Label("Quiet search").
						Placeholder("Quiet search...").
						Quiet(true),
				),
			),

		// Search sizes
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Search Sizes"),
				app.Div().Class("component-row").Body(
					app.Div().Style("display", "flex").Style("flex-direction", "column").Style("gap", "16px").Body(
						sp.Search().
							Label("Small (S)").
							Placeholder("Small search...").
							Size(sp.SearchSizeS),
						sp.Search().
							Label("Medium (M) - Default").
							Placeholder("Medium search...").
							Size(sp.SearchSizeM),
						sp.Search().
							Label("Large (L)").
							Placeholder("Large search...").
							Size(sp.SearchSizeL),
						sp.Search().
							Label("Extra Large (XL)").
							Placeholder("Extra large search...").
							Size(sp.SearchSizeXl),
					),
				),
			),

		// Multiline search
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Multiline Search"),
				app.Div().Class("component-row").Body(
					sp.Search().
						Label("Multiline search").
						Placeholder("Enter multiple lines to search...").
						Multiline(true).
						Rows(3),
				),
			),

		// Growing multiline search
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Growing Multiline Search"),
				app.Div().Class("component-row").Body(
					sp.Search().
						Label("Growing multiline search").
						Placeholder("Type multiple lines to see this grow...").
						Multiline(true).
						Grows(true),
				),
			),
	)
}
