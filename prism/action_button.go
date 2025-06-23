package prism

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// ActionButton is a thin wrapper around Spectrum's sp-action-button that
// accepts our typed Icon enum and exposes a Go-fluent API.
// It is *only a stub* – it renders a real Spectrum button but leaves many
// options unimplemented for now.

type ActionButton struct {
	app.Compo

	label   string
	icon    Icon
	handler app.EventHandler
}

func NewActionButton() *ActionButton { return &ActionButton{} }

func (b *ActionButton) Label(l string) *ActionButton { b.label = l; return b }
func (b *ActionButton) Icon(i Icon) *ActionButton    { b.icon = i; return b }
func (b *ActionButton) OnClick(h app.EventHandler) *ActionButton {
	b.handler = h
	return b
}

func (b *ActionButton) Render() app.UI {
	btn := sp.ActionButton().Label(b.label)
	if b.icon != "" {
		// the generated component expects an app.UI for the icon slot
		btn = btn.Icon(app.Text(b.icon.String()))
	}
	if b.handler != nil {
		btn = btn.OnClick(b.handler)
	}
	return btn
}
