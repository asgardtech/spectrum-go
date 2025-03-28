package sp

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// SpectrumTray represents an sp-tray component
type SpectrumTray struct {
	app.Compo

	// Properties
	open bool

	// Content
	content app.UI
	slot    string

	// Event handlers
	onClose app.EventHandler
}

// Tray creates a new tray component
func Tray() *SpectrumTray {
	return &SpectrumTray{}
}

// Open sets whether the tray is visible
func (t *SpectrumTray) Open(open bool) *SpectrumTray {
	t.open = open
	return t
}

// Content sets the content of the tray
func (t *SpectrumTray) Content(content app.UI) *SpectrumTray {
	t.content = content
	return t
}

// Slot sets the slot attribute of the tray
func (t *SpectrumTray) Slot(slot string) *SpectrumTray {
	t.slot = slot
	return t
}

// OnClose sets the close event handler
func (t *SpectrumTray) OnClose(handler app.EventHandler) *SpectrumTray {
	t.onClose = handler
	return t
}

// Render renders the tray component
func (t *SpectrumTray) Render() app.UI {
	tray := app.Elem("sp-tray")

	// Set attributes
	if t.open {
		tray = tray.Attr("open", true)
	}
	if t.slot != "" {
		tray = tray.Attr("slot", t.slot)
	}

	// Add event handlers
	if t.onClose != nil {
		tray = tray.On("close", t.onClose)
	}

	// Add content if provided
	if t.content != nil {
		tray = tray.Body(t.content)
	}

	return tray
}
