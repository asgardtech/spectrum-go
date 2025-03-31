package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type toastPage struct {
	app.Compo

	// Toast visibility state
	standardToastVisible   bool
	actionToastVisible     bool
	negativeToastVisible   bool
	positiveToastVisible   bool
	infoToastVisible       bool
	timeoutToastVisible    bool
	multiBtnToastVisible   bool
	customIconToastVisible bool
}

func newToastPage() *toastPage {
	return &toastPage{}
}

func (p *toastPage) OnNav(ctx app.Context) {
	p.initPage(ctx)
}

func (p *toastPage) initPage(ctx app.Context) {
	ctx.Page().SetTitle(defaultTitle)
	ctx.Page().SetDescription(defaultDescription)
}

func (p *toastPage) showToast(ctx app.Context, toastType string) {
	// Reset all toast visibility
	p.standardToastVisible = false
	p.actionToastVisible = false
	p.negativeToastVisible = false
	p.positiveToastVisible = false
	p.infoToastVisible = false
	p.timeoutToastVisible = false
	p.multiBtnToastVisible = false
	p.customIconToastVisible = false

	// Set the selected toast to visible
	switch toastType {
	case "standard":
		p.standardToastVisible = true
	case "action":
		p.actionToastVisible = true
	case "negative":
		p.negativeToastVisible = true
	case "positive":
		p.positiveToastVisible = true
	case "info":
		p.infoToastVisible = true
	case "timeout":
		p.timeoutToastVisible = true
		// Auto-hide the toast after timeout (shown in the toast)
		go func() {
			app.Window().Get("setTimeout").Invoke(func() {
				ctx.Dispatch(func(ctx app.Context) {
					p.timeoutToastVisible = false
					ctx.Reload()
				})
			}, 3000)
		}()
	case "multi-btn":
		p.multiBtnToastVisible = true
	case "custom-icon":
		p.customIconToastVisible = true
	}

	ctx.Reload()
}

func (p *toastPage) hideToast(ctx app.Context, toastType string) {
	// Hide the selected toast
	switch toastType {
	case "standard":
		p.standardToastVisible = false
	case "action":
		p.actionToastVisible = false
	case "negative":
		p.negativeToastVisible = false
	case "positive":
		p.positiveToastVisible = false
	case "info":
		p.infoToastVisible = false
	case "timeout":
		p.timeoutToastVisible = false
	case "multi-btn":
		p.multiBtnToastVisible = false
	case "custom-icon":
		p.customIconToastVisible = false
	}

	ctx.Reload()
}

