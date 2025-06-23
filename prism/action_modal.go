package prism

import (
	sp "github.com/asgardtech/spectrum-go"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type ActionModal struct {
	app.Compo

	title           string
	dismissible     bool
	body            app.UI
	actions         []app.UI
	onDismiss       func(app.Context)
	visible         bool
	size            string
	confirmText     string
	cancelText      string
	onConfirm       func(app.Context)
	onCancel        func(app.Context)
	confirmVariant  string
}

func NewActionModal() *ActionModal { 
	return &ActionModal{
		dismissible:    true,
		visible:        false,
		size:          "medium",
		confirmText:   "OK",
		cancelText:    "Cancel",
		confirmVariant: "primary",
	} 
}

func (m *ActionModal) Title(t string) *ActionModal                 { m.title = t; return m }
func (m *ActionModal) Dismissible(b bool) *ActionModal             { m.dismissible = b; return m }
func (m *ActionModal) Body(ui app.UI) *ActionModal                 { m.body = ui; return m }
func (m *ActionModal) Visible(v bool) *ActionModal                 { m.visible = v; return m }
func (m *ActionModal) Size(s string) *ActionModal                  { m.size = s; return m }
func (m *ActionModal) ConfirmText(t string) *ActionModal           { m.confirmText = t; return m }
func (m *ActionModal) CancelText(t string) *ActionModal            { m.cancelText = t; return m }
func (m *ActionModal) ConfirmVariant(v string) *ActionModal        { m.confirmVariant = v; return m }
func (m *ActionModal) OnDismiss(fn func(app.Context)) *ActionModal { m.onDismiss = fn; return m }
func (m *ActionModal) OnConfirm(fn func(app.Context)) *ActionModal { m.onConfirm = fn; return m }
func (m *ActionModal) OnCancel(fn func(app.Context)) *ActionModal  { m.onCancel = fn; return m }

func (m *ActionModal) AddAction(action app.UI) *ActionModal {
	m.actions = append(m.actions, action)
	return m
}

func (m *ActionModal) Render() app.UI {
	if !m.visible {
		return app.Div()
	}

	actions := m.actions
	if len(actions) == 0 {
		// Default actions
		if m.onCancel != nil {
			actions = append(actions, 
				sp.Button().
					Text(m.cancelText).
					OnClick(func(ctx app.Context, e app.Event) {
						if m.onCancel != nil {
							m.onCancel(ctx)
						}
					}),
			)
		}
		if m.onConfirm != nil {
			actions = append(actions, 
				sp.Button().
					Text(m.confirmText).
					OnClick(func(ctx app.Context, e app.Event) {
						if m.onConfirm != nil {
							m.onConfirm(ctx)
						}
					}),
			)
		}
	}

	return app.If(m.visible, func() app.UI {
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
			OnClick(func(ctx app.Context, e app.Event) {
				if m.dismissible && m.onDismiss != nil {
					m.onDismiss(ctx)
				}
			}).
			Body(
				app.Div().
					Class("spectrum-Dialog").
					Style("background", "white").
					Style("border-radius", "8px").
					Style("padding", "24px").
					Style("max-width", "500px").
					Style("width", "90%").
					OnClick(func(ctx app.Context, e app.Event) {
						e.PreventDefault()
					}).
					Body(
						app.H2().
							Class("spectrum-Dialog-title").
							Text(m.title),
						app.Div().
							Class("spectrum-Dialog-content").
							Body(m.body),
						app.Div().
							Class("spectrum-Dialog-actions").
							Style("display", "flex").
							Style("gap", "8px").
							Style("justify-content", "flex-end").
							Style("margin-top", "16px").
							Body(actions...),
					),
			)
	})
}
