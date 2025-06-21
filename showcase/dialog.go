package main

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/asgardtech/spectrum-go/prism"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type dialogPage struct {
	app.Compo
	dialogOpen      bool
	wrapperOpen     bool
	confirmOpen     bool
	fullscreenOpen  bool
	takoverOpen     bool
	errorDialogOpen bool
}

func newDialogPage() *dialogPage {
	return &dialogPage{}
}

func (p *dialogPage) OnMount(ctx app.Context) {
	ctx.SetState("dialogOpen", p.dialogOpen)
	ctx.SetState("wrapperOpen", p.wrapperOpen)
	ctx.SetState("confirmOpen", p.confirmOpen)
	ctx.SetState("fullscreenOpen", p.fullscreenOpen)
	ctx.SetState("takoverOpen", p.takoverOpen)
	ctx.SetState("errorDialogOpen", p.errorDialogOpen)
}

func (p *dialogPage) updateState(ctx app.Context) {
	ctx.Reload()
}

func (p *dialogPage) Render() app.UI {
	return prism.NewLayout().Content(
		app.H1().Text("Dialog Component"),
		app.P().Text("Dialogs display important information that users need to acknowledge. They appear over the interface and block further interactions."),

		// Basic Dialog
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Basic Dialog"),
				app.Div().Class("component-row").Body(
					sp.Button().
						Text("Open Dialog").
						Variant(sp.ButtonVariantPrimary).
						OnClick(func(ctx app.Context, e app.Event) {
							p.dialogOpen = true

						}),
					func() app.UI {
						if p.dialogOpen {
							return sp.Dialog().
								SizeM().
								SetDismissable().
								Heading(app.H2().Text("Basic Dialog")).
								Body(
									app.P().Text("This is a basic dialog with dismissable set to true. Click the X to close it."),
								).
								OnClose(func(ctx app.Context, e app.Event) {
									p.dialogOpen = false

								})
						}
						return nil
					}(),
				),
			),

		// Dialog Wrapper
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Dialog Wrapper"),
				app.P().Text("Dialog Wrapper provides a simpler API for common dialog use cases."),
				app.Div().Class("component-row").Body(
					sp.Button().
						Text("Open Dialog Wrapper").
						Variant(sp.ButtonVariantPrimary).
						OnClick(func(ctx app.Context, e app.Event) {
							p.wrapperOpen = true

						}),
					func() app.UI {
						if p.wrapperOpen {
							return sp.DialogWrapper().
								SizeM().
								SetDismissable().
								SetOpen().
								Headline("Dialog Wrapper").
								Confirmlabel("OK").
								Cancellabel("Cancel").
								Text("This is a dialog wrapper that provides buttons and simplified management.").
								OnConfirm(func(ctx app.Context, e app.Event) {
									p.wrapperOpen = false

								}).
								OnCancel(func(ctx app.Context, e app.Event) {
									p.wrapperOpen = false

								}).
								OnClose(func(ctx app.Context, e app.Event) {
									p.wrapperOpen = false

								})
						}
						return nil
					}(),
				),
			),

		// Confirmation Dialog
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Confirmation Dialog"),
				app.Div().Class("component-row").Body(
					sp.Button().
						Text("Confirmation Dialog").
						Variant(sp.ButtonVariantPrimary).
						OnClick(func(ctx app.Context, e app.Event) {
							p.confirmOpen = true
						}),
					func() app.UI {
						if p.confirmOpen {
							return sp.Dialog().
								SizeM().
								SetDismissable().
								Heading(app.H2().Text("Confirm Action")).
								Body(
									app.P().Text("Are you sure you want to proceed with this action?"),
								).
								Button(
									app.Div().Body(
										sp.ButtonGroup().
											Body(
												sp.Button().
													Text("Cancel").
													Variant(sp.ButtonVariantSecondary).
													OnClick(func(ctx app.Context, e app.Event) {
														p.confirmOpen = false

													}),
												sp.Button().
													Text("Confirm").
													Variant(sp.ButtonVariantPrimary).
													OnClick(func(ctx app.Context, e app.Event) {
														p.confirmOpen = false

													}),
											),
									),
								).
								OnClose(func(ctx app.Context, e app.Event) {
									p.confirmOpen = false

								})
						}
						return nil
					}(),
				),
			),

		// Dialog Sizes
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Dialog Sizes"),
				app.Div().Class("component-row").Body(
					sp.Button().
						Text("Small Dialog").
						Variant(sp.ButtonVariantPrimary).
						Style("margin-right", "10px").
						OnClick(func(ctx app.Context, e app.Event) {
							ctx.Defer(func(context app.Context) {
								context.JSSrc().Set("currentDialog", context.JSSrc().Get("document").Call("querySelector", "sp-dialog[size=s]"))
								context.JSSrc().Call("setTimeout", app.FuncOf(func(this app.Value, args []app.Value) interface{} {
									context.JSSrc().Get("currentDialog").Set("style", "display: block;")
									return nil
								}), 50)
							})
						}),
					sp.Button().
						Text("Medium Dialog").
						Variant(sp.ButtonVariantPrimary).
						Style("margin-right", "10px").
						OnClick(func(ctx app.Context, e app.Event) {
							ctx.Defer(func(context app.Context) {
								context.JSSrc().Set("currentDialog", context.JSSrc().Get("document").Call("querySelector", "sp-dialog[size=m]"))
								context.JSSrc().Call("setTimeout", app.FuncOf(func(this app.Value, args []app.Value) interface{} {
									context.JSSrc().Get("currentDialog").Set("style", "display: block;")
									return nil
								}), 50)
							})
						}),
					sp.Button().
						Text("Large Dialog").
						Variant(sp.ButtonVariantPrimary).
						OnClick(func(ctx app.Context, e app.Event) {
							ctx.Defer(func(context app.Context) {
								context.JSSrc().Set("currentDialog", context.JSSrc().Get("document").Call("querySelector", "sp-dialog[size=l]"))
								context.JSSrc().Call("setTimeout", app.FuncOf(func(this app.Value, args []app.Value) interface{} {
									context.JSSrc().Get("currentDialog").Set("style", "display: block;")
									return nil
								}), 50)
							})
						}),
					sp.Dialog().
						SizeS().
						SetDismissable().
						Style("display", "none").
						Heading(app.H2().Text("Small Dialog")).
						Body(
							app.P().Text("This is a small dialog. It's good for simple messages and confirmations."),
						),
					sp.Dialog().
						SizeM().
						SetDismissable().
						Style("display", "none").
						Heading(app.H2().Text("Medium Dialog")).
						Body(
							app.P().Text("This is a medium dialog. It's the default size and works well for most content."),
						),
					sp.Dialog().
						SizeL().
						SetDismissable().
						Style("display", "none").
						Heading(app.H2().Text("Large Dialog")).
						Body(
							app.P().Text("This is a large dialog. Use it when you need to display a lot of content."),
						),
				),
			),

		// Fullscreen Dialogs
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Fullscreen Dialogs"),
				app.Div().Class("component-row").Body(
					sp.Button().
						Text("Fullscreen Dialog").
						Variant(sp.ButtonVariantPrimary).
						Style("margin-right", "10px").
						OnClick(func(ctx app.Context, e app.Event) {
							p.fullscreenOpen = true

						}),
					sp.Button().
						Text("Fullscreen Takeover").
						Variant(sp.ButtonVariantPrimary).
						OnClick(func(ctx app.Context, e app.Event) {
							p.takoverOpen = true

						}),
					func() app.UI {
						if p.fullscreenOpen {
							return sp.Dialog().
								SizeL().
								SetDismissable().
								ModeFullscreen().
								Heading(app.H2().Text("Fullscreen Dialog")).
								Body(
									app.P().Text("This dialog is displayed in fullscreen mode."),
								).
								Button(
									sp.Button().
										Text("Close").
										Variant(sp.ButtonVariantPrimary).
										OnClick(func(ctx app.Context, e app.Event) {
											p.fullscreenOpen = false

										}),
								).
								OnClose(func(ctx app.Context, e app.Event) {
									p.fullscreenOpen = false

								})
						}
						return nil
					}(),
					func() app.UI {
						if p.takoverOpen {
							return sp.Dialog().
								SizeL().
								ModeFullscreentakeover().
								Heading(app.H2().Text("Fullscreen Takeover")).
								Body(
									app.P().Text("This dialog takes over the entire screen. Note that the dismissable attribute is not honored in takeover mode."),
								).
								Button(
									sp.Button().
										Text("Close").
										Variant(sp.ButtonVariantPrimary).
										OnClick(func(ctx app.Context, e app.Event) {
											p.takoverOpen = false

										}),
								)
						}
						return nil
					}(),
				),
			),

		// Error Dialog
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Error Dialog"),
				app.Div().Class("component-row").Body(
					sp.Button().
						Text("Error Dialog").
						Variant(sp.ButtonVariantNegative).
						OnClick(func(ctx app.Context, e app.Event) {
							p.errorDialogOpen = true

						}),
					func() app.UI {
						if p.errorDialogOpen {
							return sp.Dialog().
								SizeM().
								SetDismissable().
								SetError().
								Heading(app.H2().Text("Error")).
								Body(
									app.P().Text("An error has occurred. Please try again later."),
								).
								Button(
									sp.Button().
										Text("OK").
										Variant(sp.ButtonVariantPrimary).
										OnClick(func(ctx app.Context, e app.Event) {
											p.errorDialogOpen = false

										}),
								).
								OnClose(func(ctx app.Context, e app.Event) {
									p.errorDialogOpen = false

								})
						}
						return nil
					}(),
				),
			),

		// Dialog with Hero
		app.Div().
			Class("showcase-section").
			Body(
				app.H2().Text("Dialog with Hero"),
				app.Div().Class("component-row").Body(
					sp.Button().
						Text("Dialog with Hero").
						Variant(sp.ButtonVariantPrimary).
						OnClick(func(ctx app.Context, e app.Event) {
							ctx.Defer(func(context app.Context) {
								context.JSSrc().Set("heroDialog", context.JSSrc().Get("document").Call("querySelector", "sp-dialog.hero-dialog"))
								context.JSSrc().Call("setTimeout", app.FuncOf(func(this app.Value, args []app.Value) interface{} {
									context.JSSrc().Get("heroDialog").Set("style", "display: block;")
									return nil
								}), 50)
							})
						}),
					sp.Dialog().
						SizeL().
						SetDismissable().
						SetNodivider().
						Class("hero-dialog").
						Style("display", "none").
						Hero(
							app.Div().
								Style("background-image", "url(https://picsum.photos/1400/260)").
								Style("height", "200px"),
						).
						Heading(app.H2().Text("Dialog with Hero Image")).
						Body(
							app.P().Text("This dialog includes a hero image at the top. The noDivider attribute is set to true to remove the divider between the header and content."),
						),
				),
			),
	)
}
