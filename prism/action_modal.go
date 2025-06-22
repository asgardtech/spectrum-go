package prism

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// ActionModal is a basic modal wrapper component (stub).

type ActionModal struct {
	app.Compo

	title       string
	dismissible bool
	body        app.UI
	onDismiss   func(app.Context)
}

func ActionModal() *ActionModal { return &ActionModal{dismissible: true} }

func (m *ActionModal) Title(t string) *ActionModal                 { m.title = t; return m }
func (m *ActionModal) Dismissible(b bool) *ActionModal             { m.dismissible = b; return m }
func (m *ActionModal) Body(ui app.UI) *ActionModal                 { m.body = ui; return m }
func (m *ActionModal) OnDismiss(fn func(app.Context)) *ActionModal { m.onDismiss = fn; return m }

func (m *ActionModal) Render() app.UI {
	// Real implementation will use sp.Dialog etc.
	return app.Div().Text("[Modal] " + m.title).Body(m.body)
}
