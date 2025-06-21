package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/asgardtech/spectrum-go/prism"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type menuPage struct {
	app.Compo
}

func newMenuPage() *menuPage {
	return &menuPage{}
}

func (p *menuPage) Render() app.UI {
	return prism.NewLayout().Content(
		app.H1().Text("Menu Component"),
		app.P().Text("Menus display a list of options or commands, typically triggered from a button."),

		// Basic menu
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Basic Menu"),
				app.Div().Class("component-row").Style("height", "320px").Body(
					sp.Menu().
						Style("width", "200px").
						Body(
							sp.MenuItem().Label("Menu Item 1").Text("Menu Item 1"),
							sp.MenuItem().Label("Menu Item 2").Text("Menu Item 2"),
							sp.MenuItem().Label("Menu Item 3").Text("Menu Item 3"),
							sp.MenuItem().Label("Menu Item 4").Text("Menu Item 4"),
						),
				),
			),

		// Menu with dividers and groups
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Menu with Groups and Dividers"),
				app.Div().Class("component-row").Style("height", "400px").Body(
					sp.Menu().
						Style("width", "200px").
						Body(
							app.Div().
								Attr("slot", "header").
								Text("Group 1"),
							sp.MenuItem().Label("Item 1.1").Text("Item 1.1"),
							sp.MenuItem().Label("Item 1.2").Text("Item 1.2"),
							sp.Divider(),
							app.Div().
								Attr("slot", "header").
								Text("Group 2"),
							sp.MenuItem().Label("Item 2.1").Text("Item 2.1"),
							sp.MenuItem().Label("Item 2.2").Text("Item 2.2"),
						),
				),
			),

		// Menu items with icons
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Menu Items with Icons"),
				app.Div().Class("component-row").Style("height", "320px").Body(
					sp.Menu().
						Style("width", "200px").
						Body(
							sp.MenuItem().Label("New").Text("New").Icon(sp.Icon().Name("ui:Add")),
							sp.MenuItem().Label("Open").Text("Open").Icon(sp.Icon().Name("ui:FolderOpen")),
							sp.MenuItem().Label("Save").Text("Save").Icon(sp.Icon().Name("ui:SaveFloppy")),
							sp.MenuItem().Label("Delete").Text("Delete").Icon(sp.Icon().Name("ui:Delete")),
						),
				),
			),

		// Menu with selected and disabled items
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Selected and Disabled Items"),
				app.Div().Class("component-row").Style("height", "320px").Body(
					sp.Menu().
						Style("width", "200px").
						Body(
							sp.MenuItem().Label("Regular Item").Text("Regular Item"),
							sp.MenuItem().Label("Selected Item").Text("Selected Item").Selected(true),
							sp.MenuItem().Label("Disabled Item").Text("Disabled Item").Disabled(true),
							sp.MenuItem().Label("Selected & Disabled").Text("Selected & Disabled").Selected(true).Disabled(true),
						),
				),
			),

		// Menu items with values and descriptions
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Menu Items with Values and Descriptions"),
				app.Div().Class("component-row").Style("height", "400px").Body(
					sp.Menu().
						Style("width", "250px").
						Body(
							sp.MenuItem().
								Label("Item with Value").
								Text("Item with Value").
								ValueContent(app.Text("value1")),
							sp.MenuItem().
								Label("Item with Description").
								Text("Item with Description"),
							app.Div().
								Attr("slot", "description").
								Text("This is a description for the menu item."),
							sp.MenuItem().
								Label("Item with Value & Description").
								Text("Item with Value & Description"),
							app.Span().
								Attr("slot", "value").
								Text("value2"),
							app.Div().
								Attr("slot", "description").
								Text("Item with both value and description."),
						),
				),
			),

		// Nested menu items
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Nested Menu Items"),
				app.Div().Class("component-row").Style("height", "400px").Body(
					sp.Menu().
						Style("width", "250px").
						Body(
							sp.MenuItem().Label("Level 1 Item 1").Text("Level 1 Item 1"),
							sp.MenuItem().
								Label("Level 1 Item 2").
								Text("Level 1 Item 2").
								SetHasSubmenu(),
							app.Div().
								Attr("slot", "submenu").
								Body(
									sp.Menu().Body(
										sp.MenuItem().Label("Level 2 Item 1").Text("Level 2 Item 1"),
										sp.MenuItem().Label("Level 2 Item 2").Text("Level 2 Item 2"),
										sp.MenuItem().
											Label("Level 2 Item 3").
											Text("Level 2 Item 3").
											SetHasSubmenu(),
										app.Div().
											Attr("slot", "submenu").
											Body(
												sp.Menu().Body(
													sp.MenuItem().Label("Level 3 Item 1").Text("Level 3 Item 1"),
													sp.MenuItem().Label("Level 3 Item 2").Text("Level 3 Item 2"),
												),
											),
									),
								),
							sp.MenuItem().Label("Level 1 Item 3").Text("Level 1 Item 3"),
						),
				),
			),

		// Menu with action menu
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Action Menu"),
				app.Div().Class("component-row").Body(
					sp.ActionMenu().
						Label("Action Menu").
						Body(
							sp.MenuItem().Label("Action 1").Text("Action 1"),
							sp.MenuItem().Label("Action 2").Text("Action 2"),
							sp.MenuItem().Label("Action 3").Text("Action 3"),
						),
					sp.ActionMenu().
						Label("Action Menu with Icon").
						Icon(sp.Icon().Name("ui:MoreVertical")).
						Body(
							sp.MenuItem().Label("Action 1").Text("Action 1"),
							sp.MenuItem().Label("Action 2").Text("Action 2"),
							sp.MenuItem().Label("Action 3").Text("Action 3"),
						),
				),
			),
	)
}
