package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/asgardtech/spectrum-go/prism"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type coachmarkPage struct {
	app.Compo
}

func newCoachmarkPage() *coachmarkPage {
	return &coachmarkPage{}
}

func (p *coachmarkPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *coachmarkPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle(defaultTitle)
	ctx.Page().SetDescription(defaultDescription)
}

func (p *coachmarkPage) Render() app.UI {
	return prism.NewLayout().Content(
		app.H1().Text("Coachmark Component"),
		app.P().Text("Coachmark is a temporary message that educates users through new or unfamiliar product experiences. They can be chained into a sequence to form a tour."),

		// Basic Coachmark
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Basic Coachmark"),
				app.P().Text("A simple coachmark with text-only content."),
				app.Div().Body(
					sp.Coachmark().
						Open(true).
						Body(
							app.Div().
								Attr("slot", "title").
								Text("Welcome to Spectrum Go!"),
							app.Div().
								Attr("slot", "content").
								Text("This is a basic coachmark that helps guide users through new features."),
						),
				),
			),

		// Coachmark with Action Menu
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Coachmark with Action Menu"),
				app.P().Text("Coachmarks can include an action menu for tour navigation."),
				app.Div().Body(
					sp.Coachmark().
						Open(true).
						Currentstep(2).
						Totalsteps(8).
						Primarycta("Next").
						Secondarycta("Previous").
						Body(
							app.Div().
								Attr("slot", "title").
								Text("Tour Navigation"),
							app.Div().
								Attr("slot", "content").
								Text("This coachmark demonstrates navigation controls and an action menu."),
							app.Div().
								Attr("slot", "actions").
								Body(
									sp.ActionMenu().
										Label("More Actions").
										Placement("bottom-end").
										Quiet(true).
										Body(
											sp.MenuItem().
												Text("Skip tour"),
											sp.MenuItem().
												Text("Restart tour"),
										),
								),
						),
				),
			),

		// User Action Dependent Coachmark
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("User Action Dependent Coachmark"),
				app.P().Text("This coachmark guides users based on their interactions and remains until a specific action is taken."),
				app.Div().Body(
					sp.Coachmark().
						Open(true).
						Currentstep(1).
						Totalsteps(3).
						Primarycta("Complete Action").
						Secondarycta("Previous").
						Body(
							app.Div().
								Attr("slot", "title").
								Text("Interactive Guide"),
							app.Div().
								Attr("slot", "content").
								Text("This coachmark will remain until you complete the required action."),
							app.Div().
								Attr("slot", "actions").
								Body(
									sp.ActionMenu().
										Label("More Actions").
										Placement("bottom-end").
										Quiet(true).
										Body(
											sp.MenuItem().
												Text("Skip tour"),
											sp.MenuItem().
												Text("Restart tour"),
										),
								),
						),
				),
			),

		// Coachmark with Media
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Coachmark with Media"),
				app.P().Text("Coachmarks can include images or GIFs to demonstrate features."),
				app.Div().Body(
					sp.Coachmark().
						Open(true).
						Currentstep(1).
						Totalsteps(4).
						Primarycta("Next").
						Secondarycta("Previous").
						Src("https://picsum.photos/id/237/200/300").
						Mediatype("image").
						Body(
							app.Div().
								Attr("slot", "title").
								Text("Feature Demonstration"),
							app.Div().
								Attr("slot", "content").
								Text("This coachmark includes an image to help illustrate the feature."),
							app.Div().
								Attr("slot", "actions").
								Body(
									sp.ActionMenu().
										Label("More Actions").
										Placement("bottom-end").
										Quiet(true).
										Body(
											sp.MenuItem().
												Text("Skip tour"),
											sp.MenuItem().
												Text("Restart tour"),
										),
								),
						),
				),
			),

		// Coachmark with Shortcut Keys
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Coachmark with Shortcut Keys"),
				app.P().Text("Coachmarks can display keyboard shortcuts and modifier keys."),
				app.Div().Body(
					sp.Coachmark().
						Open(true).
						Currentstep(1).
						Totalsteps(2).
						Primarycta("Next").
						Secondarycta("Previous").
						Modifierkeys([]string{"⇧ Shift", "⌘"}).
						Body(
							app.Div().
								Attr("slot", "title").
								Text("Keyboard Shortcuts"),
							app.Div().
								Attr("slot", "content").
								Text("This coachmark shows keyboard shortcuts for quick actions."),
							app.Div().
								Attr("slot", "actions").
								Body(
									sp.ActionMenu().
										Label("More Actions").
										Placement("bottom-end").
										Quiet(true).
										Body(
											sp.MenuItem().
												Text("Skip tour"),
											sp.MenuItem().
												Text("Restart tour"),
										),
								),
						),
				),
			),
	)
}