func (p *toastPage) Render() app.UI {
	return newPage().Content(
		app.H1().Text("Toast Component"),
		app.P().Text("Toast components display brief, temporary notifications that don't require user action. They are useful for showing success messages, errors, or informational updates."),

		// Toast Container
		app.Div().
			ID("toast-container").
			Style("position", "fixed").
			Style("bottom", "20px").
			Style("right", "20px").
			Style("z-index", "1000").
			Body(
				// Standard Toast
				app.If(p.standardToastVisible,
					func() app.UI {
						return sp.Toast().
							Open(true).
							Style("margin-bottom", "10px").
							Text("This is a standard toast notification.").
							OnClose(func(ctx app.Context, e app.Event) {
								p.hideToast(ctx, "standard")
							})
					},
				),

				// Toast with Action Button
				app.If(p.actionToastVisible,
					func() app.UI {
						return app.Div().Body(
							sp.Toast().
								Open(true).
								Style("margin-bottom", "10px").
								Text("This toast has an action button.").
								OnClose(func(ctx app.Context, e app.Event) {
									p.hideToast(ctx, "action")
								}),
							app.Div().
								Attr("slot", "action").
								Body(
									sp.Button().
										Variant("secondary").
										Treatment("outline").
										Text("Action").
										OnClick(func(ctx app.Context, e app.Event) {
											p.hideToast(ctx, "action")
										}),
								),
						)
					},
				),

				// Negative Toast
				app.If(p.negativeToastVisible,
					func() app.UI {
						return sp.Toast().
							Open(true).
							Variant("negative").
							Style("margin-bottom", "10px").
							Text("This is a negative toast notification.").
							OnClose(func(ctx app.Context, e app.Event) {
								p.hideToast(ctx, "negative")
							})
					},
				),

				// Positive Toast
				app.If(p.positiveToastVisible,
					func() app.UI {
						return sp.Toast().
							Open(true).
							Variant("positive").
							Style("margin-bottom", "10px").
							Text("This is a positive toast notification.").
							OnClose(func(ctx app.Context, e app.Event) {
								p.hideToast(ctx, "positive")
							})
					},
				),

				// Info Toast
				app.If(p.infoToastVisible,
					func() app.UI {
						return sp.Toast().
							Open(true).
							Variant("info").
							Style("margin-bottom", "10px").
							Text("This is an info toast notification.").
							OnClose(func(ctx app.Context, e app.Event) {
								p.hideToast(ctx, "info")
							})
					},
				),

				// Toast with Timeout
				app.If(p.timeoutToastVisible,
					func() app.UI {
						return sp.Toast().
							Open(true).
							Timeout(3000).
							Style("margin-bottom", "10px").
							Text("This toast will disappear after 3 seconds.").
							OnClose(func(ctx app.Context, e app.Event) {
								p.hideToast(ctx, "timeout")
							})
					},
				),

				// Toast with Multiple Buttons
				app.If(p.multiBtnToastVisible,
					func() app.UI {
						return app.Div().Body(
							sp.Toast().
								Open(true).
								Variant("info").
								Style("margin-bottom", "10px").
								Text("This toast has multiple action buttons.").
								OnClose(func(ctx app.Context, e app.Event) {
									p.hideToast(ctx, "multi-btn")
								}),
							app.Div().
								Attr("slot", "action").
								Style("display", "flex").
								Style("gap", "8px").
								Body(
									sp.Button().
										Variant("secondary").
										Treatment("outline").
										Text("Accept").
										OnClick(func(ctx app.Context, e app.Event) {
											p.hideToast(ctx, "multi-btn")
										}),
									sp.Button().
										Variant("secondary").
										Treatment("outline").
										Text("Decline").
										OnClick(func(ctx app.Context, e app.Event) {
											p.hideToast(ctx, "multi-btn")
										}),
								),
						)
					},
				),

				// Toast with Custom Icon
				app.If(p.customIconToastVisible,
					func() app.UI {
						return sp.Toast().
							Open(true).
							Style("margin-bottom", "10px").
							IconLabel("Custom icon").
							Text("This toast has a custom icon.").
							OnClose(func(ctx app.Context, e app.Event) {
								p.hideToast(ctx, "custom-icon")
							})
					},
				),
			),

		// Toast Examples
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Toast Examples"),
				app.P().Text("Click the buttons below to display different types of toast notifications."),
				app.Div().
					Style("display", "flex").
					Style("flex-wrap", "wrap").
					Style("gap", "10px").
					Body(
						sp.Button().
							Variant("primary").
							Text("Standard Toast").
							OnClick(func(ctx app.Context, e app.Event) {
								p.showToast(ctx, "standard")
							}),
						sp.Button().
							Variant("primary").
							Text("Toast with Action").
							OnClick(func(ctx app.Context, e app.Event) {
								p.showToast(ctx, "action")
							}),
						sp.Button().
							Variant("negative").
							Text("Negative Toast").
							OnClick(func(ctx app.Context, e app.Event) {
								p.showToast(ctx, "negative")
							}),
						sp.Button().
							Variant("positive").
							Text("Positive Toast").
							OnClick(func(ctx app.Context, e app.Event) {
								p.showToast(ctx, "positive")
							}),
						sp.Button().
							Variant("primary").
							Text("Info Toast").
							OnClick(func(ctx app.Context, e app.Event) {
								p.showToast(ctx, "info")
							}),
						sp.Button().
							Variant("primary").
							Text("Toast with Timeout").
							OnClick(func(ctx app.Context, e app.Event) {
								p.showToast(ctx, "timeout")
							}),
						sp.Button().
							Variant("primary").
							Text("Toast with Multiple Buttons").
							OnClick(func(ctx app.Context, e app.Event) {
								p.showToast(ctx, "multi-btn")
							}),
						sp.Button().
							Variant("primary").
							Text("Toast with Custom Icon").
							OnClick(func(ctx app.Context, e app.Event) {
								p.showToast(ctx, "custom-icon")
							}),
					),
			),

		// Description of Toast Variants
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Toast Variants"),
				app.P().Text("Toasts come in several variants to indicate different types of messages:"),
				app.Ul().
					Body(
						app.Li().Text("Standard: Default toast used for general notifications."),
						app.Li().Text("Negative: Used for error messages or warnings."),
						app.Li().Text("Positive: Used for success messages or confirmations."),
						app.Li().Text("Info: Used for informational messages."),
					),
			),

		// Toast Features
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Toast Features"),
				app.P().Text("Toasts support various features:"),
				app.Ul().
					Body(
						app.Li().Text("Action Buttons: Allow users to take immediate action."),
						app.Li().Text("Auto-dismissal: Toast can automatically disappear after a set timeout."),
						app.Li().Text("Custom Icons: Toast can display different icons based on context."),
						app.Li().Text("Manual Dismissal: Users can close the toast using the close button."),
					),
			),
	)
}
