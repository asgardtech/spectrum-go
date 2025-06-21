package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/asgardtech/spectrum-go/prism"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type sidenavPage struct {
	app.Compo
}

func newSidenavPage() *sidenavPage {
	return &sidenavPage{}
}

func (p *sidenavPage) Render() app.UI {
	return prism.NewLayout().Content(
		app.H1().Text("Sidenav Component"),
		app.P().Text("Sidenav provides hierarchical navigation to different pages of your application."),

		// Basic sidenav
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Basic Sidenav"),
				app.Div().Class("component-row").Style("height", "320px").Body(
					sp.Sidenav().
						Label("Main Navigation").
						Style("width", "250px").
						Body(
							sp.SidenavItem().
								Label("Home").
								Value("home").
								Text("Home"),
							sp.SidenavItem().
								Label("Projects").
								Value("projects").
								Text("Projects"),
							sp.SidenavItem().
								Label("Tasks").
								Value("tasks").
								Text("Tasks"),
							sp.SidenavItem().
								Label("Reports").
								Value("reports").
								Text("Reports"),
							sp.SidenavItem().
								Label("Settings").
								Value("settings").
								Text("Settings"),
						),
				),
			),

		// Sidenav with selected item
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Sidenav with Selected Item"),
				app.Div().Class("component-row").Style("height", "320px").Body(
					sp.Sidenav().
						Label("Main Navigation").
						Value("projects").
						Style("width", "250px").
						Body(
							sp.SidenavItem().
								Label("Home").
								Value("home").
								Text("Home"),
							sp.SidenavItem().
								Label("Projects").
								Value("projects").
								Text("Projects"),
							sp.SidenavItem().
								Label("Tasks").
								Value("tasks").
								Text("Tasks"),
							sp.SidenavItem().
								Label("Reports").
								Value("reports").
								Text("Reports"),
							sp.SidenavItem().
								Label("Settings").
								Value("settings").
								Text("Settings"),
						),
				),
			),

		// Sidenav with links
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Sidenav with Links"),
				app.Div().Class("component-row").Style("height", "320px").Body(
					sp.Sidenav().
						Label("Site Navigation").
						Style("width", "250px").
						Body(
							sp.SidenavItem().
								Label("Home").
								Value("home").
								Text("Home").
								Href("#/home"),
							sp.SidenavItem().
								Label("Projects").
								Value("projects").
								Text("Projects").
								Href("#/projects"),
							sp.SidenavItem().
								Label("Tasks").
								Value("tasks").
								Text("Tasks").
								Href("#/tasks"),
							sp.SidenavItem().
								Label("GitHub").
								Value("github").
								Text("GitHub").
								Href("https://github.com").
								Target("_blank"),
							sp.SidenavItem().
								Label("Documentation").
								Value("docs").
								Text("Documentation").
								Href("https://example.com/docs").
								Target("_blank"),
						),
				),
			),

		// Sidenav with heading
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Sidenav with Heading"),
				app.Div().Class("component-row").Style("height", "350px").Body(
					app.Div().
						Style("width", "250px").
						Body(
							sp.SidenavHeading().
								Text("PROJECT NAVIGATION"),
							sp.Sidenav().
								Label("Project Navigation").
								Body(
									sp.SidenavItem().
										Label("Dashboard").
										Value("dashboard").
										Text("Dashboard"),
									sp.SidenavItem().
										Label("Team").
										Value("team").
										Text("Team"),
									sp.SidenavItem().
										Label("Documents").
										Value("documents").
										Text("Documents"),
								),
							sp.SidenavHeading().
								Text("SYSTEM"),
							sp.Sidenav().
								Label("System Navigation").
								Body(
									sp.SidenavItem().
										Label("Settings").
										Value("settings").
										Text("Settings"),
									sp.SidenavItem().
										Label("Help").
										Value("help").
										Text("Help"),
								),
						),
				),
			),

		// Sidenav with icons
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Sidenav with Icons"),
				app.Div().Class("component-row").Style("height", "320px").Body(
					sp.Sidenav().
						Label("Main Navigation").
						Style("width", "250px").
						Body(
							sp.SidenavItem().
								Label("Home").
								Value("home").
								Text("Home"),
							app.Elem("sp-icon").
								Attr("slot", "icon").
								Attr("name", "ui:Home"),
							sp.SidenavItem().
								Label("Projects").
								Value("projects").
								Text("Projects"),
							app.Elem("sp-icon").
								Attr("slot", "icon").
								Attr("name", "ui:FolderOpen"),
							sp.SidenavItem().
								Label("Tasks").
								Value("tasks").
								Text("Tasks"),
							app.Elem("sp-icon").
								Attr("slot", "icon").
								Attr("name", "ui:CheckmarkCircle"),
							sp.SidenavItem().
								Label("Reports").
								Value("reports").
								Text("Reports"),
							app.Elem("sp-icon").
								Attr("slot", "icon").
								Attr("name", "ui:GraphTrend"),
							sp.SidenavItem().
								Label("Settings").
								Value("settings").
								Text("Settings"),
							app.Elem("sp-icon").
								Attr("slot", "icon").
								Attr("name", "ui:Settings"),
						),
				),
			),

		// Multilevel sidenav
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Multilevel Sidenav"),
				app.Div().Class("component-row").Style("height", "400px").Body(
					sp.Sidenav().
						Label("Multilevel Navigation").
						Style("width", "250px").
						Variant(sp.SidenavVariantMultilevel).
						Body(
							sp.SidenavItem().
								Label("Dashboard").
								Value("dashboard").
								Text("Dashboard"),
							sp.SidenavItem().
								Label("Projects").
								Value("projects").
								SetExpanded().
								Text("Projects"),
							sp.Sidenav().
								Label("Project Submenu").
								Body(
									sp.SidenavItem().
										Label("Active").
										Value("active").
										Text("Active"),
									sp.SidenavItem().
										Label("Archived").
										Value("archived").
										Text("Archived"),
								),
							sp.SidenavItem().
								Label("Tasks").
								Value("tasks").
								Text("Tasks"),
							sp.Sidenav().
								Label("Task Submenu").
								Body(
									sp.SidenavItem().
										Label("Assigned to me").
										Value("my-tasks").
										Text("Assigned to me"),
									sp.SidenavItem().
										Label("All tasks").
										Value("all-tasks").
										Text("All tasks"),
									sp.SidenavItem().
										Label("Completed").
										Value("completed").
										Text("Completed"),
								),
							sp.SidenavItem().
								Label("Settings").
								Value("settings").
								Text("Settings"),
						),
				),
			),

		// Disabled items
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Disabled Items"),
				app.Div().Class("component-row").Style("height", "320px").Body(
					sp.Sidenav().
						Label("Main Navigation").
						Style("width", "250px").
						Body(
							sp.SidenavItem().
								Label("Home").
								Value("home").
								Text("Home"),
							sp.SidenavItem().
								Label("Projects").
								Value("projects").
								Text("Projects"),
							sp.SidenavItem().
								Label("Tasks").
								Value("tasks").
								Text("Tasks").
								Disabled(true),
							sp.SidenavItem().
								Label("Reports").
								Value("reports").
								Text("Reports").
								Disabled(true),
							sp.SidenavItem().
								Label("Settings").
								Value("settings").
								Text("Settings"),
						),
				),
			),
	)
}
