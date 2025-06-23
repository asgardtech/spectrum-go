package prism

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type ConfirmDialog struct {
	app.Compo

	title           string
	message         string
	confirmText     string
	cancelText      string
	confirmVariant  string
	visible         bool
	onConfirm       func(app.Context)
	onCancel        func(app.Context)
	onDismiss       func(app.Context)
	destructive     bool
}

func NewConfirmDialog() *ConfirmDialog {
	return &ConfirmDialog{
		confirmText:    "Confirm",
		cancelText:     "Cancel",
		confirmVariant: "primary",
		visible:        false,
		destructive:    false,
	}
}

func (cd *ConfirmDialog) Title(t string) *ConfirmDialog                    { cd.title = t; return cd }
func (cd *ConfirmDialog) Message(m string) *ConfirmDialog                 { cd.message = m; return cd }
func (cd *ConfirmDialog) ConfirmText(t string) *ConfirmDialog             { cd.confirmText = t; return cd }
func (cd *ConfirmDialog) CancelText(t string) *ConfirmDialog              { cd.cancelText = t; return cd }
func (cd *ConfirmDialog) ConfirmVariant(v string) *ConfirmDialog          { cd.confirmVariant = v; return cd }
func (cd *ConfirmDialog) Visible(v bool) *ConfirmDialog                   { cd.visible = v; return cd }
func (cd *ConfirmDialog) Destructive(d bool) *ConfirmDialog               { cd.destructive = d; return cd }
func (cd *ConfirmDialog) OnConfirm(fn func(app.Context)) *ConfirmDialog   { cd.onConfirm = fn; return cd }
func (cd *ConfirmDialog) OnCancel(fn func(app.Context)) *ConfirmDialog    { cd.onCancel = fn; return cd }
func (cd *ConfirmDialog) OnDismiss(fn func(app.Context)) *ConfirmDialog   { cd.onDismiss = fn; return cd }

func (cd *ConfirmDialog) Show() *ConfirmDialog {
	cd.visible = true
	return cd
}

func (cd *ConfirmDialog) Hide() *ConfirmDialog {
	cd.visible = false
	return cd
}

func (cd *ConfirmDialog) Render() app.UI {
	if !cd.visible {
		return app.Div()
	}

	return app.Div().
		Class("spectrum-Modal").
		Style("position", "fixed").
		Style("top", "0").
		Style("left", "0").
		Style("width", "100%").
		Style("height", "100%").
		Style("background-color", "rgba(0,0,0,0.5)").
		Style("display", "flex").
		Style("align-items", "center").
		Style("justify-content", "center").
		Style("z-index", "1000").
		Body(
			app.Div().
				Class("spectrum-Dialog").
				Style("background", "white").
				Style("border-radius", "8px").
				Style("padding", "24px").
				Style("max-width", "400px").
				Style("width", "90%").
				Body(
					app.H2().
						Class("spectrum-Dialog-title").
						Text(cd.title),
					app.P().
						Text(cd.message),
					app.Div().
						Class("spectrum-Dialog-actions").
						Style("display", "flex").
						Style("gap", "8px").
						Style("justify-content", "flex-end").
						Style("margin-top", "16px").
						Body(
							sp.Button().
								Text(cd.cancelText).
								OnClick(func(ctx app.Context, e app.Event) {
									cd.visible = false
									if cd.onCancel != nil {
										cd.onCancel(ctx)
									}
								}),
							sp.Button().
								Text(cd.confirmText).
								OnClick(func(ctx app.Context, e app.Event) {
									cd.visible = false
									if cd.onConfirm != nil {
										cd.onConfirm(ctx)
									}
								}),
						),
				),
		)
}

// Helper functions for common confirm dialogs
func ShowDeleteConfirm(itemName string, onConfirm func(app.Context)) *ConfirmDialog {
	return NewConfirmDialog().
		Title("Delete " + itemName).
		Message("Are you sure you want to delete this " + itemName + "? This action cannot be undone.").
		ConfirmText("Delete").
		CancelText("Cancel").
		Destructive(true).
		OnConfirm(onConfirm).
		Show()
}

func ShowActionConfirm(title, message string, onConfirm func(app.Context)) *ConfirmDialog {
	return NewConfirmDialog().
		Title(title).
		Message(message).
		OnConfirm(onConfirm).
		Show()
}

